# TokenRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticAssociation** | **bool** | Specifies if the relationship is implicitly/explicitly associated. | 
**Balance** | **int64** | For FUNGIBLE_COMMON, the balance that the account holds in the smallest denomination. For NON_FUNGIBLE_UNIQUE, the number of NFTs held by the account. | 
**CreatedTimestamp** | **string** |  | 
**FreezeStatus** | **string** | The Freeze status of the account. | 
**KycStatus** | **string** | The KYC status of the account. | 
**TokenId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 

## Methods

### NewTokenRelationship

`func NewTokenRelationship(automaticAssociation bool, balance int64, createdTimestamp string, freezeStatus string, kycStatus string, tokenId NullableString, ) *TokenRelationship`

NewTokenRelationship instantiates a new TokenRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRelationshipWithDefaults

`func NewTokenRelationshipWithDefaults() *TokenRelationship`

NewTokenRelationshipWithDefaults instantiates a new TokenRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticAssociation

`func (o *TokenRelationship) GetAutomaticAssociation() bool`

GetAutomaticAssociation returns the AutomaticAssociation field if non-nil, zero value otherwise.

### GetAutomaticAssociationOk

`func (o *TokenRelationship) GetAutomaticAssociationOk() (*bool, bool)`

GetAutomaticAssociationOk returns a tuple with the AutomaticAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticAssociation

`func (o *TokenRelationship) SetAutomaticAssociation(v bool)`

SetAutomaticAssociation sets AutomaticAssociation field to given value.


### GetBalance

`func (o *TokenRelationship) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *TokenRelationship) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *TokenRelationship) SetBalance(v int64)`

SetBalance sets Balance field to given value.


### GetCreatedTimestamp

`func (o *TokenRelationship) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *TokenRelationship) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *TokenRelationship) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetFreezeStatus

`func (o *TokenRelationship) GetFreezeStatus() string`

GetFreezeStatus returns the FreezeStatus field if non-nil, zero value otherwise.

### GetFreezeStatusOk

`func (o *TokenRelationship) GetFreezeStatusOk() (*string, bool)`

GetFreezeStatusOk returns a tuple with the FreezeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezeStatus

`func (o *TokenRelationship) SetFreezeStatus(v string)`

SetFreezeStatus sets FreezeStatus field to given value.


### GetKycStatus

`func (o *TokenRelationship) GetKycStatus() string`

GetKycStatus returns the KycStatus field if non-nil, zero value otherwise.

### GetKycStatusOk

`func (o *TokenRelationship) GetKycStatusOk() (*string, bool)`

GetKycStatusOk returns a tuple with the KycStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycStatus

`func (o *TokenRelationship) SetKycStatus(v string)`

SetKycStatus sets KycStatus field to given value.


### GetTokenId

`func (o *TokenRelationship) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenRelationship) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenRelationship) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### SetTokenIdNil

`func (o *TokenRelationship) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *TokenRelationship) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


