package mutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalFloat64_Create(t *testing.T) {
	opFloat64 := OptionalFloat64Create()
	assert.IsType(t, opFloat64, &OptionalFloat64{})
}

func TestOptionalFloat64_IsPresent(t *testing.T) {
	opFloat64 := OptionalFloat64Create()
	assert.False(t, opFloat64.IsPresent())
	opFloat64.ValueSet(0.12)
	assert.True(t, opFloat64.IsPresent())
}

func TestOptionalFloat64_ValueGet(t *testing.T) {
	valueExpected := 1.1
	opFloat64 := OptionalFloat64Create()
	_, err1 := opFloat64.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opFloat64.ValueSet(valueExpected)
	valueGot, err2 := opFloat64.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalFloat64_ValueSet(t *testing.T) {
	valueExpected := float64(2.3)
	opFloat64 := OptionalFloat64Create()
	opFloat64.ValueSet(valueExpected)
	valueGot, err2 := opFloat64.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
	valueExpected2 := float64(3.1)
	opFloat64.ValueSet(valueExpected2)
	valueGot2, _ := opFloat64.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalFloat64_MarshalJSON(t *testing.T) {
	opFloat64 := OptionalFloat64Create()

	jsonValueGot1, jsonValueErr1 := opFloat64.MarshalJSON()
	assert.Equal(t, jsonValueGot1, []byte("null"))
	assert.Nil(t, jsonValueErr1)

	opFloat64.ValueSet(float64(6.12))
	jsonValueGot2, jsonValueErr2 := opFloat64.MarshalJSON()
	assert.Equal(t, jsonValueGot2, []byte("6.12"))
	assert.Nil(t, jsonValueErr2)
}

func TestOptionalFloat64_UnmarshalJSON(t *testing.T) {
	opFloat64 := OptionalFloat64Create()

	err1 := opFloat64.UnmarshalJSON([]byte("asd"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opFloat64.IsPresent())

	err2 := opFloat64.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opFloat64.IsPresent())

	err3 := opFloat64.UnmarshalJSON([]byte("7.23"))
	assert.Nil(t, err3)
	assert.True(t, opFloat64.IsPresent())
	valueGot1, _ := opFloat64.ValueGet()
	assert.Equal(t, valueGot1, float64(7.23))

	err4 := opFloat64.UnmarshalJSON([]byte("8.34"))
	assert.Nil(t, err4)
	assert.True(t, opFloat64.IsPresent())
	valueGot2, _ := opFloat64.ValueGet()
	assert.Equal(t, valueGot2, float64(8.34))
}
