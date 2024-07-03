package images

type BaseImage struct {
}

type IImageCompressAble interface {
	Compress(imgDatas []byte, quality int, speed int) ([]byte, error)
}
