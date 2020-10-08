package api

import (
	"io/ioutil"
	"mime/multipart"
	"strconv"

	"github.com/777777miSSU7777777/go-ass/helper"
)

type StorageManager struct {
	storageLocation string
}

func NewStorageManager(storageLocation string) StorageManager {
	return StorageManager{storageLocation}
}

func (storageManager StorageManager) UploadTrack(file *multipart.File) helper.UploadTrackCallback {
	return func(id int64) error {
		fileBytes, err := ioutil.ReadAll(*file)
		if err != nil {
			return err
		}

		dirPath := storageManager.storageLocation + "/" + strconv.FormatInt(id, 10)
		mp3DirPath := dirPath + "/mp3"
		hlsDirPath := dirPath + "/hls"
		mp3FileName := strconv.FormatInt(id, 10) + ".mp3"
		err = helper.SaveMP3File(mp3DirPath, mp3FileName, fileBytes)
		if err != nil {
			helper.DeleteDir(dirPath)
			return err
		}

		mp3FilePath := mp3DirPath + "/" + mp3FileName
		err = helper.TranscodeToHLS(hlsDirPath, mp3FilePath)
		if err != nil {
			helper.DeleteDir(dirPath)
			return err
		}

		return nil
	}
}

func (storageManager StorageManager) DeleteTrack(id int64) helper.DeleteTrackCallback {
	return func() error {
		dirPath := storageManager.storageLocation + "/" + strconv.FormatInt(id, 10)
		err := helper.DeleteDir(dirPath)
		if err != nil {
			return err
		}

		return nil
	}
}
