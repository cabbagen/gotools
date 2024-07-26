package gotools

import (
	"encoding/base64"
	"os"
	"testing"
)

func TestImageCompress(t *testing.T) {
	file, error := os.ReadFile("./images/JnD5BF9Uq-1200x800.webp")

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
