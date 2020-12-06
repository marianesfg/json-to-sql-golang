package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	destionationpath = "./script.sql"
	sourcepath       = "./teste.json"
)

// struct Card
type Card struct {
	IDCLIENT    int
	CPF         string
	DESCRIPTION string
}

func main() {

	//Ler arquivo JSON
	f, err := ioutil.ReadFile(sourcepath)
	if err != nil {
		log.Fatal(err)
	}

	//desserializar JSON
	var datalist []Card
	err = json.Unmarshal(f, &datalist)
	if err != nil {
		log.Fatal(err)
	}

	sql := "INSERT INTO TableJSON VALUES ("

	arquivo, err := os.Create(destionationpath)
	if err != nil {
		log.Fatal(err)
	}

	// defer: existe na maioria das outras linguagens
	// Um dos principais usos de uma instrução defer é o da limpeza de recursos, como:
	// .arquivos abertos
	// .conexões de banco de dados
	// Garante que o arquivo sera fechado apos o uso:
	defer arquivo.Close() // add em 1 pilha (LIFO)

	// Vai escrever cada linha do arquivo
	escritor := bufio.NewWriter(arquivo)

	for i := 0; i < len(datalist); i++ {
		if datalist[i].IDCLIENT == 0 && datalist[i].CPF == "" && datalist[i].DESCRIPTION == "" {
			result := fmt.Sprintf("%s%d%s", "Linha ", i+1, " desconsiderada por estar em branco.")
			log.Println(result)
		} else if len(datalist[i].DESCRIPTION) > 4000 {
			result := fmt.Sprintf("%s%d%s", "Linha ", i, " desconsiderada pela descrição ter mais que 4000 caracteres.")
			log.Println(result)
		} else {
			result := fmt.Sprintf("%s%d%s%s%s%s%s%s%s%s%s", sql, datalist[i].IDCLIENT, ",", "'",
				strings.ReplaceAll(strings.ReplaceAll(datalist[i].CPF, ".", ""), "-", ""),
				"'", ",", "'",
				strings.ReplaceAll(strings.ReplaceAll(datalist[i].DESCRIPTION, "\n", `\n`), "\r", `\r`),
				"'", ")")
			//fmt.Println(result)
			fmt.Fprintln(escritor, result)
		}
	}

	err = escritor.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
