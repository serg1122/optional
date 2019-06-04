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

func OptionalFloat32Create() *OptionalFloat32 {
	return &OptionalFloat32{
		isPresent: false,
	}
}

func (o OptionalFloat32) IsPresent() bool {
	return o.isPresent
}

func (o OptionalFloat32) ValueGet() (float32, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return 0.0, optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalFloat32) ValueSet(value float32) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
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
		return optional.ErrorValueIsPresentCreate()
	}
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value float32
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if err2 := o.ValueSet(value); err2 != nil {
		return err2
	}
	return nil
}
