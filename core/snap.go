package core

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Card struct {
	Id          int    `json:"id"`
	DefId       string `json:"defId"`
	Cost        int    `json:"cost"`
	Power       int    `json:"power"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Series      int    `json:"series"`
}

type Location struct {
	Id          int    `json:"id"`
	DefId       string `json:"defId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var cardMap map[string]Card
var locationMap map[string]Location

func snapHandler(s *discordgo.Session, e *discordgo.InteractionCreate) {
	switch e.ApplicationCommandData().Options[0].Name {
	case "card":
		snapCardHandler(s, e)
	case "location":
		snapLocationHandler(s, e)
	}
}

func snapLocationHandler(s *discordgo.Session, e *discordgo.InteractionCreate) {
	var err error
	if locationMap == nil {
		locationMap, err = getLocations()
		if err != nil {
			s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Failed to load locations",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			slog.Error("Failed to load locations", "error", err)
			return
		}
	}

	if len(locationMap) == 0 {
		s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Locations not loaded",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	var location Location
	var ok bool
	lowerParam := strings.ToLower(e.ApplicationCommandData().Options[0].Options[0].StringValue())
	if location, ok = locationMap[lowerParam]; !ok {
		s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Location not found",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       location.Name,
					Description: htmlToMarkdown(location.Description),
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: fmt.Sprintf("https://snapjson.untapped.gg/art/loc/%s.webp", location.DefId),
					},
				},
			},
		},
	})
}

func snapCardHandler(s *discordgo.Session, e *discordgo.InteractionCreate) {
	var err error
	if cardMap == nil {
		cardMap, err = getCards()
		if err != nil {
			s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Failed to load cards",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			slog.Error("Failed to load cards", "error", err)
			return
		}
	}

	if len(cardMap) == 0 {
		s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Cards not loaded",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	var card Card
	var ok bool
	lowerParam := strings.ToLower(e.ApplicationCommandData().Options[0].Options[0].StringValue())
	if card, ok = cardMap[lowerParam]; !ok {
		s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Card not found",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	var series string
	switch card.Series {
	case 0:
		series = "Unobtainable"
	case 1:
		series = "Starter"
	case 2:
		series = "0 [Level 1-14]"
	case 3:
		series = "Recruit Season"
	case 4:
		series = "1 [Level 18-214]"
	case 5:
		series = "2 [Level 222-486]"
	case 6:
		series = "3 [Level 486+]"
	case 7:
		series = "4"
	case 8:
		series = "5"
	case 9:
		series = "Season Pass"
	}

	s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       card.Name,
					Description: htmlToMarkdown(card.Description),
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: fmt.Sprintf("https://snapjson.untapped.gg/art/render/framebreak/common/512/%s.webp", card.DefId),
					},
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Cost",
							Value:  fmt.Sprintf("%d", card.Cost),
							Inline: true,
						},
						{
							Name:   "Power",
							Value:  fmt.Sprintf("%d", card.Power),
							Inline: true,
						},
						{
							Name:   "Series",
							Value:  series,
							Inline: true,
						},
					},
				},
			},
		},
	})
}

func getCards() (map[string]Card, error) {
	url := "https://snapjson.untapped.gg/v2/latest/en/cards.json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "snap.fan")
	req.Header.Set("Priority", "u=0, i")
	req.Header.Set("Referer", "https://snap.fan/")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:128.0) Gecko/20100101 Firefox/128.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get cards: %s", resp.Status)
	}

	var cards []Card
	err = json.NewDecoder(resp.Body).Decode(&cards)
	if err != nil {
		return nil, err
	}

	cm := make(map[string]Card, len(cards))
	for _, card := range cards {
		cm[strings.ToLower(card.Name)] = card
	}
	return cm, nil
}

func getLocations() (map[string]Location, error) {
	url := "https://snapjson.untapped.gg/v2/latest/en/locations.json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "snap.fan")
	req.Header.Set("Priority", "u=0, i")
	req.Header.Set("Referer", "https://snap.fan/")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:128.0) Gecko/20100101 Firefox/128.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get locations: %s", resp.Status)
	}

	var locations []Location
	err = json.NewDecoder(resp.Body).Decode(&locations)
	if err != nil {
		return nil, err
	}

	lm := make(map[string]Location, len(locations))
	for _, card := range locations {
		lm[strings.ToLower(card.Name)] = card
	}
	return lm, nil
}

func htmlToMarkdown(text string) string {
	text = strings.ReplaceAll(text, "<br>", "\n")
	text = strings.ReplaceAll(text, "<br/>", "\n")
	text = strings.ReplaceAll(text, "<br />", "\n")
	text = strings.ReplaceAll(text, "<p>", "")
	text = strings.ReplaceAll(text, "</p>", "")
	text = strings.ReplaceAll(text, "<strong>", "**")
	text = strings.ReplaceAll(text, "<b>", "**")
	text = strings.ReplaceAll(text, "</strong>", "**")
	text = strings.ReplaceAll(text, "</b>", "**")
	text = strings.ReplaceAll(text, "<em>", "*")
	text = strings.ReplaceAll(text, "<i>", "*")
	text = strings.ReplaceAll(text, "</em>", "*")
	text = strings.ReplaceAll(text, "</i>", "*")
	text = strings.ReplaceAll(text, "<a href=\"", "[")
	text = strings.ReplaceAll(text, "\">", "](")
	text = strings.ReplaceAll(text, "</a>", ")")
	text = strings.ReplaceAll(text, "<ul>", "")
	text = strings.ReplaceAll(text, "</ul>", "")
	text = strings.ReplaceAll(text, "<li>", "- ")
	text = strings.ReplaceAll(text, "</li>", "")
	return text
}
