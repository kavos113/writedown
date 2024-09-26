package repository

import (
	"log"
	"strings"
)

func (pr pagesRepository) CreatePage(p Page) (Page, error) {
	p.CreatorUserID = userID[p.CreatorName]
	res, err := db.Exec("INSERT INTO pages (parent_id, name, body, path, created_at, updated_at, creator_user_id) VALUES (?, ?, ?, ?, ?, ?, ?)", p.ParentID, p.Name, p.Body, p.Path, p.CreatedAt, p.UpdatedAt, p.CreatorUserID)
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

func (pr pagesRepository) GetPageByID(id int) (Page, error) {
	p := Page{}
	err := db.Get(&p, "SELECT * FROM pages WHERE id = ?", id)
	if err != nil {
		log.Printf("Error getting page: %v", err)
		return Page{}, err
	}

	p.CreatorName = userName[p.CreatorUserID]

	return p, nil
}

func (pr pagesRepository) UpdatePageByID(id int, p Page) (Page, error) {
	path, err := pr.GetPagePath(id)
	if err != nil {
		log.Printf("Error getting page path: %v", err)
		return Page{}, err
	}

	newPath := path[:strings.LastIndex(path, "/")+1] + p.Name

	p.CreatorUserID = userID[p.CreatorName]

	_, err = db.Exec("UPDATE pages SET parent_id = ?, name = ?, body = ?, path = ?, updated_at = ?, creator_user_id = ? WHERE id = ?", p.ParentID, p.Name, p.Body, newPath, p.UpdatedAt, p.CreatorUserID, id)
	if err != nil {
		log.Printf("Error updating page: %v", err)
		return Page{}, err
	}

	return p, nil
}

func (pr pagesRepository) DeletePageByID(id int) error {
	_, err := db.Exec("DELETE FROM pages WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting page: %v", err)
		return err
	}

	return nil
}

func (pr pagesRepository) GetPagePath(id int) (string, error) {
	var path string
	err := db.Get(&path, "SELECT path FROM pages WHERE id = ?", id)
	if err != nil {
		log.Printf("Error getting page path: %v", err)
		return "", err
	}

	return path, nil
}

func (pr pagesRepository) GetChildrenPages(id int) ([]PageAbstract, error) {
	var pages []PageAbstract
	err := db.Select(&pages, "SELECT id, path, name FROM pages WHERE parent_id = ?", id)
	if err != nil {
		log.Printf("Error getting children pages: %v", err)
		return nil, err
	}

	return pages, nil
}
