package services

import (
	"spe/models"
	"sync"
	"time"
)

var (
	registros   = map[int][]models.Ponto{}
	registrosMu sync.RWMutex
	idCounter   = 1
)

type PontoService struct{}

// Create registra um novo ponto de um bolsista.
func (PontoService) Create(bolsistaID int) models.Ponto {
	registrosMu.Lock()
	defer registrosMu.Unlock()

	novoPonto := models.Ponto{
		Id:         idCounter,
		BolsistaID: bolsistaID,
		Timestamp:  time.Now(),
	}

	if pontos := registros[bolsistaID]; len(pontos) > 0 {
		ultimo := pontos[len(pontos)-1]
		novoPonto.Tipo = !ultimo.Tipo
	} else {
		novoPonto.Tipo = true
	}

	idCounter++
	registros[bolsistaID] = append(registros[bolsistaID], novoPonto)
	return novoPonto
}

// RetrieveAll retorna todos os registros de ponto de um bolsista.
func (PontoService) RetrieveAll(bolsistaID int) []models.Ponto {
	registrosMu.RLock()
	defer registrosMu.RUnlock()

	return registros[bolsistaID]
}

// RetrieveLast retorna o último registro de ponto de um bolsista.
func (PontoService) RetrieveLast(bolsistaID int) *models.Ponto {
	registrosMu.RLock()
	defer registrosMu.RUnlock()

	pontos := registros[bolsistaID]
	if len(pontos) == 0 {
		return nil
	}

	return &pontos[len(pontos)-1]
}
