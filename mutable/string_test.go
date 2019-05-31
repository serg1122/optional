package mutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalString_Create(t *testing.T) {
	opStr := OptionalStringCreate()
	assert.IsType(t, opStr, &OptionalString{})
}

func TestOptionalString_IsPresent(t *testing.T) {
	opStr := OptionalStringCreate()
	assert.False(t, opStr.IsPresnt())
	opStr.ValueSet("asd")
	assert.True(t, opStr.IsPresnt())
}

func TestOptionalString_ValueGet(t *testing.T) {
	valueExpected := "ValueGet expected value"
	opStr := OptionalStringCreate()
	_, err1 := opStr.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())
	opStr.ValueSet(valueExpected)
	valueGot, err2 := opStr.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalString_ValueSet(t *testing.T) {
	valueExpected := "ValueSet expected value"
	opStr := OptionalStringCreate()
	opStr.ValueSet(valueExpected)
	valueGot1, _ := opStr.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)
	valueExpected2 := "asd"
	opStr.ValueSet(valueExpected2)
	valueGot2, _ := opStr.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalString_MarshalJSON(t *testing.T) {
	opString := OptionalStringCreate()

	bytesGot1, errGot1 := opString.MarshalJSON()
	assert.False(t, opString.IsPresnt())
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	valueExpected := "megastring"
	opString.ValueSet(valueExpected)
	bytesGot2, errGot2 := opString.MarshalJSON()
	assert.True(t, opString.IsPresnt())
	assert.Equal(t, bytesGot2, []byte(`"`+valueExpected+`"`))
	assert.Nil(t, errGot2)
}

func TestOptionalString_UnmarshalJSON(t *testing.T) {
	opString := OptionalStringCreate()

	err1 := opString.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err1)
	assert.False(t, opString.IsPresnt())

	valueExpected1 := "megastring-1"
	err2 := opString.UnmarshalJSON([]byte(`"` + valueExpected1 + `"`))
	assert.Nil(t, err2)
	assert.True(t, opString.IsPresnt())
	valueGot1, _ := opString.ValueGet()
	assert.Equal(t, valueGot1, valueExpected1)

	valueExpected2 := "megastring-2"
	err3 := opString.UnmarshalJSON([]byte(`"` + valueExpected2 + `"`))
	assert.Nil(t, err3)
	assert.True(t, opString.IsPresnt())
	valueGot2, _ := opString.ValueGet()
	assert.Equal(t, valueGot2, valueExpected2)
}
