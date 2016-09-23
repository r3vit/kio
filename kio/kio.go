// Package kio (Keep It Organized - inspired to Boilr) let manage simple template creation logic
package kio

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

type Kio struct {
	name      string
	input     string
	output    string
	reference map[string]string
	FuncMap   template.FuncMap
}

func NewKio(name, input, output string) *Kio {
	return &Kio{
		name:      name,
		input:     input,
		output:    output,
		reference: make(map[string]string),
	}

}

//Execute apply all the funcMap to the template
func (kio *Kio) Execute() {
	funcMap := template.FuncMap{
		"choose": kio.templateChoose,
		"isset":  kio.templateIsSet,
		"custom": kio.templateCustom,
	}
	kio.FuncMap = funcMap

	templateText := kio.load(kio.input)

	tmpl, err := template.New(kio.name).Funcs(kio.FuncMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	var newFile *os.File
	newFile, err = os.Create(kio.output)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	err = tmpl.Execute(newFile, kio.name)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

}

//load file info
func (kio *Kio) load(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

//List all the references
func (kio *Kio) List() {
	color.Set(color.FgCyan)
	fmt.Println("references List:")
	color.Unset()

	for i, v := range kio.reference {
		fmt.Println(" reference[", i, "]", "=", "\"", v, "\"")
	}
}
