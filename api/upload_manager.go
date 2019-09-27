package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

var FileUploadError = "FILE UPLOAD ERROR"
var FileReadError = "FILE READ ERROR"
var FileWriteError = "FILE WRITE ERROR"
var ParseFormError = "PARSE FORM ERROR"
var HLSError = "HLS ERROR"

type UploadManager struct {
	baseLocation string
}

func NewUploadManager(base string) UploadManager {
	return UploadManager{base}
}

func (m UploadManager) Upload(w http.ResponseWriter, r *http.Request, id int64) error {
	r.ParseMultipartForm(20 << 20)
	formfile, _, err := r.FormFile("audiofile")
	if err != nil {
		writeError(w, 400, FileUploadError, fmt.Errorf("error while uploading file: %v", err))
		return err
	}
	defer formfile.Close()

	fileBytes, err := ioutil.ReadAll(formfile)
	if err != nil {
		writeError(w, 400, FileReadError, fmt.Errorf("error while reading file: %v", err))
		return err
	}

	err = os.MkdirAll(fmt.Sprintf("%s/%d/mp3", m.baseLocation, id), 0755)
	if err != nil {
		writeError(w, 500, FileWriteError, fmt.Errorf("error while writing file: %v", err))
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s/%d/mp3/audio%d.mp3", m.baseLocation, id, id))
	if err != nil {
		writeError(w, 500, FileWriteError, fmt.Errorf("error while writing file: %v", err))
		return err
	}
	defer file.Close()

	_, err = file.Write(fileBytes)
	if err != nil {
		writeError(w, 500, FileWriteError, fmt.Errorf("error while writing file: %v", err))
		return err
	}

	err = m.transcodeToHLS(id)
	if err != nil {
		writeError(w, 500, HLSError, fmt.Errorf("error while transcoding to hls: %v", err))
		return err
	}

	return nil
}

func (m UploadManager) transcodeToHLS(id int64) error {
	err := os.MkdirAll(fmt.Sprintf("%s/%d/hls", m.baseLocation, id), 0755)
	if err != nil {
		return err
	}

	m3u8Dst := fmt.Sprintf("%s/%d/hls/audio%d.m3u8", m.baseLocation, id, id)
	segDst := fmt.Sprintf("%s/%d/hls/seg%s.ts", m.baseLocation, id, "%02d")
	cmd := exec.Command("ffmpeg",
		"-i", fmt.Sprintf("%s/%d/mp3/audio%d.mp3", m.baseLocation, id, id),
		"-vn", "-ac", "2",
		"-acodec", "aac",
		"-f", "segment",
		"-segment_format", "mpegts",
		"-segment_time", "10",
		"-segment_list", m3u8Dst, segDst)

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
