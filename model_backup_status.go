/*
Camunda Management API

Manage Camunda Clusters and API Clients via API.

API version: 1.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// BackupStatus the model 'BackupStatus'
type BackupStatus string

// List of BackupStatus
const (
	IN_PROGRESS BackupStatus = "In progress"
	FAILED BackupStatus = "Failed"
	COMPLETE BackupStatus = "Complete"
	MINUS BackupStatus = "-"
)

// All allowed values of BackupStatus enum
var AllowedBackupStatusEnumValues = []BackupStatus{
	"In progress",
	"Failed",
	"Complete",
	"-",
}

func (v *BackupStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackupStatus(value)
	for _, existing := range AllowedBackupStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackupStatus", value)
}

// NewBackupStatusFromValue returns a pointer to a valid BackupStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackupStatusFromValue(v string) (*BackupStatus, error) {
	ev := BackupStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackupStatus: valid values are %v", v, AllowedBackupStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackupStatus) IsValid() bool {
	for _, existing := range AllowedBackupStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupStatus value
func (v BackupStatus) Ptr() *BackupStatus {
	return &v
}

type NullableBackupStatus struct {
	value *BackupStatus
	isSet bool
}

func (v NullableBackupStatus) Get() *BackupStatus {
	return v.value
}

func (v *NullableBackupStatus) Set(val *BackupStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupStatus(val *BackupStatus) *NullableBackupStatus {
	return &NullableBackupStatus{value: val, isSet: true}
}

func (v NullableBackupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

