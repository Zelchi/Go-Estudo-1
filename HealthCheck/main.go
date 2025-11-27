package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func print(message any) {
	fmt.Println(message)
}

const monitoramentos = 3

func main() {
	for {
		showMenu()
		switch readKey() {
		case 1:
			startMonitoring()
		case 2:
			readLogs()
		case 3:
			limpaLogs()
		case 0:
			exitProgram()
		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}
}

func showMenu() {
	fmt.Println("|----------------------|")
	fmt.Println("| 1 - Monitoramento    |")
	fmt.Println("| 2 - Exibir logs      |")
	fmt.Println("| 3 - Limpar logs      |")
	fmt.Println("| 0 - Sair do programa |")
	fmt.Println("|----------------------|")
}

func readKey() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func startMonitoring() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Iniciando monitoramento...")
	fmt.Println("--------------------------")

	urls := leArquivoDeTexto()

	for _, url := range urls {

		for range monitoramentos {
			now := time.Now()
			res, err := http.Get(url)
			since := time.Since(now)
			ms := int(since / time.Millisecond)
			msStr := fmt.Sprintf("%03d", ms)

			if err != nil {
				fmt.Println("Erro ao acessar:", url, "->", err)
				registraLog(url, msStr, false)
			} else {
				if res.StatusCode == 200 {
					fmt.Println("["+msStr+"ms]", "Site:", url, "carregado com sucesso!")
					registraLog(url, msStr, true)
				} else {
					fmt.Println("Site:", url, "retornou status:", res.StatusCode)
					registraLog(url, msStr, false)
				}
				res.Body.Close()
			}

			time.Sleep(3 * time.Second)
		}
	}

	fmt.Println("Monitoramento concluído.")
}

func exitProgram() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Saindo do programa...")
	time.Sleep(2 * time.Second)
	os.Exit(0)
}

func leArquivoDeTexto() []string {

	var listaDeUrls []string

	arq, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao ler o arquivo:", err)
	}

	scanner := bufio.NewScanner(arq)
	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text())
		if linha != "" {
			print(linha)
			listaDeUrls = append(listaDeUrls, linha)
		}
	}

	arq.Close()

	return listaDeUrls
}

func registraLog(site string, ms string, status bool) {
	arq, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo de log:", err)
		return
	}

	strStatus := strconv.FormatBool(status)
	timeNow := time.Now().Format("02/01/2006 15:04:05")

	arq.WriteString(fmt.Sprintf("[%v]|[%vms]|[%v]|[%v]\n", timeNow, ms, strStatus, site))
	defer arq.Close()
}

func readLogs() []string {
	fmt.Print("\033[H\033[2J")

	var listaDeLogs []string

	arq, err := os.Open("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao ler o arquivo:", err)
	}

	scanner := bufio.NewScanner(arq)
	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text())
		if linha != "" {
			print(linha)
			listaDeLogs = append(listaDeLogs, linha)
		}
	}

	defer arq.Close()
	return listaDeLogs
}

func limpaLogs() {
	fmt.Print("\033[H\033[2J")

	err := os.Remove("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro ao limpar os logs:", err)
		return
	}
	fmt.Println("Logs limpos com sucesso.")
}
