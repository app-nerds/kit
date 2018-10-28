package database

import (
	"io"
	"path/filepath"

	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"

	"github.com/globalsign/mgo"
)

/*
A MongoUploader uploads files to a MongoDB database. This struct
satisfies the DatabaseUploader interface
*/
type MongoUploader struct {
	DB DocumentDatabase
}

/*
Upload uploads a file into the MongoDB GridFS system
*/
func (u *MongoUploader) Upload(reader io.Reader, name, path string) (*DatabaseUploadResponse, error) {
	var err error
	var file *mgo.GridFile

	result := &DatabaseUploadResponse{}
	totalBytesWritten := 0
	bytesWritten := 0
	bytesRead := 0

	buffer := make([]byte, 2048)
	name = u.sanitizeFileName(name)

	/*
	 * Create the file in GridFS
	 */
	if file, err = u.DB.GridFS(path).Create(name); err != nil {
		return result, errors.Wrapf(err, "Error uploading file '%s' to MongoDB GridFS", name)
	}

	defer file.Close()

	/*
	 * Read bytes from the file, write bytes to GridFS
	 */
	for {
		bytesRead, err = reader.Read(buffer)

		if err == nil {
			if bytesRead > 0 {
				if bytesWritten, err = file.Write(buffer); err != nil {
					result.BytesWritten = bytesWritten
					return result, errors.Wrapf(err, "Error writing file bytes to GridFS")
				}

				totalBytesWritten += bytesWritten
			}
		}

		if err == io.EOF {
			result.BytesWritten = totalBytesWritten
			result.FileID = (file.Id()).(bson.ObjectId)
			result.FileName = name

			return result, nil
		}
	}
}

func (u *MongoUploader) sanitizeFileName(name string) string {
	return filepath.Base(filepath.Clean(name))
}
