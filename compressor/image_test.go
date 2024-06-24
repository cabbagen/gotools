package compressor

import (
	"os"
	"testing"
)

func TestNewImageCompressor(t *testing.T) {
	file, error := os.ReadFile("../example/691668.jpg")

	if error != nil {
		t.Errorf("TestNewImageCompressor error: %s", error.Error())
		return
	}

	imgDatas, error := NewImageCompressor("jpg").Compress(file, 45, 10)

	if error != nil {
		t.Errorf("TestNewImageCompressor error: %s", error.Error())
		return
	}

	if error := os.WriteFile("../example/news-45.jpg", imgDatas, os.ModePerm); error != nil {
		t.Errorf("TestNewImageCompressor error: %s", error.Error())
		return

	}
	t.Logf("运行成功")
}
