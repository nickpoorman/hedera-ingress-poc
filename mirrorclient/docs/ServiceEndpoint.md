# ServiceEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddressV4** | **string** |  | 
**Port** | **int32** |  | 

## Methods

### NewServiceEndpoint

`func NewServiceEndpoint(ipAddressV4 string, port int32, ) *ServiceEndpoint`

NewServiceEndpoint instantiates a new ServiceEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceEndpointWithDefaults

`func NewServiceEndpointWithDefaults() *ServiceEndpoint`

NewServiceEndpointWithDefaults instantiates a new ServiceEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddressV4

`func (o *ServiceEndpoint) GetIpAddressV4() string`

GetIpAddressV4 returns the IpAddressV4 field if non-nil, zero value otherwise.

### GetIpAddressV4Ok

`func (o *ServiceEndpoint) GetIpAddressV4Ok() (*string, bool)`

GetIpAddressV4Ok returns a tuple with the IpAddressV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressV4

`func (o *ServiceEndpoint) SetIpAddressV4(v string)`

SetIpAddressV4 sets IpAddressV4 field to given value.


### GetPort

`func (o *ServiceEndpoint) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ServiceEndpoint) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ServiceEndpoint) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


