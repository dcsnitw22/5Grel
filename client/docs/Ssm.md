# Ssm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceIpAddr** | [**NullableIpAddr**](IpAddr.md) |  | 
**DestIpAddr** | [**NullableIpAddr**](IpAddr.md) |  | 

## Methods

### NewSsm

`func NewSsm(sourceIpAddr NullableIpAddr, destIpAddr NullableIpAddr, ) *Ssm`

NewSsm instantiates a new Ssm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsmWithDefaults

`func NewSsmWithDefaults() *Ssm`

NewSsmWithDefaults instantiates a new Ssm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceIpAddr

`func (o *Ssm) GetSourceIpAddr() IpAddr`

GetSourceIpAddr returns the SourceIpAddr field if non-nil, zero value otherwise.

### GetSourceIpAddrOk

`func (o *Ssm) GetSourceIpAddrOk() (*IpAddr, bool)`

GetSourceIpAddrOk returns a tuple with the SourceIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpAddr

`func (o *Ssm) SetSourceIpAddr(v IpAddr)`

SetSourceIpAddr sets SourceIpAddr field to given value.


### SetSourceIpAddrNil

`func (o *Ssm) SetSourceIpAddrNil(b bool)`

 SetSourceIpAddrNil sets the value for SourceIpAddr to be an explicit nil

### UnsetSourceIpAddr
`func (o *Ssm) UnsetSourceIpAddr()`

UnsetSourceIpAddr ensures that no value is present for SourceIpAddr, not even an explicit nil
### GetDestIpAddr

`func (o *Ssm) GetDestIpAddr() IpAddr`

GetDestIpAddr returns the DestIpAddr field if non-nil, zero value otherwise.

### GetDestIpAddrOk

`func (o *Ssm) GetDestIpAddrOk() (*IpAddr, bool)`

GetDestIpAddrOk returns a tuple with the DestIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestIpAddr

`func (o *Ssm) SetDestIpAddr(v IpAddr)`

SetDestIpAddr sets DestIpAddr field to given value.


### SetDestIpAddrNil

`func (o *Ssm) SetDestIpAddrNil(b bool)`

 SetDestIpAddrNil sets the value for DestIpAddr to be an explicit nil

### UnsetDestIpAddr
`func (o *Ssm) UnsetDestIpAddr()`

UnsetDestIpAddr ensures that no value is present for DestIpAddr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


