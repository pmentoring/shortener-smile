package models

type Link struct {
	Id              int
	Title           string
	FullLink        string
	ShortenLinkCode string
	Follows         int
	ShortenLink     string
}
