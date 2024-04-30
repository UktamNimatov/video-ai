package transcription

import (
	"encoding/json"
	"log"
	"net/http"
	"video-transcription/shared"
	"video-transcription/config"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func TranscribeVideo(cttx *gin.Context) {
			// Get the video file from the request
			file, header, err := cttx.Request.FormFile("video")
			if err != nil {
				cttx.JSON(400, gin.H{"error": "video file is required"})
				// return nil, err
			}
			defer file.Close()
	
			// Create a new Resty client
			client := resty.New()
	
			// Set the API key
			apiKey := config.Cfg.API_KEY
			client.SetAuthToken(apiKey)
	
			// Set the API endpoint
			endpoint := config.Cfg.TRANSCRIPTION_ENDPOINT
	
			// Create a new request
			fileName := header.Filename
			resp, err := client.R().
				SetHeader("Content-Type", "application/form-data").
				SetFormData(map[string]string{
					"model": config.Cfg.MODEL,
				}).
				SetFileReader("file", fileName, file).
				Post(endpoint)
			if err != nil {
				log.Printf("Failed to send request: %v\n", err)
				cttx.JSON(500, gin.H{"error": "failed to translate video"})
				// return nil, err
			}
	
			// Check if the request was successful
			if resp.StatusCode() != 200 {
				log.Printf("Failed to get transcript: %s\n", resp.Status())
				cttx.JSON(500, gin.H{"error": "failed to transcript video"})
				// return nil, err
			}
	
			// Parse the response body
			var transcriptionResponse shared.TranscriptionResponse
			err = json.Unmarshal(resp.Body(), &transcriptionResponse)
			if err != nil {
				log.Printf("Failed to parse response: %v\n", err)
				cttx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
				// return nil, err
			}
	
			// Print and return the transcription text
			transcriptionText := transcriptionResponse.Text
			log.Println("Transcription:", transcriptionText)
	
			cttx.JSON(http.StatusOK, gin.H{"transcription": transcriptionText})
			// return &shared.TranscriptionResponse{Text: transcriptionText}, nil
}