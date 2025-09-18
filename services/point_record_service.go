package services

import (
	"spe/models"
	"spe/repository"

	"errors"
	"time"
)

type PointRecordService struct{}

func (PointRecordService) GetMyPoints(uid int64) ([]models.PointRecord, error) {
	s, exists := repository.ScholarshipRepository{}.FindByUserID(uid)
	if !exists {
		return nil, errors.New("scholarship not found.")
	}

	prs := repository.PointRecordRepository{}.FindAllByScholarshipID(s.ID)
	return prs, nil
}

func (PointRecordService) RegisterMyPoint(uid int64) (*models.PointRecord, error) {
	s, exists := repository.ScholarshipRepository{}.FindByUserID(uid)
	if !exists {
		return nil, errors.New("scholarship not found.")
	}

	pr := models.PointRecord{
		ScholarshipID: s.ID,
		RecordedAt:    time.Now(),
	}

	last, exists := repository.PointRecordRepository{}.FindLastByScholarshipID(s.ID)
	if exists {
		pr.Type = nextType(last)
	} else {
		pr.Type = models.Entry
	}

	pr = repository.PointRecordRepository{}.Save(pr)
	return &pr, nil
}

func (PointRecordService) GetMyLastPoint(uid int64) (*models.PointRecord, error) {
	s, exists := repository.ScholarshipRepository{}.FindByUserID(uid)
	if !exists {
		return nil, errors.New("scholarship not found")
	}

	last, exists := repository.PointRecordRepository{}.FindLastByScholarshipID(s.ID)
	if !exists {
		return nil, nil
	}

	return last, nil
}

func (PointRecordService) GetPointsByScholarshipID(sid int64) ([]models.PointRecord, error) {
	// Todo
	return []models.PointRecord{}, nil
}

func (PointRecordService) GetLastByScholarshipID(sid int64) (*models.PointRecord, error) {
	// Todo
	return &models.PointRecord{}, nil
}

func nextType(last *models.PointRecord) models.PointRecordType {
	if last == nil || last.Type == models.Exit {
		return models.Entry
	}
	return models.Exit
}
