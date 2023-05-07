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

// checks if the ClusterGeneration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterGeneration{}

// ClusterGeneration The version of Camunda running on your cluster.
type ClusterGeneration struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

type _ClusterGeneration ClusterGeneration

// NewClusterGeneration instantiates a new ClusterGeneration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterGeneration(name string, uuid string) *ClusterGeneration {
	this := ClusterGeneration{}
	this.Name = name
	this.Uuid = uuid
	return &this
}

// NewClusterGenerationWithDefaults instantiates a new ClusterGeneration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterGenerationWithDefaults() *ClusterGeneration {
	this := ClusterGeneration{}
	return &this
}

// GetName returns the Name field value
func (o *ClusterGeneration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterGeneration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterGeneration) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value
func (o *ClusterGeneration) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ClusterGeneration) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ClusterGeneration) SetUuid(v string) {
	o.Uuid = v
}

func (o ClusterGeneration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterGeneration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

func (o *ClusterGeneration) UnmarshalJSON(data []byte) (err error) {
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

	varClusterGeneration := _ClusterGeneration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClusterGeneration)

	if err != nil {
		return err
	}

	*o = ClusterGeneration(varClusterGeneration)

	return err
}

type NullableClusterGeneration struct {
	value *ClusterGeneration
	isSet bool
}

func (v NullableClusterGeneration) Get() *ClusterGeneration {
	return v.value
}

func (v *NullableClusterGeneration) Set(val *ClusterGeneration) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterGeneration) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterGeneration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterGeneration(val *ClusterGeneration) *NullableClusterGeneration {
	return &NullableClusterGeneration{value: val, isSet: true}
}

func (v NullableClusterGeneration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterGeneration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


