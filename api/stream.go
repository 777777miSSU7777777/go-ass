package api

import (
	"github.com/gofiber/fiber/v2"
)

type StreamAPI struct {
	storageManager StorageManager
}

func NewStreamAPI(storageManager StorageManager) StreamAPI {
	return StreamAPI{storageManager}
}

func (streamAPI StreamAPI) Stream(ctx *fiber.Ctx) error {
	trackID := ctx.Params("trackId")

	seg := ctx.Params("seg")

	if seg == "" {
		quality := ctx.Params("quality")

		if quality == "" {
			streamAPI.storageManager.ServeMasterM3u8(ctx, trackID)
			return nil
		}

		streamAPI.storageManager.ServeQualityM3u8(ctx, trackID, quality)
		return nil
	}

	streamAPI.storageManager.ServeTs(ctx, trackID, seg)
	return nil
}

func (streamAPI StreamAPI) Download(ctx *fiber.Ctx) error {
	trackID := ctx.Params("trackId")

	streamAPI.storageManager.ServeMp3(ctx, trackID)
	return nil
}
