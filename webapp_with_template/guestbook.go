package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signature      []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// getStrings возвращает сегмент строк, прочитанный из fileName,
// по одной строке на каждую строку файла.
func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file) // создается сканер для содержимого файла
	for scanner.Scan() {              // для каждой строки
		lines = append(lines, scanner.Text()) // текс присоединяется к слайсу
	}
	check(scanner.Err())
	return lines
}

// viewHandler читает записи гостевой книги и выводит их вместе со счетчиком записей.
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("/home/rendle/GolandProjects/github.com/rendley/headfirstgo/webapp_with_template/signatures.txt")
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("/home/rendle/GolandProjects/github.com/rendley/headfirstgo/webapp_with_template/view.html")
	check(err)
	guestbook := Guestbook{SignatureCount: len(signatures), Signature: signatures}

	err = html.Execute(writer, guestbook)
	check(err)
	//placeholder := []byte("signature list goes here")
	//_, err := writer.Write(placeholder)
	//check(err)
}

// newHandler отображает форму для ввода записи.
func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("/home/rendle/GolandProjects/github.com/rendley/headfirstgo/webapp_with_template/new.html")
	check(err)
	err = html.Execute(writer, nil)
}

// createHandler получает запрос POST с добавляемой записью
// и присоединяет ее к файлу signatures.
func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature") // получаем значение формы "signature в new.html"
	//fmt.Println(signature)
	//_, err := writer.Write([]byte(signature)) //Записываем значение поля в ответ и отдаем по урл
	//check(err)
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE // параметры открытия файла
	file, err := os.OpenFile("/home/rendle/GolandProjects/github.com/rendley/headfirstgo/webapp_with_template/signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature) // записываем текст в новой строке
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound) //перенаправляем после отправки формы
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
