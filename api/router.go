package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupAPIRouter(app *fiber.App, api API) {
	publicAPI := app.Group("/api")
	publicAPI.Get("/audio/tracks", api.GetAllTracks)
	publicAPI.Get("/audio/tracks/:trackId", api.GetTrackByID)
	publicAPI.Get("/audio/playlists", api.GetAllPlaylists)
	publicAPI.Get("/audio/playlists/:playlistId", api.GetPlaylistByID)

	privateAPI := app.Group("/api")
	privateAPI.Use(JWTAuthMiddleware)
	privateAPI.Post("/audio/tracks", api.AddNewTrack)
	privateAPI.Put("/audio/tracks/:trackId", api.UpdateTrackByID)
	privateAPI.Delete("/audio/tracks/:trackId", api.DeleteTrackByID)
	privateAPI.Get("/audio/user-list", api.GetUserTrackList)
	privateAPI.Patch("/audio/user-list/add", api.AddTracksToUserTrackList)
	privateAPI.Patch("/audio/user-list/delete", api.DeleteTracksFromUserList)
	privateAPI.Post("/audio/playlists", api.CreateNewPlaylist)
	privateAPI.Delete("/audio/playlists/:playlistId", api.DeletePlaylistByID)
	privateAPI.Patch("audio/playlists/:playlistId/add", api.AddTracksToPlaylist)
	privateAPI.Patch("/audio/playlists/:playlistId/delete", api.DeleteTracksFromPlaylist)
	privateAPI.Get("/audio/user-playlists", api.GetUserPlaylists)
	privateAPI.Patch("/audio/user-playlists/add", api.AddPlaylistsToUserList)
	privateAPI.Patch("/audio/user-playlists/delete", api.DeletePlaylistsFromUserList)
}

func SetupAuthRouter(app *fiber.App, api API) {
	authAPI := app.Group("/auth")
	authAPI.Post("/signup", api.SignUp)
	authAPI.Post("/signin", api.SignIn)
	authAPI.Post("/refresh-token", api.RefreshToken)
	authAPI.Delete("/signout", api.SignOut)
}

func SetupStreamRouter(app *fiber.App, streamAPI StreamAPI) {
	streamRouter := app.Group("/media")
	streamRouter.Get("/:trackId/stream/", streamAPI.Stream)
	streamRouter.Get("/:trackId/stream/:quality", streamAPI.Stream)
	streamRouter.Get("/:trackId/stream/:seg", streamAPI.Stream)
	streamRouter.Get("/:trackId/download", streamAPI.Download)
}
