package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const DELAY_RANDOM = 5
const URLS_FILE = "urls.txt"
const URL_RANDOM_STATUS = "https://random-status-code.herokuapp.com"
const LOG_FILE = "logs.txt"

func main() {
	screenClear()
	// for vazio é o mesmo que um while true(loop infinito), Go não tem while
	for {
		executeOption(showMenu())
		screenClear()
	}
}

func showMenu() int {
	fmt.Println("------------")
	fmt.Println("*** Menu ***")
	fmt.Println("------------")
	fmt.Println("1 - Verificar internet")
	fmt.Println("2 - Verificar alguma url(exemplo http://www.eprogramar.com.br)")
	fmt.Println("3 - Testar Http Status random")
	fmt.Println("4 - Listar urls de teste(exemplo http://www.eprogramar.com.br)")
	fmt.Println("5 - Adicionar uma nova URL de teste")
	fmt.Println("------------")
	fmt.Println("0 - Sair do Programa")

	var option int
	fmt.Scan(&option)
	return option
}

func executeOption(option int) {
	switch option {
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(-1)
	case 1:
		verifyInternet()
	case 2:
		verifySomeUrl()
	case 3:
		verifyStatusRandom()
	case 4:
		showUrls()
	case 5:
		addNewUrl()
	default:
		fmt.Println("Opção inválida!")
	}
}

// Testando uma lista de urls que estão em um arquivo texto
func verifyInternet() {
	fmt.Println("Start test...")
	urls := loadFile()
	for index, url := range urls {
		message, _ := verifyUrl(url)
		fmt.Println("[", index, "]", "Internet OK, sent request to ", message)
	}
	fmt.Println("Test finished!")
	zeroToBack()
}

// Testando uma url passada por parametro pelo console, através do Scan
func verifySomeUrl() {
	var url string
	fmt.Println("Digite a url que deve ser testada: ")
	fmt.Scan(&url)
	message, _ := verifyUrl(url)
	fmt.Println(message)
	zeroToBack()
}

// Testando a url de status random com um loop infinito até que o status seja >= 500
func verifyStatusRandom() {
	for {
		message, status := verifyUrl(URL_RANDOM_STATUS)
		fmt.Println(message)

		if status >= 500 {
			break
		}

		time.Sleep(DELAY_RANDOM * time.Second) // aguarda 5 segundos
	}
	zeroToBack()
}

// Executando uma requisição de GET para um site externo
func verifyUrl(url string) (msg string, status int) {
	resp, err := http.Get(url)
	if err != nil {
		return err.Error(), resp.StatusCode
	}

	var message string
	if resp.StatusCode == 200 {
		message = "Status OK, sent request to " + url
	} else {
		message = "Status Error " + strconv.Itoa(resp.StatusCode) + " sent request to " + url
	}

	writeLog(message)
	return message, resp.StatusCode
}

// Limpando o console
func screenClear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// Listando todas as linhas de um arquivo
func showUrls() {
	file, err := ioutil.ReadFile(URLS_FILE)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	zeroToBack()
}

// Extraindo todas as linhas de um arquivo para um slice
func loadFile() []string {
	var urls []string
	loadFile, err := os.Open(URLS_FILE)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(loadFile)
	for {
		line, err := reader.ReadString('\n')
		urls = append(urls, strings.TrimSpace(line))

		if err == io.EOF {
			break
		}

	}

	loadFile.Close()
	return urls
}

func addNewUrl() {
	var url string
	fmt.Println("Digite a url que deseja adicionar: ")
	fmt.Scan(&url)

	msg, status := verifyUrl(url)
	if status != 200 {
		fmt.Println(msg)
		zeroToBack()
	}

	urlsFile, err := os.OpenFile(URLS_FILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	urlsFile.WriteString(url + "\n")
	urlsFile.Close()

}

func zeroToBack() {
	var o int
	fmt.Println("Digite 0 para voltar")
	fmt.Scan(&o)
}

func writeLog(msg string) {
	logFile, err := os.OpenFile(LOG_FILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// Doc, Go format time: https://go.dev/src/time/format.go
	logFile.WriteString(time.Now().Format("02/02/2006 15:04:05") + " - " + msg + "\n")

	logFile.Close()
}
