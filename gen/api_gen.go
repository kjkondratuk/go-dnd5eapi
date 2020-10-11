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
	templateRoot := os.Args[1]
	packageNameString := os.Args[2]
	apiName := os.Args[3]
	apiClass := os.Args[4]
	testInput := os.Args[5]
	lcApiName := strings.ToLower(string(apiName[0])) + string(apiName[1:])
	log.Println("Generating with template: " + packageNameString + " - " + apiName)

	params := struct {
		PackageName string
		Endpoint    string
		ApiName     string
		ApiClass    string
		LCApiName   string
		TestInput   string
	}{
		packageNameString,
		strings.ReplaceAll(packageNameString, "_", "-"),
		apiName,
		apiClass,
		lcApiName,
		testInput,
	}

	templateStr, err := ioutil.ReadFile(templateRoot + "/api_template.tpl")
	apiFile := template.Must(template.New("").Parse(string(templateStr)))

	f, err := os.Create(packageNameString + "_api.go")
	if err != nil {
		log.Fatal("Error writing file " + packageNameString + "_api.go")
		os.Exit(-1)
	}
	apiFile.Execute(f, params)

	testTemplateStr, err := ioutil.ReadFile(templateRoot + "/api_template_test.tpl")
	testFile := template.Must(template.New("").Parse(string(testTemplateStr)))

	tf, err := os.Create(packageNameString + "_api_test.go")
	if err != nil {
		log.Fatal("Error writing file " + packageNameString + "_api_test.go")
		os.Exit(-1)
	}
	testFile.Execute(tf, params)

	//log.Println("Generating... " + os.Args[1])
}
