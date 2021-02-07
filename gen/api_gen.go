// +build ignore

package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

func main() {
	// templateRoot := os.Args[1]
	// Arch: %s OS: %s Line: %s Dollar: %s
	log.Printf("File: %s Package: %s",
		// os.Getenv("GOARCH"),
		// os.Getenv("GOOS"),
		os.Getenv("GOFILE"),
		// os.Getenv("GOLINE"),
		os.Getenv("GOPACKAGE"),
		// os.Getenv("DOLLAR"),
	)
	p, _ := os.Getwd()
	parts, _ := path.Split(p)
	// log.Println(parts)
	templateRoot := parts + "gen"
	log.Println(templateRoot)
	packageNameString := os.Getenv("GOPACKAGE")
	apiName := os.Args[1]
	apiClass := apiName + "Detail"
	testInput := os.Args[2]
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

	log.Println("Reading: " + templateRoot + "/api_template.tmpl")
	templateStr, err := ioutil.ReadFile(templateRoot + "/api_template.tmpl")
	apiFile := template.Must(template.New("").Parse(string(templateStr)))

	f, err := os.Create(packageNameString + "_api.go")
	if err != nil {
		log.Fatal("Error writing file " + packageNameString + "_api.go")
		os.Exit(-1)
	}
	apiFile.Execute(f, params)

	testTemplateStr, err := ioutil.ReadFile(templateRoot + "/api_template_test.tmpl")
	testFile := template.Must(template.New("").Parse(string(testTemplateStr)))

	tf, err := os.Create(packageNameString + "_api_test.go")
	defer tf.Close()
	if err != nil {
		log.Fatal("Error writing file " + packageNameString + "_api_test.go")
		os.Exit(-1)
	}
	testFile.Execute(tf, params)

	//log.Println("Generating... " + os.Args[1])
}
