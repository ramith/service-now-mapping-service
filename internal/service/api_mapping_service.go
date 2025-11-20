package service

import (
	"errors"

	"github.com/ramith/service-now-mapping-service/internal/models"
	"github.com/ramith/service-now-mapping-service/internal/repository"
)

type APIMappingService interface {
	CreateMapping(req *models.CreateAPIMapping) (*models.APIMapping, error)
	GetMappingBySysID(sysID string) (*models.APIMapping, error)
	GetMappingByAPIID(apiID string) (*models.APIMapping, error)
	GetMappingByNameAndVersion(apiName, apiVersion string) (*models.APIMapping, error)
	GetAllMappings() ([]models.APIMapping, error)
	UpdateMapping(id uint, req *models.UpdateAPIMapping) (*models.APIMapping, error)
	DeleteMapping(id uint) error
}

type apiMappingService struct {
	repo repository.APIMappingRepository
}

func NewAPIMappingService(repo repository.APIMappingRepository) APIMappingService {
	return &apiMappingService{repo: repo}
}

func (s *apiMappingService) CreateMapping(req *models.CreateAPIMapping) (*models.APIMapping, error) {
	if req.SysID == "" || req.APIID == "" || req.APIName == "" || req.APIVersion == "" {
		return nil, errors.New("all fields are required")
	}

	mapping := &models.APIMapping{
		SysID:      req.SysID,
		APIID:      req.APIID,
		APIName:    req.APIName,
		APIVersion: req.APIVersion,
	}

	if err := s.repo.Create(mapping); err != nil {
		return nil, err
	}
	return mapping, nil
}

func (s *apiMappingService) GetMappingBySysID(sysID string) (*models.APIMapping, error) {
	if sysID == "" {
		return nil, errors.New("sys_id cannot be empty")
	}
	return s.repo.GetBySysID(sysID)
}

func (s *apiMappingService) GetMappingByAPIID(apiID string) (*models.APIMapping, error) {
	if apiID == "" {
		return nil, errors.New("api_id cannot be empty")
	}
	return s.repo.GetByAPIID(apiID)
}

func (s *apiMappingService) GetMappingByNameAndVersion(apiName, apiVersion string) (*models.APIMapping, error) {
	if apiName == "" || apiVersion == "" {
		return nil, errors.New("both api_name and api_version are required")
	}
	return s.repo.GetByNameAndVersion(apiName, apiVersion)
}

func (s *apiMappingService) GetAllMappings() ([]models.APIMapping, error) {
	return s.repo.GetAll()
}

func (s *apiMappingService) UpdateMapping(id uint, req *models.UpdateAPIMapping) (*models.APIMapping, error) {
	// Get existing mapping
	mapping, err := s.repo.GetBySysID("")
	if err != nil {
		// Try to find by ID - for simplicity, get all and filter
		mappings, err := s.repo.GetAll()
		if err != nil {
			return nil, err
		}

		var found *models.APIMapping
		for i := range mappings {
			if mappings[i].ID == id {
				found = &mappings[i]
				break
			}
		}

		if found == nil {
			return nil, errors.New("mapping not found")
		}
		mapping = found
	}

	// Update fields
	if req.APIID != "" {
		mapping.APIID = req.APIID
	}
	if req.APIName != "" {
		mapping.APIName = req.APIName
	}
	if req.APIVersion != "" {
		mapping.APIVersion = req.APIVersion
	}

	if err := s.repo.Update(mapping); err != nil {
		return nil, err
	}
	return mapping, nil
}

func (s *apiMappingService) DeleteMapping(id uint) error {
	return s.repo.Delete(id)
}
