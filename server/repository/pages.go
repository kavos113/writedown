package repository

import "time"

type Page struct {
	ID            int       `db:"id"`
	ParentID      int       `db:"parent_id"`
	Name          string    `db:"name"`
	Body          string    `db:"body"`
	Path          string    `db:"path"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	CreatorName   string
	CreatorUserID int `db:"creator_user_id"`
}

type PageAbstract struct {
	ID   int    `db:"id"`
	Path string `db:"path"`
	Name string `db:"name"`
}

type PagesRepository interface {
	CreatePage(p Page) (Page, error)
	GetPageByID(id int) (Page, error)
	UpdatePageByID(id int, p Page) (Page, error)
	DeletePageByID(id int) error

	GetPagePath(id int) (string, error)
	GetChildrenPages(id int) ([]PageAbstract, error)
}

type pagesRepository struct{}

func NewPagesRepository() PagesRepository {
	return pagesRepository{}
}
