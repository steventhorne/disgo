package core

import "github.com/bwmarrin/discordgo"

type commandType int

const (
	commandTypeChatInput commandType = iota
	commandTypeUser
	commandTypeMessage
)

type commandOptionType int

const (
	commandOptionTypeSubcommand commandOptionType = iota
	commandOptionTypeSubcommandGroup
	commandOptionTypeString
	commandOptionTypeInteger
	commandOptionTypeBoolean
	commandOptionTypeUser
	commandOptionTypeChannel
	commandOptionTypeRole
	commandOptionTypeMentionable
	commandOptionTypeNumber
	commandOptionTypeAttachment
)

type command struct {
	discordgo.ApplicationCommand

	Handler func(s *discordgo.Session, e *discordgo.InteractionCreate)
}

type commandOption struct {
	Name        string                `json:"name"`
	OptionType  commandOptionType     `json:"type"`
	Description string                `json:"description"`
	Required    bool                  `json:"required"`
	Choices     []commandOptionChoice `json:"choices"`
}

type commandOptionChoice struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
