package utils

import (
	"strings"
	"time"
)

var StartTime time.Time

func ValidarCPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	return len(cpf) == 11
}

func ValidarCNPJ(cnpj string) bool {
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")

	return len(cnpj) == 14

}
