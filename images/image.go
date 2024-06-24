package images

type BaseImage struct {
}

type IImageCompressAble interface {
	Compress(imgDatas []byte, quality int, speed int) ([]byte, error)
}

func NewImageCompressor(imageType string) IImageCompressAble {
	if imageType == "jpg" {
		return JPGImage{}
	}
	if imageType == "png" {
		return JPGImage{}
	}
	return nil
}
