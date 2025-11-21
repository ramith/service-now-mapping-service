package repository

import (
	"errors"

	"github.com/ramith/service-now-mapping-service/internal/models"
	"gorm.io/gorm"
)

type APIMappingRepository interface {
	Create(mapping *models.APIMapping) error
	GetBySysID(sysID string) (*models.APIMapping, error)
	GetByAPIID(apiID string) (*models.APIMapping, error)
	GetByNameAndVersion(apiName, apiVersion string) (*models.APIMapping, error)
	GetAll() ([]models.APIMapping, error)
	Update(mapping *models.APIMapping) error
	UpdateByAPIID(apiID string, mapping *models.APIMapping) error
	Delete(id uint) error
}

type apiMappingRepository struct {
	db *gorm.DB
}

func NewAPIMappingRepository(db *gorm.DB) APIMappingRepository {
	return &apiMappingRepository{db: db}
}

func (r *apiMappingRepository) Create(mapping *models.APIMapping) error {
	return r.db.Create(mapping).Error
}

func (r *apiMappingRepository) GetBySysID(sysID string) (*models.APIMapping, error) {
	var mapping models.APIMapping
	err := r.db.Where("sys_id = ?", sysID).First(&mapping).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("mapping not found")
		}
		return nil, err
	}
	return &mapping, nil
}

func (r *apiMappingRepository) GetByAPIID(apiID string) (*models.APIMapping, error) {
	var mapping models.APIMapping
	err := r.db.Where("api_id = ?", apiID).First(&mapping).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("mapping not found")
		}
		return nil, err
	}
	return &mapping, nil
}

func (r *apiMappingRepository) GetByNameAndVersion(apiName, apiVersion string) (*models.APIMapping, error) {
	var mapping models.APIMapping
	err := r.db.Where("api_name = ? AND api_version = ?", apiName, apiVersion).First(&mapping).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("mapping not found")
		}
		return nil, err
	}
	return &mapping, nil
}

func (r *apiMappingRepository) GetAll() ([]models.APIMapping, error) {
	var mappings []models.APIMapping
	err := r.db.Find(&mappings).Error
	return mappings, err
}

func (r *apiMappingRepository) Update(mapping *models.APIMapping) error {
	return r.db.Save(mapping).Error
}

func (r *apiMappingRepository) UpdateByAPIID(apiID string, mapping *models.APIMapping) error {
	return r.db.Model(&models.APIMapping{}).Where("api_id = ?", apiID).Updates(mapping).Error
}

func (r *apiMappingRepository) Delete(id uint) error {
	return r.db.Delete(&models.APIMapping{}, id).Error
}
