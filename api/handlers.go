package api

import (
	"github.com/777777miSSU7777777/go-ass/model"

	"github.com/777777miSSU7777777/go-ass/service"
	"github.com/gofiber/fiber/v2"
)

type API struct {
	svc            service.Service
	storageManager StorageManager
}

func NewAPI(svc service.Service, storageManager StorageManager) API {
	return API{svc, storageManager}
}

func (api API) AddNewTrack(ctx *fiber.Ctx) error {
	artistID := ctx.FormValue("artistID")
	title := ctx.FormValue("title")
	genreID := ctx.FormValue("genreID")

	userID := ctx.Context().UserValue("userID").(string)

	fileHeader, err := ctx.FormFile("audiofile")
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	file, err := fileHeader.Open()
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	defer file.Close()

	uploadTrack := api.storageManager.UploadTrack(&file)

	newTrack, err := api.svc.AddNewTrack(title, artistID, genreID, userID, uploadTrack)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": newTrack,
	})
	return nil
}

func (api API) GetAllTracks(ctx *fiber.Ctx) error {
	allTracks, err := api.svc.GetAllTracks()
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": allTracks,
	})
	return nil
}

func (api API) GetTrackByID(ctx *fiber.Ctx) error {
	trackID := ctx.Params("trackId")

	track, err := api.svc.GetTrackByID(trackID)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":   false,
		"data": track,
	})
	return nil
}

func (api API) UpdateTrackByID(ctx *fiber.Ctx) error {
	trackID := ctx.Params("trackId")

	var req model.UpdateTrackByIDRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	updatedTrack, err := api.svc.UpdateTrackByID(trackID, req.TrackTitle, req.ArtistID, req.GenreID)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": updatedTrack,
	})
	return nil
}

func (api API) DeleteTrackByID(ctx *fiber.Ctx) error {
	trackID := ctx.Params("trackId")

	deleteTrack := api.storageManager.DeleteTrack(trackID)

	err := api.svc.DeleteTrackByID(trackID, deleteTrack)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok": true,
	})
	return nil
}

func (api API) SignUp(ctx *fiber.Ctx) error {
	var req model.SignUpRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	err = api.svc.SignUp(req.Email, req.Username, req.Password)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok": true,
	})
	return nil
}

func (api API) SignIn(ctx *fiber.Ctx) error {
	var req model.SignInRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	accessToken, refreshToken, err := api.svc.SignIn(req.Email, req.Password)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	res := model.SignInResponse{AccessToken: accessToken, RefreshToken: refreshToken}
	ctx.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": res,
	})
	return nil
}

func (api API) RefreshToken(ctx *fiber.Ctx) error {
	var req model.RefreshTokenRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	accessToken, refreshToken, err := api.svc.RefreshToken(req.RefreshToken)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	res := model.RefreshTokenResponse{AccessToken: accessToken, RefreshToken: refreshToken}
	ctx.Status(200).JSON(fiber.Map{
		"ok":   false,
		"data": res,
	})
	return nil
}

func (api API) SignOut(ctx *fiber.Ctx) error {
	var req model.SignOutRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	err = api.svc.SignOut(req.RefreshToken)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok": true,
	})
	return nil
}

func (api API) GetUserTrackList(ctx *fiber.Ctx) error {
	userID := ctx.Context().UserValue("userID").(string)

	userTrackList, err := api.svc.GetUserTrackList(userID)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":   false,
		"data": userTrackList,
	})
	return nil
}

func (api API) AddTracksToUserTrackList(ctx *fiber.Ctx) error {
	userID := ctx.Context().UserValue("userID").(string)

	var req model.AddTracksToUserListRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	err = api.svc.AddTracksToUserTrackList(userID, req.TrackList...)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok": true,
	})
	return nil
}

func (api API) DeleteTracksFromUserList(ctx *fiber.Ctx) error {
	userID := ctx.Context().UserValue("userID").(string)

	var req model.DeleteTracksFromUserListRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	err = api.svc.DeleteTracksFromUserTrackList(userID, req.TrackList...)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok": true,
	})
	return nil
}

func (api API) GetAllPlaylists(ctx *fiber.Ctx) error {
	allPlaylists, err := api.svc.GetAllPlaylists()
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": allPlaylists,
	})
	return nil
}

func (api API) GetUserPlaylists(ctx *fiber.Ctx) error {
	userID := ctx.Context().UserValue("userID").(string)

	userPlaylists, err := api.svc.GetUserPlaylists(userID)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": userPlaylists,
	})
	return nil
}

func (api API) CreateNewPlaylist(ctx *fiber.Ctx) error {
	userID := ctx.Context().UserValue("userID").(string)

	var req model.CreateNewPlaylistRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	newPlaylist, err := api.svc.CreateNewPlaylist(req.PlaylistTitle, userID, req.TrackList)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": newPlaylist,
	})
	return nil
}

func (api API) GetPlaylistByID(ctx *fiber.Ctx) error {
	playlistID := ctx.Params("playlistId")

	playlist, err := api.svc.GetPlaylistByID(playlistID)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": playlist,
	})
	return nil
}

func (api API) DeletePlaylistByID(ctx *fiber.Ctx) error {
	playlistID := ctx.Params("playlistId")

	userID := ctx.Context().UserValue("userID").(string)

	err := api.svc.DeletePlaylistByID(playlistID, userID)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok": true,
	})
	return nil
}

func (api API) AddTracksToPlaylist(ctx *fiber.Ctx) error {
	playlistID := ctx.Params("playlistID")

	var req model.AddTracksToPlaylistRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	userID := ctx.Context().UserValue("userID").(string)

	err = api.svc.AddTracksToPlaylist(userID, playlistID, req.TrackList)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok": true,
	})
	return nil
}

func (api API) DeleteTracksFromPlaylist(ctx *fiber.Ctx) error {
	playlistID := ctx.Params("playlistID")

	var req model.DeleteTracksFromPlaylistRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	userID := ctx.Context().UserValue("userID").(string)

	err = api.svc.DeleteTracksFromPlaylist(userID, playlistID, req.TrackList)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok": true,
	})
	return nil
}

func (api API) AddPlaylistsToUserList(ctx *fiber.Ctx) error {
	var req model.AddPlaylistsToUserListRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	userID := ctx.Context().UserValue("userID").(string)

	err = api.svc.AddPlaylistsToUserList(userID, req.Playlists...)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok": true,
	})
	return nil
}

func (api API) DeletePlaylistsFromUserList(ctx *fiber.Ctx) error {
	var req model.DeletePlaylistsFromUserListRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	userID := ctx.Context().UserValue("userID").(string)

	err = api.svc.DeletePlaylistsFromUserList(userID, req.Playlists...)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return err
	}

	ctx.Status(200).JSON(fiber.Map{
		"ok": true,
	})
	return nil
}
