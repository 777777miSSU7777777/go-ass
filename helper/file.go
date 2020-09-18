package helper

type UploadTrackCallback func(id int64) error

type DeleteTrackCallback func(id int64) error