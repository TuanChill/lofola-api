package helpers

import (
	"github.com/disintegration/imaging"
)

func ReSizeImageForAvatar(inputPath string, outputPath string) error {
	img, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	// resize the img to width 180px maintaining the aspect ratio (auto height)
	img = imaging.Resize(img, 180, 0, imaging.Lanczos)

	return imaging.Save(img, outputPath)
}

func ReduceImageQuality(inputPath string, outputPath string, qualityPercent int) error {
	img, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	// reduce the img quality to 90%
	return imaging.Save(img, outputPath, imaging.JPEGQuality(qualityPercent))
}
