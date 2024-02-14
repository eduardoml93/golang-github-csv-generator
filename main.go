package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	// URL do arquivo CSV
	url := "https://github.com/datasciencedojo/datasets/blob/master/titanic.csv?raw=True"

	// Baixando o arquivo CSV
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao baixar o arquivo:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Lendo o corpo da resposta como string
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		os.Exit(1)
	}

	// Decodificando a string de bytes para string
	bodyString := string(body)

	// Dividindo o conteúdo em linhas
	linhas := strings.Split(bodyString, "\n")

	// Imprimindo as linhas
	for _, linha := range linhas {
		fmt.Println(linha)
	}

	// Salvando o arquivo CSV
	err = salvarCSV("titanic_from_github.csv", linhas)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo CSV:", err)
		os.Exit(1)
	}

	fmt.Println("Arquivo CSV salvo com sucesso!")
}

func salvarCSV(nomeArquivo string, conteudo []string) error {
	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	for _, linha := range conteudo {
		_, err := arquivo.WriteString(linha + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
