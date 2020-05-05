/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
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
	"github.com/pkg/errors"
)

// ErrImageFormatNotFound is used when the provided image format is unsupported
var ErrImageFormatNotFound = fmt.Errorf("Image format not found")

/*
IImageCropper is an interface for cropping images
*/
type IImageCropper interface {
	Crop(imageBytes []byte, cropOptions *CropOptions) (*bytes.Buffer, error)
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
func (ic *ImageCropper) Crop(imageBytes []byte, cropOptions *CropOptions) (*bytes.Buffer, error) {
	var err error
	var originalImage image.Image
	var formatName string
	var croppedImage image.Image
	var anchorMode cutter.AnchorMode
	var options cutter.Option

	result := new(bytes.Buffer)

	if originalImage, formatName, err = ic.decodeImage(imageBytes); err != nil {
		return result, errors.Wrapf(err, "Error decoding image in Crop")
	}

	if cropOptions.AnchorMode == CropAnchorModeTopLeft {
		anchorMode = cutter.TopLeft
	} else {
		anchorMode = cutter.Centered
	}

	if cropOptions.UseRatio {
		options = cutter.Ratio
	}

	if croppedImage, err = cutter.Crop(
		originalImage,
		cutter.Config{
			Width:   cropOptions.Width,
			Height:  cropOptions.Height,
			Anchor:  cropOptions.Anchor,
			Mode:    anchorMode,
			Options: options,
		},
	); err != nil {
		return result, errors.Wrapf(err, "Error cropping image")
	}

	if strings.Contains(formatName, "jpg") || strings.Contains(formatName, "jpeg") {
		if err = jpeg.Encode(result, croppedImage, nil); err != nil {
			return result, errors.Wrapf(err, "Unable to encode cropped image back to JPG")
		}

		return result, nil
	}

	if strings.Contains(formatName, "png") {
		if err = png.Encode(result, croppedImage); err != nil {
			return result, errors.Wrapf(err, "Unable to encode cropped image back to PNG")
		}

		return result, nil
	}

	if strings.Contains(formatName, "gif") {
		if err = gif.Encode(result, croppedImage, nil); err != nil {
			return result, errors.Wrapf(err, "Unable to encode cropped image back to GIF")
		}

		return result, nil
	}

	return result, ErrImageFormatNotFound
}

func (ic *ImageCropper) decodeImage(imageBytes []byte) (image.Image, string, error) {
	var err error
	var formatName string
	var result image.Image

	if result, formatName, err = image.Decode(bytes.NewReader(imageBytes)); err != nil {
		return result, "", errors.Wrapf(err, "Error decoding image before cropping")
	}

	if _, formatName, err = image.DecodeConfig(bytes.NewReader(imageBytes)); err != nil {
		return result, "", errors.Wrapf(err, "Error determining image configuration when decoding image for cropping")
	}

	return result, formatName, nil
}
