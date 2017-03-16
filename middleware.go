import (
	"context"

	"github.com/chihaya/chihaya/bittorrent"
)

// START OMIT
func (l *Logic) HandleAnnounce(ctx context.Context, req *bittorrent.AnnounceRequest) (resp *bittorrent.AnnounceResponse, err error) {
	resp = &bittorrent.AnnounceResponse{}
	for _, h := range l.preHooks {
		if ctx, err = h.HandleAnnounce(ctx, req, resp); err != nil {
			return nil, err
// END OMIT
