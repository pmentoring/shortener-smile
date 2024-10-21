package repository

import (
	"database/sql"
	"shortener-smile/internal/shortener/domain/models"
)

type LinksRepository interface {
	GetNextId() (int, error)
	Save(*models.Link) error
	FindLinkByCode(linkCode string) (*models.Link, error)
	IncrementFollowsCounter(link *models.Link) error
}

type LinksRepositoryMysql struct {
	conn *sql.DB
}

func NewLinksRepository(conn *sql.DB) *LinksRepositoryMysql {
	return &LinksRepositoryMysql{conn: conn}
}

func (repo LinksRepositoryMysql) GetNextId() (int, error) {
	result := repo.conn.QueryRow("SELECT nextval(pg_get_serial_sequence('link', 'id'))")

	var nextId int

	err := result.Scan(&nextId)
	if err != nil {
		return 0, err
	}

	return nextId, nil
}

func (repo LinksRepositoryMysql) Save(link *models.Link) error {
	_, err := repo.conn.Exec(`
		INSERT INTO link (id, title, full_link, shorten_link_code, shorten_link) VALUES ($1, $2, $3, $4, $5);
		`,
		link.Id, link.Title, link.FullLink, link.ShortenLinkCode, link.ShortenLink,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo LinksRepositoryMysql) FindLinkByCode(linkCode string) (*models.Link, error) {
	row := repo.conn.QueryRow(`
		SELECT id,
			   title,
			   full_link,
			   shorten_link_code,
			   shorten_link_code
		FROM link
		WHERE shorten_link_code = $1
	`, linkCode)

	if row.Err() != nil {
		return nil, row.Err()
	}

	link := new(models.Link)
	err := row.Scan(&link.Id, &link.Title, &link.FullLink, &link.ShortenLinkCode, &link.ShortenLink)

	if err != nil {
		return nil, err
	}

	return link, nil
}

func (repo LinksRepositoryMysql) IncrementFollowsCounter(link *models.Link) error {
	_, err := repo.conn.Exec("UPDATE link SET follows = follows + 1 where id = $1", link.Id)
	if err != nil {
		return err
	}

	return nil
}
