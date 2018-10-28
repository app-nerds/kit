package database

import "github.com/globalsign/mgo/bson"

/*
A DatabaseUploadResponse is used to alert a caller information about their uploaded file.
*/
type DatabaseUploadResponse struct {
	BytesWritten int           `json:"bytesWritten"`
	FileID       bson.ObjectId `json:"fileID"`
	FileName     string        `json:"fileName"`
	Height       int           `json:"height"`
	Width        int           `json:"width"`
}
