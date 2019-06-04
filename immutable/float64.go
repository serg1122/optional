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

func OptionalFloat64Create() *OptionalFloat64 {
	return &OptionalFloat64{
		isPresent: false,
	}
}

func (o *OptionalFloat64) IsPresent() bool {
	return o.isPresent
}

func (o *OptionalFloat64) ValueGet() (float64, *optional.ErrorValueIsNotPresent) {
	if o.IsPresent() {
		return o.value, nil
	}
	return 0.0, optional.ErrorValueIsNotPresentCreate()
}

func (o *OptionalFloat64) ValueSet(value float64) *optional.ErrorValueIsPresent {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	o.value = value
	o.isPresent = true
	return nil
}

func (o *OptionalFloat64) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *OptionalFloat64) UnmarshalJSON(data []byte) error {
	if o.IsPresent() {
		return optional.ErrorValueIsPresentCreate()
	}
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	var value float64
	if err1 := json.Unmarshal(data, &value); err1 != nil {
		return err1
	}
	if err2 := o.ValueSet(value); err2 != nil {
		return err2
	}
	return nil
}
