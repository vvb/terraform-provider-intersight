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
)

// HyperflexCapabilityInfoAllOf Definition of the list of properties defined in 'hyperflex.CapabilityInfo', excluding properties defined in parent classes.
type HyperflexCapabilityInfoAllOf struct {
	CapabilityConstraints *[]HclConstraint `json:"CapabilityConstraints,omitempty"`
	// Name of the capability or feature set consisting of a collection of constraint rules and value.
	Name *string `json:"Name,omitempty"`
	// Capability Value which is valid only iff all specified constraints match.
	Value                *string                          `json:"Value,omitempty"`
	AppCatalog           *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexCapabilityInfoAllOf HyperflexCapabilityInfoAllOf

// NewHyperflexCapabilityInfoAllOf instantiates a new HyperflexCapabilityInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexCapabilityInfoAllOf() *HyperflexCapabilityInfoAllOf {
	this := HyperflexCapabilityInfoAllOf{}
	return &this
}

// NewHyperflexCapabilityInfoAllOfWithDefaults instantiates a new HyperflexCapabilityInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexCapabilityInfoAllOfWithDefaults() *HyperflexCapabilityInfoAllOf {
	this := HyperflexCapabilityInfoAllOf{}
	return &this
}

// GetCapabilityConstraints returns the CapabilityConstraints field value if set, zero value otherwise.
func (o *HyperflexCapabilityInfoAllOf) GetCapabilityConstraints() []HclConstraint {
	if o == nil || o.CapabilityConstraints == nil {
		var ret []HclConstraint
		return ret
	}
	return *o.CapabilityConstraints
}

// GetCapabilityConstraintsOk returns a tuple with the CapabilityConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCapabilityInfoAllOf) GetCapabilityConstraintsOk() (*[]HclConstraint, bool) {
	if o == nil || o.CapabilityConstraints == nil {
		return nil, false
	}
	return o.CapabilityConstraints, true
}

// HasCapabilityConstraints returns a boolean if a field has been set.
func (o *HyperflexCapabilityInfoAllOf) HasCapabilityConstraints() bool {
	if o != nil && o.CapabilityConstraints != nil {
		return true
	}

	return false
}

// SetCapabilityConstraints gets a reference to the given []HclConstraint and assigns it to the CapabilityConstraints field.
func (o *HyperflexCapabilityInfoAllOf) SetCapabilityConstraints(v []HclConstraint) {
	o.CapabilityConstraints = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexCapabilityInfoAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCapabilityInfoAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexCapabilityInfoAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexCapabilityInfoAllOf) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HyperflexCapabilityInfoAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCapabilityInfoAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HyperflexCapabilityInfoAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *HyperflexCapabilityInfoAllOf) SetValue(v string) {
	o.Value = &v
}

// GetAppCatalog returns the AppCatalog field value if set, zero value otherwise.
func (o *HyperflexCapabilityInfoAllOf) GetAppCatalog() HyperflexAppCatalogRelationship {
	if o == nil || o.AppCatalog == nil {
		var ret HyperflexAppCatalogRelationship
		return ret
	}
	return *o.AppCatalog
}

// GetAppCatalogOk returns a tuple with the AppCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCapabilityInfoAllOf) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool) {
	if o == nil || o.AppCatalog == nil {
		return nil, false
	}
	return o.AppCatalog, true
}

// HasAppCatalog returns a boolean if a field has been set.
func (o *HyperflexCapabilityInfoAllOf) HasAppCatalog() bool {
	if o != nil && o.AppCatalog != nil {
		return true
	}

	return false
}

// SetAppCatalog gets a reference to the given HyperflexAppCatalogRelationship and assigns it to the AppCatalog field.
func (o *HyperflexCapabilityInfoAllOf) SetAppCatalog(v HyperflexAppCatalogRelationship) {
	o.AppCatalog = &v
}

func (o HyperflexCapabilityInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CapabilityConstraints != nil {
		toSerialize["CapabilityConstraints"] = o.CapabilityConstraints
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	if o.AppCatalog != nil {
		toSerialize["AppCatalog"] = o.AppCatalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexCapabilityInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexCapabilityInfoAllOf := _HyperflexCapabilityInfoAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexCapabilityInfoAllOf); err == nil {
		*o = HyperflexCapabilityInfoAllOf(varHyperflexCapabilityInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "CapabilityConstraints")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Value")
		delete(additionalProperties, "AppCatalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexCapabilityInfoAllOf struct {
	value *HyperflexCapabilityInfoAllOf
	isSet bool
}

func (v NullableHyperflexCapabilityInfoAllOf) Get() *HyperflexCapabilityInfoAllOf {
	return v.value
}

func (v *NullableHyperflexCapabilityInfoAllOf) Set(val *HyperflexCapabilityInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexCapabilityInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexCapabilityInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexCapabilityInfoAllOf(val *HyperflexCapabilityInfoAllOf) *NullableHyperflexCapabilityInfoAllOf {
	return &NullableHyperflexCapabilityInfoAllOf{value: val, isSet: true}
}

func (v NullableHyperflexCapabilityInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexCapabilityInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}