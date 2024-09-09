package gotools

import (
	"github.com/pdfcpu/pdfcpu/pkg/cli"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

// 检验 PDF 文件
func PDFValidateFile(files []string) error {
	if _, error := cli.Process(cli.ValidateCommand(files, nil)); error != nil {
		return error
	}
	return nil
}

// 创建 pdf 文件
func PDFCreateFile(inFilePDF, inFileJSON, outFilePDF string) error {
	if _, error := cli.Process(cli.CreateCommand(inFilePDF, inFileJSON, outFilePDF, nil)); error != nil {
		return error
	}
	return nil
}

// 合并 pdf 文件
func PDFMergeFiles(inFilesPDF []string, outFilePDF string) error {
	if _, error := cli.Process(cli.MergeCreateCommand(inFilesPDF, outFilePDF, true, nil)); error != nil {
		return error
	}
	return nil
}

// 图片转 pdf 文件
func PDFImportImages(inFilesIMG []string, outFilePDF, description string) error {
	imgImport, error := pdfcpu.ParseImportDetails(description, types.POINTS)

	if error != nil {
		return error
	}

	if _, error := cli.Process(cli.ImportImagesCommand(inFilesIMG, outFilePDF, imgImport, nil)); error != nil {
		return error
	}

	return nil
}

// pdf 导出图片
func PDFExportImages(inFile, outDir string, pageSelection []string) error {
	if _, error := cli.Process(cli.ExtractImagesCommand(inFile, outDir, pageSelection, nil)); error != nil {
		return error
	}

	return nil
}
