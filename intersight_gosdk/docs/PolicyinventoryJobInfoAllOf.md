# PolicyinventoryJobInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policyinventory.JobInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policyinventory.JobInfo"]
**ExecutionStatus** | Pointer to **string** | Execution status of the inventory job. * &#x60;Scheduled&#x60; - Inventory job is marked as scheduled. * &#x60;Completed&#x60; - Inventory job is marked as completed. * &#x60;Error&#x60; - Inventory job has errored out. | [optional] [readonly] [default to "Scheduled"]
**LastScheduledTime** | Pointer to [**time.Time**](time.Time.md) | Last scheduled time of the inventory job. | [optional] [readonly] 
**PolicyId** | Pointer to **string** | Policy ID for the inventory job. | [optional] [readonly] 
**PolicyName** | Pointer to **string** | Policy name for the inventory job. | [optional] [readonly] 

## Methods

### NewPolicyinventoryJobInfoAllOf

`func NewPolicyinventoryJobInfoAllOf(classId string, objectType string, ) *PolicyinventoryJobInfoAllOf`

NewPolicyinventoryJobInfoAllOf instantiates a new PolicyinventoryJobInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyinventoryJobInfoAllOfWithDefaults

`func NewPolicyinventoryJobInfoAllOfWithDefaults() *PolicyinventoryJobInfoAllOf`

NewPolicyinventoryJobInfoAllOfWithDefaults instantiates a new PolicyinventoryJobInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyinventoryJobInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyinventoryJobInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyinventoryJobInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyinventoryJobInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyinventoryJobInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyinventoryJobInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExecutionStatus

`func (o *PolicyinventoryJobInfoAllOf) GetExecutionStatus() string`

GetExecutionStatus returns the ExecutionStatus field if non-nil, zero value otherwise.

### GetExecutionStatusOk

`func (o *PolicyinventoryJobInfoAllOf) GetExecutionStatusOk() (*string, bool)`

GetExecutionStatusOk returns a tuple with the ExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStatus

`func (o *PolicyinventoryJobInfoAllOf) SetExecutionStatus(v string)`

SetExecutionStatus sets ExecutionStatus field to given value.

### HasExecutionStatus

`func (o *PolicyinventoryJobInfoAllOf) HasExecutionStatus() bool`

HasExecutionStatus returns a boolean if a field has been set.

### GetLastScheduledTime

`func (o *PolicyinventoryJobInfoAllOf) GetLastScheduledTime() time.Time`

GetLastScheduledTime returns the LastScheduledTime field if non-nil, zero value otherwise.

### GetLastScheduledTimeOk

`func (o *PolicyinventoryJobInfoAllOf) GetLastScheduledTimeOk() (*time.Time, bool)`

GetLastScheduledTimeOk returns a tuple with the LastScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScheduledTime

`func (o *PolicyinventoryJobInfoAllOf) SetLastScheduledTime(v time.Time)`

SetLastScheduledTime sets LastScheduledTime field to given value.

### HasLastScheduledTime

`func (o *PolicyinventoryJobInfoAllOf) HasLastScheduledTime() bool`

HasLastScheduledTime returns a boolean if a field has been set.

### GetPolicyId

`func (o *PolicyinventoryJobInfoAllOf) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyinventoryJobInfoAllOf) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyinventoryJobInfoAllOf) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PolicyinventoryJobInfoAllOf) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *PolicyinventoryJobInfoAllOf) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PolicyinventoryJobInfoAllOf) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PolicyinventoryJobInfoAllOf) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PolicyinventoryJobInfoAllOf) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


