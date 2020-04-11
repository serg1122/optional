package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalInt64_(t *testing.T) {
	opInt64 := CreateOptionalInt64()
	assert.IsType(t, opInt64, &OptionalInt64{})
}

func TestOptionalInt64_IsPresent(t *testing.T) {
	opInt64 := CreateOptionalInt64()
	assert.False(t, opInt64.IsPresent())
	opInt64.SetValue(int64(1))
	assert.True(t, opInt64.IsPresent())
}

func TestOptionalInt64_GetValue(t *testing.T) {
	valueExpected := int64(2)
	opInt64 := CreateOptionalInt64()
	_, err1 := opInt64.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())
	opInt64.SetValue(valueExpected)
	valueGot, err2 := opInt64.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalInt64_SetValue(t *testing.T) {
	valueExpected := int64(4)
	opInt64 := CreateOptionalInt64()
	err1 := opInt64.SetValue(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opInt64.GetValue()
	assert.Equal(t, valueGot1, valueExpected)
	err2 := opInt64.SetValue(int64(5))
	assert.IsType(t, err2, optional.CreateErrorValueIsPresent())
	valueGot2, _ := opInt64.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptinalInt64_MarshalJSON(t *testing.T) {
	opInt64 := CreateOptionalInt64()

	valueGot1, err1 := opInt64.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opInt64.SetValue(int64(5))
	valueGot2, err2 := opInt64.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalInt64_UnmarshalJSON(t *testing.T) {
	opInt64 := CreateOptionalInt64()

	err1 := opInt64.UnmarshalJSON([]byte("qwe"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opInt64.IsPresent())

	err2 := opInt64.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opInt64.IsPresent())

	valueExpected := int64(6)

	err3 := opInt64.UnmarshalJSON([]byte("6"))
	valueGot1, _ := opInt64.GetValue()
	assert.True(t, opInt64.IsPresent())
	assert.Equal(t, valueGot1, valueExpected)
	assert.Nil(t, err3)

	err4 := opInt64.UnmarshalJSON([]byte("6"))
	valueGot2, _ := opInt64.GetValue()
	assert.Equal(t, err4, optional.CreateErrorValueIsPresent())
	assert.Equal(t, valueGot2, valueExpected)
}
