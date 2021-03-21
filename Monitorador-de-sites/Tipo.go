package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	mostrarIntro()
	for {
		mostrarMenu()
		command := lerComando()
		switch command {
		case 1:
			monitorar()
		case 2:
			fmt.Println("Exibindo Logs...")
			mostarLogs()
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando Desconhecido!")
			main()
		}
	}
}

func mostrarIntro() {
	fmt.Printf("Insira seu nome: ")
	var nome string
	fmt.Scan(&nome)
	versao := 1.1
	fmt.Println("Olá sr(a).", nome)
	fmt.Println("Versão do programa:", versao)
}

func mostrarMenu() {
	fmt.Println("\n1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando escolhido:", comando)
	return comando
}

func monitorar() {
	fmt.Println("Iniciando Monitoramento...")
	sites := lerArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println(i+1, "site")
			testar(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println(" ")
	}
}

func testar(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, " foi carregado com sucesso!")
		registrarLog(site, true)
	} else {
		fmt.Println("Site: ", site, " não foi carregado com sucesso! Status Code: ", resp.StatusCode)
		registrarLog(site, false)
	}
}

func lerArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registrarLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func mostarLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	fmt.Println(string(arquivo))
}
