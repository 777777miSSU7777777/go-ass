package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StreamAPI struct {
	storageManager StorageManager
}

func NewStreamAPI(storageManager StorageManager) StreamAPI {
	return StreamAPI{storageManager}
}

func (streamAPI StreamAPI) Stream(ctx *fiber.Ctx) error {
	trackID, err := strconv.ParseInt(ctx.Params("trackId"), 10, 64)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

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
	trackID, err := strconv.ParseInt(ctx.Params("trackId"), 10, 64)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	streamAPI.storageManager.ServeMp3(ctx, trackID)
	return nil
}
