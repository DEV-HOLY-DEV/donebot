package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SendDiscordNotification sends a notification to a Discord webhook.
func SendDiscordNotification(webhookURL string, cmd string, success bool) error {
	status := "Success"
	if !success {
		status = "Failed"
	}
	message := fmt.Sprintf("Command `%s` finished with status: %s", cmd, status)

	payload := map[string]string{
		"content": message,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload: %w", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error sending to Discord: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("discord returned status: %s", resp.Status)
	}

	return nil
}
