package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	locationPtr := flag.String("l", "", "the workspace directory")
	namePtr := flag.String("n", "", "the project name")
	languagePtr := flag.String("lang", "", "the project language")
	licensePtr := flag.String("license", "", "the project license")
	flag.Parse()
	fmt.Println("location:", *locationPtr)
	fmt.Println("name:", *namePtr)
	fmt.Println("lang:", *languagePtr)
	fmt.Println("license:", *licensePtr)

	createBaseDirectory(*namePtr, *locationPtr)
	retrieveGitignore(*languagePtr)
	retrieveLicense(*licensePtr)

	fmt.Println("Hello, World")
}

func createBaseDirectory(name string, location string) (string, error) {
	log.Println("Creating base directory", location+"/"+name)
	return "test", nil
}

func retrieveGitignore(language string) (string, error) {
	log.Println("Retrieving gitignore template for", language)
	return "gitignore test", nil
}

func retrieveLicense(license string) (string, error) {
	log.Println("Retrieving license template for", license)
	return "license test", nil
}
