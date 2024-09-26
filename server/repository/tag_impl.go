package repository

import "log"

func (tr tagsRepository) CreateTag(t Tag) (Tag, error) {
	res, err := db.Exec("INSERT INTO tags (page_id, name) VALUES (?, ?)", t.PageID, t.Name)
	if err != nil {
		log.Printf("Error creating tag: %v", err)
		return Tag{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return Tag{}, err
	}

	ret := Tag{
		ID:     int(id),
		PageID: t.PageID,
		Name:   t.Name,
	}

	return ret, nil
}

func (tr tagsRepository) GetTagsByPageID(pageID int) ([]Tag, error) {
	var tags []Tag
	err := db.Select(&tags, "SELECT * FROM tags WHERE page_id = ?", pageID)
	if err != nil {
		log.Printf("Error getting tags by page ID: %v", err)
		return nil, err
	}

	return tags, nil
}

func (tr tagsRepository) DeleteTagByTagID(tagID int) error {
	_, err := db.Exec("DELETE FROM tags WHERE id = ?", tagID)
	if err != nil {
		log.Printf("Error deleting tag by tag ID: %v", err)
		return err
	}

	return nil
}

func (tr tagsRepository) GetPagesByTagID(tagID int) ([]PageAbstract, error) {
	var pages []PageAbstract
	err := db.Select(&pages, "SELECT id, path, name FROM pages WHERE id IN (SELECT page_id FROM tags WHERE id = ?)", tagID)
	if err != nil {
		log.Printf("Error getting pages by tag ID: %v", err)
		return nil, err
	}

	return pages, nil
}
