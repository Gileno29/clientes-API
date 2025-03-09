package utils

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Gileno29/clientes-API/models"
	"gorm.io/gorm"
)

var StartTime time.Time

// pagina para referencia e consulta a respeito de calculo do digito verificador de cpf/cnpj
// URL: https://www.devmedia.com.br/validando-o-cpf-em-uma-aplicacao-java/22374

func ClearNumber(n string) string {
	newString := n
	newString = strings.ReplaceAll(newString, ".", "")
	newString = strings.ReplaceAll(newString, "-", "")
	newString = strings.ReplaceAll(newString, "/", "")
	return newString
}

func ValidaDocumento(cpfcnpj string) bool {
	cpfcnpj = ClearNumber(cpfcnpj)
	if len(cpfcnpj) < 11 || len(cpfcnpj) > 14 {
		return false
	}

	if len(cpfcnpj) == 11 {
		return ValidarCPF(cpfcnpj)
	}
	if len(cpfcnpj) == 14 {
		return ValidarCNPJ(cpfcnpj)

	}

	return false
}

func todosDigitosIguais(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

func calcularDigitoVerificadorCNPJ(numero string, pesos []int) int {
	soma := 0
	for i, r := range numero {
		digito, _ := strconv.Atoi(string(r))
		soma += digito * pesos[i]
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto

}

func calcularDigitoVerificador(numero string, pesoInicial int) int {
	soma := 0
	for i, r := range numero {
		digito, _ := strconv.Atoi(string(r))
		soma += digito * (pesoInicial - i)
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}

func ValidarCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	// Verifica se todos os dígitos são iguais
	if todosDigitosIguais(cpf) {
		return false
	}

	// Calcula o primeiro dígito verificador
	primeiroDigito := calcularDigitoVerificador(cpf[:9], 10)

	// Calcula o segundo dígito verificador
	segundoDigito := calcularDigitoVerificador(cpf[:10], 11)

	// Verifica se os dígitos calculados são iguais aos informados
	return cpf[9:11] == strconv.Itoa(primeiroDigito)+strconv.Itoa(segundoDigito)
}

func ValidarCNPJ(cnpj string) bool {
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")

	if len(cnpj) != 14 {
		return false
	}

	if todosDigitosIguais(cnpj) {
		return false
	}

	// Calcula o primeiro dígito verificador
	primeiroDigito := calcularDigitoVerificadorCNPJ(cnpj[:12], []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})

	// Calcula o segundo dígito verificador
	segundoDigito := calcularDigitoVerificadorCNPJ(cnpj[:13], []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})

	// Verifica se os dígitos calculados são iguais aos informados
	return cnpj[12:14] == strconv.Itoa(primeiroDigito)+strconv.Itoa(segundoDigito)
}

func VerificarTabelaClientes(db *gorm.DB) error {

	if !db.Migrator().HasTable(&models.Cliente{}) {
		log.Println("Tabela de clientes não encontrada. Criando tabela...")

		err := db.AutoMigrate(&models.Cliente{})
		if err != nil {
			log.Fatalf("Erro ao criar tabela de clientes: %v", err)
			return err
		}

		log.Println("Tabela de clientes criada com sucesso!")
	} else {
		log.Println("Tabela de clientes já existe.")
	}

	return nil
}
