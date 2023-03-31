package native1

import (
	"github.com/pinpr/backend-shared-kit/vld"
)

type (
	ImageMimeType string
	AudioMimeType string
	VideoMimeType string
)

const (
	ImageMimeTypeAVIF ImageMimeType = "image/avif"         // avif
	ImageMimeTypeGIF  ImageMimeType = "image/gif"          // gif
	ImageMimeTypeHEIC ImageMimeType = "image/heic"         // heic
	ImageMimeTypeJPEG ImageMimeType = "image/jpeg"         // jpeg jpg
	ImageMimeTypePNG  ImageMimeType = "image/png"          // png
	ImageMimeTypeSVG  ImageMimeType = "image/svg+xml"      // svg svgz
	ImageMimeTypeTIFF ImageMimeType = "image/tiff"         // tif tiff
	ImageMimeTypeWBMP ImageMimeType = "image/vnd.wap.wbmp" // wbmp
	ImageMimeTypeWEBP ImageMimeType = "image/webp"         // webp
	ImageMimeTypeICO  ImageMimeType = "image/x-icon"       // ico
	ImageMimeTypeJNG  ImageMimeType = "image/x-jng"        // jng
	ImageMimeTypeBMP  ImageMimeType = "image/x-ms-bmp"     // bmp

	AudioMimeTypeMID AudioMimeType = "audio/midi"        // mid midi kar
	AudioMimeTypeMP3 AudioMimeType = "audio/mpeg"        // mp3
	AudioMimeTypeOGG AudioMimeType = "audio/ogg"         // ogg
	AudioMimeTypeM4A AudioMimeType = "audio/x-m4a"       // m4a
	AudioMimeTypeRA  AudioMimeType = "audio/x-realaudio" // ra

	VideoMimeType3GP  VideoMimeType = "video/3gpp"      // 3gpp 3gp
	VideoMimeTypeTS   VideoMimeType = "video/mp2t"      // ts
	VideoMimeTypeMP4  VideoMimeType = "video/mp4"       // mp4
	VideoMimeTypeMPEG VideoMimeType = "video/mpeg"      // mpeg mpg
	VideoMimeTypeMOV  VideoMimeType = "video/quicktime" // mov
	VideoMimeTypeWEBM VideoMimeType = "video/webm"      // webm
	VideoMimeTypeFLV  VideoMimeType = "video/x-flv"     // flv
	VideoMimeTypeM4V  VideoMimeType = "video/x-m4v"     // m4v
	VideoMimeTypeMNG  VideoMimeType = "video/x-mng"     // mng
	VideoMimeTypeASF  VideoMimeType = "video/x-ms-asf"  // asx asf
	VideoMimeTypeWMV  VideoMimeType = "video/x-ms-wmv"  // wmv
	VideoMimeTypeAVI  VideoMimeType = "video/x-msvideo" // avi
)

var (
	AllImageMimeTypes = []ImageMimeType{
		ImageMimeTypeAVIF,
		ImageMimeTypeGIF,
		ImageMimeTypeHEIC,
		ImageMimeTypeJPEG,
		ImageMimeTypePNG,
		ImageMimeTypeSVG,
		ImageMimeTypeTIFF,
		ImageMimeTypeWBMP,
		ImageMimeTypeWEBP,
		ImageMimeTypeICO,
		ImageMimeTypeJNG,
		ImageMimeTypeBMP,
	}

	AllAudioMimeTypes = []AudioMimeType{
		AudioMimeTypeMID,
		AudioMimeTypeMP3,
		AudioMimeTypeOGG,
		AudioMimeTypeM4A,
		AudioMimeTypeRA,
	}

	AllVideoMimeTypes = []VideoMimeType{
		VideoMimeType3GP,
		VideoMimeTypeTS,
		VideoMimeTypeASF,
		VideoMimeTypeAVI,
		VideoMimeTypeMP4,
		VideoMimeTypeMPEG,
		VideoMimeTypeMOV,
		VideoMimeTypeWEBM,
		VideoMimeTypeFLV,
		VideoMimeTypeM4V,
		VideoMimeTypeMNG,
		VideoMimeTypeWMV,
	}
)

func (e ImageMimeType) IsValid() error {
	return vld.EnumIsValid(&e, AllImageMimeTypes)
}

func (e AudioMimeType) IsValid() error {
	return vld.EnumIsValid(&e, AllAudioMimeTypes)
}

func (e VideoMimeType) IsValid() error {
	return vld.EnumIsValid(&e, AllVideoMimeTypes)
}
