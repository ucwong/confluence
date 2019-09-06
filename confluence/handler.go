package confluence

import (
	"net/http"
	"sync"
	"time"

	"github.com/anacrolix/torrent"
)

type Handler struct {
	TC             *torrent.Client
	TorrentGrace   time.Duration
	OnTorrentGrace func(t *torrent.Torrent)
	mux            http.ServeMux
	initMuxOnce    sync.Once
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.initMux()
	h.mux.ServeHTTP(w, r)
}
