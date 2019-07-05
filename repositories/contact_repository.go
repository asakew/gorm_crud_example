package repositories

import (
	"github.com/herusdianto/gorm_crud_example/models"
	"github.com/jinzhu/gorm"
)

type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) Save(contact *models.Contact) RepositoryResult {
	err := r.db.Save(contact).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: contact}
}

func (r *ContactRepository) FindAll() RepositoryResult {
	var contacts models.Contacts

	err := r.db.Find(&contacts).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &contacts}
}

func (r *ContactRepository) FindOneById(id string) RepositoryResult {
	var contact models.Contact

	err := r.db.Where(&models.Contact{ID: id}).Take(&contact).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &contact}
}

func (r *ContactRepository) DeleteOneById(id string) RepositoryResult {
	err := r.db.Delete(&models.Contact{ID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *ContactRepository) DeleteByIds(ids *[]string) RepositoryResult {
	err := r.db.Where("ID IN (?)", *ids).Delete(&models.Contacts{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
