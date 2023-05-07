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

// checks if the IpAllowListBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpAllowListBody{}

// IpAllowListBody struct for IpAllowListBody
type IpAllowListBody struct {
	Ipallowlist []ClusterIpwhitelistInner `json:"ipallowlist"`
}

type _IpAllowListBody IpAllowListBody

// NewIpAllowListBody instantiates a new IpAllowListBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpAllowListBody(ipallowlist []ClusterIpwhitelistInner) *IpAllowListBody {
	this := IpAllowListBody{}
	this.Ipallowlist = ipallowlist
	return &this
}

// NewIpAllowListBodyWithDefaults instantiates a new IpAllowListBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpAllowListBodyWithDefaults() *IpAllowListBody {
	this := IpAllowListBody{}
	return &this
}

// GetIpallowlist returns the Ipallowlist field value
func (o *IpAllowListBody) GetIpallowlist() []ClusterIpwhitelistInner {
	if o == nil {
		var ret []ClusterIpwhitelistInner
		return ret
	}

	return o.Ipallowlist
}

// GetIpallowlistOk returns a tuple with the Ipallowlist field value
// and a boolean to check if the value has been set.
func (o *IpAllowListBody) GetIpallowlistOk() ([]ClusterIpwhitelistInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipallowlist, true
}

// SetIpallowlist sets field value
func (o *IpAllowListBody) SetIpallowlist(v []ClusterIpwhitelistInner) {
	o.Ipallowlist = v
}

func (o IpAllowListBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpAllowListBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ipallowlist"] = o.Ipallowlist
	return toSerialize, nil
}

func (o *IpAllowListBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ipallowlist",
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

	varIpAllowListBody := _IpAllowListBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIpAllowListBody)

	if err != nil {
		return err
	}

	*o = IpAllowListBody(varIpAllowListBody)

	return err
}

type NullableIpAllowListBody struct {
	value *IpAllowListBody
	isSet bool
}

func (v NullableIpAllowListBody) Get() *IpAllowListBody {
	return v.value
}

func (v *NullableIpAllowListBody) Set(val *IpAllowListBody) {
	v.value = val
	v.isSet = true
}

func (v NullableIpAllowListBody) IsSet() bool {
	return v.isSet
}

func (v *NullableIpAllowListBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpAllowListBody(val *IpAllowListBody) *NullableIpAllowListBody {
	return &NullableIpAllowListBody{value: val, isSet: true}
}

func (v NullableIpAllowListBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpAllowListBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


