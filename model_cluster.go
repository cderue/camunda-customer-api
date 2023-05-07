/*
Camunda Management API

Manage Camunda Clusters and API Clients via API.

API version: 1.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the Cluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cluster{}

// Cluster Describing a Camunda cluster running in your organization. For reference, use the UUID.
type Cluster struct {
	// The ID used in all further API operations referencing your cluster.
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	OwnerId string `json:"ownerId"`
	Created time.Time `json:"created"`
	PlanType ClusterPlanType `json:"planType"`
	Region ClusterRegion `json:"region"`
	BackupRegion *ClusterBackupRegion `json:"backupRegion,omitempty"`
	Generation ClusterGeneration `json:"generation"`
	Channel ClusterChannel `json:"channel"`
	// The cluster will automatically be updated by Camunda when a new C8 version is released If the cluster you want to create is running on a non-stable Channel passing this property will be ignored!
	AutoUpdate bool `json:"autoUpdate"`
	// DEPRECATED: this field is going to be removed in June 2025, please use ipallowlist instead  the IP Whitelist rules for your cluster - will only be returned if your organization has the feature enabled and the client you are using has the permission to see it.
	Ipwhitelist []ClusterIpwhitelistInner `json:"ipwhitelist,omitempty"`
	// the IP Allowlist rules for your cluster - will only be returned if your organization has the feature enabled and the client you are using has the permission to see it.
	Ipallowlist []ClusterIpwhitelistInner `json:"ipallowlist,omitempty"`
	Status ClusterStatus `json:"status"`
	Links ClusterLinks `json:"links"`
	Labels []string `json:"labels,omitempty"`
}

type _Cluster Cluster

// NewCluster instantiates a new Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster(uuid string, name string, ownerId string, created time.Time, planType ClusterPlanType, region ClusterRegion, generation ClusterGeneration, channel ClusterChannel, autoUpdate bool, status ClusterStatus, links ClusterLinks) *Cluster {
	this := Cluster{}
	this.Uuid = uuid
	this.Name = name
	this.OwnerId = ownerId
	this.Created = created
	this.PlanType = planType
	this.Region = region
	this.Generation = generation
	this.Channel = channel
	this.AutoUpdate = autoUpdate
	this.Status = status
	this.Links = links
	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *Cluster) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *Cluster) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *Cluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Cluster) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value
func (o *Cluster) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *Cluster) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetCreated returns the Created field value
func (o *Cluster) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Cluster) SetCreated(v time.Time) {
	o.Created = v
}

// GetPlanType returns the PlanType field value
func (o *Cluster) GetPlanType() ClusterPlanType {
	if o == nil {
		var ret ClusterPlanType
		return ret
	}

	return o.PlanType
}

// GetPlanTypeOk returns a tuple with the PlanType field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetPlanTypeOk() (*ClusterPlanType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanType, true
}

// SetPlanType sets field value
func (o *Cluster) SetPlanType(v ClusterPlanType) {
	o.PlanType = v
}

// GetRegion returns the Region field value
func (o *Cluster) GetRegion() ClusterRegion {
	if o == nil {
		var ret ClusterRegion
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetRegionOk() (*ClusterRegion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *Cluster) SetRegion(v ClusterRegion) {
	o.Region = v
}

// GetBackupRegion returns the BackupRegion field value if set, zero value otherwise.
func (o *Cluster) GetBackupRegion() ClusterBackupRegion {
	if o == nil || IsNil(o.BackupRegion) {
		var ret ClusterBackupRegion
		return ret
	}
	return *o.BackupRegion
}

// GetBackupRegionOk returns a tuple with the BackupRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetBackupRegionOk() (*ClusterBackupRegion, bool) {
	if o == nil || IsNil(o.BackupRegion) {
		return nil, false
	}
	return o.BackupRegion, true
}

// HasBackupRegion returns a boolean if a field has been set.
func (o *Cluster) HasBackupRegion() bool {
	if o != nil && !IsNil(o.BackupRegion) {
		return true
	}

	return false
}

// SetBackupRegion gets a reference to the given ClusterBackupRegion and assigns it to the BackupRegion field.
func (o *Cluster) SetBackupRegion(v ClusterBackupRegion) {
	o.BackupRegion = &v
}

// GetGeneration returns the Generation field value
func (o *Cluster) GetGeneration() ClusterGeneration {
	if o == nil {
		var ret ClusterGeneration
		return ret
	}

	return o.Generation
}

// GetGenerationOk returns a tuple with the Generation field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetGenerationOk() (*ClusterGeneration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Generation, true
}

// SetGeneration sets field value
func (o *Cluster) SetGeneration(v ClusterGeneration) {
	o.Generation = v
}

// GetChannel returns the Channel field value
func (o *Cluster) GetChannel() ClusterChannel {
	if o == nil {
		var ret ClusterChannel
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetChannelOk() (*ClusterChannel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *Cluster) SetChannel(v ClusterChannel) {
	o.Channel = v
}

// GetAutoUpdate returns the AutoUpdate field value
func (o *Cluster) GetAutoUpdate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoUpdate
}

// GetAutoUpdateOk returns a tuple with the AutoUpdate field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetAutoUpdateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoUpdate, true
}

// SetAutoUpdate sets field value
func (o *Cluster) SetAutoUpdate(v bool) {
	o.AutoUpdate = v
}

// GetIpwhitelist returns the Ipwhitelist field value if set, zero value otherwise.
func (o *Cluster) GetIpwhitelist() []ClusterIpwhitelistInner {
	if o == nil || IsNil(o.Ipwhitelist) {
		var ret []ClusterIpwhitelistInner
		return ret
	}
	return o.Ipwhitelist
}

// GetIpwhitelistOk returns a tuple with the Ipwhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetIpwhitelistOk() ([]ClusterIpwhitelistInner, bool) {
	if o == nil || IsNil(o.Ipwhitelist) {
		return nil, false
	}
	return o.Ipwhitelist, true
}

// HasIpwhitelist returns a boolean if a field has been set.
func (o *Cluster) HasIpwhitelist() bool {
	if o != nil && !IsNil(o.Ipwhitelist) {
		return true
	}

	return false
}

// SetIpwhitelist gets a reference to the given []ClusterIpwhitelistInner and assigns it to the Ipwhitelist field.
func (o *Cluster) SetIpwhitelist(v []ClusterIpwhitelistInner) {
	o.Ipwhitelist = v
}

// GetIpallowlist returns the Ipallowlist field value if set, zero value otherwise.
func (o *Cluster) GetIpallowlist() []ClusterIpwhitelistInner {
	if o == nil || IsNil(o.Ipallowlist) {
		var ret []ClusterIpwhitelistInner
		return ret
	}
	return o.Ipallowlist
}

// GetIpallowlistOk returns a tuple with the Ipallowlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetIpallowlistOk() ([]ClusterIpwhitelistInner, bool) {
	if o == nil || IsNil(o.Ipallowlist) {
		return nil, false
	}
	return o.Ipallowlist, true
}

// HasIpallowlist returns a boolean if a field has been set.
func (o *Cluster) HasIpallowlist() bool {
	if o != nil && !IsNil(o.Ipallowlist) {
		return true
	}

	return false
}

// SetIpallowlist gets a reference to the given []ClusterIpwhitelistInner and assigns it to the Ipallowlist field.
func (o *Cluster) SetIpallowlist(v []ClusterIpwhitelistInner) {
	o.Ipallowlist = v
}

// GetStatus returns the Status field value
func (o *Cluster) GetStatus() ClusterStatus {
	if o == nil {
		var ret ClusterStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetStatusOk() (*ClusterStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Cluster) SetStatus(v ClusterStatus) {
	o.Status = v
}

// GetLinks returns the Links field value
func (o *Cluster) GetLinks() ClusterLinks {
	if o == nil {
		var ret ClusterLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetLinksOk() (*ClusterLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *Cluster) SetLinks(v ClusterLinks) {
	o.Links = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Cluster) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Cluster) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *Cluster) SetLabels(v []string) {
	o.Labels = v
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["name"] = o.Name
	toSerialize["ownerId"] = o.OwnerId
	toSerialize["created"] = o.Created
	toSerialize["planType"] = o.PlanType
	toSerialize["region"] = o.Region
	if !IsNil(o.BackupRegion) {
		toSerialize["backupRegion"] = o.BackupRegion
	}
	toSerialize["generation"] = o.Generation
	toSerialize["channel"] = o.Channel
	toSerialize["autoUpdate"] = o.AutoUpdate
	if !IsNil(o.Ipwhitelist) {
		toSerialize["ipwhitelist"] = o.Ipwhitelist
	}
	if !IsNil(o.Ipallowlist) {
		toSerialize["ipallowlist"] = o.Ipallowlist
	}
	toSerialize["status"] = o.Status
	toSerialize["links"] = o.Links
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

func (o *Cluster) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
		"name",
		"ownerId",
		"created",
		"planType",
		"region",
		"generation",
		"channel",
		"autoUpdate",
		"status",
		"links",
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

	varCluster := _Cluster{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCluster)

	if err != nil {
		return err
	}

	*o = Cluster(varCluster)

	return err
}

type NullableCluster struct {
	value *Cluster
	isSet bool
}

func (v NullableCluster) Get() *Cluster {
	return v.value
}

func (v *NullableCluster) Set(val *Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCluster(val *Cluster) *NullableCluster {
	return &NullableCluster{value: val, isSet: true}
}

func (v NullableCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


