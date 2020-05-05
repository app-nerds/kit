/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package images

/*
ImageSize describes generic image sizes
*/
type ImageSize string

const (
	DEFAULT   ImageSize = "default"
	THUMBNAIL ImageSize = "thumbnail"
	SMALL     ImageSize = "small"
	MEDIUM    ImageSize = "medium"
	LARGE     ImageSize = "large"
)

func (is ImageSize) String() string {
	return string(is)
}
