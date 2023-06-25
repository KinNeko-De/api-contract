package protobuf

import (
	google "github.com/google/uuid"
)

func ToUuid(protobufUuid *Uuid) (google.UUID, error) {
	var parsed, parsedError = google.Parse(protobufUuid.Value)
	return parsed, parsedError
}

func ToProtobuf(googleUuid google.UUID) (*Uuid, error) {
	parsed := &Uuid{Value: googleUuid.String()}
	return parsed, nil
}
