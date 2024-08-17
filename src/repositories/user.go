package repositories

import (
	"api-devbook/src/models"
	"database/sql"
	"fmt"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (u userRepository) Create(user models.User) (uint64, error) {
	statement, err := u.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil

}

func (u userRepository) Fetch(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) //%nameOrNick%
	line, err := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer line.Close()

	var users []models.User

	for line.Next() {
		var user models.User

		if err = line.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u userRepository) Find(userId uint64) (models.User, error) {
	line, err := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		userId,
	)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (u userRepository) Update(userId uint64, user models.User) error {
	statament, err := u.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statament.Close()

	_, err = statament.Exec(user.Name, user.Nick, user.Email, userId)
	if err != nil {
		return err
	}

	return nil
}

func (u userRepository) Delete(userId uint64) error {
	statement, err := u.db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userId)
	if err != nil {
		return err
	}

	return nil
}
