package transcription

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"video-transcription/config"
	"video-transcription/shared"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func TranscribeVideo(cttx *gin.Context) {
	// Get the video file from the request
	file, header, err := cttx.Request.FormFile("video")
	if err != nil {
		cttx.JSON(400, gin.H{"error": "video file is required"})
		return
	}
	defer file.Close()

	// Create a new Resty client
	client := resty.New()

	// Set the API key
	apiKey := config.Cfg.API_KEY
	client.SetAuthToken(apiKey)

	transcriptionResp, err := sendTranscriptionRequest(client, header.Filename, file)
	if result := errorAndStatusCheck(err, transcriptionResp); result.Error != "" {
		cttx.JSON(400, gin.H{"result": result})
		return
	}

	// Parse the response body
	var transcriptionResponse shared.Transcription
	parseResponse(transcriptionResp, cttx, &transcriptionResponse)

	// Print and return the transcription text
	segments := transcriptionResponse.Segments
	log.Println("Transcription:", segments)

	completionResponse, err := sendContentGenerationRequest(client, transcriptionResponse.Text)
	errorAndStatusCheck(err, completionResponse)
	if result := errorAndStatusCheck(err, completionResponse); result.Error != "" {
		cttx.JSON(400, gin.H{"result": result})
		return
	}

	var chatCompletion shared.ChatCompletion
	parseResponse(completionResponse, cttx, &chatCompletion)

	cttx.JSON(http.StatusOK, gin.H{"result": createResult(&transcriptionResponse, chatCompletion)})
}

func sendTranscriptionRequest(client *resty.Client, fileName string, file multipart.File) (*resty.Response, error) {
	// Set the API transcriptionEndpoint
	transcriptionEndpoint := config.Cfg.TRANSCRIPTION_ENDPOINT

	// Create a new request
	return client.R().
		SetHeader("Content-Type", "application/form-data").
		SetFormData(map[string]string{
			"model":           config.Cfg.MODEL,
			"response_format": config.Cfg.RESPONSE_FORMAT,
		}).
		SetFormDataFromValues(url.Values{
			"timestamp_granularities": []string{config.Cfg.TIMESTAMP_GRANULARITIES},
		}).
		SetFileReader("file", fileName, file).
		Post(transcriptionEndpoint)
}

func sendContentGenerationRequest(client *resty.Client, transcriptionText string) (*resty.Response, error) {
	return client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model": "gpt-4-turbo",
			"messages": []map[string]interface{}{
				{
					"role": "user",
					"content": []map[string]interface{}{
						{
							"type": "text",
							"text": config.Cfg.COMPLETION_MESSAGE,
						},
						{
							"type": "text",
							"text": transcriptionText,
						},
					},
				},
			},
			"max_tokens": 300,
		}).
		Post(config.Cfg.COMPLETION_ENDPOINT)
}

func errorAndStatusCheck(err error, response *resty.Response) shared.Result {
	var result shared.Result
	if err != nil {
		log.Printf("Failed to send request: %v\n", err)
		result.Error = "Error occured"
	}

	// Check if the request was successful
	if response.StatusCode() != 200 {
		log.Printf("Failed to get transcript: %s\n", response.Status())
		result.Error = "Error occured"
	}
	return result
}

func parseResponse(response *resty.Response, cttx *gin.Context, v any) {
	err := json.Unmarshal(response.Body(), v)
	if err != nil {
		log.Printf("Failed to parse response: %v\n", err)
		cttx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
	}
}

func createResult(transcriptionResponse *shared.Transcription, completionResponse shared.ChatCompletion) shared.Result {
	var result shared.Result
	videoContent := completionResponse.Choices[0].Message.Content
	contentItself, hashTags, _ := strings.Cut(videoContent, "###")
	result.Transcription = transcriptionResponse.Segments
	result.Content = contentItself
	result.HashTags = hashTags
	return result
}
