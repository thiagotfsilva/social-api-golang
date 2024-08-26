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
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
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
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
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
		"select id, name, nick, email, createdAt from users where id = ?",
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
		"update users set name = ?, nick = ?, email = ? where id = ?",
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
	statement, err := u.db.Prepare("delete from users where id = ?")
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

func (u userRepository) FindByEmail(email string) (models.User, error) {
	line, err := u.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (u userRepository) FollowUser(userId, followerId uint64) error {
	statement, err := u.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userId, followerId)
	if err != nil {
		return err
	}

	return nil
}

func (u userRepository) UnfollowUser(userId, followerId uint64) error {
	statement, err := u.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userId, followerId)
	if err != nil {
		return err
	}

	return nil
}

// GetFollowers traz todos os seguidores de um usuário
func (u userRepository) GetFollowers(userId uint64) ([]models.User, error) {
	line, err := u.db.Query(`
    select u.id, u.name, u.nick, u.email, u.createdAt
    from users u inner join followers s on u.id = s.follower_id
    where s.user_id = ?
  `, userId)
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

// GetFollowing traz todos os usuários que um determinado usuário está seguindo
func (u userRepository) GetFollowing(userId uint64) ([]models.User, error) {
	line, err := u.db.Query(`
    select u.id, u.name, u.nick, u.email, u.createdAt
    from users u inner join followers s on u.id = s.user_id
    where s.follower_id = ?
  `, userId)
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

	return users, err
}

func (u userRepository) GetUserPassword(userId uint64) (string, error) {
	line, err := u.db.Query("select password from users where id = ?", userId)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil

}

func (u userRepository) UpdatePassword(userId uint64, password string) error {
	statement, err := u.db.Prepare(
		"update users set password = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userId); err != nil {
		return err
	}

	return nil
}
