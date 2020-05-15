package recordviewings

import (
	"context"

	goes "github.com/jetbasrawi/go.geteventstore"

	vidtuts "github.com/dkimot/video-tutorials"
	"github.com/dkimot/video-tutorials/pkg"
)

type actions struct {
	es goes.Client
}

func newActions(es goes.Client) *actions {
	return &actions{es: es}
}

func (a *actions) recordViewing(ctx context.Context, videoID string) error {

	viewedEv, err := vidtuts.NewVideoViewed(ctx, videoID)
	if err != nil {
		return err
	}

	streamName := "viewing-" + viewedEv.VideoID
	writer := a.es.NewStreamWriter(streamName)

	return a.msgStore.Write(streamName, viewedEv.Message)
}
