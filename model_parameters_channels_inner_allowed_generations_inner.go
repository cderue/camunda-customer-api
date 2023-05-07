/*
Camunda Management API

Manage Camunda Clusters and API Clients via API.

API version: 1.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ParametersChannelsInnerAllowedGenerationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParametersChannelsInnerAllowedGenerationsInner{}

// ParametersChannelsInnerAllowedGenerationsInner struct for ParametersChannelsInnerAllowedGenerationsInner
type ParametersChannelsInnerAllowedGenerationsInner struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

type _ParametersChannelsInnerAllowedGenerationsInner ParametersChannelsInnerAllowedGenerationsInner

// NewParametersChannelsInnerAllowedGenerationsInner instantiates a new ParametersChannelsInnerAllowedGenerationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParametersChannelsInnerAllowedGenerationsInner(name string, uuid string) *ParametersChannelsInnerAllowedGenerationsInner {
	this := ParametersChannelsInnerAllowedGenerationsInner{}
	this.Name = name
	this.Uuid = uuid
	return &this
}

// NewParametersChannelsInnerAllowedGenerationsInnerWithDefaults instantiates a new ParametersChannelsInnerAllowedGenerationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParametersChannelsInnerAllowedGenerationsInnerWithDefaults() *ParametersChannelsInnerAllowedGenerationsInner {
	this := ParametersChannelsInnerAllowedGenerationsInner{}
	return &this
}

// GetName returns the Name field value
func (o *ParametersChannelsInnerAllowedGenerationsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParametersChannelsInnerAllowedGenerationsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ParametersChannelsInnerAllowedGenerationsInner) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value
func (o *ParametersChannelsInnerAllowedGenerationsInner) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ParametersChannelsInnerAllowedGenerationsInner) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ParametersChannelsInnerAllowedGenerationsInner) SetUuid(v string) {
	o.Uuid = v
}

func (o ParametersChannelsInnerAllowedGenerationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParametersChannelsInnerAllowedGenerationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

func (o *ParametersChannelsInnerAllowedGenerationsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"uuid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varParametersChannelsInnerAllowedGenerationsInner := _ParametersChannelsInnerAllowedGenerationsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParametersChannelsInnerAllowedGenerationsInner)

	if err != nil {
		return err
	}

	*o = ParametersChannelsInnerAllowedGenerationsInner(varParametersChannelsInnerAllowedGenerationsInner)

	return err
}

type NullableParametersChannelsInnerAllowedGenerationsInner struct {
	value *ParametersChannelsInnerAllowedGenerationsInner
	isSet bool
}

func (v NullableParametersChannelsInnerAllowedGenerationsInner) Get() *ParametersChannelsInnerAllowedGenerationsInner {
	return v.value
}

func (v *NullableParametersChannelsInnerAllowedGenerationsInner) Set(val *ParametersChannelsInnerAllowedGenerationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableParametersChannelsInnerAllowedGenerationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableParametersChannelsInnerAllowedGenerationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParametersChannelsInnerAllowedGenerationsInner(val *ParametersChannelsInnerAllowedGenerationsInner) *NullableParametersChannelsInnerAllowedGenerationsInner {
	return &NullableParametersChannelsInnerAllowedGenerationsInner{value: val, isSet: true}
}

func (v NullableParametersChannelsInnerAllowedGenerationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParametersChannelsInnerAllowedGenerationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


