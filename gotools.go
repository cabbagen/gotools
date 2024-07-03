package gotools

import "github.com/cabbagen/gotools/images"

func NewImageCompressor(imageType string) images.IImageCompressAble {
	if imageType == "jpg" {
		return images.JPGImage{}
	}
	if imageType == "png" {
		return images.PNGImage{}
	}
	return nil
}
