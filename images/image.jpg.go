package images

import (
	"bytes"
	"image/jpeg"
)

type JPGImage struct {
	BaseImage
}

func (jpgc JPGImage) Compress(imgThunk []byte, quality int, speed int) ([]byte, error) {
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
