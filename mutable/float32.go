package mutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalFloat32 struct {
	isPresent bool
	value     float32
}

func OptionalFloat32Create() *OptionalFloat32 {
	return &OptionalFloat32{
		isPresent: false,
	}
}

func (o OptionalFloat32) IsPresent() bool {
	return o.isPresent
}

func (o OptionalFloat32) GetValue() (float32, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return 0.0, optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalFloat32) ValueSet(value float32) {
	o.value = value
	o.isPresent = true
}

func (o OptionalFloat32) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalFloat32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value float32
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	o.ValueSet(value)
	return nil
}
