package stream

import (
	"fmt"
	"net/http"
)

type MediaManager struct {
	baseLocation string
}

func NewMediaManager(base string) MediaManager {
	return MediaManager{base}
}

func (m MediaManager) ServeM3u8(w http.ResponseWriter, r *http.Request, id int64) {
	m3u8File := fmt.Sprintf("%s/%d/hls/audio%d.m3u8", m.baseLocation, id, id)
	http.ServeFile(w, r, m3u8File)
	w.Header().Set("Content-Type", "application/x-mpegURL")
}

func (m MediaManager) ServeTs(w http.ResponseWriter, r *http.Request, id int64, seg string) {
	tsFile := fmt.Sprintf("%s/%d/hls/%s", m.baseLocation, id, seg)
	http.ServeFile(w, r, tsFile)
	w.Header().Set("Content-Type", "video/MP2T")
}

func (m MediaManager) ServerMp3(w http.ResponseWriter, r *http.Request, id int64) {
	mp3File := fmt.Sprintf("%s/%d/mp3/audio%d.mp3", m.baseLocation, id, id)
	http.ServeFile(w, r, mp3File)
	w.Header().Set("Content-Type", "audio/mpeg")
}
