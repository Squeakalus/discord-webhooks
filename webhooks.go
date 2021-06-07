package webhooks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type WebhookContent struct {
	Content   string         `json:"content,omitempty"`
	Username  string         `json:"username,omitempty"`
	AvatarURL string         `json:"avatar_url,omitempty"`
	Embeds    []DiscordEmbed `json:"embeds"`
}

type DiscordEmbed struct {
	Title       string         `json:"title,omitempty"`
	TitleURL    string         `json:"url,omitempty"`
	Description string         `json:"description,omitempty"`
	Colour      int64          `json:"color,omitempty"`
	Author      EmbedAuthor    `json:"author"`
	Footer      EmbedFooter    `json:"footer"`
	Thumbnail   EmbedThumbnail `json:"thumbnail"`
	Fields      []EmbedField   `json:"fields,omitempty"`
	Timestamp   string         `json:"timestamp,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type EmbedAuthor struct {
	Text    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type EmbedFooter struct {
	Text    string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type EmbedThumbnail struct {
	ImageURL string `json:"url,omitempty"`
}

func NewEmbed() DiscordEmbed {
	return DiscordEmbed{Footer: EmbedFooter{}}
}

func (Embed *DiscordEmbed) AddField(name, value string, inline bool) {
	embedField := EmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	}
	Embed.Fields = append(Embed.Fields, embedField)
}

func (Embed *DiscordEmbed) SetAuthor(author, url, iconURL string) {
	Embed.Author = EmbedAuthor{Text: author, URL: url, IconURL: iconURL}
}

func (Embed *DiscordEmbed) SetColour(hexCode string) {
	colourDecimal, err := strconv.ParseInt(hexCode, 16, 64)
	if err != nil {
		Embed.Colour = 0
		return
	}
	Embed.Colour = colourDecimal
}

func (Embed *DiscordEmbed) SetFooter(footer, iconURL string) {
	Embed.Footer = EmbedFooter{Text: footer, IconURL: iconURL}
}

func (Embed *DiscordEmbed) SetTimestamp() {
	Embed.Timestamp = time.Now().Format(time.RFC3339)
}

func (Embed *DiscordEmbed) SetThumbnail(imageURL string) {
	Embed.Thumbnail = EmbedThumbnail{ImageURL: imageURL}
}

func (Embed DiscordEmbed) Send(webhookURL, message, username, avatarURL string) error {
	webhookData := WebhookContent{Content: message, Username: username, AvatarURL: avatarURL, Embeds: []DiscordEmbed{Embed}}
	webhookDataBytes, err := json.Marshal(webhookData)
	if err != nil {
		return err
	}
	webhookPostResponse, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(webhookDataBytes))
	if err != nil {
		return err
	}
	if webhookPostResponse.Body == nil {
		return errors.New("webhook response body is nil")
	}
	defer webhookPostResponse.Body.Close()
	if webhookPostResponse.StatusCode != 204 {
		return fmt.Errorf("invalid webhook response status code %d", webhookPostResponse.StatusCode)
	}
	return nil
}
