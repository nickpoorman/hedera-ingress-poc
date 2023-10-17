# NetworkExchangeRateSetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentRate** | Pointer to [**ExchangeRate**](ExchangeRate.md) |  | [optional] 
**NextRate** | Pointer to [**ExchangeRate**](ExchangeRate.md) |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkExchangeRateSetResponse

`func NewNetworkExchangeRateSetResponse() *NetworkExchangeRateSetResponse`

NewNetworkExchangeRateSetResponse instantiates a new NetworkExchangeRateSetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkExchangeRateSetResponseWithDefaults

`func NewNetworkExchangeRateSetResponseWithDefaults() *NetworkExchangeRateSetResponse`

NewNetworkExchangeRateSetResponseWithDefaults instantiates a new NetworkExchangeRateSetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentRate

`func (o *NetworkExchangeRateSetResponse) GetCurrentRate() ExchangeRate`

GetCurrentRate returns the CurrentRate field if non-nil, zero value otherwise.

### GetCurrentRateOk

`func (o *NetworkExchangeRateSetResponse) GetCurrentRateOk() (*ExchangeRate, bool)`

GetCurrentRateOk returns a tuple with the CurrentRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRate

`func (o *NetworkExchangeRateSetResponse) SetCurrentRate(v ExchangeRate)`

SetCurrentRate sets CurrentRate field to given value.

### HasCurrentRate

`func (o *NetworkExchangeRateSetResponse) HasCurrentRate() bool`

HasCurrentRate returns a boolean if a field has been set.

### GetNextRate

`func (o *NetworkExchangeRateSetResponse) GetNextRate() ExchangeRate`

GetNextRate returns the NextRate field if non-nil, zero value otherwise.

### GetNextRateOk

`func (o *NetworkExchangeRateSetResponse) GetNextRateOk() (*ExchangeRate, bool)`

GetNextRateOk returns a tuple with the NextRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRate

`func (o *NetworkExchangeRateSetResponse) SetNextRate(v ExchangeRate)`

SetNextRate sets NextRate field to given value.

### HasNextRate

`func (o *NetworkExchangeRateSetResponse) HasNextRate() bool`

HasNextRate returns a boolean if a field has been set.

### GetTimestamp

`func (o *NetworkExchangeRateSetResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NetworkExchangeRateSetResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NetworkExchangeRateSetResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *NetworkExchangeRateSetResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


