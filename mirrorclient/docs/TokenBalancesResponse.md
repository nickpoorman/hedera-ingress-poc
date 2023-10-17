# TokenBalancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **NullableString** |  | [optional] 
**Balances** | Pointer to [**[]TokenDistributionInner**](TokenDistributionInner.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewTokenBalancesResponse

`func NewTokenBalancesResponse() *TokenBalancesResponse`

NewTokenBalancesResponse instantiates a new TokenBalancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenBalancesResponseWithDefaults

`func NewTokenBalancesResponseWithDefaults() *TokenBalancesResponse`

NewTokenBalancesResponseWithDefaults instantiates a new TokenBalancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *TokenBalancesResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TokenBalancesResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TokenBalancesResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TokenBalancesResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TokenBalancesResponse) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TokenBalancesResponse) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetBalances

`func (o *TokenBalancesResponse) GetBalances() []TokenDistributionInner`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *TokenBalancesResponse) GetBalancesOk() (*[]TokenDistributionInner, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *TokenBalancesResponse) SetBalances(v []TokenDistributionInner)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *TokenBalancesResponse) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetLinks

`func (o *TokenBalancesResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TokenBalancesResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TokenBalancesResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TokenBalancesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


