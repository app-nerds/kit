/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package images

import (
	"errors"
)

// ErrInvalidFileType is an error when the user uploads a bad file type
var ErrInvalidFileType = errors.New("Invalid image type. Supported image formats are JPG and PNG")
