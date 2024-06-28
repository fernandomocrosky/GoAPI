package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db: db}
}

func (u users) Create(user models.User) (uint64, error) {
	statement, err := u.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

func (u users) Get(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, err := u.db.Query(
		"SELECT id, name, nick, email, created_at FROM users WHERE name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User
		err = lines.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u users) GetByID(id uint64) (models.User, error) {
	lines, err := u.db.Query(
		"SELECT id, name, nick, email, created_at FROM users WHERE id = ?",
		id,
	)

	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err := lines.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
		return user, nil
	}

	return models.User{}, errors.New("user not found")
}
