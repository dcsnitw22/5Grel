# EpsIwkPgw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PgwFqdn** | **string** | Fully Qualified Domain Name | 
**SmfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 

## Methods

### NewEpsIwkPgw

`func NewEpsIwkPgw(pgwFqdn string, smfInstanceId string, ) *EpsIwkPgw`

NewEpsIwkPgw instantiates a new EpsIwkPgw object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpsIwkPgwWithDefaults

`func NewEpsIwkPgwWithDefaults() *EpsIwkPgw`

NewEpsIwkPgwWithDefaults instantiates a new EpsIwkPgw object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPgwFqdn

`func (o *EpsIwkPgw) GetPgwFqdn() string`

GetPgwFqdn returns the PgwFqdn field if non-nil, zero value otherwise.

### GetPgwFqdnOk

`func (o *EpsIwkPgw) GetPgwFqdnOk() (*string, bool)`

GetPgwFqdnOk returns a tuple with the PgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdn

`func (o *EpsIwkPgw) SetPgwFqdn(v string)`

SetPgwFqdn sets PgwFqdn field to given value.


### GetSmfInstanceId

`func (o *EpsIwkPgw) GetSmfInstanceId() string`

GetSmfInstanceId returns the SmfInstanceId field if non-nil, zero value otherwise.

### GetSmfInstanceIdOk

`func (o *EpsIwkPgw) GetSmfInstanceIdOk() (*string, bool)`

GetSmfInstanceIdOk returns a tuple with the SmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInstanceId

`func (o *EpsIwkPgw) SetSmfInstanceId(v string)`

SetSmfInstanceId sets SmfInstanceId field to given value.


### GetPlmnId

`func (o *EpsIwkPgw) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *EpsIwkPgw) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *EpsIwkPgw) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *EpsIwkPgw) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


