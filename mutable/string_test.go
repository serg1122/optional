package mutable

import (
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalString_(t *testing.T) {
	opStr := CreateOptionalString()
	assert.IsType(t, opStr, &OptionalString{})
}

func TestOptionalString_IsPresent(t *testing.T) {
	opStr := CreateOptionalString()
	assert.False(t, opStr.IsPresent())
	opStr.SetValue("asd")
	assert.True(t, opStr.IsPresent())
}

func TestOptionalString_GetValue(t *testing.T) {
	valueExpected := "GetValue expected value"
	opStr := CreateOptionalString()
	_, err1 := opStr.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())
	opStr.SetValue(valueExpected)
	valueGot, err2 := opStr.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalString_SetValue(t *testing.T) {
	valueExpected := "SetValue expected value"
	opStr := CreateOptionalString()
	opStr.SetValue(valueExpected)
	valueGot1, _ := opStr.GetValue()
	assert.Equal(t, valueGot1, valueExpected)
	valueExpected2 := "asd"
	opStr.SetValue(valueExpected2)
	valueGot2, _ := opStr.GetValue()
	assert.Equal(t, valueGot2, valueExpected2)
}

func TestOptionalString_MarshalJSON(t *testing.T) {
	opString := CreateOptionalString()

	bytesGot1, errGot1 := opString.MarshalJSON()
	assert.False(t, opString.IsPresent())
	assert.Equal(t, bytesGot1, []byte("null"))
	assert.Nil(t, errGot1)

	valueExpected := "megastring"
	opString.SetValue(valueExpected)
	bytesGot2, errGot2 := opString.MarshalJSON()
	assert.True(t, opString.IsPresent())
	assert.Equal(t, bytesGot2, []byte(`"`+valueExpected+`"`))
	assert.Nil(t, errGot2)
}

func TestOptionalString_UnmarshalJSON(t *testing.T) {
	opString := CreateOptionalString()

	err1 := opString.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err1)
	assert.False(t, opString.IsPresent())

	valueExpected1 := "Наш президент - вор"
	err2 := opString.UnmarshalJSON([]byte(`"` + valueExpected1 + `"`))
	assert.Nil(t, err2)
	assert.True(t, opString.IsPresent())
	valueGot1, _ := opString.GetValue()
	assert.Equal(t, valueGot1, valueExpected1)

	valueExpected2 := "megastring-123"
	err3 := opString.UnmarshalJSON([]byte(`"` + valueExpected2 + `"`))
	assert.Nil(t, err3)
	assert.True(t, opString.IsPresent())
	valueGot2, _ := opString.GetValue()
	assert.Equal(t, valueGot2, valueExpected2)
}
