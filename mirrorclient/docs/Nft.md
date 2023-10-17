# Nft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**CreatedTimestamp** | Pointer to **NullableString** |  | [optional] 
**DelegatingSpender** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Deleted** | Pointer to **bool** | whether the nft or the token it belongs to has been deleted | [optional] 
**Metadata** | Pointer to **string** | base64 encoded binary data | [optional] 
**ModifiedTimestamp** | Pointer to **NullableString** |  | [optional] 
**SerialNumber** | Pointer to **int64** |  | [optional] 
**Spender** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**TokenId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 

## Methods

### NewNft

`func NewNft() *Nft`

NewNft instantiates a new Nft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftWithDefaults

`func NewNftWithDefaults() *Nft`

NewNftWithDefaults instantiates a new Nft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Nft) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Nft) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Nft) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Nft) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *Nft) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *Nft) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetCreatedTimestamp

`func (o *Nft) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Nft) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Nft) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *Nft) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### SetCreatedTimestampNil

`func (o *Nft) SetCreatedTimestampNil(b bool)`

 SetCreatedTimestampNil sets the value for CreatedTimestamp to be an explicit nil

### UnsetCreatedTimestamp
`func (o *Nft) UnsetCreatedTimestamp()`

UnsetCreatedTimestamp ensures that no value is present for CreatedTimestamp, not even an explicit nil
### GetDelegatingSpender

`func (o *Nft) GetDelegatingSpender() string`

GetDelegatingSpender returns the DelegatingSpender field if non-nil, zero value otherwise.

### GetDelegatingSpenderOk

`func (o *Nft) GetDelegatingSpenderOk() (*string, bool)`

GetDelegatingSpenderOk returns a tuple with the DelegatingSpender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatingSpender

`func (o *Nft) SetDelegatingSpender(v string)`

SetDelegatingSpender sets DelegatingSpender field to given value.

### HasDelegatingSpender

`func (o *Nft) HasDelegatingSpender() bool`

HasDelegatingSpender returns a boolean if a field has been set.

### SetDelegatingSpenderNil

`func (o *Nft) SetDelegatingSpenderNil(b bool)`

 SetDelegatingSpenderNil sets the value for DelegatingSpender to be an explicit nil

### UnsetDelegatingSpender
`func (o *Nft) UnsetDelegatingSpender()`

UnsetDelegatingSpender ensures that no value is present for DelegatingSpender, not even an explicit nil
### GetDeleted

`func (o *Nft) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Nft) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Nft) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Nft) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetMetadata

`func (o *Nft) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Nft) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Nft) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Nft) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModifiedTimestamp

`func (o *Nft) GetModifiedTimestamp() string`

GetModifiedTimestamp returns the ModifiedTimestamp field if non-nil, zero value otherwise.

### GetModifiedTimestampOk

`func (o *Nft) GetModifiedTimestampOk() (*string, bool)`

GetModifiedTimestampOk returns a tuple with the ModifiedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTimestamp

`func (o *Nft) SetModifiedTimestamp(v string)`

SetModifiedTimestamp sets ModifiedTimestamp field to given value.

### HasModifiedTimestamp

`func (o *Nft) HasModifiedTimestamp() bool`

HasModifiedTimestamp returns a boolean if a field has been set.

### SetModifiedTimestampNil

`func (o *Nft) SetModifiedTimestampNil(b bool)`

 SetModifiedTimestampNil sets the value for ModifiedTimestamp to be an explicit nil

### UnsetModifiedTimestamp
`func (o *Nft) UnsetModifiedTimestamp()`

UnsetModifiedTimestamp ensures that no value is present for ModifiedTimestamp, not even an explicit nil
### GetSerialNumber

`func (o *Nft) GetSerialNumber() int64`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Nft) GetSerialNumberOk() (*int64, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Nft) SetSerialNumber(v int64)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Nft) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSpender

`func (o *Nft) GetSpender() string`

GetSpender returns the Spender field if non-nil, zero value otherwise.

### GetSpenderOk

`func (o *Nft) GetSpenderOk() (*string, bool)`

GetSpenderOk returns a tuple with the Spender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpender

`func (o *Nft) SetSpender(v string)`

SetSpender sets Spender field to given value.

### HasSpender

`func (o *Nft) HasSpender() bool`

HasSpender returns a boolean if a field has been set.

### SetSpenderNil

`func (o *Nft) SetSpenderNil(b bool)`

 SetSpenderNil sets the value for Spender to be an explicit nil

### UnsetSpender
`func (o *Nft) UnsetSpender()`

UnsetSpender ensures that no value is present for Spender, not even an explicit nil
### GetTokenId

`func (o *Nft) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Nft) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Nft) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *Nft) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### SetTokenIdNil

`func (o *Nft) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *Nft) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


