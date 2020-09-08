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

// SyslogLocalFileLoggingClient Local logging client on the endpoint.
type SyslogLocalFileLoggingClient struct {
	SyslogLocalClientBase
	AdditionalProperties map[string]interface{}
}

type _SyslogLocalFileLoggingClient SyslogLocalFileLoggingClient

// NewSyslogLocalFileLoggingClient instantiates a new SyslogLocalFileLoggingClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogLocalFileLoggingClient() *SyslogLocalFileLoggingClient {
	this := SyslogLocalFileLoggingClient{}
	return &this
}

// NewSyslogLocalFileLoggingClientWithDefaults instantiates a new SyslogLocalFileLoggingClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogLocalFileLoggingClientWithDefaults() *SyslogLocalFileLoggingClient {
	this := SyslogLocalFileLoggingClient{}
	return &this
}

func (o SyslogLocalFileLoggingClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSyslogLocalClientBase, errSyslogLocalClientBase := json.Marshal(o.SyslogLocalClientBase)
	if errSyslogLocalClientBase != nil {
		return []byte{}, errSyslogLocalClientBase
	}
	errSyslogLocalClientBase = json.Unmarshal([]byte(serializedSyslogLocalClientBase), &toSerialize)
	if errSyslogLocalClientBase != nil {
		return []byte{}, errSyslogLocalClientBase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SyslogLocalFileLoggingClient) UnmarshalJSON(bytes []byte) (err error) {
	type SyslogLocalFileLoggingClientWithoutEmbeddedStruct struct {
	}

	varSyslogLocalFileLoggingClientWithoutEmbeddedStruct := SyslogLocalFileLoggingClientWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSyslogLocalFileLoggingClientWithoutEmbeddedStruct)
	if err == nil {
		varSyslogLocalFileLoggingClient := _SyslogLocalFileLoggingClient{}
		*o = SyslogLocalFileLoggingClient(varSyslogLocalFileLoggingClient)
	} else {
		return err
	}

	varSyslogLocalFileLoggingClient := _SyslogLocalFileLoggingClient{}

	err = json.Unmarshal(bytes, &varSyslogLocalFileLoggingClient)
	if err == nil {
		o.SyslogLocalClientBase = varSyslogLocalFileLoggingClient.SyslogLocalClientBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectSyslogLocalClientBase := reflect.ValueOf(o.SyslogLocalClientBase)
		for i := 0; i < reflectSyslogLocalClientBase.Type().NumField(); i++ {
			t := reflectSyslogLocalClientBase.Type().Field(i)

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

type NullableSyslogLocalFileLoggingClient struct {
	value *SyslogLocalFileLoggingClient
	isSet bool
}

func (v NullableSyslogLocalFileLoggingClient) Get() *SyslogLocalFileLoggingClient {
	return v.value
}

func (v *NullableSyslogLocalFileLoggingClient) Set(val *SyslogLocalFileLoggingClient) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogLocalFileLoggingClient) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogLocalFileLoggingClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogLocalFileLoggingClient(val *SyslogLocalFileLoggingClient) *NullableSyslogLocalFileLoggingClient {
	return &NullableSyslogLocalFileLoggingClient{value: val, isSet: true}
}

func (v NullableSyslogLocalFileLoggingClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogLocalFileLoggingClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}