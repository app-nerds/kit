/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package images

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"strings"

	"github.com/nfnt/resize"
)

/*
IResizer is an interface to describe structs that resize images
*/
type IResizer interface {
	ResizeImage(source io.ReadSeeker, contentType string, imageSize ImageSize) (*bytes.Buffer, error)
}

/*
A Resizer contains methods for resizing images and re-encoding them.
*/
type Resizer struct{}

/*
ResizeImage takes a source image, content type (MIME), and a image size (
THUMBNAIL, SMALL, MEDIUM, LARGE) and resizes proportionally.
*/
func (r *Resizer) ResizeImage(source io.ReadSeeker, contentType string, imageSize ImageSize) (*bytes.Buffer, error) {
	var err error
	var result *bytes.Buffer
	var sourceImage image.Image

	if !r.isValidImageFormat(contentType) {
		return result, ErrInvalidFileType
	}

	source.Seek(0, 0)

	if sourceImage, err = r.readSourceImage(source); err != nil {
		return result, err
	}

	sourceHeight := sourceImage.Bounds().Max.Y - sourceImage.Bounds().Min.Y
	sourceWidth := sourceImage.Bounds().Max.X - sourceImage.Bounds().Min.Y

	adjustedHeight := r.calculateHeight(sourceHeight, imageSize)
	adjustedWidth := r.calculateWidth(sourceWidth, imageSize)

	percentHeightChange := r.calculateHeightChangePercentage(adjustedHeight, sourceHeight)
	percentWidthChange := r.calculateWidthChangePercentage(adjustedWidth, sourceWidth)

	percent := r.determinePercentageChangeToMake(percentWidthChange, percentHeightChange)

	newHeight := r.calculateNewHeight(sourceHeight, percent)
	newWidth := r.calculateNewWidth(sourceWidth, percent)

	return r.resizeImage(sourceImage, contentType, newWidth, newHeight)

}

func (r *Resizer) resizeImage(sourceImage image.Image, contentType string, width, height int) (*bytes.Buffer, error) {
	result := new(bytes.Buffer)
	var err error

	resizedImage := resize.Resize(uint(width), uint(height), sourceImage, resize.Lanczos3)
	encoderType := r.getEncoderType(contentType)

	if encoderType == "jpg" {
		err = jpeg.Encode(result, resizedImage, nil)
		return result, err
	}

	err = png.Encode(result, resizedImage)
	return result, err
}

func (r *Resizer) readSourceImage(sourceImage io.Reader) (image.Image, error) {
	decodedImage, _, err := image.Decode(sourceImage)
	if err != nil {
		return nil, err
	}

	return decodedImage, nil
}
func (r *Resizer) getMultiplierFromSize(imageSize ImageSize) float64 {
	if imageSize == THUMBNAIL {
		return 0.10
	}

	if imageSize == SMALL {
		return 0.25
	}

	if imageSize == MEDIUM {
		return 0.50
	}

	return 1.0
}

func (r *Resizer) calculateHeight(height int, imageSize ImageSize) float64 {
	return float64(height) * r.getMultiplierFromSize(imageSize)
}

func (r *Resizer) calculateWidth(width int, imageSize ImageSize) float64 {
	return float64(width) * r.getMultiplierFromSize(imageSize)
}

func (r *Resizer) calculateHeightChangePercentage(adjustedHeight float64, originalHeight int) float64 {
	return adjustedHeight / float64(originalHeight)
}

func (r *Resizer) calculateWidthChangePercentage(adjustedWidth float64, originalWidth int) float64 {
	return adjustedWidth / float64(originalWidth)
}

func (r *Resizer) determinePercentageChangeToMake(widthChangePercentage float64, heightChangePercentage float64) float64 {
	if heightChangePercentage < widthChangePercentage {
		return heightChangePercentage
	}

	return widthChangePercentage
}

func (r *Resizer) calculateNewHeight(originalHeight int, percentageChange float64) int {
	return int(float64(originalHeight) * percentageChange)
}

func (r *Resizer) calculateNewWidth(originalWidth int, percentageChange float64) int {
	return int(float64(originalWidth) * percentageChange)
}

func (r *Resizer) isValidImageFormat(contentType string) bool {
	validMIMETypes := []string{
		"jpg",
		"jpeg",
		"png",
	}

	for _, mimeType := range validMIMETypes {
		if strings.Contains(contentType, mimeType) {
			return true
		}
	}

	return false
}

func (r *Resizer) getEncoderType(contentType string) string {
	if strings.Contains(contentType, "jpg") || strings.Contains(contentType, "jpeg") {
		return "jpg"
	}

	return "png"
}
