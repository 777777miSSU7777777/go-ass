package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupAPIRouter(app *fiber.App, api API) {
	publicAPI := app.Group("/api")
	publicAPI.Get("/tracks", api.GetAllTracks)
	publicAPI.Get("/tracks/:trackId", api.GetTrackByID)
	publicAPI.Get("/playlists", api.GetAllPlaylists)
	publicAPI.Get("/playlists/:playlistId", api.GetPlaylistByID)

	privateAPI := app.Group("/api")
	privateAPI.Use(JWTAuthMiddleware)

	tracksAPI := privateAPI.Group("")
	tracksAPI.Post("/tracks", api.AddNewTrack)
	tracksAPI.Put("/tracks/:trackId", api.UpdateTrackByID)
	tracksAPI.Delete("/tracks/:trackId", api.DeleteTrackByID)

	userTracksAPI := privateAPI.Group("")
	userTracksAPI.Use(UserAudioMiddleware)
	userTracksAPI.Get("/user/:userId/tracks", api.GetUserTrackList)
	userTracksAPI.Post("/user/:userId/tracks", api.AddTracksToUserTrackList)
	userTracksAPI.Delete("/user/:userId/tracks", api.DeleteTracksFromUserList)

	playlistsAPI := privateAPI.Group("")
	playlistsAPI.Post("/playlists", api.CreateNewPlaylist)
	playlistsAPI.Delete("/playlists/:playlistId", api.DeletePlaylistByID)
	playlistsAPI.Post("/playlists/:playlistId/tracks", api.AddTracksToPlaylist)
	playlistsAPI.Delete("/playlists/:playlistId/tracks", api.DeleteTracksFromPlaylist)

	userPlaylistsAPI := privateAPI.Group("")
	userPlaylistsAPI.Use(UserAudioMiddleware)
	userPlaylistsAPI.Get("/user/:userId/playlists", api.GetUserPlaylists)
	userPlaylistsAPI.Post("/user/:userId/playlists", api.AddPlaylistsToUserList)
	userPlaylistsAPI.Delete("/user/:userId/playlists", api.DeletePlaylistsFromUserList)
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
