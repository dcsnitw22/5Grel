# AcknowledgeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SorMacIue** | Pointer to **string** | MAC value for protecting SOR procedure (SoR-MAC-IAUSF and SoR-XMAC-IUE). | [optional] 
**UpuMacIue** | Pointer to **string** | MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE). | [optional] 
**ProvisioningTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**SorTransparentContainer** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**UeNotReachable** | Pointer to **bool** |  | [optional] [default to false]
**UpuTransparentContainer** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 

## Methods

### NewAcknowledgeInfo

`func NewAcknowledgeInfo(provisioningTime time.Time, ) *AcknowledgeInfo`

NewAcknowledgeInfo instantiates a new AcknowledgeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcknowledgeInfoWithDefaults

`func NewAcknowledgeInfoWithDefaults() *AcknowledgeInfo`

NewAcknowledgeInfoWithDefaults instantiates a new AcknowledgeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSorMacIue

`func (o *AcknowledgeInfo) GetSorMacIue() string`

GetSorMacIue returns the SorMacIue field if non-nil, zero value otherwise.

### GetSorMacIueOk

`func (o *AcknowledgeInfo) GetSorMacIueOk() (*string, bool)`

GetSorMacIueOk returns a tuple with the SorMacIue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorMacIue

`func (o *AcknowledgeInfo) SetSorMacIue(v string)`

SetSorMacIue sets SorMacIue field to given value.

### HasSorMacIue

`func (o *AcknowledgeInfo) HasSorMacIue() bool`

HasSorMacIue returns a boolean if a field has been set.

### GetUpuMacIue

`func (o *AcknowledgeInfo) GetUpuMacIue() string`

GetUpuMacIue returns the UpuMacIue field if non-nil, zero value otherwise.

### GetUpuMacIueOk

`func (o *AcknowledgeInfo) GetUpuMacIueOk() (*string, bool)`

GetUpuMacIueOk returns a tuple with the UpuMacIue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuMacIue

`func (o *AcknowledgeInfo) SetUpuMacIue(v string)`

SetUpuMacIue sets UpuMacIue field to given value.

### HasUpuMacIue

`func (o *AcknowledgeInfo) HasUpuMacIue() bool`

HasUpuMacIue returns a boolean if a field has been set.

### GetProvisioningTime

`func (o *AcknowledgeInfo) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *AcknowledgeInfo) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *AcknowledgeInfo) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.


### GetSorTransparentContainer

`func (o *AcknowledgeInfo) GetSorTransparentContainer() string`

GetSorTransparentContainer returns the SorTransparentContainer field if non-nil, zero value otherwise.

### GetSorTransparentContainerOk

`func (o *AcknowledgeInfo) GetSorTransparentContainerOk() (*string, bool)`

GetSorTransparentContainerOk returns a tuple with the SorTransparentContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorTransparentContainer

`func (o *AcknowledgeInfo) SetSorTransparentContainer(v string)`

SetSorTransparentContainer sets SorTransparentContainer field to given value.

### HasSorTransparentContainer

`func (o *AcknowledgeInfo) HasSorTransparentContainer() bool`

HasSorTransparentContainer returns a boolean if a field has been set.

### GetUeNotReachable

`func (o *AcknowledgeInfo) GetUeNotReachable() bool`

GetUeNotReachable returns the UeNotReachable field if non-nil, zero value otherwise.

### GetUeNotReachableOk

`func (o *AcknowledgeInfo) GetUeNotReachableOk() (*bool, bool)`

GetUeNotReachableOk returns a tuple with the UeNotReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeNotReachable

`func (o *AcknowledgeInfo) SetUeNotReachable(v bool)`

SetUeNotReachable sets UeNotReachable field to given value.

### HasUeNotReachable

`func (o *AcknowledgeInfo) HasUeNotReachable() bool`

HasUeNotReachable returns a boolean if a field has been set.

### GetUpuTransparentContainer

`func (o *AcknowledgeInfo) GetUpuTransparentContainer() string`

GetUpuTransparentContainer returns the UpuTransparentContainer field if non-nil, zero value otherwise.

### GetUpuTransparentContainerOk

`func (o *AcknowledgeInfo) GetUpuTransparentContainerOk() (*string, bool)`

GetUpuTransparentContainerOk returns a tuple with the UpuTransparentContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuTransparentContainer

`func (o *AcknowledgeInfo) SetUpuTransparentContainer(v string)`

SetUpuTransparentContainer sets UpuTransparentContainer field to given value.

### HasUpuTransparentContainer

`func (o *AcknowledgeInfo) HasUpuTransparentContainer() bool`

HasUpuTransparentContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


