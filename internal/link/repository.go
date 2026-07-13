package link

import "link-shortener/pkg/db"

type LinkRepository struct {
	Database *db.DB
}

func NewLinkRepository(database *db.DB) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link)  error {
	result := repo.Database.DB.Create(link)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
