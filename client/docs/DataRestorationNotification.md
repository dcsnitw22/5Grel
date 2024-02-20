# DataRestorationNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastReplicationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RecoveryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**SupiRanges** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**GpsiRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 
**SNssaiList** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**DnnList** | Pointer to **[]string** |  | [optional] 
**UdmGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 

## Methods

### NewDataRestorationNotification

`func NewDataRestorationNotification() *DataRestorationNotification`

NewDataRestorationNotification instantiates a new DataRestorationNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataRestorationNotificationWithDefaults

`func NewDataRestorationNotificationWithDefaults() *DataRestorationNotification`

NewDataRestorationNotificationWithDefaults instantiates a new DataRestorationNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastReplicationTime

`func (o *DataRestorationNotification) GetLastReplicationTime() time.Time`

GetLastReplicationTime returns the LastReplicationTime field if non-nil, zero value otherwise.

### GetLastReplicationTimeOk

`func (o *DataRestorationNotification) GetLastReplicationTimeOk() (*time.Time, bool)`

GetLastReplicationTimeOk returns a tuple with the LastReplicationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplicationTime

`func (o *DataRestorationNotification) SetLastReplicationTime(v time.Time)`

SetLastReplicationTime sets LastReplicationTime field to given value.

### HasLastReplicationTime

`func (o *DataRestorationNotification) HasLastReplicationTime() bool`

HasLastReplicationTime returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *DataRestorationNotification) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *DataRestorationNotification) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *DataRestorationNotification) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *DataRestorationNotification) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetPlmnId

`func (o *DataRestorationNotification) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *DataRestorationNotification) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *DataRestorationNotification) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *DataRestorationNotification) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetSupiRanges

`func (o *DataRestorationNotification) GetSupiRanges() []SupiRange`

GetSupiRanges returns the SupiRanges field if non-nil, zero value otherwise.

### GetSupiRangesOk

`func (o *DataRestorationNotification) GetSupiRangesOk() (*[]SupiRange, bool)`

GetSupiRangesOk returns a tuple with the SupiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRanges

`func (o *DataRestorationNotification) SetSupiRanges(v []SupiRange)`

SetSupiRanges sets SupiRanges field to given value.

### HasSupiRanges

`func (o *DataRestorationNotification) HasSupiRanges() bool`

HasSupiRanges returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *DataRestorationNotification) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *DataRestorationNotification) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *DataRestorationNotification) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *DataRestorationNotification) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.

### GetResetIds

`func (o *DataRestorationNotification) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *DataRestorationNotification) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *DataRestorationNotification) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *DataRestorationNotification) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.

### GetSNssaiList

`func (o *DataRestorationNotification) GetSNssaiList() []Snssai`

GetSNssaiList returns the SNssaiList field if non-nil, zero value otherwise.

### GetSNssaiListOk

`func (o *DataRestorationNotification) GetSNssaiListOk() (*[]Snssai, bool)`

GetSNssaiListOk returns a tuple with the SNssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiList

`func (o *DataRestorationNotification) SetSNssaiList(v []Snssai)`

SetSNssaiList sets SNssaiList field to given value.

### HasSNssaiList

`func (o *DataRestorationNotification) HasSNssaiList() bool`

HasSNssaiList returns a boolean if a field has been set.

### GetDnnList

`func (o *DataRestorationNotification) GetDnnList() []string`

GetDnnList returns the DnnList field if non-nil, zero value otherwise.

### GetDnnListOk

`func (o *DataRestorationNotification) GetDnnListOk() (*[]string, bool)`

GetDnnListOk returns a tuple with the DnnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnList

`func (o *DataRestorationNotification) SetDnnList(v []string)`

SetDnnList sets DnnList field to given value.

### HasDnnList

`func (o *DataRestorationNotification) HasDnnList() bool`

HasDnnList returns a boolean if a field has been set.

### GetUdmGroupId

`func (o *DataRestorationNotification) GetUdmGroupId() string`

GetUdmGroupId returns the UdmGroupId field if non-nil, zero value otherwise.

### GetUdmGroupIdOk

`func (o *DataRestorationNotification) GetUdmGroupIdOk() (*string, bool)`

GetUdmGroupIdOk returns a tuple with the UdmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmGroupId

`func (o *DataRestorationNotification) SetUdmGroupId(v string)`

SetUdmGroupId sets UdmGroupId field to given value.

### HasUdmGroupId

`func (o *DataRestorationNotification) HasUdmGroupId() bool`

HasUdmGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


