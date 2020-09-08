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

// TelemetryDruidPeriodGranularityAllOf struct for TelemetryDruidPeriodGranularityAllOf
type TelemetryDruidPeriodGranularityAllOf struct {
	// The period in ISO 8601 format. Examples are P2W, P3M, PT1H30M, PT0.750S.
	Period string `json:"period"`
	// An optional value specifying the time zone. Standard [IANA time zones](http://joda-time.sourceforge.net/timezones.html) are supported.
	TimeZone             *string `json:"timeZone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidPeriodGranularityAllOf TelemetryDruidPeriodGranularityAllOf

// NewTelemetryDruidPeriodGranularityAllOf instantiates a new TelemetryDruidPeriodGranularityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidPeriodGranularityAllOf(period string) *TelemetryDruidPeriodGranularityAllOf {
	this := TelemetryDruidPeriodGranularityAllOf{}
	this.Period = period
	return &this
}

// NewTelemetryDruidPeriodGranularityAllOfWithDefaults instantiates a new TelemetryDruidPeriodGranularityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidPeriodGranularityAllOfWithDefaults() *TelemetryDruidPeriodGranularityAllOf {
	this := TelemetryDruidPeriodGranularityAllOf{}
	return &this
}

// GetPeriod returns the Period field value
func (o *TelemetryDruidPeriodGranularityAllOf) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidPeriodGranularityAllOf) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *TelemetryDruidPeriodGranularityAllOf) SetPeriod(v string) {
	o.Period = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *TelemetryDruidPeriodGranularityAllOf) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidPeriodGranularityAllOf) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *TelemetryDruidPeriodGranularityAllOf) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *TelemetryDruidPeriodGranularityAllOf) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o TelemetryDruidPeriodGranularityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["period"] = o.Period
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidPeriodGranularityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidPeriodGranularityAllOf := _TelemetryDruidPeriodGranularityAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidPeriodGranularityAllOf); err == nil {
		*o = TelemetryDruidPeriodGranularityAllOf(varTelemetryDruidPeriodGranularityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "period")
		delete(additionalProperties, "timeZone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidPeriodGranularityAllOf struct {
	value *TelemetryDruidPeriodGranularityAllOf
	isSet bool
}

func (v NullableTelemetryDruidPeriodGranularityAllOf) Get() *TelemetryDruidPeriodGranularityAllOf {
	return v.value
}

func (v *NullableTelemetryDruidPeriodGranularityAllOf) Set(val *TelemetryDruidPeriodGranularityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidPeriodGranularityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidPeriodGranularityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidPeriodGranularityAllOf(val *TelemetryDruidPeriodGranularityAllOf) *NullableTelemetryDruidPeriodGranularityAllOf {
	return &NullableTelemetryDruidPeriodGranularityAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidPeriodGranularityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidPeriodGranularityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}