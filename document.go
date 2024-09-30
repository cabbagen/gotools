package gotools

import (
	"errors"
	"fmt"
	"image/jpeg"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gen2brain/go-fitz"
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

// pdf 检出图片
func PDFExtractImages(inFile, outDir string, pageSelection []string) error {
	return iCliExec(cli.ExtractImagesCommand(inFile, outDir, pageSelection, nil))
}

// pdf 页面导出图片
func PDFExportImages(inFile, outDir string, pageSelection []string) error {
	document, error := fitz.New(inFile)

	if error != nil {
		return error
	}

	defer document.Close()

	// 不传为全部截取
	if len(pageSelection) == 0 {
		pageSelection = append(pageSelection, fmt.Sprintf("1-%d", document.NumPage()))
	}

	selection := MapBySlice[string, int](strings.Split(pageSelection[0], "-"), func(pageText string, index int) int {
		if page, error := strconv.Atoi(pageText); error == nil {
			return page
		}
		return -1
	})

	if IsExistBySlice(selection, -1) || len(selection) != 2 {
		return errors.New("pageSelection error. example: 1-4")
	}

	for index := 0; index < document.NumPage(); index++ {
		if selection[0] > index+1 || index+1 > selection[1] {
			continue
		}

		image, error := document.Image(index)

		if error != nil {
			return error
		}

		file, error := os.Create(path.Join(outDir, fmt.Sprintf("%s-%d.jpg", strings.TrimSuffix(path.Base(inFile), path.Ext(path.Base(inFile))), index+1)))

		if error != nil {
			return error
		}

		if error := jpeg.Encode(file, image, &jpeg.Options{Quality: jpeg.DefaultQuality}); error != nil {
			return error
		}

		file.Close()
	}

	return nil
}

// pdf 添加水印

// 提取指定范围页面

// 删除页面

// 页面排序
