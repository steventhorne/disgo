package core

import (
	"errors"
	"log/slog"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	ErrSessionNotOpen = errors.New("session not open")
)

type Bot struct {
	session *discordgo.Session

	router   router
	handlers []func()
}

// NewBot creates a new Bot.
func NewBot() *Bot {
	return &Bot{
		router:   newRouter(),
		handlers: make([]func(), 2),
	}
}

func (b *Bot) Close() {
	for _, h := range b.handlers {
		if h != nil {
			h()
		}
	}

	if b.session != nil {
		b.session.Close()
	}
}

func (b *Bot) Open() {
	sess, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		panic(err)
	}
	err = sess.Open()
	if err != nil {
		panic(err)
	}

	b.session = sess
}

func (b *Bot) OpenAndRun() {
	if b.session == nil {
		b.Open()
	}

	b.handlers[0] = b.session.AddHandlerOnce(b.router.ready)
	b.handlers[1] = b.session.AddHandler(b.router.interactionCreate)
}

func (b *Bot) RegisterCommands() []error {
	discordAppId := os.Getenv("DISCORD_APPLICATION_ID")
	errs := make([]error, 0, len(b.router.commands))

	existingCmds := make(map[string]string, len(b.router.commands))
	cmds, err := b.session.ApplicationCommands(discordAppId, "")
	if err != nil {
		return []error{err}
	}
	for _, c := range cmds {
		existingCmds[c.Name] = c.ID
	}

	for _, c := range b.router.commands {
		if existingID, ok := existingCmds[c.Name]; ok {
			slog.Info("Updating command", "name", c.Name)
			_, err = b.session.ApplicationCommandEdit(discordAppId, "", existingID, &c.ApplicationCommand)
			if err != nil {
				errs = append(errs, err)
			}
			delete(existingCmds, c.Name)
			continue
		}
		slog.Info("Creating command", "name", c.Name)
		_, err = b.session.ApplicationCommandCreate(discordAppId, "", &c.ApplicationCommand)
		if err != nil {
			errs = append(errs, err)
		}
		delete(existingCmds, c.Name)
	}

	for name, ID := range existingCmds {
		slog.Info("Deleting command", "name", name)
		err = b.session.ApplicationCommandDelete(discordAppId, "", ID)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
