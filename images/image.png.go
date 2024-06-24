package images

import (
	"bytes"
	"image/png"

	"github.com/joway/libimagequant-go/pngquant"
)

type PNGImage struct {
	BaseImage
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
