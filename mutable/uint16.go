package mutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalUint16 struct {
	isPresent bool
	value     uint16
}

func OptionalUint16Create() *OptionalUint16 {
	return &OptionalUint16{
		isPresent: false,
	}
}

func (o *OptionalUint16) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalUint16) ValueGet() (uint16, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return 0, optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalUint16) ValueSet(value uint16) {
	o.value = value
	o.isPresent = true
}

func (o *OptionalUint16) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalUint16) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value uint16
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	o.ValueSet(value)
	return nil
}
