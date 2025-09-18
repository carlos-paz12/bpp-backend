package repository

import (
	"spe/models"
)

var (
	pointRecordsIDs int64 = 1
	pointRecords          = map[int64][]models.PointRecord{}
)

type PointRecordRepository struct{}

func (PointRecordRepository) Save(pr models.PointRecord) models.PointRecord {
	pr.ID = pointRecordsIDs
	pointRecordsIDs++
	pointRecords[pr.ScholarshipID] = append(pointRecords[pr.ScholarshipID], pr)
	return pr
}

func (PointRecordRepository) FindAllByScholarshipID(sid int64) []models.PointRecord {
	return pointRecords[sid]
}

func (PointRecordRepository) FindLastByScholarshipID(sid int64) (*models.PointRecord, bool) {
	records := pointRecords[sid]
	if len(records) == 0 {
		return &models.PointRecord{}, false
	}
	return &records[len(records)-1], true
}
