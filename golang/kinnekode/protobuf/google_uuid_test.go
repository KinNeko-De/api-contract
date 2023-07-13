package protobuf

import (
	"testing"

	google "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var NotNullableTestData = []struct {
	desc     string
	input    string
	expected google.UUID
}{
	{
		desc:     "Formated well",
		input:    "ed420000-0000-0000-0000-000000000000",
		expected: google.MustParse("ed420000-0000-0000-0000-000000000000"),
	},
	{
		desc:     "No minus and capital",
		input:    "ED420000000000000000000000000000",
		expected: google.MustParse("ed420000000000000000000000000000"),
	},
	{
		desc:     "No minus and lower",
		input:    "ed420000000000000000000000000000",
		expected: google.MustParse("ed420000000000000000000000000000"),
	},
}

var invalidTestData = []struct {
	desc  string
	input string
}{
	{
		desc:  "Empty",
		input: "",
	},
	{
		desc:  "Int id",
		input: "12",
	},
}

func TestToUuid(t *testing.T) {
	for _, tt := range NotNullableTestData {
		t.Run(tt.desc, func(t *testing.T) {
			expected := tt.expected
			toParse := Uuid{Value: tt.input}
			actual, err := ToUuid(&toParse)
			assert.Nil(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestToUuid_Invalid(t *testing.T) {
	expected := google.Nil

	for _, tt := range invalidTestData {
		t.Run(tt.desc, func(t *testing.T) {
			toParse := Uuid{Value: tt.input}
			actual, err := ToUuid(&toParse)
			assert.NotNil(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestToProtobuf(t *testing.T) {
	uuidString := "cd673f27-d8ed-446c-b257-e96835895cfa"
	expected := Uuid{Value: uuidString}
	toParse := google.MustParse(uuidString)

	actual, err := ToProtobuf(toParse)
	assert.Nil(t, err)
	assert.Equal(t, &expected, actual)
}
