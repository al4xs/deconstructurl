package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Verifica se os argumentos da linha de comando são fornecidos corretamente
	if len(os.Args) != 1 {
		fmt.Println("+=======================================================+")
		fmt.Println("+==            	  Create By Al4xs                     ==+")
		fmt.Println("+== this script unchecks the url, printing all paths  ==+")
		fmt.Println("+=======================================================+")
		fmt.Println(" ")
		fmt.Println("Usage: file-crawler.txt | deconstructurl")
		return
	}

	// Scanner para ler linhas da entrada padrão (pipe)
	scanner := bufio.NewScanner(os.Stdin)

	// Itera sobre cada linha da entrada padrão
	for scanner.Scan() {
		url := scanner.Text()

		// Divide a URL em partes usando "/"
		parts := strings.Split(url, "/")

		// Itera sobre as partes e imprime os resultados
		for i := 3; i <= 30 && i <= len(parts); i++ {
			result := strings.Join(parts[:i], "/")
			fmt.Println(result)
		}
	}

	// Verifica erros durante a leitura da entrada padrão
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler a entrada padrão:", err)
	}
}
