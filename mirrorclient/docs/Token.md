# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**Symbol** | **string** |  | 
**AdminKey** | [**NullableKey**](Key.md) |  | 
**Type** | **string** |  | 

## Methods

### NewToken

`func NewToken(tokenId NullableString, symbol string, adminKey NullableKey, type_ string, ) *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *Token) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Token) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Token) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### SetTokenIdNil

`func (o *Token) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *Token) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
### GetSymbol

`func (o *Token) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Token) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Token) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetAdminKey

`func (o *Token) GetAdminKey() Key`

GetAdminKey returns the AdminKey field if non-nil, zero value otherwise.

### GetAdminKeyOk

`func (o *Token) GetAdminKeyOk() (*Key, bool)`

GetAdminKeyOk returns a tuple with the AdminKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminKey

`func (o *Token) SetAdminKey(v Key)`

SetAdminKey sets AdminKey field to given value.


### SetAdminKeyNil

`func (o *Token) SetAdminKeyNil(b bool)`

 SetAdminKeyNil sets the value for AdminKey to be an explicit nil

### UnsetAdminKey
`func (o *Token) UnsetAdminKey()`

UnsetAdminKey ensures that no value is present for AdminKey, not even an explicit nil
### GetType

`func (o *Token) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Token) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Token) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


