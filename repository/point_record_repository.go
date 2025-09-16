package repository

import (
	"errors"
	"spe/models"
	"time"
)

var (
	pointRecordsIDs int64 = 1
	records               = map[int64][]models.PointRecord{}
)

// PointRecordRepository gerencia registros de ponto em memória.
type PointRecordRepository struct{}

// Create adiciona um novo registro de ponto para um bolsista.
// O tipo do ponto (Entry/Exit) é alternado em relação ao último registro.
func (PointRecordRepository) Create(scholarshipID int64) models.PointRecord {
	newPointRecord := models.PointRecord{
		ID:            pointRecordsIDs,
		ScholarshipID: scholarshipID,
		RecordedAt:    time.Now(),
	}

	pointRecords := records[scholarshipID]

	if len(pointRecords) > 0 {
		last := pointRecords[len(pointRecords)-1]
		if last.Type == models.Entry {
			newPointRecord.Type = models.Exit
		} else {
			newPointRecord.Type = models.Entry
		}
	} else {
		newPointRecord.Type = models.Entry
	}

	pointRecordsIDs++
	records[scholarshipID] = append(records[scholarshipID], newPointRecord)
	return newPointRecord
}

// FindAllWhereScholarshipID retorna todos os registros de ponto de um bolsista.
func (PointRecordRepository) FindAllByScholarshipID(scholarshipID int64) []models.PointRecord {
	return records[scholarshipID]
}

// FindLastByScholarshipID retorna o último registro de ponto de um bolsista.
// Retorna erro se não houver registros.
func (PointRecordRepository) FindLastByScholarshipID(scholarshipID int64) (*models.PointRecord, error) {
	pointRecords := records[scholarshipID]
	if len(pointRecords) == 0 {
		return nil, errors.New("Nenhum registro encontrado.")
	}

	last := pointRecords[len(pointRecords)-1]
	return &last, nil
}
