# PgwInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**PgwFqdn** | **string** | Fully Qualified Domain Name | 
**PgwIpAddr** | Pointer to [**NullableIpAddress**](IpAddress.md) |  | [optional] 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**EpdgInd** | Pointer to **bool** |  | [optional] [default to false]
**PcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**RegistrationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**WildcardInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewPgwInfo

`func NewPgwInfo(dnn string, pgwFqdn string, ) *PgwInfo`

NewPgwInfo instantiates a new PgwInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPgwInfoWithDefaults

`func NewPgwInfoWithDefaults() *PgwInfo`

NewPgwInfoWithDefaults instantiates a new PgwInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *PgwInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PgwInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PgwInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetPgwFqdn

`func (o *PgwInfo) GetPgwFqdn() string`

GetPgwFqdn returns the PgwFqdn field if non-nil, zero value otherwise.

### GetPgwFqdnOk

`func (o *PgwInfo) GetPgwFqdnOk() (*string, bool)`

GetPgwFqdnOk returns a tuple with the PgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdn

`func (o *PgwInfo) SetPgwFqdn(v string)`

SetPgwFqdn sets PgwFqdn field to given value.


### GetPgwIpAddr

`func (o *PgwInfo) GetPgwIpAddr() IpAddress`

GetPgwIpAddr returns the PgwIpAddr field if non-nil, zero value otherwise.

### GetPgwIpAddrOk

`func (o *PgwInfo) GetPgwIpAddrOk() (*IpAddress, bool)`

GetPgwIpAddrOk returns a tuple with the PgwIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwIpAddr

`func (o *PgwInfo) SetPgwIpAddr(v IpAddress)`

SetPgwIpAddr sets PgwIpAddr field to given value.

### HasPgwIpAddr

`func (o *PgwInfo) HasPgwIpAddr() bool`

HasPgwIpAddr returns a boolean if a field has been set.

### SetPgwIpAddrNil

`func (o *PgwInfo) SetPgwIpAddrNil(b bool)`

 SetPgwIpAddrNil sets the value for PgwIpAddr to be an explicit nil

### UnsetPgwIpAddr
`func (o *PgwInfo) UnsetPgwIpAddr()`

UnsetPgwIpAddr ensures that no value is present for PgwIpAddr, not even an explicit nil
### GetPlmnId

`func (o *PgwInfo) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *PgwInfo) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *PgwInfo) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *PgwInfo) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetEpdgInd

`func (o *PgwInfo) GetEpdgInd() bool`

GetEpdgInd returns the EpdgInd field if non-nil, zero value otherwise.

### GetEpdgIndOk

`func (o *PgwInfo) GetEpdgIndOk() (*bool, bool)`

GetEpdgIndOk returns a tuple with the EpdgInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpdgInd

`func (o *PgwInfo) SetEpdgInd(v bool)`

SetEpdgInd sets EpdgInd field to given value.

### HasEpdgInd

`func (o *PgwInfo) HasEpdgInd() bool`

HasEpdgInd returns a boolean if a field has been set.

### GetPcfId

`func (o *PgwInfo) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *PgwInfo) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *PgwInfo) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *PgwInfo) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.

### GetRegistrationTime

`func (o *PgwInfo) GetRegistrationTime() time.Time`

GetRegistrationTime returns the RegistrationTime field if non-nil, zero value otherwise.

### GetRegistrationTimeOk

`func (o *PgwInfo) GetRegistrationTimeOk() (*time.Time, bool)`

GetRegistrationTimeOk returns a tuple with the RegistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTime

`func (o *PgwInfo) SetRegistrationTime(v time.Time)`

SetRegistrationTime sets RegistrationTime field to given value.

### HasRegistrationTime

`func (o *PgwInfo) HasRegistrationTime() bool`

HasRegistrationTime returns a boolean if a field has been set.

### GetWildcardInd

`func (o *PgwInfo) GetWildcardInd() bool`

GetWildcardInd returns the WildcardInd field if non-nil, zero value otherwise.

### GetWildcardIndOk

`func (o *PgwInfo) GetWildcardIndOk() (*bool, bool)`

GetWildcardIndOk returns a tuple with the WildcardInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardInd

`func (o *PgwInfo) SetWildcardInd(v bool)`

SetWildcardInd sets WildcardInd field to given value.

### HasWildcardInd

`func (o *PgwInfo) HasWildcardInd() bool`

HasWildcardInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


