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
	statement, err := p.db.Prepare(
		"insert into publications (title, content, author_id) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(
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

func (p PublicationRepository) FindById(publicationId uint64) (models.Publication, error) {
	line, err := p.db.Query(`
    select p.*, u.nick from
    publications p inner join users u
    on u.id = p.author_id where p.id = ?`,
		publicationId,
	)
	if err != nil {
		return models.Publication{}, err
	}
	defer line.Close()

	var publication models.Publication

	if line.Next() {
		// Scan segue a ordem da query
		if err = line.Scan(
			&publication.Id,
			&publication.Title,
			&publication.Content,
			&publication.AuthorId,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNickName,
		); err != nil {
			return models.Publication{}, err
		}
	}

	return publication, nil
}

func (p PublicationRepository) Fetch(userId uint64) ([]models.Publication, error) {
	lines, err := p.db.Query(`
    select distinct p.*, u.nick from publications p
    inner join users u on u.id = p.author_id
    inner join followers f on p.author_id = f.user_id
    where u.id = ? or f.follower_id = ?
    order by 1 desc`,
		userId,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if err = lines.Scan(
			&publication.Id,
			&publication.Title,
			&publication.Content,
			&publication.AuthorId,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNickName,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (p PublicationRepository) Update(publicationId uint64, publication models.Publication) error {
	statement, err := p.db.Prepare(
		"update publications set title = ?, content = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Content, publicationId); err != nil {
		return err
	}

	return nil
}

func (p PublicationRepository) Delete(publicationId uint64) error {
	statement, err := p.db.Prepare(
		"delete from publications where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationId); err != nil {
		return err
	}

	return nil
}

func (p PublicationRepository) FindPublicationByUser(userId uint64) ([]models.Publication, error) {
	lines, err := p.db.Query(`
    select p.*, u.nick from publications p 
    join users u on u.id = p.author_id
    where p.author_id = ?`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if err = lines.Scan(
			&publication.Id,
			&publication.Title,
			&publication.Content,
			&publication.AuthorId,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNickName,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (p PublicationRepository) LikePublication(publicationId uint64) error {
	statement, err := p.db.Prepare("update publications set likes = likes + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationId); err != nil {
		return err
	}

	return nil
}
