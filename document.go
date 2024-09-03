package gotools

import (
	"github.com/pdfcpu/pdfcpu/pkg/cli"
)

// 检验 PDF 文件
func PDFValidateFile(files []string) error {
	if _, error := cli.Process(cli.ValidateCommand(files, nil)); error != nil {
		return error
	}
	return nil
}

// 创建 demo pdf 文件
func PDFCreateFile(inFilePDF, inFileJSON, outFilePDF string) error {
	if _, error := cli.Process(cli.CreateCommand(inFilePDF, inFileJSON, outFilePDF, nil)); error != nil {
		return error
	}
	return nil
}

// 图片转 pdf 文件
func ImageToPDFFile() {

}
