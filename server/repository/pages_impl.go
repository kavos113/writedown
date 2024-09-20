package repository

import (
	"log"
)

func (r Repository) CreatePage(p Page) (Page, error) {
	res, err := r.db.Exec("INSERT INTO pages (parent_id, title, body, path, created_at, updated_at, creator_name) VALUES (?, ?, ?, ?, ?, ?, ?)", p.ParentID, p.Title, p.Body, p.Path, p.CreatedAt, p.UpdatedAt, p.CreatorName)
	if err != nil {
		log.Printf("Error creating page: %v", err)
		return Page{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return Page{}, err
	}

	ret := Page{
		ID:          int(id),
		ParentID:    p.ParentID,
		Title:       p.Title,
		Body:        p.Body,
		Path:        p.Path,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		CreatorName: p.CreatorName,
	}

	return ret, nil
}

func (r Repository) GetPageByID(id int) (Page, error) {
	p := Page{}
	err := r.db.Get(&p, "SELECT * FROM pages WHERE id = ?", id)
	if err != nil {
		log.Printf("Error getting page: %v", err)
		return Page{}, err
	}

	return p, nil
}

func (r Repository) UpdatePageByID(id int, p Page) (Page, error) {
	_, err := r.db.Exec("UPDATE pages SET parent_id = ?, title = ?, body = ?, path = ?, updated_at = ?, creator_name = ? WHERE id = ?", p.ParentID, p.Title, p.Body, p.Path, p.UpdatedAt, p.CreatorName, id)
	if err != nil {
		log.Printf("Error updating page: %v", err)
		return Page{}, err
	}

	return p, nil
}

func (r Repository) DeletePageByID(id int) error {
	_, err := r.db.Exec("DELETE FROM pages WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting page: %v", err)
		return err
	}

	return nil
}

func (r Repository) GetPagePath(id int) (string, error) {
	var path string
	err := r.db.Get(&path, "SELECT path FROM pages WHERE id = ?", id)
	if err != nil {
		log.Printf("Error getting page path: %v", err)
		return "", err
	}

	return path, nil
}

func (r Repository) GetChildrenPages(id int) ([]Page, error) {
	pages := []Page{}
	err := r.db.Select(&pages, "SELECT * FROM pages WHERE parent_id = ?", id)
	if err != nil {
		log.Printf("Error getting children pages: %v", err)
		return nil, err
	}

	return pages, nil
}
