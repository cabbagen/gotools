/**
 * 部分 golang 依赖 C++ 库，需要提前配置环境
 * macOS：
 * brew install webp tesseract tesseract-lang // 英文：eng、简体中文：chi_sim、简体中文反向：chi_sim_vert、繁体中文：chi_tra、繁体中文反向：chi_tra_vert => 默认英文
 *
 * for libraries：github.com/kolesa-team/go-webp github.com/otiai10/gosseract/v2
 */
package gotools

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/joway/libimagequant-go/pngquant"
	"github.com/kolesa-team/go-webp/decoder"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"github.com/nfnt/resize"
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
	if imageType == "webp" {
		return imageCompressWithWebp(imgDatas, quality)
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

// brew install webp
// webp 图片压缩
func imageCompressWithWebp(imgThunk []byte, quality int) ([]byte, error) {
	img, error := webp.Decode(bytes.NewReader(imgThunk), &decoder.Options{})

	if error != nil {
		return nil, error
	}

	options, error := encoder.NewLossyEncoderOptions(encoder.PresetDefault, float32(quality))

	if error != nil {
		return nil, error
	}

	imgBuf := bytes.Buffer{}

	if error := webp.Encode(&imgBuf, img, options); error != nil {
		return nil, error
	}

	return imgBuf.Bytes(), nil
}

// 图片转换
func ImageConvert(sourceType, targetType string, imgDatas []byte) ([]byte, error) {
	typeKey := fmt.Sprintf("%s-%s", sourceType, targetType)

	typeFuncMap := map[string]func(imageThunk []byte) ([]byte, error){
		"jpg-png":  imageConvertJPGToPNG,
		"jpg-webp": imageConvertJPGToWEBP,
		"png-jpg":  imageConvertPNGToJPG,
		"png-webp": imageConvertPNGToWEBP,
		"webp-jpg": imageConvertWEBPToJPG,
		"webp-png": imageConvertWEBPToPNG,
	}

	if imageConvertFunc, ok := typeFuncMap[typeKey]; ok {
		return imageConvertFunc(imgDatas)
	}
	return []byte{}, errors.New("sourceType, targetType 传参错误")
}

// 图片格式转换: jpg => png
func imageConvertJPGToPNG(imgThunk []byte) ([]byte, error) {
	img, error := jpeg.Decode(bytes.NewReader(imgThunk))

	if error != nil {
		return nil, error
	}

	imgBuf := bytes.Buffer{}

	if error := png.Encode(&imgBuf, img); error != nil {
		return nil, error
	}

	return imgBuf.Bytes(), nil
}

// 图片格式转换: jpg => webp
func imageConvertJPGToWEBP(imgThunk []byte) ([]byte, error) {
	img, error := jpeg.Decode(bytes.NewReader(imgThunk))

	if error != nil {
		return nil, error
	}

	imgBuf := bytes.Buffer{}

	options, error := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)

	if error != nil {
		return nil, error
	}

	if error := webp.Encode(&imgBuf, img, options); error != nil {
		return nil, error
	}

	return imgBuf.Bytes(), nil
}

// 图片格式转换: png => jpg
func imageConvertPNGToJPG(imgThunk []byte) ([]byte, error) {
	img, error := png.Decode(bytes.NewReader(imgThunk))

	if error != nil {
		return nil, error
	}

	imgBuf := bytes.Buffer{}

	if error := jpeg.Encode(&imgBuf, img, &jpeg.Options{Quality: 100}); error != nil {
		return nil, error
	}

	return imgBuf.Bytes(), nil
}

// 图片格式转换: png => webp
func imageConvertPNGToWEBP(imgThunk []byte) ([]byte, error) {
	img, error := png.Decode(bytes.NewReader(imgThunk))

	if error != nil {
		return nil, error
	}

	imgBuf := bytes.Buffer{}

	options, error := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)

	if error != nil {
		return nil, error
	}

	if error := webp.Encode(&imgBuf, img, options); error != nil {
		return nil, error
	}

	return imgBuf.Bytes(), nil
}

// 图片格式转换: webp => jpg
func imageConvertWEBPToJPG(imgThunk []byte) ([]byte, error) {
	img, error := webp.Decode(bytes.NewReader(imgThunk), &decoder.Options{})

	if error != nil {
		return nil, error
	}

	imgBuf := bytes.Buffer{}

	if error := jpeg.Encode(&imgBuf, img, &jpeg.Options{Quality: 100}); error != nil {
		return nil, error
	}

	return imgBuf.Bytes(), nil
}

// 图片格式转换: webp => png
func imageConvertWEBPToPNG(imgThunk []byte) ([]byte, error) {
	img, error := webp.Decode(bytes.NewReader(imgThunk), &decoder.Options{})

	if error != nil {
		return nil, error
	}

	imgBuf := bytes.Buffer{}

	if error := png.Encode(&imgBuf, img); error != nil {
		return nil, error
	}

	return imgBuf.Bytes(), nil
}

// 转换图片大小
func ImageResize(width, height uint, imgDatas []byte) (image.Image, error) {
	img, _, error := image.Decode(bytes.NewBuffer(imgDatas))

	if error != nil {
		return nil, error
	}
	return resize.Resize(width, height, img, resize.Lanczos3), nil
}

// 图片 OCR 识别
func ImageChunkOCRToText(imageChunk []byte, languages ...string) (string, error) {
	client := gosseract.NewClient()

	defer client.Close()

	// 修改识别语言
	if len(languages) > 0 {
		client.SetLanguage(languages...)
	}

	if error := client.SetImageFromBytes(imageChunk); error != nil {
		return "", error
	}

	return client.Text()
}
