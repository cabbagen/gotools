package gotools

import (
	"os"
	"testing"

	edgettstool "github.com/cabbagen/edge-tts-tool"
)

func TestGenerateTTS(t *testing.T) {
	thunk, error := GenerateTTS(edgettstool.DEFAULT_LANG, edgettstool.DEFAULT_VOICE, edgettstool.DEFAULT_VOLUME, "你好啊")

	if error != nil {
		t.Errorf("转换: %s\n", error.Error())
		return
	}

	if error := os.WriteFile("./example-1.mp3", thunk, 0777); error != nil {
		t.Errorf("写入失败: %s\n", error.Error())
		return
	}
	t.Logf("写入成功 \n")
}

func TestGenerateTTSFile(t *testing.T) {
	if error := GenerateTTSFile(edgettstool.DEFAULT_LANG, edgettstool.DEFAULT_VOICE, edgettstool.DEFAULT_VOLUME, "你好啊", "./example-2.mp3", 0777); error != nil {
		t.Errorf("写入失败: %s\n", error.Error())
		return
	}
	t.Logf("写入成功 \n")
}

func TestGetTTSVoices(t *testing.T) {
	voices, error := edgettstool.GetVoiceList()

	if error != nil {
		t.Logf("获取声音失败: %s\n", error.Error())
		return
	}
	for _, voice := range voices {
		t.Logf("Local: %s, ShortName: %s, Gender: %s\n", voice.Locale, voice.ShortName, voice.Gender)
	}

	t.Logf("获取成功")
}
