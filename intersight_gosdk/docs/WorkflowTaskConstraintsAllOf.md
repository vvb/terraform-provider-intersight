# WorkflowTaskConstraintsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TaskConstraints"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TaskConstraints"]
**TargetDataType** | Pointer to **interface{}** | List of property constraints that helps to narrow down task implementations based on target device input. | [optional] 

## Methods

### NewWorkflowTaskConstraintsAllOf

`func NewWorkflowTaskConstraintsAllOf(classId string, objectType string, ) *WorkflowTaskConstraintsAllOf`

NewWorkflowTaskConstraintsAllOf instantiates a new WorkflowTaskConstraintsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskConstraintsAllOfWithDefaults

`func NewWorkflowTaskConstraintsAllOfWithDefaults() *WorkflowTaskConstraintsAllOf`

NewWorkflowTaskConstraintsAllOfWithDefaults instantiates a new WorkflowTaskConstraintsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTaskConstraintsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTaskConstraintsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTaskConstraintsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTaskConstraintsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTaskConstraintsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTaskConstraintsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTargetDataType

`func (o *WorkflowTaskConstraintsAllOf) GetTargetDataType() interface{}`

GetTargetDataType returns the TargetDataType field if non-nil, zero value otherwise.

### GetTargetDataTypeOk

`func (o *WorkflowTaskConstraintsAllOf) GetTargetDataTypeOk() (*interface{}, bool)`

GetTargetDataTypeOk returns a tuple with the TargetDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDataType

`func (o *WorkflowTaskConstraintsAllOf) SetTargetDataType(v interface{})`

SetTargetDataType sets TargetDataType field to given value.

### HasTargetDataType

`func (o *WorkflowTaskConstraintsAllOf) HasTargetDataType() bool`

HasTargetDataType returns a boolean if a field has been set.

### SetTargetDataTypeNil

`func (o *WorkflowTaskConstraintsAllOf) SetTargetDataTypeNil(b bool)`

 SetTargetDataTypeNil sets the value for TargetDataType to be an explicit nil

### UnsetTargetDataType
`func (o *WorkflowTaskConstraintsAllOf) UnsetTargetDataType()`

UnsetTargetDataType ensures that no value is present for TargetDataType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


