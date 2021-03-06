package mutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalUint32 struct {
	isPresent bool
	value     uint32
}

func CreateOptionalUint32() *OptionalUint32 {
	return &OptionalUint32{
		isPresent: false,
	}
}

func (o OptionalUint32) IsPresent() bool {
	return o.isPresent
}

func (o OptionalUint32) GetValue() (uint32, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return uint32(0), optional.CreateErrorValueIsNotPresent()
}

func (o *OptionalUint32) SetValue(value uint32) {
	o.value = value
	o.isPresent = true
}

func (o OptionalUint32) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalUint32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value uint32
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	o.SetValue(value)
	return nil
}
