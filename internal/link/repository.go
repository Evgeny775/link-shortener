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

func (repo *LinkRepository) Create(link *Link) error {
	result := repo.Database.DB.Create(link)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)

	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) IsHashExist(hash string) (bool,error){
	var count int64 

	err := repo.Database.DB.Model(&Link{}).
	Where("hash = ?", hash).
	Limit(1).
	Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil 
}
