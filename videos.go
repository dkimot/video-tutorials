package vidtuts

import (
	"context"

	"github.com/dkimot/video-tutorials/pkg"
)

type VideoViewed struct {
	Message pkg.Message
	VideoID string
}

func NewVideoViewed(ctx context.Context, videoID string) (*VideoViewed, error) {
	payload := map[string]string{
		"videoID": videoID,
	}
	msgType := (*VideoViewed)(nil)
	uMsg, err := pkg.NewUserMessage(ctx, msgType, nil, payload)
	if err != nil {
		return nil, err
	}

	return &VideoViewed{
		Message: uMsg,
		VideoID: videoID,
	}, nil
}

func (v *VideoViewed) Type() string {
	return "video_viewed"
}
