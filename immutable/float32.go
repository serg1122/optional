package immutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalFloat32 struct {
	isPresent bool
	value     float32
}

func CreateOptionalFloat32() *OptionalFloat32 {
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
	return 0.0, optional.CreateErrorValueIsNotPresent()
}

func (o *OptionalFloat32) SetValue(value float32) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.CreateErrorValueIsPresent()
	}
	o.value = value
	o.isPresent = true
	return nil
}

func (o OptionalFloat32) MarshalJSON() ([]byte, error) {

	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalFloat32) UnmarshalJSON(data []byte) error {

	if o.IsPresent() {
		return optional.CreateErrorValueIsPresent()
	}
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value float32
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if err2 := o.SetValue(value); err2 != nil {
		return err2
	}
	return nil
}
