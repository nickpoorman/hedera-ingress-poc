# TokenInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminKey** | Pointer to [**NullableKey**](Key.md) |  | [optional] 
**AutoRenewAccount** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**AutoRenewPeriod** | Pointer to **NullableInt64** |  | [optional] 
**CreatedTimestamp** | Pointer to **string** |  | [optional] 
**Decimals** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 
**ExpiryTimestamp** | Pointer to **NullableInt64** |  | [optional] 
**FeeScheduleKey** | Pointer to [**NullableKey**](Key.md) |  | [optional] 
**FreezeDefault** | Pointer to **bool** |  | [optional] 
**FreezeKey** | Pointer to [**NullableKey**](Key.md) |  | [optional] 
**InitialSupply** | Pointer to **string** |  | [optional] 
**KycKey** | Pointer to [**NullableKey**](Key.md) |  | [optional] 
**MaxSupply** | Pointer to **string** |  | [optional] 
**ModifiedTimestamp** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**PauseKey** | Pointer to [**NullableKey**](Key.md) |  | [optional] 
**PauseStatus** | Pointer to **string** |  | [optional] 
**SupplyKey** | Pointer to [**NullableKey**](Key.md) |  | [optional] 
**SupplyType** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TokenId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**TotalSupply** | Pointer to **string** |  | [optional] 
**TreasuryAccountId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**WipeKey** | Pointer to [**NullableKey**](Key.md) |  | [optional] 
**CustomFees** | Pointer to [**CustomFees**](CustomFees.md) |  | [optional] 

## Methods

### NewTokenInfo

`func NewTokenInfo() *TokenInfo`

NewTokenInfo instantiates a new TokenInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenInfoWithDefaults

`func NewTokenInfoWithDefaults() *TokenInfo`

NewTokenInfoWithDefaults instantiates a new TokenInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminKey

`func (o *TokenInfo) GetAdminKey() Key`

GetAdminKey returns the AdminKey field if non-nil, zero value otherwise.

### GetAdminKeyOk

`func (o *TokenInfo) GetAdminKeyOk() (*Key, bool)`

GetAdminKeyOk returns a tuple with the AdminKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminKey

`func (o *TokenInfo) SetAdminKey(v Key)`

SetAdminKey sets AdminKey field to given value.

### HasAdminKey

`func (o *TokenInfo) HasAdminKey() bool`

HasAdminKey returns a boolean if a field has been set.

### SetAdminKeyNil

`func (o *TokenInfo) SetAdminKeyNil(b bool)`

 SetAdminKeyNil sets the value for AdminKey to be an explicit nil

### UnsetAdminKey
`func (o *TokenInfo) UnsetAdminKey()`

UnsetAdminKey ensures that no value is present for AdminKey, not even an explicit nil
### GetAutoRenewAccount

`func (o *TokenInfo) GetAutoRenewAccount() string`

GetAutoRenewAccount returns the AutoRenewAccount field if non-nil, zero value otherwise.

### GetAutoRenewAccountOk

`func (o *TokenInfo) GetAutoRenewAccountOk() (*string, bool)`

GetAutoRenewAccountOk returns a tuple with the AutoRenewAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewAccount

`func (o *TokenInfo) SetAutoRenewAccount(v string)`

SetAutoRenewAccount sets AutoRenewAccount field to given value.

### HasAutoRenewAccount

`func (o *TokenInfo) HasAutoRenewAccount() bool`

HasAutoRenewAccount returns a boolean if a field has been set.

### SetAutoRenewAccountNil

`func (o *TokenInfo) SetAutoRenewAccountNil(b bool)`

 SetAutoRenewAccountNil sets the value for AutoRenewAccount to be an explicit nil

### UnsetAutoRenewAccount
`func (o *TokenInfo) UnsetAutoRenewAccount()`

UnsetAutoRenewAccount ensures that no value is present for AutoRenewAccount, not even an explicit nil
### GetAutoRenewPeriod

`func (o *TokenInfo) GetAutoRenewPeriod() int64`

GetAutoRenewPeriod returns the AutoRenewPeriod field if non-nil, zero value otherwise.

### GetAutoRenewPeriodOk

`func (o *TokenInfo) GetAutoRenewPeriodOk() (*int64, bool)`

GetAutoRenewPeriodOk returns a tuple with the AutoRenewPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewPeriod

`func (o *TokenInfo) SetAutoRenewPeriod(v int64)`

SetAutoRenewPeriod sets AutoRenewPeriod field to given value.

### HasAutoRenewPeriod

`func (o *TokenInfo) HasAutoRenewPeriod() bool`

HasAutoRenewPeriod returns a boolean if a field has been set.

### SetAutoRenewPeriodNil

`func (o *TokenInfo) SetAutoRenewPeriodNil(b bool)`

 SetAutoRenewPeriodNil sets the value for AutoRenewPeriod to be an explicit nil

### UnsetAutoRenewPeriod
`func (o *TokenInfo) UnsetAutoRenewPeriod()`

UnsetAutoRenewPeriod ensures that no value is present for AutoRenewPeriod, not even an explicit nil
### GetCreatedTimestamp

`func (o *TokenInfo) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *TokenInfo) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *TokenInfo) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *TokenInfo) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetDecimals

`func (o *TokenInfo) GetDecimals() string`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenInfo) GetDecimalsOk() (*string, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenInfo) SetDecimals(v string)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *TokenInfo) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetDeleted

`func (o *TokenInfo) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *TokenInfo) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *TokenInfo) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *TokenInfo) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *TokenInfo) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *TokenInfo) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetExpiryTimestamp

`func (o *TokenInfo) GetExpiryTimestamp() int64`

GetExpiryTimestamp returns the ExpiryTimestamp field if non-nil, zero value otherwise.

### GetExpiryTimestampOk

`func (o *TokenInfo) GetExpiryTimestampOk() (*int64, bool)`

GetExpiryTimestampOk returns a tuple with the ExpiryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimestamp

`func (o *TokenInfo) SetExpiryTimestamp(v int64)`

SetExpiryTimestamp sets ExpiryTimestamp field to given value.

### HasExpiryTimestamp

`func (o *TokenInfo) HasExpiryTimestamp() bool`

HasExpiryTimestamp returns a boolean if a field has been set.

### SetExpiryTimestampNil

`func (o *TokenInfo) SetExpiryTimestampNil(b bool)`

 SetExpiryTimestampNil sets the value for ExpiryTimestamp to be an explicit nil

### UnsetExpiryTimestamp
`func (o *TokenInfo) UnsetExpiryTimestamp()`

UnsetExpiryTimestamp ensures that no value is present for ExpiryTimestamp, not even an explicit nil
### GetFeeScheduleKey

`func (o *TokenInfo) GetFeeScheduleKey() Key`

GetFeeScheduleKey returns the FeeScheduleKey field if non-nil, zero value otherwise.

### GetFeeScheduleKeyOk

`func (o *TokenInfo) GetFeeScheduleKeyOk() (*Key, bool)`

GetFeeScheduleKeyOk returns a tuple with the FeeScheduleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeScheduleKey

`func (o *TokenInfo) SetFeeScheduleKey(v Key)`

SetFeeScheduleKey sets FeeScheduleKey field to given value.

### HasFeeScheduleKey

`func (o *TokenInfo) HasFeeScheduleKey() bool`

HasFeeScheduleKey returns a boolean if a field has been set.

### SetFeeScheduleKeyNil

`func (o *TokenInfo) SetFeeScheduleKeyNil(b bool)`

 SetFeeScheduleKeyNil sets the value for FeeScheduleKey to be an explicit nil

### UnsetFeeScheduleKey
`func (o *TokenInfo) UnsetFeeScheduleKey()`

UnsetFeeScheduleKey ensures that no value is present for FeeScheduleKey, not even an explicit nil
### GetFreezeDefault

`func (o *TokenInfo) GetFreezeDefault() bool`

GetFreezeDefault returns the FreezeDefault field if non-nil, zero value otherwise.

### GetFreezeDefaultOk

`func (o *TokenInfo) GetFreezeDefaultOk() (*bool, bool)`

GetFreezeDefaultOk returns a tuple with the FreezeDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezeDefault

`func (o *TokenInfo) SetFreezeDefault(v bool)`

SetFreezeDefault sets FreezeDefault field to given value.

### HasFreezeDefault

`func (o *TokenInfo) HasFreezeDefault() bool`

HasFreezeDefault returns a boolean if a field has been set.

### GetFreezeKey

`func (o *TokenInfo) GetFreezeKey() Key`

GetFreezeKey returns the FreezeKey field if non-nil, zero value otherwise.

### GetFreezeKeyOk

`func (o *TokenInfo) GetFreezeKeyOk() (*Key, bool)`

GetFreezeKeyOk returns a tuple with the FreezeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezeKey

`func (o *TokenInfo) SetFreezeKey(v Key)`

SetFreezeKey sets FreezeKey field to given value.

### HasFreezeKey

`func (o *TokenInfo) HasFreezeKey() bool`

HasFreezeKey returns a boolean if a field has been set.

### SetFreezeKeyNil

`func (o *TokenInfo) SetFreezeKeyNil(b bool)`

 SetFreezeKeyNil sets the value for FreezeKey to be an explicit nil

### UnsetFreezeKey
`func (o *TokenInfo) UnsetFreezeKey()`

UnsetFreezeKey ensures that no value is present for FreezeKey, not even an explicit nil
### GetInitialSupply

`func (o *TokenInfo) GetInitialSupply() string`

GetInitialSupply returns the InitialSupply field if non-nil, zero value otherwise.

### GetInitialSupplyOk

`func (o *TokenInfo) GetInitialSupplyOk() (*string, bool)`

GetInitialSupplyOk returns a tuple with the InitialSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSupply

`func (o *TokenInfo) SetInitialSupply(v string)`

SetInitialSupply sets InitialSupply field to given value.

### HasInitialSupply

`func (o *TokenInfo) HasInitialSupply() bool`

HasInitialSupply returns a boolean if a field has been set.

### GetKycKey

`func (o *TokenInfo) GetKycKey() Key`

GetKycKey returns the KycKey field if non-nil, zero value otherwise.

### GetKycKeyOk

`func (o *TokenInfo) GetKycKeyOk() (*Key, bool)`

GetKycKeyOk returns a tuple with the KycKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycKey

`func (o *TokenInfo) SetKycKey(v Key)`

SetKycKey sets KycKey field to given value.

### HasKycKey

`func (o *TokenInfo) HasKycKey() bool`

HasKycKey returns a boolean if a field has been set.

### SetKycKeyNil

`func (o *TokenInfo) SetKycKeyNil(b bool)`

 SetKycKeyNil sets the value for KycKey to be an explicit nil

### UnsetKycKey
`func (o *TokenInfo) UnsetKycKey()`

UnsetKycKey ensures that no value is present for KycKey, not even an explicit nil
### GetMaxSupply

`func (o *TokenInfo) GetMaxSupply() string`

GetMaxSupply returns the MaxSupply field if non-nil, zero value otherwise.

### GetMaxSupplyOk

`func (o *TokenInfo) GetMaxSupplyOk() (*string, bool)`

GetMaxSupplyOk returns a tuple with the MaxSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSupply

`func (o *TokenInfo) SetMaxSupply(v string)`

SetMaxSupply sets MaxSupply field to given value.

### HasMaxSupply

`func (o *TokenInfo) HasMaxSupply() bool`

HasMaxSupply returns a boolean if a field has been set.

### GetModifiedTimestamp

`func (o *TokenInfo) GetModifiedTimestamp() string`

GetModifiedTimestamp returns the ModifiedTimestamp field if non-nil, zero value otherwise.

### GetModifiedTimestampOk

`func (o *TokenInfo) GetModifiedTimestampOk() (*string, bool)`

GetModifiedTimestampOk returns a tuple with the ModifiedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTimestamp

`func (o *TokenInfo) SetModifiedTimestamp(v string)`

SetModifiedTimestamp sets ModifiedTimestamp field to given value.

### HasModifiedTimestamp

`func (o *TokenInfo) HasModifiedTimestamp() bool`

HasModifiedTimestamp returns a boolean if a field has been set.

### GetName

`func (o *TokenInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMemo

`func (o *TokenInfo) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TokenInfo) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TokenInfo) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TokenInfo) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPauseKey

`func (o *TokenInfo) GetPauseKey() Key`

GetPauseKey returns the PauseKey field if non-nil, zero value otherwise.

### GetPauseKeyOk

`func (o *TokenInfo) GetPauseKeyOk() (*Key, bool)`

GetPauseKeyOk returns a tuple with the PauseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseKey

`func (o *TokenInfo) SetPauseKey(v Key)`

SetPauseKey sets PauseKey field to given value.

### HasPauseKey

`func (o *TokenInfo) HasPauseKey() bool`

HasPauseKey returns a boolean if a field has been set.

### SetPauseKeyNil

`func (o *TokenInfo) SetPauseKeyNil(b bool)`

 SetPauseKeyNil sets the value for PauseKey to be an explicit nil

### UnsetPauseKey
`func (o *TokenInfo) UnsetPauseKey()`

UnsetPauseKey ensures that no value is present for PauseKey, not even an explicit nil
### GetPauseStatus

`func (o *TokenInfo) GetPauseStatus() string`

GetPauseStatus returns the PauseStatus field if non-nil, zero value otherwise.

### GetPauseStatusOk

`func (o *TokenInfo) GetPauseStatusOk() (*string, bool)`

GetPauseStatusOk returns a tuple with the PauseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseStatus

`func (o *TokenInfo) SetPauseStatus(v string)`

SetPauseStatus sets PauseStatus field to given value.

### HasPauseStatus

`func (o *TokenInfo) HasPauseStatus() bool`

HasPauseStatus returns a boolean if a field has been set.

### GetSupplyKey

`func (o *TokenInfo) GetSupplyKey() Key`

GetSupplyKey returns the SupplyKey field if non-nil, zero value otherwise.

### GetSupplyKeyOk

`func (o *TokenInfo) GetSupplyKeyOk() (*Key, bool)`

GetSupplyKeyOk returns a tuple with the SupplyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplyKey

`func (o *TokenInfo) SetSupplyKey(v Key)`

SetSupplyKey sets SupplyKey field to given value.

### HasSupplyKey

`func (o *TokenInfo) HasSupplyKey() bool`

HasSupplyKey returns a boolean if a field has been set.

### SetSupplyKeyNil

`func (o *TokenInfo) SetSupplyKeyNil(b bool)`

 SetSupplyKeyNil sets the value for SupplyKey to be an explicit nil

### UnsetSupplyKey
`func (o *TokenInfo) UnsetSupplyKey()`

UnsetSupplyKey ensures that no value is present for SupplyKey, not even an explicit nil
### GetSupplyType

`func (o *TokenInfo) GetSupplyType() string`

GetSupplyType returns the SupplyType field if non-nil, zero value otherwise.

### GetSupplyTypeOk

`func (o *TokenInfo) GetSupplyTypeOk() (*string, bool)`

GetSupplyTypeOk returns a tuple with the SupplyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplyType

`func (o *TokenInfo) SetSupplyType(v string)`

SetSupplyType sets SupplyType field to given value.

### HasSupplyType

`func (o *TokenInfo) HasSupplyType() bool`

HasSupplyType returns a boolean if a field has been set.

### GetSymbol

`func (o *TokenInfo) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenInfo) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenInfo) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TokenInfo) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTokenId

`func (o *TokenInfo) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenInfo) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenInfo) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TokenInfo) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### SetTokenIdNil

`func (o *TokenInfo) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *TokenInfo) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
### GetTotalSupply

`func (o *TokenInfo) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *TokenInfo) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *TokenInfo) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.

### HasTotalSupply

`func (o *TokenInfo) HasTotalSupply() bool`

HasTotalSupply returns a boolean if a field has been set.

### GetTreasuryAccountId

`func (o *TokenInfo) GetTreasuryAccountId() string`

GetTreasuryAccountId returns the TreasuryAccountId field if non-nil, zero value otherwise.

### GetTreasuryAccountIdOk

`func (o *TokenInfo) GetTreasuryAccountIdOk() (*string, bool)`

GetTreasuryAccountIdOk returns a tuple with the TreasuryAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreasuryAccountId

`func (o *TokenInfo) SetTreasuryAccountId(v string)`

SetTreasuryAccountId sets TreasuryAccountId field to given value.

### HasTreasuryAccountId

`func (o *TokenInfo) HasTreasuryAccountId() bool`

HasTreasuryAccountId returns a boolean if a field has been set.

### SetTreasuryAccountIdNil

`func (o *TokenInfo) SetTreasuryAccountIdNil(b bool)`

 SetTreasuryAccountIdNil sets the value for TreasuryAccountId to be an explicit nil

### UnsetTreasuryAccountId
`func (o *TokenInfo) UnsetTreasuryAccountId()`

UnsetTreasuryAccountId ensures that no value is present for TreasuryAccountId, not even an explicit nil
### GetType

`func (o *TokenInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TokenInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWipeKey

`func (o *TokenInfo) GetWipeKey() Key`

GetWipeKey returns the WipeKey field if non-nil, zero value otherwise.

### GetWipeKeyOk

`func (o *TokenInfo) GetWipeKeyOk() (*Key, bool)`

GetWipeKeyOk returns a tuple with the WipeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWipeKey

`func (o *TokenInfo) SetWipeKey(v Key)`

SetWipeKey sets WipeKey field to given value.

### HasWipeKey

`func (o *TokenInfo) HasWipeKey() bool`

HasWipeKey returns a boolean if a field has been set.

### SetWipeKeyNil

`func (o *TokenInfo) SetWipeKeyNil(b bool)`

 SetWipeKeyNil sets the value for WipeKey to be an explicit nil

### UnsetWipeKey
`func (o *TokenInfo) UnsetWipeKey()`

UnsetWipeKey ensures that no value is present for WipeKey, not even an explicit nil
### GetCustomFees

`func (o *TokenInfo) GetCustomFees() CustomFees`

GetCustomFees returns the CustomFees field if non-nil, zero value otherwise.

### GetCustomFeesOk

`func (o *TokenInfo) GetCustomFeesOk() (*CustomFees, bool)`

GetCustomFeesOk returns a tuple with the CustomFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFees

`func (o *TokenInfo) SetCustomFees(v CustomFees)`

SetCustomFees sets CustomFees field to given value.

### HasCustomFees

`func (o *TokenInfo) HasCustomFees() bool`

HasCustomFees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


