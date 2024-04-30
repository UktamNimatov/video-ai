package ginreg

import (
	"github.com/gin-gonic/gin"
	"video-transcription/config"
	"video-transcription/transcription"
)

func GinInit() {
	//Initialize the gin router
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		video := v1.Group("/video")
		{
			video.POST("/transcribe", transcription.TranscribeVideo)
		}
	}

	router.Run(config.Cfg.SERVER_URL)
}
