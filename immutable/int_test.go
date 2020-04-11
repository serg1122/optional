package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalInt_Create(t *testing.T) {
	opInt := OptionalIntCreate()
	assert.IsType(t, opInt, &OptionalInt{})
}

func TestOptionalInt_IsPresent(t *testing.T) {
	opInt := OptionalIntCreate()
	assert.False(t, opInt.IsPresent())
	opInt.ValueSet(123)
	assert.True(t, opInt.IsPresent())
}

func TestOptionalInt_GetValue(t *testing.T) {
	valueExpected := 345
	opInt := OptionalIntCreate()
	_, err1 := opInt.GetValue()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opInt.ValueSet(valueExpected)
	valueGot, err2 := opInt.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptinoalInt_ValueSet(t *testing.T) {
	valueExpected := 567
	opInt := OptionalIntCreate()
	err1 := opInt.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot, _ := opInt.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	err3 := opInt.ValueSet(789)
	assert.IsType(t, err3, optional.ErrorValueIsPresentCreate())
}

func TestOptinalInt_MarshalJSON(t *testing.T) {
	opInt := OptionalIntCreate()

	valueGot1, err1 := opInt.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opInt.ValueSet(int(5))
	valueGot2, err2 := opInt.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalInt_UnmarshalJSON(t *testing.T) {
	opInt := OptionalIntCreate()

	err1 := opInt.UnmarshalJSON([]byte("qwe"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opInt.IsPresent())

	err2 := opInt.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opInt.IsPresent())

	valueExpected := int(6)

	err3 := opInt.UnmarshalJSON([]byte("6"))
	valueGot1, _ := opInt.GetValue()
	assert.True(t, opInt.IsPresent())
	assert.Equal(t, valueGot1, valueExpected)
	assert.Nil(t, err3)

	err4 := opInt.UnmarshalJSON([]byte("6"))
	valueGot2, _ := opInt.GetValue()
	assert.Equal(t, err4, optional.ErrorValueIsPresentCreate())
	assert.Equal(t, valueGot2, valueExpected)
}
