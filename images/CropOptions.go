/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package images

import (
	"image"
)

/*
CropOptions describes how an image is cropped. Width/Height defines how big to make
the image. AnchorMode describes how to apply the crop anchor point. Valid options are
top-left, centered. Anchor describes the X/Y offset from top left if the anchor mode
is top-left. UseRatio, when true, makes width/height a ratio-based instead of pixels.
*/
type CropOptions struct {
	Anchor     image.Point
	AnchorMode CropAnchorMode
	Height     int
	UseRatio   bool
	Width      int
}

/*
NewCropOptions creates a new structure with default values
filled in
*/
func NewCropOptions() *CropOptions {
	return &CropOptions{
		Anchor:     image.Point{X: 0, Y: 0},
		AnchorMode: CropAnchorModeTopLeft,
		Height:     0,
		UseRatio:   false,
		Width:      0,
	}
}
