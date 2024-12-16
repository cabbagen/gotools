package gotools

import (
	"os"

	edgettstool "github.com/cabbagen/edge-tts-tool"
)

func GenerateTTS(lang, voice, volume, text string) ([]byte, error) {
	return edgettstool.NewCommunicate(lang, voice, volume).HandleGenerateTTS(text)
}

func GenerateTTSFile(lang, voice, volume, text, filePath string, mode os.FileMode) error {
	return edgettstool.NewCommunicate(lang, voice, volume).HandleSaveTTSFile(text, filePath, mode)
}

func GetTTSVoices() ([]edgettstool.Voice, error) {
	return edgettstool.GetVoiceList()
}
