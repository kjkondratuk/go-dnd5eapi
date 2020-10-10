// +build ignore

package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	packageNameString := os.Args[1]
	apiName := os.Args[2]
	testInput := os.Args[3]
	lcApiName := strings.ToLower(string(apiName[0])) + string(apiName[1:])
	log.Println("Generating with template: " + packageNameString + " - " + apiName)

	params := struct {
		PackageName string
		Endpoint    string
		ApiName     string
		LCApiName   string
		TestInput   string
	}{
		packageNameString,
		strings.ReplaceAll(packageNameString, "_", "-"),
		apiName,
		lcApiName,
		testInput,
	}

	templateStr, err := ioutil.ReadFile("../gen/api_template.tpl")
	apiFile := template.Must(template.New("").Parse(string(templateStr)))

	f, err := os.Create(packageNameString + "_api.go")
	if err != nil {
		log.Fatal("Error writing file " + packageNameString + "_api.go")
		os.Exit(-1)
	}
	apiFile.Execute(f, params)

	testTemplateStr, err := ioutil.ReadFile("../gen/api_template_test.tpl")
	testFile := template.Must(template.New("").Parse(string(testTemplateStr)))

	tf, err := os.Create(packageNameString + "_api_test.go")
	if err != nil {
		log.Fatal("Error writing file " + packageNameString + "_api_test.go")
		os.Exit(-1)
	}
	testFile.Execute(tf, params)

	//log.Println("Generating... " + os.Args[1])
}
