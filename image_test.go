package gotools

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"os"
	"testing"
)

func TestImageCompress(t *testing.T) {
	file, error := os.ReadFile("./resources/images/simple.jpg")

	if error != nil {
		t.Logf("读取文件错误: %s\n", error.Error())
		return
	}

	image, error := ImageCompress("webp", file, 20, 40)

	if error != nil {
		t.Logf("文件压缩错误: %s\n", error.Error())
		return
	}

	t.Logf("图片压缩成功\ndata:image/webp;base64,%s", base64.RawStdEncoding.EncodeToString(image))
}

func TestImageConvert(t *testing.T) {
	file, error := os.ReadFile("./resources/images/ex.webp")

	if error != nil {
		t.Logf("读取文件错误: %s\n", error.Error())
		return
	}

	image, error := ImageConvert("webp", "png", file)

	if error != nil {
		t.Logf("文件转换错误: %s\n", error.Error())
		return
	}

	t.Logf("图片转换成功\ndata:image/png;base64,%s", base64.RawStdEncoding.EncodeToString(image))
}

func TestImageResize(t *testing.T) {
	file, error := os.ReadFile("./resources/images/pexels.jpg")

	if error != nil {
		t.Logf("读取文件错误: %s\n", error.Error())
		return
	}

	image, error := ImageResize(200, 200, file)

	if error != nil {
		t.Logf("转换文件错误: %s\n", error.Error())
		return
	}

	imgBuf := bytes.Buffer{}

	if error := jpeg.Encode(&imgBuf, image, &jpeg.Options{Quality: 70}); error != nil {
		t.Logf("写入文件错误: %s\n", error.Error())
		return
	}

	t.Logf("图片转换成功\ndata:image/jpeg;base64,%s", base64.RawStdEncoding.EncodeToString(imgBuf.Bytes()))
}
