/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionossdk

import (
	"encoding/json"
)

// SnapshotProperties struct for SnapshotProperties
type SnapshotProperties struct {
	// A name of that resource
	Name *string `json:"name,omitempty"`
	// Human readable description
	Description *string `json:"description,omitempty"`
	// Location of that image/snapshot.
	Location *string `json:"location,omitempty"`
	// The size of the image in GB
	Size *float32 `json:"size,omitempty"`
	// Boolean value representing if the snapshot requires extra protection e.g. two factor protection
	SecAuthProtection *bool `json:"secAuthProtection,omitempty"`
	// Is capable of CPU hot plug (no reboot required)
	CpuHotPlug *bool `json:"cpuHotPlug,omitempty"`
	// Is capable of CPU hot unplug (no reboot required)
	CpuHotUnplug *bool `json:"cpuHotUnplug,omitempty"`
	// Is capable of memory hot plug (no reboot required)
	RamHotPlug *bool `json:"ramHotPlug,omitempty"`
	// Is capable of memory hot unplug (no reboot required)
	RamHotUnplug *bool `json:"ramHotUnplug,omitempty"`
	// Is capable of nic hot plug (no reboot required)
	NicHotPlug *bool `json:"nicHotPlug,omitempty"`
	// Is capable of nic hot unplug (no reboot required)
	NicHotUnplug *bool `json:"nicHotUnplug,omitempty"`
	// Is capable of Virt-IO drive hot plug (no reboot required)
	DiscVirtioHotPlug *bool `json:"discVirtioHotPlug,omitempty"`
	// Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.
	DiscVirtioHotUnplug *bool `json:"discVirtioHotUnplug,omitempty"`
	// Is capable of SCSI drive hot plug (no reboot required)
	DiscScsiHotPlug *bool `json:"discScsiHotPlug,omitempty"`
	// Is capable of SCSI drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.
	DiscScsiHotUnplug *bool `json:"discScsiHotUnplug,omitempty"`
	// OS type of this Snapshot
	LicenceType *string `json:"licenceType,omitempty"`
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SnapshotProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *SnapshotProperties) SetName(v string) {
	o.Name = &v
}

// HasName returns a boolean if a field has been set.
func (o *SnapshotProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SnapshotProperties) GetDescription() *string {
	if o == nil {
		return nil
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description, true
}

// SetDescription sets field value
func (o *SnapshotProperties) SetDescription(v string) {
	o.Description = &v
}

// HasDescription returns a boolean if a field has been set.
func (o *SnapshotProperties) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// GetLocation returns the Location field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SnapshotProperties) GetLocation() *string {
	if o == nil {
		return nil
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location, true
}

// SetLocation sets field value
func (o *SnapshotProperties) SetLocation(v string) {
	o.Location = &v
}

// HasLocation returns a boolean if a field has been set.
func (o *SnapshotProperties) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// GetSize returns the Size field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *SnapshotProperties) GetSize() *float32 {
	if o == nil {
		return nil
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size, true
}

// SetSize sets field value
func (o *SnapshotProperties) SetSize(v float32) {
	o.Size = &v
}

// HasSize returns a boolean if a field has been set.
func (o *SnapshotProperties) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// GetSecAuthProtection returns the SecAuthProtection field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *SnapshotProperties) GetSecAuthProtection() *bool {
	if o == nil {
		return nil
	}

	return o.SecAuthProtection
}

// GetSecAuthProtectionOk returns a tuple with the SecAuthProtection field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetSecAuthProtectionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecAuthProtection, true
}

// SetSecAuthProtection sets field value
func (o *SnapshotProperties) SetSecAuthProtection(v bool) {
	o.SecAuthProtection = &v
}

// HasSecAuthProtection returns a boolean if a field has been set.
func (o *SnapshotProperties) HasSecAuthProtection() bool {
	if o != nil && o.SecAuthProtection != nil {
		return true
	}

	return false
}

// GetCpuHotPlug returns the CpuHotPlug field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *SnapshotProperties) GetCpuHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.CpuHotPlug
}

// GetCpuHotPlugOk returns a tuple with the CpuHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetCpuHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuHotPlug, true
}

// SetCpuHotPlug sets field value
func (o *SnapshotProperties) SetCpuHotPlug(v bool) {
	o.CpuHotPlug = &v
}

// HasCpuHotPlug returns a boolean if a field has been set.
func (o *SnapshotProperties) HasCpuHotPlug() bool {
	if o != nil && o.CpuHotPlug != nil {
		return true
	}

	return false
}

// GetCpuHotUnplug returns the CpuHotUnplug field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *SnapshotProperties) GetCpuHotUnplug() *bool {
	if o == nil {
		return nil
	}

	return o.CpuHotUnplug
}

// GetCpuHotUnplugOk returns a tuple with the CpuHotUnplug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetCpuHotUnplugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuHotUnplug, true
}

// SetCpuHotUnplug sets field value
func (o *SnapshotProperties) SetCpuHotUnplug(v bool) {
	o.CpuHotUnplug = &v
}

// HasCpuHotUnplug returns a boolean if a field has been set.
func (o *SnapshotProperties) HasCpuHotUnplug() bool {
	if o != nil && o.CpuHotUnplug != nil {
		return true
	}

	return false
}

// GetRamHotPlug returns the RamHotPlug field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *SnapshotProperties) GetRamHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.RamHotPlug
}

// GetRamHotPlugOk returns a tuple with the RamHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetRamHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RamHotPlug, true
}

// SetRamHotPlug sets field value
func (o *SnapshotProperties) SetRamHotPlug(v bool) {
	o.RamHotPlug = &v
}

// HasRamHotPlug returns a boolean if a field has been set.
func (o *SnapshotProperties) HasRamHotPlug() bool {
	if o != nil && o.RamHotPlug != nil {
		return true
	}

	return false
}

// GetRamHotUnplug returns the RamHotUnplug field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *SnapshotProperties) GetRamHotUnplug() *bool {
	if o == nil {
		return nil
	}

	return o.RamHotUnplug
}

// GetRamHotUnplugOk returns a tuple with the RamHotUnplug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetRamHotUnplugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RamHotUnplug, true
}

// SetRamHotUnplug sets field value
func (o *SnapshotProperties) SetRamHotUnplug(v bool) {
	o.RamHotUnplug = &v
}

// HasRamHotUnplug returns a boolean if a field has been set.
func (o *SnapshotProperties) HasRamHotUnplug() bool {
	if o != nil && o.RamHotUnplug != nil {
		return true
	}

	return false
}

// GetNicHotPlug returns the NicHotPlug field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *SnapshotProperties) GetNicHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.NicHotPlug
}

// GetNicHotPlugOk returns a tuple with the NicHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetNicHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NicHotPlug, true
}

// SetNicHotPlug sets field value
func (o *SnapshotProperties) SetNicHotPlug(v bool) {
	o.NicHotPlug = &v
}

// HasNicHotPlug returns a boolean if a field has been set.
func (o *SnapshotProperties) HasNicHotPlug() bool {
	if o != nil && o.NicHotPlug != nil {
		return true
	}

	return false
}

// GetNicHotUnplug returns the NicHotUnplug field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *SnapshotProperties) GetNicHotUnplug() *bool {
	if o == nil {
		return nil
	}

	return o.NicHotUnplug
}

// GetNicHotUnplugOk returns a tuple with the NicHotUnplug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetNicHotUnplugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NicHotUnplug, true
}

// SetNicHotUnplug sets field value
func (o *SnapshotProperties) SetNicHotUnplug(v bool) {
	o.NicHotUnplug = &v
}

// HasNicHotUnplug returns a boolean if a field has been set.
func (o *SnapshotProperties) HasNicHotUnplug() bool {
	if o != nil && o.NicHotUnplug != nil {
		return true
	}

	return false
}

// GetDiscVirtioHotPlug returns the DiscVirtioHotPlug field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *SnapshotProperties) GetDiscVirtioHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.DiscVirtioHotPlug
}

// GetDiscVirtioHotPlugOk returns a tuple with the DiscVirtioHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetDiscVirtioHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscVirtioHotPlug, true
}

// SetDiscVirtioHotPlug sets field value
func (o *SnapshotProperties) SetDiscVirtioHotPlug(v bool) {
	o.DiscVirtioHotPlug = &v
}

// HasDiscVirtioHotPlug returns a boolean if a field has been set.
func (o *SnapshotProperties) HasDiscVirtioHotPlug() bool {
	if o != nil && o.DiscVirtioHotPlug != nil {
		return true
	}

	return false
}

// GetDiscVirtioHotUnplug returns the DiscVirtioHotUnplug field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *SnapshotProperties) GetDiscVirtioHotUnplug() *bool {
	if o == nil {
		return nil
	}

	return o.DiscVirtioHotUnplug
}

// GetDiscVirtioHotUnplugOk returns a tuple with the DiscVirtioHotUnplug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetDiscVirtioHotUnplugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscVirtioHotUnplug, true
}

// SetDiscVirtioHotUnplug sets field value
func (o *SnapshotProperties) SetDiscVirtioHotUnplug(v bool) {
	o.DiscVirtioHotUnplug = &v
}

// HasDiscVirtioHotUnplug returns a boolean if a field has been set.
func (o *SnapshotProperties) HasDiscVirtioHotUnplug() bool {
	if o != nil && o.DiscVirtioHotUnplug != nil {
		return true
	}

	return false
}

// GetDiscScsiHotPlug returns the DiscScsiHotPlug field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *SnapshotProperties) GetDiscScsiHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.DiscScsiHotPlug
}

// GetDiscScsiHotPlugOk returns a tuple with the DiscScsiHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetDiscScsiHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscScsiHotPlug, true
}

// SetDiscScsiHotPlug sets field value
func (o *SnapshotProperties) SetDiscScsiHotPlug(v bool) {
	o.DiscScsiHotPlug = &v
}

// HasDiscScsiHotPlug returns a boolean if a field has been set.
func (o *SnapshotProperties) HasDiscScsiHotPlug() bool {
	if o != nil && o.DiscScsiHotPlug != nil {
		return true
	}

	return false
}

// GetDiscScsiHotUnplug returns the DiscScsiHotUnplug field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *SnapshotProperties) GetDiscScsiHotUnplug() *bool {
	if o == nil {
		return nil
	}

	return o.DiscScsiHotUnplug
}

// GetDiscScsiHotUnplugOk returns a tuple with the DiscScsiHotUnplug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetDiscScsiHotUnplugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscScsiHotUnplug, true
}

// SetDiscScsiHotUnplug sets field value
func (o *SnapshotProperties) SetDiscScsiHotUnplug(v bool) {
	o.DiscScsiHotUnplug = &v
}

// HasDiscScsiHotUnplug returns a boolean if a field has been set.
func (o *SnapshotProperties) HasDiscScsiHotUnplug() bool {
	if o != nil && o.DiscScsiHotUnplug != nil {
		return true
	}

	return false
}

// GetLicenceType returns the LicenceType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SnapshotProperties) GetLicenceType() *string {
	if o == nil {
		return nil
	}

	return o.LicenceType
}

// GetLicenceTypeOk returns a tuple with the LicenceType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotProperties) GetLicenceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LicenceType, true
}

// SetLicenceType sets field value
func (o *SnapshotProperties) SetLicenceType(v string) {
	o.LicenceType = &v
}

// HasLicenceType returns a boolean if a field has been set.
func (o *SnapshotProperties) HasLicenceType() bool {
	if o != nil && o.LicenceType != nil {
		return true
	}

	return false
}

func (o SnapshotProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	if o.Location != nil {
		toSerialize["location"] = o.Location
	}

	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	if o.SecAuthProtection != nil {
		toSerialize["secAuthProtection"] = o.SecAuthProtection
	}

	if o.CpuHotPlug != nil {
		toSerialize["cpuHotPlug"] = o.CpuHotPlug
	}

	if o.CpuHotUnplug != nil {
		toSerialize["cpuHotUnplug"] = o.CpuHotUnplug
	}

	if o.RamHotPlug != nil {
		toSerialize["ramHotPlug"] = o.RamHotPlug
	}

	if o.RamHotUnplug != nil {
		toSerialize["ramHotUnplug"] = o.RamHotUnplug
	}

	if o.NicHotPlug != nil {
		toSerialize["nicHotPlug"] = o.NicHotPlug
	}

	if o.NicHotUnplug != nil {
		toSerialize["nicHotUnplug"] = o.NicHotUnplug
	}

	if o.DiscVirtioHotPlug != nil {
		toSerialize["discVirtioHotPlug"] = o.DiscVirtioHotPlug
	}

	if o.DiscVirtioHotUnplug != nil {
		toSerialize["discVirtioHotUnplug"] = o.DiscVirtioHotUnplug
	}

	if o.DiscScsiHotPlug != nil {
		toSerialize["discScsiHotPlug"] = o.DiscScsiHotPlug
	}

	if o.DiscScsiHotUnplug != nil {
		toSerialize["discScsiHotUnplug"] = o.DiscScsiHotUnplug
	}

	if o.LicenceType != nil {
		toSerialize["licenceType"] = o.LicenceType
	}

	return json.Marshal(toSerialize)
}

type NullableSnapshotProperties struct {
	value *SnapshotProperties
	isSet bool
}

func (v NullableSnapshotProperties) Get() *SnapshotProperties {
	return v.value
}

func (v *NullableSnapshotProperties) Set(val *SnapshotProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotProperties(val *SnapshotProperties) *NullableSnapshotProperties {
	return &NullableSnapshotProperties{value: val, isSet: true}
}

func (v NullableSnapshotProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
