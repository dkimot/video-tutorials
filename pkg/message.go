package pkg

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type ctxKey int

const (
	CtxKeyTraceID = ctxKey(0)
	CtxKeyUserID  = ctxKey(1)
)

var (
	ErrNoUserIDFound = errors.New("no user id found for that context")
)

type Message struct {
	ID       string
	Type     string
	Metadata map[string]string
	Payload  map[string]string
}

type TypedMessage interface {
	Type() string
}

func NewMessage(ctx context.Context, typedMsg TypedMessage, meta, payload map[string]string) Message {
	traceMeta := map[string]string{
		"traceID": ctx.Value(CtxKeyTraceID).(string),
	}

	return Message{
		ID:       uuid.New().String(),
		Type:     typedMsg.Type(),
		Metadata: mergeStrMap(traceMeta, meta),
		Payload:  payload,
	}
}

func NewUserMessage(ctx context.Context, typedMsg TypedMessage, meta, payload map[string]string) (Message, error) {
	userID, ok := ctx.Value(CtxKeyUserID).(string)
	if !ok {
		return Message{}, ErrNoUserIDFound
	}

	userMeta := map[string]string{
		"userID": userID,
	}

	return NewMessage(ctx, typedMsg, mergeStrMap(userMeta, meta), payload), nil
}

func mergeStrMap(a, b map[string]string) map[string]string {
	c := a
	for k, v := range b {
		c[k] = v
	}
	return c
}
