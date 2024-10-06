package repositories

import (
	"api-devbook/src/models"
	"database/sql"
)

type PublicationRepository struct {
	db *sql.DB
}

func NewPublicationRepository(db *sql.DB) *PublicationRepository {
	return &PublicationRepository{db}
}

func (p PublicationRepository) Create(publication models.Publication) (uint64, error) {
	statemante, err := p.db.Prepare(
		"insert into publications (title, content, author_id) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statemante.Close()

	result, err := statemante.Exec(
		publication.Title,
		publication.Content,
		publication.AuthorId,
	)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}
