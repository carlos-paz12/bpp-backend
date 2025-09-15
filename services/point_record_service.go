package services

import (
	"spe/models"
)

type PointRecordService struct{}

func (PointRecordService) Create(scholarshipID int64) (models.PointRecord, error) {
	// TODO
	return models.PointRecord{}, nil
}

func (PointRecordService) FindAllWhereScholarshipId(scholarshipID int64) ([]models.PointRecord, error) {
	// TODO
	return []models.PointRecord{}, nil
}

func (PointRecordService) FindLastWhereScholarshipId(scholarshipID int64) (*models.PointRecord, error) {
	// TODO
	return &models.PointRecord{}, nil
}
