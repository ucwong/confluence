package confluence

import (
	"net/http"
	"time"

	"github.com/anacrolix/missinggo/refclose"
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
)

var (
	torrentClientContextKey     = new(byte)
	torrentContextKey           = new(byte)
	torrentCloseGraceContextKey = new(byte)
	torrentRefs                 refclose.RefPool
)

func torrentClientForRequest(r *http.Request) *torrent.Client {
	return r.Context().Value(torrentClientContextKey).(*torrent.Client)
}

func torrentForRequest(r *http.Request) *torrent.Torrent {
	ih := r.Context().Value(torrentContextKey).(*refclose.Ref).Key().(metainfo.Hash)
	t, ok := torrentClientForRequest(r).Torrent(ih)
	if !ok {
		panic(ih)
	}
	return t
}

func torrentCloseGraceForRequest(r *http.Request) time.Duration {
	return r.Context().Value(torrentCloseGraceContextKey).(time.Duration)
}