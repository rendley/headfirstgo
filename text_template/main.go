package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Part struct {
	Name  string
	Count int
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	text := "Here's my template!\n"
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, nil)
	check(err)
	fmt.Println("#####################")
	templateText := "Template start\nAction: {{.}}\nTemplate end\n"
	tmplt, err := template.New("test2").Parse(templateText)
	err = tmplt.Execute(os.Stdout, "ABC")
	check(err)
	err = tmplt.Execute(os.Stdout, 42)
	check(err)
	err = tmplt.Execute(os.Stdout, true)
	check(err)
	fmt.Println("#####################")

	executeTemplate("Dot is: {{.}}!\n", "ABC")
	executeTemplate("Dot is: {{.}}!\n", 123.5)
	fmt.Println("#####################")
	executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", true)
	executeTemplate("Start {{if .}}Dot is true!{{end}} finish\n", false)
	fmt.Println("####################")
	templateTextLoop := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"
	executeTemplate(templateTextLoop, []string{"do", "re", "mi"})
	templateTextLoop = "Prices:\n{{range .}}${{.}}\n{{end}}"
	executeTemplate(templateTextLoop, []float64{1.25, 0.99, 27, 33.2})
	fmt.Println("####################")
	executeTemplate(templateTextLoop, []float64{})
	executeTemplate(templateTextLoop, nil)
	fmt.Println("####################")

	templateTextStruct := "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(templateTextStruct, Part{Name: "Fuses", Count: 5})
	executeTemplate(templateTextStruct, Part{Name: "Cables", Count: 2})

	fmt.Println("####################")
	templateTextStruct = "Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	subscriber := Subscriber{"Aman Singh", 4.99, true}
	executeTemplate(templateTextStruct, subscriber)
	subscriber = Subscriber{"Joy Carr", 5.05, false}
	executeTemplate(templateTextStruct, subscriber)
}
