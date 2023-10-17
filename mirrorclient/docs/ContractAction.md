# ContractAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallDepth** | Pointer to **int32** | The nesting depth of the call | [optional] 
**CallOperationType** | Pointer to **string** | The type of the call operation | [optional] 
**CallType** | Pointer to **string** | The type of the call | [optional] 
**Caller** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**CallerType** | Pointer to **string** | The entity type of the caller | [optional] 
**From** | Pointer to **string** | The EVM address of the caller | [optional] 
**Gas** | Pointer to **int64** | Gas cost in tinybars | [optional] 
**GasUsed** | Pointer to **int64** | Gas used in tinybars | [optional] 
**Index** | Pointer to **int32** | The position of the action within the ordered list of actions | [optional] 
**Input** | Pointer to **NullableString** | The hex encoded input data | [optional] 
**Recipient** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**RecipientType** | Pointer to **NullableString** | The entity type of the recipient | [optional] 
**ResultData** | Pointer to **NullableString** | The hex encoded result data | [optional] 
**ResultDataType** | Pointer to **string** | The type of the result data | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**To** | Pointer to **Nullable*os.File** | A network entity encoded as an EVM address in hex. | [optional] 
**Value** | Pointer to **int64** | The value of the transaction in tinybars | [optional] 

## Methods

### NewContractAction

`func NewContractAction() *ContractAction`

NewContractAction instantiates a new ContractAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractActionWithDefaults

`func NewContractActionWithDefaults() *ContractAction`

NewContractActionWithDefaults instantiates a new ContractAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallDepth

`func (o *ContractAction) GetCallDepth() int32`

GetCallDepth returns the CallDepth field if non-nil, zero value otherwise.

### GetCallDepthOk

`func (o *ContractAction) GetCallDepthOk() (*int32, bool)`

GetCallDepthOk returns a tuple with the CallDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallDepth

`func (o *ContractAction) SetCallDepth(v int32)`

SetCallDepth sets CallDepth field to given value.

### HasCallDepth

`func (o *ContractAction) HasCallDepth() bool`

HasCallDepth returns a boolean if a field has been set.

### GetCallOperationType

`func (o *ContractAction) GetCallOperationType() string`

GetCallOperationType returns the CallOperationType field if non-nil, zero value otherwise.

### GetCallOperationTypeOk

`func (o *ContractAction) GetCallOperationTypeOk() (*string, bool)`

GetCallOperationTypeOk returns a tuple with the CallOperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallOperationType

`func (o *ContractAction) SetCallOperationType(v string)`

SetCallOperationType sets CallOperationType field to given value.

### HasCallOperationType

`func (o *ContractAction) HasCallOperationType() bool`

HasCallOperationType returns a boolean if a field has been set.

### GetCallType

`func (o *ContractAction) GetCallType() string`

GetCallType returns the CallType field if non-nil, zero value otherwise.

### GetCallTypeOk

`func (o *ContractAction) GetCallTypeOk() (*string, bool)`

GetCallTypeOk returns a tuple with the CallType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallType

`func (o *ContractAction) SetCallType(v string)`

SetCallType sets CallType field to given value.

### HasCallType

`func (o *ContractAction) HasCallType() bool`

HasCallType returns a boolean if a field has been set.

### GetCaller

`func (o *ContractAction) GetCaller() string`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *ContractAction) GetCallerOk() (*string, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *ContractAction) SetCaller(v string)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *ContractAction) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *ContractAction) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *ContractAction) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetCallerType

`func (o *ContractAction) GetCallerType() string`

GetCallerType returns the CallerType field if non-nil, zero value otherwise.

### GetCallerTypeOk

`func (o *ContractAction) GetCallerTypeOk() (*string, bool)`

GetCallerTypeOk returns a tuple with the CallerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerType

`func (o *ContractAction) SetCallerType(v string)`

SetCallerType sets CallerType field to given value.

### HasCallerType

`func (o *ContractAction) HasCallerType() bool`

HasCallerType returns a boolean if a field has been set.

### GetFrom

`func (o *ContractAction) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ContractAction) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ContractAction) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ContractAction) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGas

`func (o *ContractAction) GetGas() int64`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *ContractAction) GetGasOk() (*int64, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *ContractAction) SetGas(v int64)`

SetGas sets Gas field to given value.

### HasGas

`func (o *ContractAction) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasUsed

`func (o *ContractAction) GetGasUsed() int64`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ContractAction) GetGasUsedOk() (*int64, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ContractAction) SetGasUsed(v int64)`

SetGasUsed sets GasUsed field to given value.

### HasGasUsed

`func (o *ContractAction) HasGasUsed() bool`

HasGasUsed returns a boolean if a field has been set.

### GetIndex

`func (o *ContractAction) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ContractAction) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ContractAction) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ContractAction) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetInput

`func (o *ContractAction) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ContractAction) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ContractAction) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *ContractAction) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *ContractAction) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ContractAction) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetRecipient

`func (o *ContractAction) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ContractAction) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ContractAction) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *ContractAction) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *ContractAction) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *ContractAction) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
### GetRecipientType

`func (o *ContractAction) GetRecipientType() string`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *ContractAction) GetRecipientTypeOk() (*string, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *ContractAction) SetRecipientType(v string)`

SetRecipientType sets RecipientType field to given value.

### HasRecipientType

`func (o *ContractAction) HasRecipientType() bool`

HasRecipientType returns a boolean if a field has been set.

### SetRecipientTypeNil

`func (o *ContractAction) SetRecipientTypeNil(b bool)`

 SetRecipientTypeNil sets the value for RecipientType to be an explicit nil

### UnsetRecipientType
`func (o *ContractAction) UnsetRecipientType()`

UnsetRecipientType ensures that no value is present for RecipientType, not even an explicit nil
### GetResultData

`func (o *ContractAction) GetResultData() string`

GetResultData returns the ResultData field if non-nil, zero value otherwise.

### GetResultDataOk

`func (o *ContractAction) GetResultDataOk() (*string, bool)`

GetResultDataOk returns a tuple with the ResultData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultData

`func (o *ContractAction) SetResultData(v string)`

SetResultData sets ResultData field to given value.

### HasResultData

`func (o *ContractAction) HasResultData() bool`

HasResultData returns a boolean if a field has been set.

### SetResultDataNil

`func (o *ContractAction) SetResultDataNil(b bool)`

 SetResultDataNil sets the value for ResultData to be an explicit nil

### UnsetResultData
`func (o *ContractAction) UnsetResultData()`

UnsetResultData ensures that no value is present for ResultData, not even an explicit nil
### GetResultDataType

`func (o *ContractAction) GetResultDataType() string`

GetResultDataType returns the ResultDataType field if non-nil, zero value otherwise.

### GetResultDataTypeOk

`func (o *ContractAction) GetResultDataTypeOk() (*string, bool)`

GetResultDataTypeOk returns a tuple with the ResultDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultDataType

`func (o *ContractAction) SetResultDataType(v string)`

SetResultDataType sets ResultDataType field to given value.

### HasResultDataType

`func (o *ContractAction) HasResultDataType() bool`

HasResultDataType returns a boolean if a field has been set.

### GetTimestamp

`func (o *ContractAction) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ContractAction) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ContractAction) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ContractAction) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTo

`func (o *ContractAction) GetTo() *os.File`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ContractAction) GetToOk() (**os.File, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ContractAction) SetTo(v *os.File)`

SetTo sets To field to given value.

### HasTo

`func (o *ContractAction) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ContractAction) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ContractAction) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetValue

`func (o *ContractAction) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContractAction) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContractAction) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContractAction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


