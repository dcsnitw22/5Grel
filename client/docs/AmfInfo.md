# AmfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**Guami** | [**Guami**](Guami.md) |  | 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 

## Methods

### NewAmfInfo

`func NewAmfInfo(amfInstanceId string, guami Guami, ) *AmfInfo`

NewAmfInfo instantiates a new AmfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfInfoWithDefaults

`func NewAmfInfoWithDefaults() *AmfInfo`

NewAmfInfoWithDefaults instantiates a new AmfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfInstanceId

`func (o *AmfInfo) GetAmfInstanceId() string`

GetAmfInstanceId returns the AmfInstanceId field if non-nil, zero value otherwise.

### GetAmfInstanceIdOk

`func (o *AmfInfo) GetAmfInstanceIdOk() (*string, bool)`

GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInstanceId

`func (o *AmfInfo) SetAmfInstanceId(v string)`

SetAmfInstanceId sets AmfInstanceId field to given value.


### GetGuami

`func (o *AmfInfo) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *AmfInfo) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *AmfInfo) SetGuami(v Guami)`

SetGuami sets Guami field to given value.


### GetAccessType

`func (o *AmfInfo) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *AmfInfo) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *AmfInfo) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *AmfInfo) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


