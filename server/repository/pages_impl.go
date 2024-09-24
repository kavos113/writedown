package repository

import (
	"log"
	"strings"
)

func (Page) CreatePage(p Page) (Page, error) {
	res, err := db.Exec("INSERT INTO pages (parent_id, name, body, path, created_at, updated_at, creator_name) VALUES (?, ?, ?, ?, ?, ?, ?)", p.ParentID, p.Name, p.Body, p.Path, p.CreatedAt, p.UpdatedAt, p.CreatorName)
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
		Name:        p.Name,
		Body:        p.Body,
		Path:        p.Path,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		CreatorName: p.CreatorName,
	}

	return ret, nil
}

func (Page) GetPageByID(id int) (Page, error) {
	p := Page{}
	err := db.Get(&p, "SELECT * FROM pages WHERE id = ?", id)
	if err != nil {
		log.Printf("Error getting page: %v", err)
		return Page{}, err
	}

	return p, nil
}

func (Page) UpdatePageByID(id int, p Page) (Page, error) {
	path, err := p.GetPagePath(id)
	if err != nil {
		log.Printf("Error getting page path: %v", err)
		return Page{}, err
	}

	newPath := path[:strings.LastIndex(path, "/")+1] + p.Name

	_, err = db.Exec("UPDATE pages SET parent_id = ?, name = ?, body = ?, path = ?, updated_at = ?, creator_name = ? WHERE id = ?", p.ParentID, p.Name, p.Body, newPath, p.UpdatedAt, p.CreatorName, id)
	if err != nil {
		log.Printf("Error updating page: %v", err)
		return Page{}, err
	}

	return p, nil
}

func (Page) DeletePageByID(id int) error {
	_, err := db.Exec("DELETE FROM pages WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting page: %v", err)
		return err
	}

	return nil
}

func (Page) GetPagePath(id int) (string, error) {
	var path string
	err := db.Get(&path, "SELECT path FROM pages WHERE id = ?", id)
	if err != nil {
		log.Printf("Error getting page path: %v", err)
		return "", err
	}

	return path, nil
}

func (Page) GetChildrenPages(id int) ([]Page, error) {
	pages := []Page{}
	err := db.Select(&pages, "SELECT * FROM pages WHERE parent_id = ?", id)
	if err != nil {
		log.Printf("Error getting children pages: %v", err)
		return nil, err
	}

	return pages, nil
}
