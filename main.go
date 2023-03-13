package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

const (
	PythonCommentToken =  "#"
	JavascriptCommentToken = "//"
	GolangCommentToken = "//"
)

func openFile(f string){
	fileExtension := extractFileExtension(f)

	// Validate file extension
	if !validateExtension(fileExtension){
		fmt.Printf("File with extension %q not supported currently.\n", fileExtension)
		os.Exit(1)
	}

	file, err := os.OpenFile(f, os.O_RDWR, 0664) // Open for reading
	handleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	tempFile, err := os.CreateTemp("/tmp", "tmp-*" + fileExtension)
	handleError(err)
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	tempFileStat, _ := tempFile.Stat()

	for scanner.Scan(){
		text := replaceComment(scanner.Text(), "", fileExtension)
		tempFile.WriteString(text)
	}

	os.Rename("/tmp/" + tempFileStat.Name(), f)
}

func main(){
	var filePath string

	fmt.Println("Please provide the file path: ")
	fmt.Scanln(&filePath) // Full file path

	openFile(filePath)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func replaceComment(source string, destination string, ext string) string {
	commentToken := getExtensionToken(ext)

	newLine := source
	commentPattern:= fmt.Sprintf("(%s.*)", commentToken)

	regex,err := regexp.Compile(commentPattern)
	handleError(err)

	if regex.MatchString(source) {
		newLine = regex.ReplaceAllString(source, destination)
	}

	return newLine + "\n"
}

func extractFileExtension(path string) string {
	return filepath.Ext(path)
}

func validateExtension(ext string) bool {
	switch ext{
	case ".go":
		return true
	case ".py":
		return true
	case ".js":
		return true
	case ".ts":
		return true
	default:
		return false
	}
}

func getExtensionToken(ext string) string {
	switch ext{
	case ".go":
		return GolangCommentToken
	case ".py":
		return PythonCommentToken
	case ".js":
		return JavascriptCommentToken
	case ".ts":
		return JavascriptCommentToken
	default:
		return ""
	}
}