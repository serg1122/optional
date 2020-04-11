package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalInt_(t *testing.T) {
	opInt := CreateOptionalInt()
	assert.IsType(t, opInt, &OptionalInt{})
}

func TestOptionalInt_IsPresent(t *testing.T) {
	opInt := CreateOptionalInt()
	assert.False(t, opInt.IsPresent())
	opInt.SetValue(123)
	assert.True(t, opInt.IsPresent())
}

func TestOptionalInt_GetValue(t *testing.T) {
	valueExpected := 345
	opInt := CreateOptionalInt()
	_, err1 := opInt.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())
	opInt.SetValue(valueExpected)
	valueGot, err2 := opInt.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptinoalInt_SetValue(t *testing.T) {
	valueExpected := 567
	opInt := CreateOptionalInt()
	err1 := opInt.SetValue(valueExpected)
	assert.Nil(t, err1)
	valueGot, _ := opInt.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	err3 := opInt.SetValue(789)
	assert.IsType(t, err3, optional.CreateErrorValueIsPresent())
}

func TestOptinalInt_MarshalJSON(t *testing.T) {
	opInt := CreateOptionalInt()

	valueGot1, err1 := opInt.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opInt.SetValue(int(5))
	valueGot2, err2 := opInt.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalInt_UnmarshalJSON(t *testing.T) {
	opInt := CreateOptionalInt()

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
	assert.Equal(t, err4, optional.CreateErrorValueIsPresent())
	assert.Equal(t, valueGot2, valueExpected)
}
