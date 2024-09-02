package gotools

import (
	"testing"
)

// func TestGenerateTTS(t *testing.T) {
// 	filename, error := GenerateTTS("你好哦啊，见到你很高兴", "./resources/template-audio", DefaultTTSOptions)

// 	if error != nil {
// 		t.Errorf("生成语音失败: %s\n", error.Error())
// 		return
// 	}

// 	t.Logf("生成语音成功: %s\n", filename)
// }

func TestGenerateSTT(t *testing.T) {
	GenerateSTT()
}
