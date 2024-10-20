package repository

import "database/sql"

type Link struct {
	Id              int
	Title           string
	FullLink        string
	ShortenLinkCode string
	Follows         int
	ShortenLink     string
}

type LinksRepository interface {
	GetNextId() (int, error)
	Save(*Link) error
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

func (repo LinksRepositoryMysql) Save(link *Link) error {
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
