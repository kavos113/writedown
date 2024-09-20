package repository

import "time"

type Page struct {
	ID          int       `db:"id"`
	ParentID    int       `db:"parent_id"`
	Title       string    `db:"title"`
	Body        string    `db:"body"`
	Path        string    `db:"path"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	CreatorName string    `db:"creator_name"`
}

type PagesRepository interface {
	CreatePage(p Page) (Page, error)
	GetPageByID(id int) (Page, error)
	UpdatePageByID(id int, p Page) (Page, error)
	DeletePageByID(id int) error

	GetPagePath(id int) (string, error)
	GetChildrenPages(id int) ([]Page, error)
}
