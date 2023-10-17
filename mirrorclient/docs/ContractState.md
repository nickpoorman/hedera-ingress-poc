# ContractState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | ***os.File** | A network entity encoded as an EVM address in hex. | 
**ContractId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**Timestamp** | **string** |  | 
**Slot** | ***os.File** | The hex encoded storage slot. | 
**Value** | ***os.File** | The hex encoded value to the slot. &#x60;0x&#x60; implies no value written. | 

## Methods

### NewContractState

`func NewContractState(address *os.File, contractId NullableString, timestamp string, slot *os.File, value *os.File, ) *ContractState`

NewContractState instantiates a new ContractState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractStateWithDefaults

`func NewContractStateWithDefaults() *ContractState`

NewContractStateWithDefaults instantiates a new ContractState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ContractState) GetAddress() *os.File`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractState) GetAddressOk() (**os.File, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractState) SetAddress(v *os.File)`

SetAddress sets Address field to given value.


### GetContractId

`func (o *ContractState) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ContractState) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ContractState) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### SetContractIdNil

`func (o *ContractState) SetContractIdNil(b bool)`

 SetContractIdNil sets the value for ContractId to be an explicit nil

### UnsetContractId
`func (o *ContractState) UnsetContractId()`

UnsetContractId ensures that no value is present for ContractId, not even an explicit nil
### GetTimestamp

`func (o *ContractState) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ContractState) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ContractState) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetSlot

`func (o *ContractState) GetSlot() *os.File`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *ContractState) GetSlotOk() (**os.File, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *ContractState) SetSlot(v *os.File)`

SetSlot sets Slot field to given value.


### GetValue

`func (o *ContractState) GetValue() *os.File`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContractState) GetValueOk() (**os.File, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContractState) SetValue(v *os.File)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


