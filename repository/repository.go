package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/777777miSSU7777777/go-ass/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var TrackNotFoundError = errors.New("track not found error")
var TableNotFoundError = errors.New("table not found error")
var UserNotFoundError = errors.New("user not found error")
var RefreshTokenNotFoundError = errors.New("refresh token not found error")
var PlaylistNotFoundError = errors.New("playlist not found error")

var SecretKey = "NOT A SECRET KEY"

type Repository struct {
	db *mongo.Database
}

func New(db *mongo.Database) Repository {
	return Repository{db}
}

type JWTPayload struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func (r Repository) AddTrack(author, title, uploadedByID string) (string, error) {
	uploadedByObjectID, err := primitive.ObjectIDFromHex(uploadedByID)
	if err != nil {
		return "", fmt.Errorf("add track error: %v", err)
	}
	track := model.Track{Author: author, Title: title, UploadedByID: uploadedByObjectID}

	addTrackResult, err := r.db.Collection("tracks").InsertOne(context.TODO(), track)
	if err != nil {
		return "", fmt.Errorf("add track error: %v", err)
	}

	_, err = r.db.Collection("users").UpdateOne(context.TODO(), bson.M{"_id": uploadedByObjectID}, bson.M{"$push": bson.M{"tracklist": addTrackResult.InsertedID}})
	if err != nil {
		return "", fmt.Errorf("add track error: %v", err)
	}

	return addTrackResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r Repository) GetAllTracks() ([]model.Track, error) {
	getAllTracksResult, err := r.db.Collection("tracks").Find(context.TODO(), bson.D{})
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, TrackNotFoundError
		}
		return nil, fmt.Errorf("get all tracks error: %v", err)
	}

	defer getAllTracksResult.Close(context.TODO())

	allTracks := []model.Track{}
	for getAllTracksResult.Next(context.TODO()) {
		var track model.Track
		err = getAllTracksResult.Decode(&track)
		if err != nil {
			return nil, fmt.Errorf("get all tracks error: %v", err)
		}
		allTracks = append(allTracks, track)
	}

	return allTracks, nil
}

func (r Repository) GetTrackByID(trackID string) (model.Track, error) {
	trackObjectID, err := primitive.ObjectIDFromHex(trackID)
	if err != nil {
		return model.Track{}, fmt.Errorf("get track by id error: %v", err)
	}

	getTrackByIDResult := r.db.Collection("tracks").FindOne(context.TODO(), bson.M{"_id": trackObjectID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Track{}, TrackNotFoundError
		}
		return model.Track{}, fmt.Errorf("get track by id error: %v", err)
	}

	track := model.Track{}
	err = getTrackByIDResult.Decode(&track)
	if err != nil {
		return model.Track{}, fmt.Errorf("get track by id error: %v", err)
	}

	return track, nil
}

func (r Repository) GetTracksByKey(key string) ([]model.Track, error) {
	pattern := fmt.Sprintf("^.*%s.*$", key)
	keyFilter := bson.M{
		"$or": bson.A{
			bson.D{{"author", primitive.Regex{Pattern: pattern, Options: "i"}}},
			bson.D{{"title", primitive.Regex{Pattern: pattern, Options: "i"}}},
		},
	}

	getTrackByKeyResult, err := r.db.Collection("tracks").Find(context.TODO(), keyFilter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, TrackNotFoundError
		}
		return nil, fmt.Errorf("get tracks by key error: %v", err)
	}

	defer getTrackByKeyResult.Close(context.TODO())

	tracksByKey := []model.Track{}
	for getTrackByKeyResult.Next(context.TODO()) {
		var track model.Track
		err = getTrackByKeyResult.Decode(&track)
		if err != nil {
			return nil, fmt.Errorf("get tracks by key error: %v", err)
		}
		tracksByKey = append(tracksByKey, track)
	}

	if getTrackByKeyResult.Err() != nil {
		return nil, fmt.Errorf("get tracks by key error: %v", err)
	}

	return tracksByKey, nil
}

func (r Repository) UpdateTrackByID(trackID, author, title string) error {
	trackObjectID, err := primitive.ObjectIDFromHex(trackID)
	if err != nil {
		return fmt.Errorf("update track by id error: %v", err)
	}

	filter := bson.M{"_id": trackObjectID}
	update := bson.M{
		"$set": bson.M{
			"author": author,
			"title":  title,
		},
	}
	updateTrackByIDResult, err := r.db.Collection("tracks").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if updateTrackByIDResult.MatchedCount == 0 && updateTrackByIDResult.ModifiedCount == 0 {
		return TrackNotFoundError
	}

	return nil
}

func (r Repository) DeleteTrackByID(trackID string) error {
	trackObjectID, err := primitive.ObjectIDFromHex(trackID)
	if err != nil {
		return fmt.Errorf("delete track by id error: %v", err)
	}

	deleteTrackByIDResult, err := r.db.Collection("tracks").DeleteOne(context.TODO(), bson.M{"_id": trackObjectID})
	if err != nil {
		return fmt.Errorf("delete track by id error: %v", err)
	}

	if deleteTrackByIDResult.DeletedCount == 0 {
		return TrackNotFoundError
	}

	return nil
}

func (r Repository) AddUser(email, name, password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("add user error: %v", err)
	}

	user := model.User{Email: email, Name: name, Password: string(passwordHash), RefreshTokens: []string{}, TrackList: []primitive.ObjectID{}, Playlists: []primitive.ObjectID{}}

	addUserResult, err := r.db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		return "", fmt.Errorf("add user error: %v", err)
	}

	return addUserResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r Repository) GetUserByID(userID string) (model.User, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return model.User{}, fmt.Errorf("get user by id error: %v", err)
	}

	getUserByIDResult := r.db.Collection("users").FindOne(context.TODO(), bson.M{"_id": userObjectID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, UserNotFoundError
		}
		return model.User{}, fmt.Errorf("get user by id error: %v", err)
	}

	user := model.User{}
	err = getUserByIDResult.Decode(&user)
	if err != nil {
		return model.User{}, fmt.Errorf("get user by id error: %v", err)
	}

	return user, nil
}

func (r Repository) GetUserByEmail(email string) (model.User, error) {
	getUserByEmailResult := r.db.Collection("users").FindOne(context.TODO(), bson.M{"email": email})

	err := getUserByEmailResult.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, UserNotFoundError
		}
		return model.User{}, fmt.Errorf("get user by email error: %v", err)
	}

	var user model.User
	err = getUserByEmailResult.Decode(&user)
	if err != nil {
		return model.User{}, fmt.Errorf("get user by email error: %v", err)
	}

	return user, nil
}

func (r Repository) UpdateUserByID(id, name, email, password string) (model.User, error) {
	return model.User{}, nil
}

func (r Repository) DeleteUserByID(id string) (model.User, error) {
	return model.User{}, nil
}

func (r Repository) AddRefreshToken(id string) (string, error) {
	customClaims := JWTPayload{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(5184000)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", fmt.Errorf("add refresh token error: %v", err)
	}

	userObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", fmt.Errorf("add refresh token error: %v", err)
	}

	addRefreshTokenResult, err := r.db.Collection("users").UpdateOne(context.TODO(), bson.M{"_id": userObjectID}, bson.M{"$push": bson.M{"refresh_tokens": signedToken}})

	if err != nil {
		return "", fmt.Errorf("add refresh token error: %v", err)
	}
	if addRefreshTokenResult.MatchedCount == 0 || addRefreshTokenResult.ModifiedCount == 0 {
		return "", RefreshTokenNotFoundError
	}

	return signedToken, nil
}

func (r Repository) UpdateRefreshToken(token string) (string, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return "", fmt.Errorf("update refresh token error: %v", err)
	}

	payload := parsedToken.Claims.(*JWTPayload)

	customClaims := JWTPayload{
		payload.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(5184000)).Unix(),
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	signedToken, err := newToken.SignedString([]byte(SecretKey))
	if err != nil {
		return "", fmt.Errorf("update refresh token error: %v", err)
	}

	updateRefreshTokenResult, err := r.db.Collection("users").UpdateOne(context.TODO(), bson.M{"refresh_tokens": bson.M{"$elemMatch": bson.M{"$eq": token}}}, bson.M{"$set": bson.M{"refresh_tokens.$": signedToken}})

	if err != nil {
		return "", fmt.Errorf("update refresh token error: %v", err)
	}

	if updateRefreshTokenResult.MatchedCount == 0 || updateRefreshTokenResult.ModifiedCount == 0 {
		return "", RefreshTokenNotFoundError
	}

	return signedToken, nil
}

func (r Repository) DeleteRefreshToken(token string) error {
	deleteRefreshTokenResult, err := r.db.Collection("users").UpdateOne(context.TODO(), bson.M{"refresh_tokens": bson.M{"$elemMatch": bson.M{"$eq": token}}}, bson.M{"$pull": bson.M{"refresh_tokens": token}})
	if err != nil {
		return fmt.Errorf("delete refresh token error: %v", err)
	}

	if deleteRefreshTokenResult.MatchedCount == 0 || deleteRefreshTokenResult.ModifiedCount == 0 {
		return RefreshTokenNotFoundError
	}

	return nil
}

func (r Repository) GetUserTrackList(userID string) ([]model.Track, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("get user track list error: %v", err)
	}

	userResult := r.db.Collection("users").FindOne(context.TODO(), bson.M{"_id": userObjectID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, UserNotFoundError
		}
		return nil, fmt.Errorf("get user track list error: %v", err)
	}

	user := model.User{}
	err = userResult.Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("get user track list error: %v", err)
	}

	getUsertrackListResult, err := r.db.Collection("tracks").Find(context.TODO(), bson.M{"_id": bson.M{"$in": user.TrackList}})
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, TrackNotFoundError
		}
		return nil, fmt.Errorf("get user track list error: %v", err)
	}

	defer getUsertrackListResult.Close(context.TODO())

	userTrackList := []model.Track{}
	for getUsertrackListResult.Next(context.TODO()) {
		var track model.Track
		err = getUsertrackListResult.Decode(&track)
		if err != nil {
			return nil, fmt.Errorf("get user track list error: %v", err)
		}
		userTrackList = append(userTrackList, track)
	}

	return userTrackList, nil
}

func (r Repository) AddTrackToUserTrackList(userID, trackID string) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("add track to user track list error: %v", err)
	}

	trackObjectID, err := primitive.ObjectIDFromHex(trackID)
	if err != nil {
		return fmt.Errorf("add track to user track list error: %v", err)
	}

	_, err = r.db.Collection("users").UpdateOne(context.TODO(), bson.M{"_id": userObjectID}, bson.M{"$push": bson.M{"tracklist": trackObjectID}})
	if err != nil {
		return fmt.Errorf("add track to user track list error: %v", err)
	}

	return nil
}

func (r Repository) RemoveTrackFromUserTrackList(userID, trackID string) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("remove track from user track list error: %v", err)
	}

	trackObjectID, err := primitive.ObjectIDFromHex(trackID)
	if err != nil {
		return fmt.Errorf("remove track from user track list error: %v", err)
	}

	_, err = r.db.Collection("users").UpdateOne(context.TODO(), bson.M{"_id": userObjectID}, bson.M{"$pull": bson.M{"tracklist": trackObjectID}})
	if err != nil {
		return fmt.Errorf("remove track from user track list error: %v", err)
	}

	return nil
}

func (r Repository) CreateNewPlaylist(title, createdByID string, trackList []string) (string, error) {
	createdByObjectID, err := primitive.ObjectIDFromHex(createdByID)
	if err != nil {
		return "", fmt.Errorf("create new playlist error: %v", err)
	}

	trackObjectList := []primitive.ObjectID{}
	for _, trackID := range trackList {
		trackOjectID, err := primitive.ObjectIDFromHex(trackID)
		if err != nil {
			return "", fmt.Errorf("create new playlist error: %v", err)
		}
		trackObjectList = append(trackObjectList, trackOjectID)
	}

	newPlaylist := model.Playlist{Title: title, TrackList: trackObjectList, CreatedByID: createdByObjectID}

	createNewPlaylistResult, err := r.db.Collection("playlists").InsertOne(context.TODO(), newPlaylist)
	if err != nil {
		return "", fmt.Errorf("create new playlist error: %v", err)
	}

	_, err = r.db.Collection("users").UpdateOne(context.TODO(), bson.M{"_id": createdByObjectID}, bson.M{"$push": bson.M{"playlists": createNewPlaylistResult.InsertedID}})
	if err != nil {
		return "", fmt.Errorf("create new playlist error: %v", err)
	}

	return createNewPlaylistResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r Repository) GetAllPlaylists() ([]model.Playlist, [][]model.Track, error) {
	getAllPlaylistsResult, err := r.db.Collection("playlists").Find(context.TODO(), bson.D{})
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, nil, PlaylistNotFoundError
		}
		return nil, nil, fmt.Errorf("get all playlists error: %v", err)
	}

	defer getAllPlaylistsResult.Close(context.TODO())

	playlists := []model.Playlist{}
	for getAllPlaylistsResult.Next(context.TODO()) {
		var playlist model.Playlist
		err = getAllPlaylistsResult.Decode(&playlist)
		if err != nil {
			return nil, nil, fmt.Errorf("get all playlists error: %v", err)
		}
		playlists = append(playlists, playlist)
	}

	playlistsTracks := [][]model.Track{}
	for _, playlist := range playlists {
		playlistTracksResult, err := r.db.Collection("tracks").Find(context.TODO(), bson.M{"_id": bson.M{"$in": playlist.TrackList}})
		if err != nil {
			return nil, nil, fmt.Errorf("get all playlists error: %v", err)
		}

		defer playlistTracksResult.Close(context.TODO())

		playlistTracks := []model.Track{}
		for playlistTracksResult.Next(context.TODO()) {
			var track model.Track
			err = playlistTracksResult.Decode(&track)
			if err != nil {
				return nil, nil, fmt.Errorf("get all playlists error: %v", err)
			}
			playlistTracks = append(playlistTracks, track)
		}

		playlistsTracks = append(playlistsTracks, playlistTracks)
	}

	return playlists, playlistsTracks, nil
}

func (r Repository) GetPlaylistByID(playlistID string) (model.Playlist, []model.Track, error) {
	playlistObjectID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		return model.Playlist{}, nil, fmt.Errorf("get playlist by id error: %v", err)
	}

	getPlaylistByIDResult := r.db.Collection("playlists").FindOne(context.TODO(), bson.M{"_id": playlistObjectID})
	err = getPlaylistByIDResult.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Playlist{}, nil, PlaylistNotFoundError
		}
		return model.Playlist{}, nil, fmt.Errorf("get playlist by id error: %v", err)
	}

	var playlist model.Playlist
	err = getPlaylistByIDResult.Decode(&playlist)
	if err != nil {
		return model.Playlist{}, nil, fmt.Errorf("get playlist by id error: %v", err)
	}

	playlistTracksResult, err := r.db.Collection("tracks").Find(context.TODO(), bson.M{"_id": bson.M{"$in": playlist.TrackList}})
	if err != nil {
		return model.Playlist{}, nil, fmt.Errorf("get playlist by id error: %v", err)
	}

	defer playlistTracksResult.Close(context.TODO())

	playlistTracks := []model.Track{}
	for playlistTracksResult.Next(context.TODO()) {
		var track model.Track
		err = playlistTracksResult.Decode(&track)
		if err != nil {
			return model.Playlist{}, nil, fmt.Errorf("get playlist by id error: %v", err)
		}
		playlistTracks = append(playlistTracks, track)
	}

	return playlist, playlistTracks, nil
}

func (r Repository) DeletePlaylistByID(playlistID, createdByID string) error {
	playlistObjectID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		return fmt.Errorf("delete playlist by id error: %v", err)
	}

	createdByObjectID, err := primitive.ObjectIDFromHex(createdByID)
	if err != nil {
		return fmt.Errorf("delete playlist by id error: %v", err)
	}

	deletePlaylistByIDResult, err := r.db.Collection("playlists").DeleteOne(context.TODO(), bson.M{"_id": playlistObjectID, "createdByID": createdByObjectID})

	if err != nil {
		return fmt.Errorf("delete playlist by id error: %v", err)
	}

	if deletePlaylistByIDResult.DeletedCount == 0 {
		return PlaylistNotFoundError
	}

	return nil
}

func (r Repository) GetUserPlaylists(userID string) ([]model.Playlist, [][]model.Track, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, nil, fmt.Errorf("get user playlists error: %v", err)
	}

	getUserPlaylists, err := r.db.Collection("playlists").Find(context.TODO(), bson.M{"createdByID": userObjectID})
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, nil, PlaylistNotFoundError
		}
		return nil, nil, fmt.Errorf("get user playlists error: %v", err)
	}

	defer getUserPlaylists.Close(context.TODO())

	playlists := []model.Playlist{}
	for getUserPlaylists.Next(context.TODO()) {
		var playlist model.Playlist
		err = getUserPlaylists.Decode(&playlist)
		if err != nil {
			return nil, nil, fmt.Errorf("get user playlists error: %v", err)
		}
		playlists = append(playlists, playlist)
	}

	playlistsTracks := [][]model.Track{}
	for _, playlist := range playlists {
		playlistTracksResult, err := r.db.Collection("tracks").Find(context.TODO(), bson.M{"_id": bson.M{"$in": playlist.TrackList}})
		if err != nil {
			return nil, nil, fmt.Errorf("get user playlists error: %v", err)
		}

		defer playlistTracksResult.Close(context.TODO())

		playlistTracks := []model.Track{}
		for playlistTracksResult.Next(context.TODO()) {
			var track model.Track
			err = playlistTracksResult.Decode(&track)
			if err != nil {
				return nil, nil, fmt.Errorf("get user playlists error: %v", err)
			}
			playlistTracks = append(playlistTracks, track)
		}

		playlistsTracks = append(playlistsTracks, playlistTracks)
	}

	return playlists, playlistsTracks, nil
}

func (r Repository) AddTracksToPlaylist(userID, playlistID string, trackList []string) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("add tracks to playlist error: %v", err)
	}

	playlistObjectID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		return fmt.Errorf("add tracks to playlist error: %v", err)
	}

	trackObjectList := []primitive.ObjectID{}
	for _, trackID := range trackList {
		trackObjectID, err := primitive.ObjectIDFromHex(trackID)
		if err != nil {
			return fmt.Errorf("add tracks to playlist error: %v", err)
		}
		trackObjectList = append(trackObjectList, trackObjectID)
	}

	_, err = r.db.Collection("playlists").UpdateOne(context.TODO(), bson.M{"_id": playlistObjectID, "createdByID": userObjectID}, bson.M{"$push": bson.M{"tracklist": bson.M{"$each": trackObjectList}}})
	if err != nil {
		return fmt.Errorf("add tracks to playlist error: %v", err)
	}

	return nil
}

func (r Repository) RemoveTracksFromPlaylist(userID, playlistID string, trackList []string) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf(": %v", err)
	}

	playlistObjectID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		return fmt.Errorf("remove tracks from playlist error: %v", err)
	}

	trackObjectList := []primitive.ObjectID{}
	for _, trackID := range trackList {
		trackObjectID, err := primitive.ObjectIDFromHex(trackID)
		if err != nil {
			return fmt.Errorf("remove tracks from playlist error: %v", err)
		}
		trackObjectList = append(trackObjectList, trackObjectID)
	}

	_, err = r.db.Collection("playlists").UpdateOne(context.TODO(), bson.M{"_id": playlistObjectID, "createdByID": userObjectID}, bson.M{"$pull": bson.M{"tracklist": bson.M{"$in": trackObjectList}}})
	if err != nil {
		return fmt.Errorf("remove tracks from playlist error: %v", err)
	}

	return nil
}
