package Discord

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"scheduler-backup-postgresql/Config"

	"github.com/rs/zerolog/log"
)

func SendToDiscord(filePath string, env Config.Env) {

	webhookURL := env.Discord.Webhook

	// Open backup file
	file, err := os.Open(filePath)
	if err != nil {
		log.Error().Msgf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Create buffer and multipart writer
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add file to multipart form
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		log.Error().Msgf("Failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		log.Error().Msgf("Failed to copy file: %v", err)
	}

	// finish multipart writer
	err = writer.Close()
	if err != nil {
		log.Error().Msgf("Failed to close writer: %v", err)
	}

	// Request
	req, err := http.NewRequest("POST", webhookURL, &requestBody)
	if err != nil {
		log.Error().Msgf("Failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send Request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("Failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().Msgf("Failed to send backup to Discord, status code: %v", resp.StatusCode)
	}

	log.Info().Msg("Backup file sent to Discord successfully.")
}
