// Package pkg fornece utilitários para manipulação de strings.
package pkg

import "strings"

// IsEmpty verifica se a string fornecida está vazia.
func IsEmpty(value string) bool {
	return strings.TrimSpace(value) == ""
}

