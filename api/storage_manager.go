package api

import (
	"io/ioutil"
	"mime/multipart"
	"path/filepath"

	"github.com/777777miSSU7777777/go-ass/helper"

	"github.com/gofiber/fiber/v2"
)

type StorageManager struct {
	storageLocation string
}

func NewStorageManager(storageLocation string) StorageManager {
	return StorageManager{storageLocation}
}

func (storageManager StorageManager) UploadTrack(file *multipart.File) helper.UploadTrackCallback {
	return func(id string) error {
		fileBytes, err := ioutil.ReadAll(*file)
		if err != nil {
			return err
		}

		dirPath := storageManager.storageLocation + "/" + id
		mp3DirPath := dirPath + "/mp3"
		hlsDirPath := dirPath + "/hls"
		mp3FileName := id + ".mp3"
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

func (storageManager StorageManager) DeleteTrack(id string) helper.DeleteTrackCallback {
	return func() error {
		dirPath := storageManager.storageLocation + "/" + id
		err := helper.DeleteDir(dirPath)
		if err != nil {
			return err
		}

		return nil
	}
}

func (storageManager StorageManager) ServeMasterM3u8(ctx *fiber.Ctx, id string) {
	m3u8FilePath := filepath.FromSlash(storageManager.storageLocation + "/" + id + "/hls/master.m3u8")
	ctx.Set("Content-Type", "application/x-mpegURL")
	ctx.SendFile(m3u8FilePath)
}

func (storageManager StorageManager) ServeQualityM3u8(ctx *fiber.Ctx, id string, quality string) {
	m3u8FilePath := filepath.FromSlash(storageManager.storageLocation + "/" + id + "/hls/" + quality)
	ctx.Set("Content-Type", "application/x-mpegURL")
	ctx.SendFile(m3u8FilePath)
}

func (storageManager StorageManager) ServeTs(ctx *fiber.Ctx, id string, seg string) {
	tsFilePath := filepath.FromSlash(storageManager.storageLocation + "/" + id + "/hls/" + seg)
	ctx.Set("Content-Type", "video/MP2T")
	ctx.SendFile(tsFilePath)
}

func (storageManager StorageManager) ServeMp3(ctx *fiber.Ctx, id string) {
	mp3FilePath := filepath.FromSlash(storageManager.storageLocation + "/" + id + "/mp3/" + id + ".mp3")
	ctx.Set("Content-Type", "audio/mpeg")
	ctx.SendFile(mp3FilePath)
}
