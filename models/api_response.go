package models

// ApiResponse representa a estrutura padrão de resposta de todos os endpoints da API.
type ApiResponse struct {
	Message  string `json:"message"`        // Mensagem de sucesso ou informação geral sobre a requisição.
	Error    string `json:"error"`          // Mensagem de erro detalhando o motivo da falha, caso exista.
	HttpCode int    `json:"http_code"`      // Código HTTP da resposta.
	Data     any    `json:"data,omitempty"` // Dados retornados pelo endpoint. Pode ser qualquer tipo e é opcional.
}
