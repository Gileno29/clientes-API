package utils

import "testing"

func TestValidarCPF(t *testing.T) {
	tests := []struct {
		cpf      string
		esperado bool
	}{
		{"11640457410", true},  // CPF válido
		{"11144477735", true},  // CPF válido
		{"11111111111", false}, // CPF inválido (todos os dígitos iguais)
		{"12345678900", false}, // CPF inválido (dígitos verificadores incorretos)
		{"", false},            // CPF vazio
		{"123", false},         // CPF com menos de 11 dígitos
	}

	for _, tt := range tests {
		resultado := ValidarCPF(tt.cpf)
		if resultado != tt.esperado {
			t.Errorf("ValidarCPF(%s) = %v; esperado %v", tt.cpf, resultado, tt.esperado)
		}
	}
}
