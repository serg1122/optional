package immutable

import (
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

func TestOptionalUint64_ValueGet(t *testing.T) {

	opUint64 := OptionalUint64Create()

	_, err1 := opUint64.ValueGet()
	assert.IsType(t, err1, optional.ErrorValueIsNotPresentCreate())

	valueExpected := uint64(2)
	opUint64.ValueSet(valueExpected)
	valueGot, _ := opUint64.ValueGet()
	assert.Equal(t, valueGot, valueExpected)
}

func TestOptionalUint64_Valueset(t *testing.T) {

	opUint64 := OptionalUint64Create()

	valueExpected := uint64(3)

	err1 := opUint64.ValueSet(valueExpected)
	assert.Nil(t, err1)
	valueGot1, _ := opUint64.ValueGet()
	assert.Equal(t, valueGot1, valueExpected)

	err2 := opUint64.ValueSet(uint64(4))
	assert.IsType(t, err2, optional.ErrorValueIsPresentCreate())
	valueGot2, _ := opUint64.ValueGet()
	assert.Equal(t, valueGot2, valueExpected)
}
