# SorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SteeringContainer** | Pointer to [**SteeringContainer**](SteeringContainer.md) |  | [optional] 
**AckInd** | **bool** | Contains indication whether the acknowledgement from UE is needed. | 
**SorMacIausf** | Pointer to **string** | MAC value for protecting SOR procedure (SoR-MAC-IAUSF and SoR-XMAC-IUE). | [optional] 
**Countersor** | Pointer to **string** | CounterSoR. | [optional] 
**ProvisioningTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**SorTransparentContainer** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**SorCmci** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**SorSnpnSi** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**SorSnpnSiLs** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**StoreSorCmciInMe** | Pointer to **bool** |  | [optional] 
**UsimSupportOfSorCmci** | Pointer to **bool** |  | [optional] 

## Methods

### NewSorInfo

`func NewSorInfo(ackInd bool, provisioningTime time.Time, ) *SorInfo`

NewSorInfo instantiates a new SorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSorInfoWithDefaults

`func NewSorInfoWithDefaults() *SorInfo`

NewSorInfoWithDefaults instantiates a new SorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSteeringContainer

`func (o *SorInfo) GetSteeringContainer() SteeringContainer`

GetSteeringContainer returns the SteeringContainer field if non-nil, zero value otherwise.

### GetSteeringContainerOk

`func (o *SorInfo) GetSteeringContainerOk() (*SteeringContainer, bool)`

GetSteeringContainerOk returns a tuple with the SteeringContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteeringContainer

`func (o *SorInfo) SetSteeringContainer(v SteeringContainer)`

SetSteeringContainer sets SteeringContainer field to given value.

### HasSteeringContainer

`func (o *SorInfo) HasSteeringContainer() bool`

HasSteeringContainer returns a boolean if a field has been set.

### GetAckInd

`func (o *SorInfo) GetAckInd() bool`

GetAckInd returns the AckInd field if non-nil, zero value otherwise.

### GetAckIndOk

`func (o *SorInfo) GetAckIndOk() (*bool, bool)`

GetAckIndOk returns a tuple with the AckInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckInd

`func (o *SorInfo) SetAckInd(v bool)`

SetAckInd sets AckInd field to given value.


### GetSorMacIausf

`func (o *SorInfo) GetSorMacIausf() string`

GetSorMacIausf returns the SorMacIausf field if non-nil, zero value otherwise.

### GetSorMacIausfOk

`func (o *SorInfo) GetSorMacIausfOk() (*string, bool)`

GetSorMacIausfOk returns a tuple with the SorMacIausf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorMacIausf

`func (o *SorInfo) SetSorMacIausf(v string)`

SetSorMacIausf sets SorMacIausf field to given value.

### HasSorMacIausf

`func (o *SorInfo) HasSorMacIausf() bool`

HasSorMacIausf returns a boolean if a field has been set.

### GetCountersor

`func (o *SorInfo) GetCountersor() string`

GetCountersor returns the Countersor field if non-nil, zero value otherwise.

### GetCountersorOk

`func (o *SorInfo) GetCountersorOk() (*string, bool)`

GetCountersorOk returns a tuple with the Countersor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountersor

`func (o *SorInfo) SetCountersor(v string)`

SetCountersor sets Countersor field to given value.

### HasCountersor

`func (o *SorInfo) HasCountersor() bool`

HasCountersor returns a boolean if a field has been set.

### GetProvisioningTime

`func (o *SorInfo) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *SorInfo) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *SorInfo) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.


### GetSorTransparentContainer

`func (o *SorInfo) GetSorTransparentContainer() string`

GetSorTransparentContainer returns the SorTransparentContainer field if non-nil, zero value otherwise.

### GetSorTransparentContainerOk

`func (o *SorInfo) GetSorTransparentContainerOk() (*string, bool)`

GetSorTransparentContainerOk returns a tuple with the SorTransparentContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorTransparentContainer

`func (o *SorInfo) SetSorTransparentContainer(v string)`

SetSorTransparentContainer sets SorTransparentContainer field to given value.

### HasSorTransparentContainer

`func (o *SorInfo) HasSorTransparentContainer() bool`

HasSorTransparentContainer returns a boolean if a field has been set.

### GetSorCmci

`func (o *SorInfo) GetSorCmci() string`

GetSorCmci returns the SorCmci field if non-nil, zero value otherwise.

### GetSorCmciOk

`func (o *SorInfo) GetSorCmciOk() (*string, bool)`

GetSorCmciOk returns a tuple with the SorCmci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorCmci

`func (o *SorInfo) SetSorCmci(v string)`

SetSorCmci sets SorCmci field to given value.

### HasSorCmci

`func (o *SorInfo) HasSorCmci() bool`

HasSorCmci returns a boolean if a field has been set.

### GetSorSnpnSi

`func (o *SorInfo) GetSorSnpnSi() string`

GetSorSnpnSi returns the SorSnpnSi field if non-nil, zero value otherwise.

### GetSorSnpnSiOk

`func (o *SorInfo) GetSorSnpnSiOk() (*string, bool)`

GetSorSnpnSiOk returns a tuple with the SorSnpnSi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorSnpnSi

`func (o *SorInfo) SetSorSnpnSi(v string)`

SetSorSnpnSi sets SorSnpnSi field to given value.

### HasSorSnpnSi

`func (o *SorInfo) HasSorSnpnSi() bool`

HasSorSnpnSi returns a boolean if a field has been set.

### GetSorSnpnSiLs

`func (o *SorInfo) GetSorSnpnSiLs() string`

GetSorSnpnSiLs returns the SorSnpnSiLs field if non-nil, zero value otherwise.

### GetSorSnpnSiLsOk

`func (o *SorInfo) GetSorSnpnSiLsOk() (*string, bool)`

GetSorSnpnSiLsOk returns a tuple with the SorSnpnSiLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorSnpnSiLs

`func (o *SorInfo) SetSorSnpnSiLs(v string)`

SetSorSnpnSiLs sets SorSnpnSiLs field to given value.

### HasSorSnpnSiLs

`func (o *SorInfo) HasSorSnpnSiLs() bool`

HasSorSnpnSiLs returns a boolean if a field has been set.

### GetStoreSorCmciInMe

`func (o *SorInfo) GetStoreSorCmciInMe() bool`

GetStoreSorCmciInMe returns the StoreSorCmciInMe field if non-nil, zero value otherwise.

### GetStoreSorCmciInMeOk

`func (o *SorInfo) GetStoreSorCmciInMeOk() (*bool, bool)`

GetStoreSorCmciInMeOk returns a tuple with the StoreSorCmciInMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSorCmciInMe

`func (o *SorInfo) SetStoreSorCmciInMe(v bool)`

SetStoreSorCmciInMe sets StoreSorCmciInMe field to given value.

### HasStoreSorCmciInMe

`func (o *SorInfo) HasStoreSorCmciInMe() bool`

HasStoreSorCmciInMe returns a boolean if a field has been set.

### GetUsimSupportOfSorCmci

`func (o *SorInfo) GetUsimSupportOfSorCmci() bool`

GetUsimSupportOfSorCmci returns the UsimSupportOfSorCmci field if non-nil, zero value otherwise.

### GetUsimSupportOfSorCmciOk

`func (o *SorInfo) GetUsimSupportOfSorCmciOk() (*bool, bool)`

GetUsimSupportOfSorCmciOk returns a tuple with the UsimSupportOfSorCmci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsimSupportOfSorCmci

`func (o *SorInfo) SetUsimSupportOfSorCmci(v bool)`

SetUsimSupportOfSorCmci sets UsimSupportOfSorCmci field to given value.

### HasUsimSupportOfSorCmci

`func (o *SorInfo) HasUsimSupportOfSorCmci() bool`

HasUsimSupportOfSorCmci returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


