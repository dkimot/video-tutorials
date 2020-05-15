package recordviewings

import (
	"context"

	vidtuts "github.com/dkimot/video-tutorials"
	"github.com/dkimot/video-tutorials/pkg"
)

type actions struct {
	msgStore pkg.MessageStore
}

func newActions(msgStore pkg.MessageStore) *actions {
	return &actions{msgStore: msgStore}
}

func (a *actions) recordViewing(ctx context.Context, videoID string) error {

	viewedEv, err := vidtuts.NewVideoViewed(ctx, videoID)
	if err != nil {
		return err
	}

	streamName := "viewing-" + viewedEv.VideoID

	return a.msgStore.Write(streamName, viewedEv.Message, nil)
}
