package repository

type Tag struct {
	ID     int    `db:"id"`
	PageID int    `db:"page_id"`
	Name   string `db:"name"`
}

type TagsRepository interface {
	CreateTag(t Tag) (Tag, error)
	GetTagsByPageID(pageID int) ([]Tag, error)
	DeleteTagByTagID(tagID int) error
	GetPagesByTagID(tagID int) ([]PageAbstract, error)
}

type tagsRepository struct{}

func NewTagsRepository() TagsRepository {
	return tagsRepository{}
}
