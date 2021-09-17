/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package images

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"strings"

	"github.com/oliamb/cutter"
)

// ErrImageFormatNotFound is used when the provided image format is unsupported
var ErrImageFormatNotFound = fmt.Errorf("Image format not found")

/*
IImageCropper is an interface for cropping images
*/
type IImageCropper interface {
	Crop(imageBytes []byte, cropOptions CropOptions) (*bytes.Buffer, error)
}

/*
ImageCropper crops images
*/
type ImageCropper struct{}

/*
Crop takes an image reader and performs a crop based on the options
provided. It returns the newly cropped image in the form of a
PNG
*/
func (ic ImageCropper) Crop(imageBytes []byte, cropOptions CropOptions) (*bytes.Buffer, error) {
	var (
		err           error
		originalImage image.Image
		formatName    string
		croppedImage  image.Image
	)

	result := new(bytes.Buffer)

	if originalImage, formatName, err = ic.decodeImage(imageBytes); err != nil {
		return result, fmt.Errorf("Error decoding image in Crop: %w", err)
	}

	config := cutter.Config{
		Width:  cropOptions.Width,
		Height: cropOptions.Height,
	}

	if cropOptions.AnchorMode == CropAnchorModeTopLeft {
		config.Mode = cutter.TopLeft
		config.Anchor = cropOptions.Anchor
	} else {
		config.Mode = cutter.Centered
	}

	if cropOptions.UseRatio {
		config.Options = cutter.Ratio
	}

	if croppedImage, err = cutter.Crop(
		originalImage,
		config,
	); err != nil {
		return result, fmt.Errorf("Error cropping image: %w", err)
	}

	if strings.Contains(formatName, "jpg") || strings.Contains(formatName, "jpeg") {
		if err = jpeg.Encode(result, croppedImage, nil); err != nil {
			return result, fmt.Errorf("Unable to encode cropped image back to JPG: %w", err)
		}

		return result, nil
	}

	if strings.Contains(formatName, "png") {
		if err = png.Encode(result, croppedImage); err != nil {
			return result, fmt.Errorf("Unable to encode cropped image back to PNG: %w", err)
		}

		return result, nil
	}

	if strings.Contains(formatName, "gif") {
		if err = gif.Encode(result, croppedImage, nil); err != nil {
			return result, fmt.Errorf("Unable to encode cropped image back to GIF: %w", err)
		}

		return result, nil
	}

	return result, ErrImageFormatNotFound
}

func (ic ImageCropper) decodeImage(imageBytes []byte) (image.Image, string, error) {
	var (
		err        error
		formatName string
		result     image.Image
	)

	if result, formatName, err = image.Decode(bytes.NewReader(imageBytes)); err != nil {
		return result, "", fmt.Errorf("Error decoding image before cropping: %w", err)
	}

	if _, formatName, err = image.DecodeConfig(bytes.NewReader(imageBytes)); err != nil {
		return result, "", fmt.Errorf("Error determining image configuration when decoding image for cropping: %w", err)
	}

	return result, formatName, nil
}
