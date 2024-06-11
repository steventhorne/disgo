package core

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

type router struct {
	commands map[string]command
}

func newRouter() router {
	return router{
		commands: map[string]command{
			"ping": {
				ApplicationCommand: discordgo.ApplicationCommand{
					Name:        "ping",
					Description: "Ping the bot",
					Type:        discordgo.ChatApplicationCommand,
					Options:     nil,
				},
				Handler: func(s *discordgo.Session, e *discordgo.InteractionCreate) {
					s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Pong!",
							Flags:   discordgo.MessageFlagsEphemeral,
						},
					})
				},
			},
		},
	}
}

func (r *router) ready(s *discordgo.Session, e *discordgo.Ready) {
	slog.Info("Bot is ready")
}

func (r *router) interactionCreate(s *discordgo.Session, e *discordgo.InteractionCreate) {
	if cmd, ok := r.commands[e.ApplicationCommandData().Name]; ok {
		cmd.Handler(s, e)
		return
	}
}
