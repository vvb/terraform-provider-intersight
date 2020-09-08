/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VirtualizationVmwareVmMemoryShareInfo Information about the virtual machine's memory sharing and limits. For more information, see VMware documentation.
type VirtualizationVmwareVmMemoryShareInfo struct {
	MoBaseComplexType
	// Limit on the memory sharing imposed (in Mbytes).
	MemLimit *int64 `json:"MemLimit,omitempty"`
	// Limit on memory overhead imposed (in Mbytes).
	MemOverheadLimit *int64 `json:"MemOverheadLimit,omitempty"`
	// Similar to CPU reservations (Mbytes).
	MemReservation *int64 `json:"MemReservation,omitempty"`
	// Similar to CPU Shares but applicable to memory. There is no unit for this value. It is a relative measure based on the settings for other resource pools.
	MemShares            *int64 `json:"MemShares,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVmMemoryShareInfo VirtualizationVmwareVmMemoryShareInfo

// NewVirtualizationVmwareVmMemoryShareInfo instantiates a new VirtualizationVmwareVmMemoryShareInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVmMemoryShareInfo() *VirtualizationVmwareVmMemoryShareInfo {
	this := VirtualizationVmwareVmMemoryShareInfo{}
	return &this
}

// NewVirtualizationVmwareVmMemoryShareInfoWithDefaults instantiates a new VirtualizationVmwareVmMemoryShareInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVmMemoryShareInfoWithDefaults() *VirtualizationVmwareVmMemoryShareInfo {
	this := VirtualizationVmwareVmMemoryShareInfo{}
	return &this
}

// GetMemLimit returns the MemLimit field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemLimit() int64 {
	if o == nil || o.MemLimit == nil {
		var ret int64
		return ret
	}
	return *o.MemLimit
}

// GetMemLimitOk returns a tuple with the MemLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemLimitOk() (*int64, bool) {
	if o == nil || o.MemLimit == nil {
		return nil, false
	}
	return o.MemLimit, true
}

// HasMemLimit returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) HasMemLimit() bool {
	if o != nil && o.MemLimit != nil {
		return true
	}

	return false
}

// SetMemLimit gets a reference to the given int64 and assigns it to the MemLimit field.
func (o *VirtualizationVmwareVmMemoryShareInfo) SetMemLimit(v int64) {
	o.MemLimit = &v
}

// GetMemOverheadLimit returns the MemOverheadLimit field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemOverheadLimit() int64 {
	if o == nil || o.MemOverheadLimit == nil {
		var ret int64
		return ret
	}
	return *o.MemOverheadLimit
}

// GetMemOverheadLimitOk returns a tuple with the MemOverheadLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemOverheadLimitOk() (*int64, bool) {
	if o == nil || o.MemOverheadLimit == nil {
		return nil, false
	}
	return o.MemOverheadLimit, true
}

// HasMemOverheadLimit returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) HasMemOverheadLimit() bool {
	if o != nil && o.MemOverheadLimit != nil {
		return true
	}

	return false
}

// SetMemOverheadLimit gets a reference to the given int64 and assigns it to the MemOverheadLimit field.
func (o *VirtualizationVmwareVmMemoryShareInfo) SetMemOverheadLimit(v int64) {
	o.MemOverheadLimit = &v
}

// GetMemReservation returns the MemReservation field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemReservation() int64 {
	if o == nil || o.MemReservation == nil {
		var ret int64
		return ret
	}
	return *o.MemReservation
}

// GetMemReservationOk returns a tuple with the MemReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemReservationOk() (*int64, bool) {
	if o == nil || o.MemReservation == nil {
		return nil, false
	}
	return o.MemReservation, true
}

// HasMemReservation returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) HasMemReservation() bool {
	if o != nil && o.MemReservation != nil {
		return true
	}

	return false
}

// SetMemReservation gets a reference to the given int64 and assigns it to the MemReservation field.
func (o *VirtualizationVmwareVmMemoryShareInfo) SetMemReservation(v int64) {
	o.MemReservation = &v
}

// GetMemShares returns the MemShares field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemShares() int64 {
	if o == nil || o.MemShares == nil {
		var ret int64
		return ret
	}
	return *o.MemShares
}

// GetMemSharesOk returns a tuple with the MemShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemSharesOk() (*int64, bool) {
	if o == nil || o.MemShares == nil {
		return nil, false
	}
	return o.MemShares, true
}

// HasMemShares returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) HasMemShares() bool {
	if o != nil && o.MemShares != nil {
		return true
	}

	return false
}

// SetMemShares gets a reference to the given int64 and assigns it to the MemShares field.
func (o *VirtualizationVmwareVmMemoryShareInfo) SetMemShares(v int64) {
	o.MemShares = &v
}

func (o VirtualizationVmwareVmMemoryShareInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.MemLimit != nil {
		toSerialize["MemLimit"] = o.MemLimit
	}
	if o.MemOverheadLimit != nil {
		toSerialize["MemOverheadLimit"] = o.MemOverheadLimit
	}
	if o.MemReservation != nil {
		toSerialize["MemReservation"] = o.MemReservation
	}
	if o.MemShares != nil {
		toSerialize["MemShares"] = o.MemShares
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareVmMemoryShareInfo) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct struct {
		// Limit on the memory sharing imposed (in Mbytes).
		MemLimit *int64 `json:"MemLimit,omitempty"`
		// Limit on memory overhead imposed (in Mbytes).
		MemOverheadLimit *int64 `json:"MemOverheadLimit,omitempty"`
		// Similar to CPU reservations (Mbytes).
		MemReservation *int64 `json:"MemReservation,omitempty"`
		// Similar to CPU Shares but applicable to memory. There is no unit for this value. It is a relative measure based on the settings for other resource pools.
		MemShares *int64 `json:"MemShares,omitempty"`
	}

	varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct := VirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareVmMemoryShareInfo := _VirtualizationVmwareVmMemoryShareInfo{}
		varVirtualizationVmwareVmMemoryShareInfo.MemLimit = varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct.MemLimit
		varVirtualizationVmwareVmMemoryShareInfo.MemOverheadLimit = varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct.MemOverheadLimit
		varVirtualizationVmwareVmMemoryShareInfo.MemReservation = varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct.MemReservation
		varVirtualizationVmwareVmMemoryShareInfo.MemShares = varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct.MemShares
		*o = VirtualizationVmwareVmMemoryShareInfo(varVirtualizationVmwareVmMemoryShareInfo)
	} else {
		return err
	}

	varVirtualizationVmwareVmMemoryShareInfo := _VirtualizationVmwareVmMemoryShareInfo{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVmMemoryShareInfo)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationVmwareVmMemoryShareInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "MemLimit")
		delete(additionalProperties, "MemOverheadLimit")
		delete(additionalProperties, "MemReservation")
		delete(additionalProperties, "MemShares")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareVmMemoryShareInfo struct {
	value *VirtualizationVmwareVmMemoryShareInfo
	isSet bool
}

func (v NullableVirtualizationVmwareVmMemoryShareInfo) Get() *VirtualizationVmwareVmMemoryShareInfo {
	return v.value
}

func (v *NullableVirtualizationVmwareVmMemoryShareInfo) Set(val *VirtualizationVmwareVmMemoryShareInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVmMemoryShareInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVmMemoryShareInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVmMemoryShareInfo(val *VirtualizationVmwareVmMemoryShareInfo) *NullableVirtualizationVmwareVmMemoryShareInfo {
	return &NullableVirtualizationVmwareVmMemoryShareInfo{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVmMemoryShareInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVmMemoryShareInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}