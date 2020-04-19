package repository

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"github.com/777777miSSU7777777/go-ass/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var AudioNotFoundError = errors.New("audio not found error")
var TableNotFoundError = errors.New("table not found error")
var UserNotFoundError = errors.New("user not found error")

type Repository struct {
	db *mongo.Database
}

func New(db *mongo.Database) Repository {
	return Repository{db}
}

func (r Repository) AddAudio(author, title string) (string, error) {
	audio := model.Audio{Author: author, Title: title}

	result, err := r.db.Collection("audio").InsertOne(context.TODO(), audio)
	if err != nil {
		return "", fmt.Errorf("add audio error: %v", err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r Repository) GetAllAudio() ([]model.Audio, error) {
	result, err := r.db.Collection("audio").Find(context.TODO(), bson.D{})
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, AudioNotFoundError
		}
		return nil, fmt.Errorf("get all audio error: %v", err)
	}

	defer result.Close(context.TODO())

	audio := []model.Audio{}
	for result.Next(context.TODO()) {
		var track model.Audio
		err = result.Decode(&track)
		if err != nil {
			return nil, fmt.Errorf("error while decoding track: %v", err)
		}
		audio = append(audio, track)
	}

	return audio, nil
}

func (r Repository) GetAudioByID(id string) (model.Audio, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Audio{}, fmt.Errorf("error while parsing audio id: %v", err)
	}

	result := r.db.Collection("audio").FindOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Audio{}, AudioNotFoundError
		}
		return model.Audio{}, fmt.Errorf("get audio by id error: %v", err)
	}

	audio := model.Audio{}
	err = result.Decode(&audio)
	if err != nil {
		return model.Audio{}, fmt.Errorf("error while decoding audio: %v", err)
	}

	return audio, nil
}

func (r Repository) GetAudioByKey(key string) ([]model.Audio, error) {
	pattern := fmt.Sprintf("^.*%s.*$", key)
	keyFilter := bson.M{
		"$or": bson.A{
			bson.D{{"author", primitive.Regex{Pattern: pattern, Options: "i"}}},
			bson.D{{"title", primitive.Regex{Pattern: pattern, Options: "i"}}},
		},
	}

	result, err := r.db.Collection("audio").Find(context.TODO(), keyFilter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, AudioNotFoundError
		}
		return nil, err
	}

	defer result.Close(context.TODO())

	audio := []model.Audio{}
	for result.Next(context.TODO()) {
		var track model.Audio
		err = result.Decode(&track)
		if err != nil {
			return nil, fmt.Errorf("error while decoding track: %v", err)
		}
		audio = append(audio, track)
	}

	if result.Err() != nil {
		return nil, fmt.Errorf("audio scan error: %v", err)
	}

	return audio, nil
}

func (r Repository) UpdateAudioByID(id, author, title string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error while parsing audio id: %v", err)
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"author": author,
			"title":  title,
		},
	}
	result, err := r.db.Collection("audio").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 && result.ModifiedCount == 0 {
		return AudioNotFoundError
	}

	return nil
}

func (r Repository) DeleteAudioByID(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error while parsing audio id: %v", err)
	}

	result, err := r.db.Collection("audio").DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return fmt.Errorf("delete audio by id error: %v", err)
	}

	if result.DeletedCount == 0 {
		return AudioNotFoundError
	}

	return nil
}

func (r Repository) AddUser(email, name, password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("error while hashing user password: %v", err)
	}

	user := model.User{Email: email, Name: name, Password: string(passwordHash)}

	result, err := r.db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		return "", fmt.Errorf("error while adding new user: %v", err)
	}
	
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r Repository) GetUserByID(id string) (model.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.User{}, fmt.Errorf("error while parsing user id: %v", err)
	}

	result := r.db.Collection("users").FindOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, UserNotFoundError
		}
		return model.User{}, fmt.Errorf("get user by id error: %v", err)
	}

	user := model.User{}
	err = result.Decode(&user)
	if err != nil {
		return model.User{}, fmt.Errorf("error while decoding user: %v", err)
	}

	return user, nil
}

func (r Repository) UpdateUserByID(id, name, email, password string) (model.User, error) {
	return model.User{}, nil	
}

func (r Repository) DeleteUserByID(id string) (model.User, error) {
	return model.User{}, nil
}
