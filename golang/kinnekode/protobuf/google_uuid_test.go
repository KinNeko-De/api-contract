package protobuf

import (
	google "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToUuid(t *testing.T) {
	expected := google.New()
	toParse := Uuid{Value: expected.String()}
	actual, err := ToUuid(&toParse)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestToUuid_UuidNotSet(t *testing.T) {
	expected := google.Nil
	toParse := Uuid{Value: ""}
	actual, err := ToUuid(&toParse)
	assert.NotNil(t, err)
	assert.Equal(t, expected, actual)
}

func TestToProtobuf(t *testing.T) {
	uuidString := "cd673f27-d8ed-446c-b257-e96835895cfa"
	expected := Uuid{Value: uuidString}
	toParse := google.MustParse(uuidString)

	actual, err := ToProtobuf(toParse)
	assert.Nil(t, err)
	assert.Same(t, &expected, actual)
}
