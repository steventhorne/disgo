package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/steventhorne/disgo/db"
)

// CreateUser creates a new user in the database.
func CreateUser(user User) (User, error) {
	pool, err := db.GetPool()
	if err != nil {
		return User{}, err
	}

	rows, err := pool.Query(context.Background(), `
		INSERT INTO users (discord_id, discord_username, discord_tag)
		VALUES ($1, $2, $3)
		RETURNING *
	`, user.DiscordId, user.DiscordUsername, user.DiscordTag)
	if err != nil {
		return User{}, err
	}

	newUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

// GetUserById retrieves a user from the database by ID.
func GetUserById(id uuid.UUID) (User, error) {
	pool, err := db.GetPool()
	if err != nil {
		return User{}, err
	}

	rows, err := pool.Query(context.Background(), `
		SELECT * FROM users WHERE id = $1
	`, id)
	if err != nil {
		return User{}, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// GetUserByDiscordId retrieves a user from the database by Discord ID.
func GetUserByDiscordId(discordId string) (User, error) {
	pool, err := db.GetPool()
	if err != nil {
		return User{}, err
	}

	rows, err := pool.Query(context.Background(), `
		SELECT * FROM users WHERE discord_id = $1
	`, discordId)
	if err != nil {
		return User{}, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// GetUserByDiscordTag retrieves a user from the database by Discord tag.
func GetUserByDiscordTag(discordTag string) (User, error) {
	pool, err := db.GetPool()
	if err != nil {
		return User{}, err
	}

	rows, err := pool.Query(context.Background(), `
		SELECT * FROM users WHERE discord_tag = $1
	`, discordTag)
	if err != nil {
		return User{}, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// UpdateUser updates a user in the database.
func UpdateUser(user User) (User, error) {
	pool, err := db.GetPool()
	if err != nil {
		return User{}, err
	}

	rows, err := pool.Query(context.Background(), `
		UPDATE users
		SET discord_username = $1, discord_tag = $2
		WHERE id = $3
		RETURNING *
	`, user.DiscordUsername, user.DiscordTag, user.DiscordId)
	if err != nil {
		return User{}, err
	}

	updatedUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return User{}, err
	}

	return updatedUser, nil
}

// DeleteUser deletes a user from the database.
func DeleteUser(id uuid.UUID) error {
	pool, err := db.GetPool()
	if err != nil {
		return err
	}

	_, err = pool.Exec(context.Background(), `
		DELETE FROM users WHERE id = $1
	`, id)
	if err != nil {
		return err
	}

	return nil
}
