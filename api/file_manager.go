package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"mime/multipart"
)

var FileUploadError = "FILE UPLOAD ERROR"
var FileReadError = "FILE READ ERROR"
var FileWriteError = "FILE WRITE ERROR"
var ParseFormError = "PARSE FORM ERROR"
var HLSError = "HLS ERROR"
var DeleteAudioError = "DELETE AUDIO ERROR"

type FileManager struct {
	baseLocation               string
	mainManifestTemplateString string
}

func NewFileManager(base string) FileManager {
	mainManifestTemplateString := "#EXTM3U\n" +
		"#EXT-X-STREAM-INF:BANDWIDTH=64000\n" +
		"%s_64k.m3u8\n" +
		"#EXT-X-STREAM-INF:BANDWIDTH-96000\n" +
		"%s_96k.m3u8\n" +
		"#EXT-X-STREAM-INF:BANDWIDTH-128000\n" +
		"%s_128k.m3u8\n" +
		"#EXT-X-STREAM-INF:BANDWIDTH=192000\n" +
		"%s_192k.m3u8\n"

	return FileManager{base, mainManifestTemplateString}
}

func (manager FileManager) Upload(fh multipart.FileHeader, id string) error {
	formFile, err := fh.Open()
	if err != nil {
		return err
	}

	defer formFile.Close()

	fileBytes, err := ioutil.ReadAll(formFile)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(fmt.Sprintf("%s/%d/mp3", manager.baseLocation, id), 0755); err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s/%s/mp3/%s.mp3", manager.baseLocation, id, id))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(fileBytes)
	if err != nil {
		return err
	}

	if err := manager.transcodeToHLS(id); err != nil {
		return err
	}

	return nil
}

func (manager FileManager) transcodeToHLS(id string) error {
	err := os.MkdirAll(fmt.Sprintf("%s/%s/hls", manager.baseLocation, id), 0755)
	if err != nil {
		return err
	}

	cmd := exec.Command("ffmpeg",
		"-i", fmt.Sprintf("%s/%s/mp3/%s.mp3", manager.baseLocation, id, id),

		"-vn", "-ac", "2", "-acodec", "libmp3lame", "-b:a", "64k", "-map", "0:a:0",
		"-hls_playlist_type", "vod", "-hls_time", "5", "-hls_segment_filename",
		fmt.Sprintf("%s/%s/hls/seg%s_64k.ts", manager.baseLocation, id, "%02d"), fmt.Sprintf("%s/%s/hls/%s_64k.m3u8", manager.baseLocation, id, id),

		"-vn", "-ac", "2", "-acodec", "libmp3lame", "-b:a", "96k", "-map", "0:a:0",
		"-hls_playlist_type", "vod", "-hls_time", "5", "-hls_segment_filename",
		fmt.Sprintf("%s/%s/hls/seg%s_96k.ts", manager.baseLocation, id, "%02d"), fmt.Sprintf("%s/%s/hls/%s_96k.m3u8", manager.baseLocation, id, id),

		"-vn", "-ac", "2", "-acodec", "libmp3lame", "-b:a", "128k", "-map", "0:a:0",
		"-hls_playlist_type", "vod", "-hls_time", "5", "-hls_segment_filename",
		fmt.Sprintf("%s/%s/hls/seg%s_128k.ts", manager.baseLocation, id, "%02d"), fmt.Sprintf("%s/%s/hls/%s_128k.m3u8", manager.baseLocation, id, id),

		"-vn", "-ac", "2", "-acodec", "libmp3lame", "-b:a", "192k", "-map", "0:a:0",
		"-hls_playlist_type", "vod", "-hls_time", "5", "-hls_segment_filename",
		fmt.Sprintf("%s/%s/hls/seg%s_192k.ts", manager.baseLocation, id, "%02d"), fmt.Sprintf("%s/%s/hls/%s_192k.m3u8", manager.baseLocation, id, id))

	err = cmd.Run()
	if err != nil {
		return err
	}

	mainManifestContent := []byte(fmt.Sprintf(manager.mainManifestTemplateString, id, id, id, id))
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/hls/%s.m3u8", manager.baseLocation, id, id), mainManifestContent, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (manager FileManager) Delete(id string) error {
	err := os.RemoveAll(fmt.Sprintf("%s/%s", manager.baseLocation, id))
	if err != nil {
		return err
	}

	return nil
}
