import (
	"context"

	"github.com/chihaya/chihaya/bittorrent"
)

// START OMIT
type Hook interface {
	HandleAnnounce(context.Context, *bittorrent.AnnounceRequest, *bittorrent.AnnounceResponse) (context.Context, error)
	HandleScrape(context.Context, *bittorrent.ScrapeRequest, *bittorrent.ScrapeResponse) (context.Context, error)
}

// END OMIT
