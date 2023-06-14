package protobuf

import (
	"errors"
	google "github.com/google/uuid"
)

func ToUuid(protobufUuid *Uuid) (google.UUID, error) {
	var parsed, parsedError = google.Parse(protobufUuid.Value)
	return parsed, parsedError
}

func ToProtobuf(googleUuid google.UUID) (*Uuid, error) {
	// TODO is this possible?
	if googleUuid.String() == "" {
		return &Uuid{}, errors.New("uuid is invalid")
	}
	parsed := &Uuid{Value: googleUuid.String()}
	return parsed, nil
}
