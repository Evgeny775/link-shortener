package link

import (
	"link-shortener/pkg/db"

	"gorm.io/gorm/clause"
)

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

func (repo *LinkRepository) Update(link *Link) error {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *LinkRepository) Delete(id int) error {
	result := repo.Database.DB.Delete(&Link{},id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *LinkRepository) HashExistExeptID(hash string, id int) (bool, error) {
	var count int64
	err := repo.Database.DB.Model(&Link{}).
		Where("hash = ? and id <> ?", hash, id).
		Count(&count).Error

	return count > 0, err
}
