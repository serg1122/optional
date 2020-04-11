package immutable

import (
	"encoding/json"
	"testing"

	"github.com/serg1122/optional"
	"github.com/stretchr/testify/assert"
)

func TestOptionalUint64_Create(t *testing.T) {
	opUint64 := OptionalUint64Create()
	assert.IsType(t, opUint64, &OptionalUint64{})
}

func TestOptionalUint64_IsPresent(t *testing.T) {
	opUint64 := OptionalUint64Create()
	assert.False(t, opUint64.IsPresent())
	opUint64.ValueSet(uint64(1))
	assert.True(t, opUint64.IsPresent())
}

func TestOptionalUint64_GetValue(t *testing.T) {

	opUint64 := OptionalUint64Create()

	_, err1 := opUint64.GetValue()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())

	valueExpected := uint64(2)
	opUint64.ValueSet(valueExpected)
	valueGot, _ := opUint64.GetValue()
	assert.Equal(t, valueGot, valueExpected)
}

func TestOptionalUint64_Valueset(t *testing.T) {

	opUint64 := OptionalUint64Create()

	valueExpected := uint64(3)

	err1 := opUint64.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opUint64.GetValue()
	assert.Equal(t, valueGot1, valueExpected)

	err2 := opUint64.ValueSet(uint64(4))
	assert.IsType(t, err2, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opUint64.GetValue()
	assert.Equal(t, valueGot2, valueExpected)
}

func TestOptinalUint64_MarshalJSON(t *testing.T) {
	opUint64 := OptionalUint64Create()

	valueGot1, err1 := opUint64.MarshalJSON()
	assert.Equal(t, []byte("null"), valueGot1)
	assert.Nil(t, err1)

	opUint64.ValueSet(uint64(5))
	valueGot2, err2 := opUint64.MarshalJSON()
	assert.Equal(t, valueGot2, []byte("5"))
	assert.Nil(t, err2)
}

func TestOptionalUint64_UnmarshalJSON(t *testing.T) {
	opUint64 := OptionalUint64Create()

	err1 := opUint64.UnmarshalJSON([]byte("qwe"))
	assert.IsType(t, err1, &json.SyntaxError{})
	assert.False(t, opUint64.IsPresent())

	err2 := opUint64.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err2)
	assert.False(t, opUint64.IsPresent())

	valueExpected := uint64(6)

	err3 := opUint64.UnmarshalJSON([]byte("6"))
	valueGot1, _ := opUint64.GetValue()
	assert.True(t, opUint64.IsPresent())
	assert.Equal(t, valueGot1, valueExpected)
	assert.Nil(t, err3)

	err4 := opUint64.UnmarshalJSON([]byte("6"))
	valueGot2, _ := opUint64.GetValue()
	assert.Equal(t, err4, optional.ErrorValueIsPresentCreate())
	assert.Equal(t, valueGot2, valueExpected)
}
