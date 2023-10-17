# NetworkFeesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fees** | Pointer to [**[]NetworkFee**](NetworkFee.md) |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkFeesResponse

`func NewNetworkFeesResponse() *NetworkFeesResponse`

NewNetworkFeesResponse instantiates a new NetworkFeesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFeesResponseWithDefaults

`func NewNetworkFeesResponseWithDefaults() *NetworkFeesResponse`

NewNetworkFeesResponseWithDefaults instantiates a new NetworkFeesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFees

`func (o *NetworkFeesResponse) GetFees() []NetworkFee`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *NetworkFeesResponse) GetFeesOk() (*[]NetworkFee, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *NetworkFeesResponse) SetFees(v []NetworkFee)`

SetFees sets Fees field to given value.

### HasFees

`func (o *NetworkFeesResponse) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetTimestamp

`func (o *NetworkFeesResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NetworkFeesResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NetworkFeesResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *NetworkFeesResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


