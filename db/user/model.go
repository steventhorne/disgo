package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id              uuid.UUID `db:"id" json:"id"`
	DiscordId       string    `db:"discord_id" json:"discordId"`
	DiscordUsername string    `db:"discord_username" json:"discordUsername"`
	DiscordTag      string    `db:"discord_tag" json:"discordTag"`
	Created         time.Time `db:"created" json:"created"`
	Modified        time.Time `db:"modified" json:"modified"`
}
