package mutable

import (
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalUint32 struct {
	isPresent bool
	value     uint32
}

func OptionalUint32Create() *OptionalUint32 {
	return &OptionalUint32{
		isPresent: false,
	}
}

func (o *OptionalUint32) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalUint32) ValueGet() (uint32, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return uint32(0), optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalUint32) ValueSet(value uint32) {
	o.value = value
	o.isPresent = true
}

func (o *OptionalUint32) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalUint32) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var value uint32
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	o.ValueSet(value)
	return nil
}
