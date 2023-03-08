package native1

import (
	"github.com/pinpr/backend-shared-kit/vld"
)

type (
	ImageMimeType string
	VideoMimeType string
)

const (
	ImageMimeTypeBMP  ImageMimeType = "image/bmp"
	ImageMimeTypeGIF  ImageMimeType = "image/gif"
	ImageMimeTypeHEIC ImageMimeType = "image/heic"
	ImageMimeTypeJPEG ImageMimeType = "image/jpeg"
	ImageMimeTypePNG  ImageMimeType = "image/png"
	ImageMimeTypeSVG  ImageMimeType = "image/svg+xml"
	ImageMimeTypeTIFF ImageMimeType = "image/tiff"
	ImageMimeTypeWEBP ImageMimeType = "image/webp"

	VideoMimeTypeASF  VideoMimeType = "video/x-ms-asf"
	VideoMimeTypeAVI  VideoMimeType = "video/x-msvideo"
	VideoMimeTypeMP4  VideoMimeType = "video/mp4"
	VideoMimeTypeMPEG VideoMimeType = "video/mpeg"
	VideoMimeTypeMOV  VideoMimeType = "video/quicktime"
	VideoMimeTypeOGG  VideoMimeType = "video/ogg"
	VideoMimeTypeWEBM VideoMimeType = "video/webm"
	VideoMimeTypeWMV  VideoMimeType = "video/x-ms-wmv"
)

var (
	AllImageMimeTypes = []ImageMimeType{
		ImageMimeTypeBMP,
		ImageMimeTypeGIF,
		ImageMimeTypeHEIC,
		ImageMimeTypeJPEG,
		ImageMimeTypePNG,
		ImageMimeTypeSVG,
		ImageMimeTypeTIFF,
		ImageMimeTypeWEBP,
	}

	AllVideoMimeTypes = []VideoMimeType{
		VideoMimeTypeASF,
		VideoMimeTypeAVI,
		VideoMimeTypeMP4,
		VideoMimeTypeMPEG,
		VideoMimeTypeMOV,
		VideoMimeTypeOGG,
		VideoMimeTypeWEBM,
		VideoMimeTypeWMV,
	}
)

func (e ImageMimeType) IsValid() error {
	return vld.EnumIsValid(&e, AllImageMimeTypes)
}

func (e VideoMimeType) IsValid() error {
	return vld.EnumIsValid(&e, AllVideoMimeTypes)
}
