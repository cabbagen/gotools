package gotools

import (
	edge_tts_go "github.com/pp-group/edge-tts-go"
	"github.com/pp-group/edge-tts-go/biz/service/tts/edge"
)

var DefaultTTSOptions []edge.Option = []edge.Option{
	edge.WithVoice("zh-CN-XiaoyiNeural"), edge.WithRate("+0%"), edge.WithVolume("+0%"), edge.WithPitch("+0Hz"),
}

// 生成 tts 参数
func CreateTTSOptions(voice, rate, volume, pitch string) []edge.Option {
	return []edge.Option{
		edge.WithVoice(voice), edge.WithRate(rate), edge.WithVolume(volume), edge.WithPitch(pitch),
	}
}

// tts 服务
func GenerateTTS(text, output string, options []edge.Option) (string, error) {
	communicate, error := edge.NewCommunicate(text, options...)

	if error != nil {
		return "", error
	}

	speech, error := edge_tts_go.NewLocalSpeech(communicate, output)

	if error != nil {
		return "", error
	}

	_, callback := speech.GenTTS()

	if error := callback(); error != nil {
		return "", error
	}

	return speech.URL(speech.FileName)
}
