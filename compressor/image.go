package compressor

import (
	"bytes"
	"image/jpeg"
	"image/png"

	"github.com/joway/libimagequant-go/pngquant"
)

type IImageCompressAble interface {
	Compress(imgDatas []byte, quality int, speed int) ([]byte, error)
}

func NewImageCompressor(imageType string) IImageCompressAble {
	if imageType == "jpg" {
		return JPGCompressor{}
	}
	if imageType == "png" {
		return PNGCompressor{}
	}
	return nil
}

type JPGCompressor struct {
}

func (jpgc JPGCompressor) Compress(imgDatas []byte, quality int, speed int) ([]byte, error) {
	img, error := jpeg.Decode(bytes.NewReader(imgDatas))

	if error != nil {
		return nil, error
	}

	imgBuf := bytes.Buffer{}

	if error := jpeg.Encode(&imgBuf, img, &jpeg.Options{Quality: quality}); error != nil {
		return nil, error
	}

	if imgBuf.Len() > len(imgDatas) {
		return imgDatas, nil
	}
	return imgBuf.Bytes(), nil
}

type PNGCompressor struct {
}

func (pngc PNGCompressor) Compress(imgDatas []byte, quality int, speed int) ([]byte, error) {
	img, error := png.Decode(bytes.NewReader(imgDatas))

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
