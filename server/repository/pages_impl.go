package repository

import "log"

func (r Repository) CreatePage(p Page) (int, error) {
	res, err := r.db.Exec("INSERT INTO pages (parent_id, title, body, path, created_at, updated_at, creator_name) VALUES (?, ?, ?, ?, ?, ?, ?)", p.ParentID, p.Title, p.Body, p.Path, p.CreatedAt, p.UpdatedAt, p.CreatorName)
	if err != nil {
		log.Printf("Error creating page: %v", err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return 0, err
	}

	return int(id), nil
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

func (r Repository) UpdatePageByID(id int, p Page) error {
	_, err := r.db.Exec("UPDATE pages SET parent_id = ?, title = ?, body = ?, path = ?, updated_at = ?, creator_name = ? WHERE id = ?", p.ParentID, p.Title, p.Body, p.Path, p.UpdatedAt, p.CreatorName, id)
	if err != nil {
		log.Printf("Error updating page: %v", err)
		return err
	}

	return nil
}

func (r Repository) DeletePageByID(id int) error {
	_, err := r.db.Exec("DELETE FROM pages WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting page: %v", err)
		return err
	}

	return nil
}
