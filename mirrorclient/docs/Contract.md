# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminKey** | Pointer to [**NullableKey**](Key.md) |  | [optional] 
**AutoRenewAccount** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**AutoRenewPeriod** | Pointer to **NullableInt64** |  | [optional] 
**ContractId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**CreatedTimestamp** | Pointer to **NullableString** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**EvmAddress** | Pointer to ***os.File** | A network entity encoded as an EVM address in hex. | [optional] 
**ExpirationTimestamp** | Pointer to **NullableString** |  | [optional] 
**FileId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**MaxAutomaticTokenAssociations** | Pointer to **NullableInt32** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**ObtainerId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**PermanentRemoval** | Pointer to **NullableBool** |  | [optional] 
**ProxyAccountId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Timestamp** | Pointer to [**TimestampRange**](TimestampRange.md) |  | [optional] 

## Methods

### NewContract

`func NewContract() *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminKey

`func (o *Contract) GetAdminKey() Key`

GetAdminKey returns the AdminKey field if non-nil, zero value otherwise.

### GetAdminKeyOk

`func (o *Contract) GetAdminKeyOk() (*Key, bool)`

GetAdminKeyOk returns a tuple with the AdminKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminKey

`func (o *Contract) SetAdminKey(v Key)`

SetAdminKey sets AdminKey field to given value.

### HasAdminKey

`func (o *Contract) HasAdminKey() bool`

HasAdminKey returns a boolean if a field has been set.

### SetAdminKeyNil

`func (o *Contract) SetAdminKeyNil(b bool)`

 SetAdminKeyNil sets the value for AdminKey to be an explicit nil

### UnsetAdminKey
`func (o *Contract) UnsetAdminKey()`

UnsetAdminKey ensures that no value is present for AdminKey, not even an explicit nil
### GetAutoRenewAccount

`func (o *Contract) GetAutoRenewAccount() string`

GetAutoRenewAccount returns the AutoRenewAccount field if non-nil, zero value otherwise.

### GetAutoRenewAccountOk

`func (o *Contract) GetAutoRenewAccountOk() (*string, bool)`

GetAutoRenewAccountOk returns a tuple with the AutoRenewAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewAccount

`func (o *Contract) SetAutoRenewAccount(v string)`

SetAutoRenewAccount sets AutoRenewAccount field to given value.

### HasAutoRenewAccount

`func (o *Contract) HasAutoRenewAccount() bool`

HasAutoRenewAccount returns a boolean if a field has been set.

### SetAutoRenewAccountNil

`func (o *Contract) SetAutoRenewAccountNil(b bool)`

 SetAutoRenewAccountNil sets the value for AutoRenewAccount to be an explicit nil

### UnsetAutoRenewAccount
`func (o *Contract) UnsetAutoRenewAccount()`

UnsetAutoRenewAccount ensures that no value is present for AutoRenewAccount, not even an explicit nil
### GetAutoRenewPeriod

`func (o *Contract) GetAutoRenewPeriod() int64`

GetAutoRenewPeriod returns the AutoRenewPeriod field if non-nil, zero value otherwise.

### GetAutoRenewPeriodOk

`func (o *Contract) GetAutoRenewPeriodOk() (*int64, bool)`

GetAutoRenewPeriodOk returns a tuple with the AutoRenewPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewPeriod

`func (o *Contract) SetAutoRenewPeriod(v int64)`

SetAutoRenewPeriod sets AutoRenewPeriod field to given value.

### HasAutoRenewPeriod

`func (o *Contract) HasAutoRenewPeriod() bool`

HasAutoRenewPeriod returns a boolean if a field has been set.

### SetAutoRenewPeriodNil

`func (o *Contract) SetAutoRenewPeriodNil(b bool)`

 SetAutoRenewPeriodNil sets the value for AutoRenewPeriod to be an explicit nil

### UnsetAutoRenewPeriod
`func (o *Contract) UnsetAutoRenewPeriod()`

UnsetAutoRenewPeriod ensures that no value is present for AutoRenewPeriod, not even an explicit nil
### GetContractId

`func (o *Contract) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *Contract) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *Contract) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *Contract) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### SetContractIdNil

`func (o *Contract) SetContractIdNil(b bool)`

 SetContractIdNil sets the value for ContractId to be an explicit nil

### UnsetContractId
`func (o *Contract) UnsetContractId()`

UnsetContractId ensures that no value is present for ContractId, not even an explicit nil
### GetCreatedTimestamp

`func (o *Contract) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Contract) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Contract) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *Contract) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### SetCreatedTimestampNil

`func (o *Contract) SetCreatedTimestampNil(b bool)`

 SetCreatedTimestampNil sets the value for CreatedTimestamp to be an explicit nil

### UnsetCreatedTimestamp
`func (o *Contract) UnsetCreatedTimestamp()`

UnsetCreatedTimestamp ensures that no value is present for CreatedTimestamp, not even an explicit nil
### GetDeleted

`func (o *Contract) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Contract) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Contract) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Contract) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEvmAddress

`func (o *Contract) GetEvmAddress() *os.File`

GetEvmAddress returns the EvmAddress field if non-nil, zero value otherwise.

### GetEvmAddressOk

`func (o *Contract) GetEvmAddressOk() (**os.File, bool)`

GetEvmAddressOk returns a tuple with the EvmAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvmAddress

`func (o *Contract) SetEvmAddress(v *os.File)`

SetEvmAddress sets EvmAddress field to given value.

### HasEvmAddress

`func (o *Contract) HasEvmAddress() bool`

HasEvmAddress returns a boolean if a field has been set.

### GetExpirationTimestamp

`func (o *Contract) GetExpirationTimestamp() string`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *Contract) GetExpirationTimestampOk() (*string, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *Contract) SetExpirationTimestamp(v string)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *Contract) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### SetExpirationTimestampNil

`func (o *Contract) SetExpirationTimestampNil(b bool)`

 SetExpirationTimestampNil sets the value for ExpirationTimestamp to be an explicit nil

### UnsetExpirationTimestamp
`func (o *Contract) UnsetExpirationTimestamp()`

UnsetExpirationTimestamp ensures that no value is present for ExpirationTimestamp, not even an explicit nil
### GetFileId

`func (o *Contract) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *Contract) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *Contract) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *Contract) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *Contract) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *Contract) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetMaxAutomaticTokenAssociations

`func (o *Contract) GetMaxAutomaticTokenAssociations() int32`

GetMaxAutomaticTokenAssociations returns the MaxAutomaticTokenAssociations field if non-nil, zero value otherwise.

### GetMaxAutomaticTokenAssociationsOk

`func (o *Contract) GetMaxAutomaticTokenAssociationsOk() (*int32, bool)`

GetMaxAutomaticTokenAssociationsOk returns a tuple with the MaxAutomaticTokenAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAutomaticTokenAssociations

`func (o *Contract) SetMaxAutomaticTokenAssociations(v int32)`

SetMaxAutomaticTokenAssociations sets MaxAutomaticTokenAssociations field to given value.

### HasMaxAutomaticTokenAssociations

`func (o *Contract) HasMaxAutomaticTokenAssociations() bool`

HasMaxAutomaticTokenAssociations returns a boolean if a field has been set.

### SetMaxAutomaticTokenAssociationsNil

`func (o *Contract) SetMaxAutomaticTokenAssociationsNil(b bool)`

 SetMaxAutomaticTokenAssociationsNil sets the value for MaxAutomaticTokenAssociations to be an explicit nil

### UnsetMaxAutomaticTokenAssociations
`func (o *Contract) UnsetMaxAutomaticTokenAssociations()`

UnsetMaxAutomaticTokenAssociations ensures that no value is present for MaxAutomaticTokenAssociations, not even an explicit nil
### GetMemo

`func (o *Contract) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *Contract) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *Contract) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *Contract) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetObtainerId

`func (o *Contract) GetObtainerId() string`

GetObtainerId returns the ObtainerId field if non-nil, zero value otherwise.

### GetObtainerIdOk

`func (o *Contract) GetObtainerIdOk() (*string, bool)`

GetObtainerIdOk returns a tuple with the ObtainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObtainerId

`func (o *Contract) SetObtainerId(v string)`

SetObtainerId sets ObtainerId field to given value.

### HasObtainerId

`func (o *Contract) HasObtainerId() bool`

HasObtainerId returns a boolean if a field has been set.

### SetObtainerIdNil

`func (o *Contract) SetObtainerIdNil(b bool)`

 SetObtainerIdNil sets the value for ObtainerId to be an explicit nil

### UnsetObtainerId
`func (o *Contract) UnsetObtainerId()`

UnsetObtainerId ensures that no value is present for ObtainerId, not even an explicit nil
### GetPermanentRemoval

`func (o *Contract) GetPermanentRemoval() bool`

GetPermanentRemoval returns the PermanentRemoval field if non-nil, zero value otherwise.

### GetPermanentRemovalOk

`func (o *Contract) GetPermanentRemovalOk() (*bool, bool)`

GetPermanentRemovalOk returns a tuple with the PermanentRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentRemoval

`func (o *Contract) SetPermanentRemoval(v bool)`

SetPermanentRemoval sets PermanentRemoval field to given value.

### HasPermanentRemoval

`func (o *Contract) HasPermanentRemoval() bool`

HasPermanentRemoval returns a boolean if a field has been set.

### SetPermanentRemovalNil

`func (o *Contract) SetPermanentRemovalNil(b bool)`

 SetPermanentRemovalNil sets the value for PermanentRemoval to be an explicit nil

### UnsetPermanentRemoval
`func (o *Contract) UnsetPermanentRemoval()`

UnsetPermanentRemoval ensures that no value is present for PermanentRemoval, not even an explicit nil
### GetProxyAccountId

`func (o *Contract) GetProxyAccountId() string`

GetProxyAccountId returns the ProxyAccountId field if non-nil, zero value otherwise.

### GetProxyAccountIdOk

`func (o *Contract) GetProxyAccountIdOk() (*string, bool)`

GetProxyAccountIdOk returns a tuple with the ProxyAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAccountId

`func (o *Contract) SetProxyAccountId(v string)`

SetProxyAccountId sets ProxyAccountId field to given value.

### HasProxyAccountId

`func (o *Contract) HasProxyAccountId() bool`

HasProxyAccountId returns a boolean if a field has been set.

### SetProxyAccountIdNil

`func (o *Contract) SetProxyAccountIdNil(b bool)`

 SetProxyAccountIdNil sets the value for ProxyAccountId to be an explicit nil

### UnsetProxyAccountId
`func (o *Contract) UnsetProxyAccountId()`

UnsetProxyAccountId ensures that no value is present for ProxyAccountId, not even an explicit nil
### GetTimestamp

`func (o *Contract) GetTimestamp() TimestampRange`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Contract) GetTimestampOk() (*TimestampRange, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Contract) SetTimestamp(v TimestampRange)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Contract) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


