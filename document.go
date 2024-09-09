package gotools

import (
	"github.com/pdfcpu/pdfcpu/pkg/cli"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

// 执行 cli 命令
func iCliExec(command *cli.Command) error {
	if _, error := cli.Process(command); error != nil {
		return error
	}
	return nil
}

// 检验 PDF 文件
func PDFValidateFile(files []string) error {
	return iCliExec(cli.ValidateCommand(files, nil))
}

// 创建 pdf 文件
func PDFCreateFile(inFilePDF, inFileJSON, outFilePDF string) error {
	return iCliExec(cli.CreateCommand(inFilePDF, inFileJSON, outFilePDF, nil))
}

// 合并 pdf 文件
func PDFMergeFiles(inFilesPDF []string, outFilePDF string) error {
	return iCliExec(cli.MergeCreateCommand(inFilesPDF, outFilePDF, true, nil))
}

// 图片转 pdf 文件
func PDFImportImages(inFilesIMG []string, outFilePDF, description string) error {
	imgImport, error := pdfcpu.ParseImportDetails(description, types.POINTS)

	if error != nil {
		return error
	}
	return iCliExec(cli.ImportImagesCommand(inFilesIMG, outFilePDF, imgImport, nil))
}

// pdf 导出图片
func PDFExportImages(inFile, outDir string, pageSelection []string) error {
	return iCliExec(cli.ExtractImagesCommand(inFile, outDir, pageSelection, nil))
}
