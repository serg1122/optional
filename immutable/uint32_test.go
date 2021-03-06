package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalUint32_(t *testing.T) {
	opUint32 := CreateOptionalUint32()
	assert.IsType(t, opUint32, &OptionalUint32{})
}

func TestOptionalUint32_IsPresent(t *testing.T) {
	opUint32 := CreateOptionalUint32()
	assert.False(t, opUint32.IsPresent())
	opUint32.SetValue(uint32(1))
	assert.True(t, opUint32.IsPresent())
}

func TestOptionalUint32_GetValue(t *testing.T) {
	opUint32 := CreateOptionalUint32()

	_, err1 := opUint32.GetValue()
	assert.IsType(t, err1, optional.CreateErrorValueIsNotPresent())

	valueExpected := uint32(2)
	opUint32.SetValue(valueExpected)
	valueGot, err2 := opUint32.GetValue()
	assert.Equal(t, valueGot, valueExpected)
	assert.Nil(t, err2)
}

func TestOptionalUint32_SetValue(t *testing.T) {
	opUint32 := CreateOptionalUint32()

	valueExpected := uint32(3)
	err1 := opUint32.SetValue(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opUint32.GetValue()
	assert.Equal(t, valueGot1, valueExpected)

	err2 := opUint32.SetValue(uint32(4))
	assert.IsType(t, err2, optional.CreateErrorValueIsPresent())
	valueGot2, _ := opUint32.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptinalUint32_MarshalJSON(t *testing.T) {
	opUint32 := CreateOptionalUint32()

	valueGot1, err1 := opUint32.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opUint32.SetValue(uint32(5))
	valueGot2, err2 := opUint32.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalUint32_UnmarshalJSON(t *testing.T) {
	opUint32 := CreateOptionalUint32()

	err1 := opUint32.UnmarshalJSON([]byte("qwe"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opUint32.IsPresent())

	err2 := opUint32.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opUint32.IsPresent())

	valueExpected := uint32(6)

	err3 := opUint32.UnmarshalJSON([]byte("6"))
	valueGot1, _ := opUint32.GetValue()
	assert.True(t, opUint32.IsPresent())
	assert.Equal(t, valueGot1, valueExpected)
	assert.Nil(t, err3)

	err4 := opUint32.UnmarshalJSON([]byte("6"))
	valueGot2, _ := opUint32.GetValue()
	assert.Equal(t, err4, optional.CreateErrorValueIsPresent())
	assert.Equal(t, valueGot2, valueExpected)
}
