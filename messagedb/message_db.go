package messagedb

import (
	"encoding/json"

	"github.com/go-pg/pg/v9"

	"github.com/dkimot/video-tutorials/pkg"
)

const writeFunctionSQL = `SELECT message_store.write_message($1, $2, $3, $4, $5, $6)`

type Options struct {
	Addr string
	Pass string
	Port string
	User string
}

func New(opts Options) pkg.MessageStore {
	conn := pg.Connect(&pg.Options{
		Addr:     opts.Addr,
		Password: opts.Pass,
		User:     opts.User,
	})

	return &messageDB{
		db: conn,
	}
}

type messageDB struct {
	db *pg.DB
}

func (msgDB *messageDB) Write(streamName string, msg pkg.Message, expVersion string) error {
	payloadStr, err := json.Marshal(msg.Payload)
	if err != nil {
		return err
	}

	metadataStr, err := json.Marshal(msg.Payload)
	if err != nil {
		return err
	}

	values := []string{
		msg.ID,
		streamName,
		msg.Type,
		string(payloadStr),
		string(metadataStr),
	}

	_, err = msgDB.db.Exec(writeFunctionSQL, values)
	return err
}

func (msgDB *messageDB) Close() error {
	return msgDB.db.Close()
}
