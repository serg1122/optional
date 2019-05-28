package immutable

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
	valueExpected := 2.3
	opFloat64 := OptionalFloat64Create()
	err1 := opFloat64.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot, err2 := opFloat64.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
	err3 := opFloat64.ValueSet(3.1)
	assert.IsType(t, err3, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opFloat64.ValueGet()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptionalFloat64_MarshalJSON(t *testing.T) {
	opFloat64 := OptionalFloat64Create()

	value1, err1 := opFloat64.MarshalJSON()
	assert.Equal(t, []byte("null"), value1)
	assert.Nil(t, err1)

	opFloat64.ValueSet(float64(1.23))
	value2, err2 := opFloat64.MarshalJSON()
	assert.Equal(t, []byte("1.23"), value2)
	assert.Nil(t, err2)
}

func TestOptionalFloat64_UnmarshalJSON(t *testing.T) {
	opFloat64 := OptionalFloat64Create()

	err1 := opFloat64.UnmarshalJSON([]byte("asd"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opFloat64.IsPresent())

	err2 := opFloat64.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opFloat64.IsPresent())

	err3 := opFloat64.UnmarshalJSON([]byte("3.45"))
	assert.Nil(t, err3)
	assert.True(t, opFloat64.IsPresent())
	valueGot1, _ := opFloat64.ValueGet()
	assert.Equal(t, float64(3.45), valueGot1)

	err4 := opFloat64.UnmarshalJSON([]byte("4.56"))
	assert.IsType(t, err4, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opFloat64.ValueGet()
	assert.Equal(t, valueGot2, float64(3.45))
}
