package immutable

import (
	"bytes"
	"encoding/json"

	"github.com/serg1122/optional"
)

type OptionalFloat64 struct {
	isPresent bool
	value     float64
}

func CreateOptionalFloat64() *OptionalFloat64 {
	return &OptionalFloat64{
		isPresent: false,
	}
}

func (o OptionalFloat64) IsPresent() bool {
	return o.isPresent
}

func (o OptionalFloat64) GetValue() (float64, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return 0.0, optional.CreateErrorValueIsNotPresent()
}

func (o *OptionalFloat64) SetValue(value float64) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.CreateErrorValueIsPresent()
	}
	o.value = value
	o.isPresent = true
	return nil
}

func (o OptionalFloat64) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalFloat64) UnmarshalJSON(data []byte) error {
	if o.IsPresent() {
		return optional.CreateErrorValueIsPresent()
	}
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value float64
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	if err2 := o.SetValue(value); err2 != nil {
		return err2
	}
	return nil
}
