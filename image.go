package gotools

import (
	"bytes"
	"image/jpeg"
	"image/png"

	"github.com/joway/libimagequant-go/pngquant"
)

// 可压缩
type IImageCompressAble interface {
	Compress(imgDatas []byte, quality int, speed int) ([]byte, error)
}

// 压缩文件
func NewImageCompressor(imageType string) IImageCompressAble {
	if imageType == "jpg" {
		return JPGImage{}
	}
	if imageType == "png" {
		return PNGImage{}
	}
	return nil
}

// jpg 图片
type JPGImage struct {
}

func (jpgc JPGImage) Compress(imgThunk []byte, quality int, speed int) ([]byte, error) {
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

// png 图片
type PNGImage struct {
}

func (pngc PNGImage) Compress(imgThunk []byte, quality int, speed int) ([]byte, error) {
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
