package helper

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type UploadTrackCallback func(id string) error

type DeleteTrackCallback func() error

var MasterManifestTemplate = "#EXTM3U\n" +
	"#EXT-X-STREAM-INF:BANDWIDTH=64000\n" +
	"64k.m3u8\n" +
	"#EXT-X-STREAM-INF:BANDWIDTH=96000\n" +
	"96k.m3u8\n" +
	"#EXT-X-STREAM-INF:BANDWIDTH=128000\n" +
	"128k.m3u8\n" +
	"#EXT-X-STREAM-INF:BANDWIDTH=192000\n" +
	"192k.m3u8\n"

func SaveMP3File(dirPath string, filename string, fileBytes []byte) error {
	err := os.MkdirAll(dirPath, 755)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.FromSlash(dirPath + "/" + filename))
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(fileBytes)
	if err != nil {
		return err
	}

	return nil
}

func TranscodeToHLS(dirPath string, inputFilePath string) error {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return err
	}

	cmd := exec.Command("ffmpeg",
		"-i", filepath.FromSlash(inputFilePath),

		"-vn", "-ac", "2", "-acodec", "libmp3lame", "-b:a", "64k", "-map", "0:a:0",
		"-hls_playlist_type", "vod", "-hls_time", "5", "-hls_segment_filename",
		filepath.FromSlash(dirPath+"/seg%02d_64k.ts"), filepath.FromSlash(dirPath+"/64k.m3u8"),

		"-vn", "-ac", "2", "-acodec", "libmp3lame", "-b:a", "96k", "-map", "0:a:0",
		"-hls_playlist_type", "vod", "-hls_time", "5", "-hls_segment_filename",
		filepath.FromSlash(dirPath+"/seg%02d_96k.ts"), filepath.FromSlash(dirPath+"/96k.m3u8"),

		"-vn", "-ac", "2", "-acodec", "libmp3lame", "-b:a", "128k", "-map", "0:a:0",
		"-hls_playlist_type", "vod", "-hls_time", "5", "-hls_segment_filename",
		filepath.FromSlash(dirPath+"/seg%02d_128k.ts"), filepath.FromSlash(dirPath+"/128k.m3u8"),

		"-vn", "-ac", "2", "-acodec", "libmp3lame", "-b:a", "192k", "-map", "0:a:0",
		"-hls_playlist_type", "vod", "-hls_time", "5", "-hls_segment_filename",
		filepath.FromSlash(dirPath+"/seg%02d_192k.ts"), filepath.FromSlash(dirPath+"/192k.m3u8"))

	err = cmd.Run()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.FromSlash(dirPath+"/master.m3u8"), []byte(MasterManifestTemplate), 0755)
	if err != nil {
		return err
	}

	return nil
}

func DeleteDir(dirPath string) error {
	err := os.RemoveAll(filepath.FromSlash(dirPath))
	if err != nil {
		return err
	}

	return nil
}
