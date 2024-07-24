package gotools

import (
	"bytes"
	"image/jpeg"
	"image/png"

	"github.com/joway/libimagequant-go/pngquant"
	gosseract "github.com/otiai10/gosseract/v2"
)

// 图片压缩
func ImageCompress(imageType string, imgDatas []byte, quality, speed int) ([]byte, error) {
	if imageType == "jpg" {
		return imageCompressWithJPG(imgDatas, quality)
	}
	if imageType == "png" {
		return imageCompressWithPNG(imgDatas, quality, speed)
	}
	return []byte{}, nil
}

// jpg 图片压缩
func imageCompressWithJPG(imgThunk []byte, quality int) ([]byte, error) {
	img, error := jpeg.Decode(bytes.NewReader(imgThunk))

	if error != nil {
		return nil, error
	}

	imgBuf := bytes.Buffer{}

	if error := jpeg.Encode(&imgBuf, img, &jpeg.Options{Quality: quality}); error != nil {
		return nil, error
	}

	if imgBuf.Len() > len(imgThunk) {
		return imgThunk, nil
	}

	return imgBuf.Bytes(), nil
}

// png 图片压缩
func imageCompressWithPNG(imgThunk []byte, quality, speed int) ([]byte, error) {
	img, error := png.Decode(bytes.NewReader(imgThunk))

	if error != nil {
		return nil, error
	}

	outputImg, error := pngquant.Compress(img, quality, speed)

	imgBuf := bytes.Buffer{}

	if error := png.Encode(&imgBuf, outputImg); error != nil {
		return nil, error
	}

	return imgBuf.Bytes(), nil
}

// 图片 OCR - 文件流
// 英文：eng、简体中文：chi_sim、简体中文反向：chi_sim_vert、繁体中文：chi_tra、繁体中文反向：chi_tra_vert
func ImageChunkOCRToText(imageChunk []byte, languages ...string) (string, error) {
	client := gosseract.NewClient()

	defer client.Close()

	// 默认英文，
	if len(languages) > 0 {
		client.SetLanguage(languages...)
	}

	if error := client.SetImageFromBytes(imageChunk); error != nil {
		return "", error
	}

	return client.Text()
}
