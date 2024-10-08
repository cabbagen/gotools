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

func TestPDFImportImages(t *testing.T) {
	images := []string{
		"./resources/images/ducati.png", "./resources/images/pexels.jpg", "./resources/images/simple.jpg",
	}
	if error := PDFImportImages(images, "./resources/pdfs/images.pdf", "pos:c, scale:0.9 rel"); error != nil {
		log.Fatalf("create images pdf file error: %s\n", error.Error())
		return
	}
	log.Printf("create success")
}

func TestPDFExtractImages(t *testing.T) {
	if error := PDFExtractImages("./resources/pdfs/images.pdf", "./resources/images/", []string{"1-2"}); error != nil {
		log.Fatalf("extract images error: %s\n", error.Error())
		return
	}
	log.Printf("extract success")
}

func TestPDFExportImages(t *testing.T) {
	if error := PDFExportImages("./resources/pdfs/example.pdf", "./resources/images/", []string{}); error != nil {
		log.Fatalf("export images error: %s\n", error.Error())
		return
	}
	log.Printf("export success")
}
