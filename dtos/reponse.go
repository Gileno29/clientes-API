package dtos

type ClienteResponse struct {
	Documento   string `json:"documento"`
	RazaoSocial string `json:"razaosocial"`
	Blocklist   bool   `json:"blocklist"`
}

type ListarClientesResponse struct {
	Page     int               `json:"page"`
	Limit    int               `json:"limit"`
	Total    int64             `json:"total"`
	Clientes []ClienteResponse `json:"clientes"`
}

type ResponseStatus struct {
	Uptime   float64 `json:"uptime"`
	Requests int     `json:"requests"`
}

type ResponseErro struct {
	Mensagem string `json:"mensagem"`
}

type AtualizaClienteRequest struct {
	RazaoSocial *string `json:"razaosocial"`
	Blocklist   *bool   `json:"blocklist"`
}
