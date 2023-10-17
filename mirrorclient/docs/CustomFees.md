# CustomFees

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTimestamp** | Pointer to **string** |  | [optional] 
**FixedFees** | Pointer to [**[]FixedFee**](FixedFee.md) |  | [optional] 
**FractionalFees** | Pointer to [**[]FractionalFee**](FractionalFee.md) |  | [optional] 
**RoyaltyFees** | Pointer to [**[]RoyaltyFee**](RoyaltyFee.md) |  | [optional] 

## Methods

### NewCustomFees

`func NewCustomFees() *CustomFees`

NewCustomFees instantiates a new CustomFees object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFeesWithDefaults

`func NewCustomFeesWithDefaults() *CustomFees`

NewCustomFeesWithDefaults instantiates a new CustomFees object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTimestamp

`func (o *CustomFees) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *CustomFees) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *CustomFees) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *CustomFees) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetFixedFees

`func (o *CustomFees) GetFixedFees() []FixedFee`

GetFixedFees returns the FixedFees field if non-nil, zero value otherwise.

### GetFixedFeesOk

`func (o *CustomFees) GetFixedFeesOk() (*[]FixedFee, bool)`

GetFixedFeesOk returns a tuple with the FixedFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedFees

`func (o *CustomFees) SetFixedFees(v []FixedFee)`

SetFixedFees sets FixedFees field to given value.

### HasFixedFees

`func (o *CustomFees) HasFixedFees() bool`

HasFixedFees returns a boolean if a field has been set.

### GetFractionalFees

`func (o *CustomFees) GetFractionalFees() []FractionalFee`

GetFractionalFees returns the FractionalFees field if non-nil, zero value otherwise.

### GetFractionalFeesOk

`func (o *CustomFees) GetFractionalFeesOk() (*[]FractionalFee, bool)`

GetFractionalFeesOk returns a tuple with the FractionalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFractionalFees

`func (o *CustomFees) SetFractionalFees(v []FractionalFee)`

SetFractionalFees sets FractionalFees field to given value.

### HasFractionalFees

`func (o *CustomFees) HasFractionalFees() bool`

HasFractionalFees returns a boolean if a field has been set.

### GetRoyaltyFees

`func (o *CustomFees) GetRoyaltyFees() []RoyaltyFee`

GetRoyaltyFees returns the RoyaltyFees field if non-nil, zero value otherwise.

### GetRoyaltyFeesOk

`func (o *CustomFees) GetRoyaltyFeesOk() (*[]RoyaltyFee, bool)`

GetRoyaltyFeesOk returns a tuple with the RoyaltyFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoyaltyFees

`func (o *CustomFees) SetRoyaltyFees(v []RoyaltyFee)`

SetRoyaltyFees sets RoyaltyFees field to given value.

### HasRoyaltyFees

`func (o *CustomFees) HasRoyaltyFees() bool`

HasRoyaltyFees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


