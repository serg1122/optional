package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalInt64_Create(t *testing.T) {
	opInt64 := OptionalInt64Create()
	assert.IsType(t, opInt64, &OptionalInt64{})
}

func TestOptionalInt64_IsPresent(t *testing.T) {
	opInt64 := OptionalInt64Create()
	assert.False(t, opInt64.IsPresent())
	opInt64.ValueSet(int64(1))
	assert.True(t, opInt64.IsPresent())
}

func TestOptionalInt64_ValueGet(t *testing.T) {
	valueExpected := int64(2)
	opInt64 := OptionalInt64Create()
	_, err1 := opInt64.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opInt64.ValueSet(valueExpected)
	valueGot, err2 := opInt64.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalInt64_ValueSet(t *testing.T) {
	valueExpected := int64(4)
	opInt64 := OptionalInt64Create()
	err1 := opInt64.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opInt64.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)
	err2 := opInt64.ValueSet(int64(5))
	assert.IsType(t, err2, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opInt64.ValueGet()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptinalInt64_MarshalJSON(t *testing.T) {
	opInt64 := OptionalInt64Create()

	valueGot1, err1 := opInt64.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opInt64.ValueSet(int64(5))
	valueGot2, err2 := opInt64.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalInt64_UnmarshalJSON(t *testing.T) {
	opInt64 := OptionalInt64Create()

	err1 := opInt64.UnmarshalJSON([]byte("qwe"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opInt64.IsPresent())

	err2 := opInt64.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opInt64.IsPresent())

	valueExpected := int64(6)

	err3 := opInt64.UnmarshalJSON([]byte("6"))
	valueGot1, _ := opInt64.ValueGet()
	assert.True(t, opInt64.IsPresent())
	assert.Equal(t, valueGot1, valueExpected)
	assert.Nil(t, err3)

	err4 := opInt64.UnmarshalJSON([]byte("6"))
	valueGot2, _ := opInt64.ValueGet()
	assert.Equal(t, err4, optional.ErrorValueIsPresentCreate())
	assert.Equal(t, valueGot2, valueExpected)
}
