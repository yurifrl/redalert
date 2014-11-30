package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type SlackWebhook struct {
	url string
}

func (a SlackWebhook) Trigger(event *Event) error {

	message := SlackPayload{
		Channel:   "#general",
		Username:  "redalert",
		Text:      event.ShortMessage(),
		Parse:     "full",
		IconEmoji: ":rocket:",
	}

	buf, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(a.url, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Not OK")
	}

	event.Server.log.Println(white, "Slack alert successfully triggered.", reset)
	return nil
}

type SlackPayload struct {
	Channel   string `json:"channel"`
	Username  string `json:"username,omitempty"`
	Text      string `json:"text"`
	Parse     string `json:"parse"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}
