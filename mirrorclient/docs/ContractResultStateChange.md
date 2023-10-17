# ContractResultStateChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to ***os.File** | A network entity encoded as an EVM address in hex. | [optional] 
**ContractId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Slot** | Pointer to ***os.File** | The hex encoded storage slot changed. | [optional] 
**ValueRead** | Pointer to ***os.File** | The hex encoded value read from the storage slot. | [optional] 
**ValueWritten** | Pointer to **Nullable*os.File** | The hex encoded value written to the slot. &#x60;null&#x60; implies no value written. | [optional] 

## Methods

### NewContractResultStateChange

`func NewContractResultStateChange() *ContractResultStateChange`

NewContractResultStateChange instantiates a new ContractResultStateChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractResultStateChangeWithDefaults

`func NewContractResultStateChangeWithDefaults() *ContractResultStateChange`

NewContractResultStateChangeWithDefaults instantiates a new ContractResultStateChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ContractResultStateChange) GetAddress() *os.File`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractResultStateChange) GetAddressOk() (**os.File, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractResultStateChange) SetAddress(v *os.File)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ContractResultStateChange) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetContractId

`func (o *ContractResultStateChange) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ContractResultStateChange) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ContractResultStateChange) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *ContractResultStateChange) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### SetContractIdNil

`func (o *ContractResultStateChange) SetContractIdNil(b bool)`

 SetContractIdNil sets the value for ContractId to be an explicit nil

### UnsetContractId
`func (o *ContractResultStateChange) UnsetContractId()`

UnsetContractId ensures that no value is present for ContractId, not even an explicit nil
### GetSlot

`func (o *ContractResultStateChange) GetSlot() *os.File`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *ContractResultStateChange) GetSlotOk() (**os.File, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *ContractResultStateChange) SetSlot(v *os.File)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *ContractResultStateChange) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetValueRead

`func (o *ContractResultStateChange) GetValueRead() *os.File`

GetValueRead returns the ValueRead field if non-nil, zero value otherwise.

### GetValueReadOk

`func (o *ContractResultStateChange) GetValueReadOk() (**os.File, bool)`

GetValueReadOk returns a tuple with the ValueRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRead

`func (o *ContractResultStateChange) SetValueRead(v *os.File)`

SetValueRead sets ValueRead field to given value.

### HasValueRead

`func (o *ContractResultStateChange) HasValueRead() bool`

HasValueRead returns a boolean if a field has been set.

### GetValueWritten

`func (o *ContractResultStateChange) GetValueWritten() *os.File`

GetValueWritten returns the ValueWritten field if non-nil, zero value otherwise.

### GetValueWrittenOk

`func (o *ContractResultStateChange) GetValueWrittenOk() (**os.File, bool)`

GetValueWrittenOk returns a tuple with the ValueWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueWritten

`func (o *ContractResultStateChange) SetValueWritten(v *os.File)`

SetValueWritten sets ValueWritten field to given value.

### HasValueWritten

`func (o *ContractResultStateChange) HasValueWritten() bool`

HasValueWritten returns a boolean if a field has been set.

### SetValueWrittenNil

`func (o *ContractResultStateChange) SetValueWrittenNil(b bool)`

 SetValueWrittenNil sets the value for ValueWritten to be an explicit nil

### UnsetValueWritten
`func (o *ContractResultStateChange) UnsetValueWritten()`

UnsetValueWritten ensures that no value is present for ValueWritten, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


