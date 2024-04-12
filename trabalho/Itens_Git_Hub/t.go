package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Abrir o arquivo SRT
	arquivoSRT := "legendas1.srt"
	arquivo, err := os.Open(arquivoSRT)
	if err != nil {
		panic(err)
	}
	defer arquivo.Close()

	// Contagem da palavra
	palavraAlvo := "i" // Palavra que você quer contar
	contagem := 0

	// Ler o arquivo linha por linha
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := scanner.Text()

		// Contar a ocorrência da palavra na linha (case insensitive)
		contagem += strings.Count(strings.ToLower(linha), palavraAlvo)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("A palavra \"%s\" ocorre %d vezes no arquivo SRT.\n", palavraAlvo, contagem)
}
