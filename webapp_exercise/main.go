package main

import (
	"html/template"
	"log"
	"os"
)

type Invoice struct {
	Name    string
	Paid    bool
	Charges []float64
	Total   float64
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	html, err := template.ParseFiles("/home/rendle/GolandProjects/github.com/rendley/headfirstgo/webapp_exercise/bill.html")
	check(err)
	bill := Invoice{
		Name:    "Mary Gibbs",
		Paid:    true,
		Charges: []float64{23.19, 1.13, 42.79},
		Total:   67.11,
	}
	err = html.Execute(os.Stdout, bill)
	check(err)
}
