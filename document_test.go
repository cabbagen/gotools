package gotools

import (
	"log"
	"testing"
)

func TestPDFValidateFile(t *testing.T) {
	if error := PDFValidateFile([]string{"./resources/pdfs/out.pdf"}); error != nil {
		log.Fatalf("validate pdf file error: %s\n", error.Error())
		return
	}
	log.Println("validate success")
}

func TestPDFCreateFile(t *testing.T) {
	if error := PDFCreateFile("", "./resources/json/createPdf.json", "./resources/pdfs/example.pdf"); error != nil {
		log.Fatalf("create pdf file error: %s\n", error.Error())
		return
	}
	log.Printf("create success")
}
