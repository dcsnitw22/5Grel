# QmcConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QoeReference** | **string** | String containing MCC (3 digits), MNC (2 or 3 digits)  and QMC ID (3 octets, encoded as 6 hexadecimal digits).  | 
**ServiceType** | Pointer to [**QoeServiceType**](QoeServiceType.md) |  | [optional] 
**SliceScope** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**AreaScope** | Pointer to [**QmcAreaScope**](QmcAreaScope.md) |  | [optional] 
**QoeCollectionEntityAddress** | Pointer to [**NullableIpAddr**](IpAddr.md) |  | [optional] 
**QoeTarget** | Pointer to [**QoeTarget**](QoeTarget.md) |  | [optional] 
**MdtAlignmentInfo** | Pointer to **interface{}** | String containing: - Trace Reference: MCC (3 digits), MNC (2 or 3 digits),    Trace ID (3 octets, encoded as 6 hexadecimal digits) - Trace Recording Session Reference (2 octets, encoded as 4 hexadecimal digits)  | [optional] 
**AvailableRanVisibleQoeMetrics** | Pointer to [**[]AvailableRanVisibleQoeMetric**](AvailableRanVisibleQoeMetric.md) |  | [optional] 
**ContainerForAppLayerMeasConfig** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**MbsCommunicationServiceType** | Pointer to [**MbsServiceType**](MbsServiceType.md) |  | [optional] 

## Methods

### NewQmcConfigInfo

`func NewQmcConfigInfo(qoeReference string, ) *QmcConfigInfo`

NewQmcConfigInfo instantiates a new QmcConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQmcConfigInfoWithDefaults

`func NewQmcConfigInfoWithDefaults() *QmcConfigInfo`

NewQmcConfigInfoWithDefaults instantiates a new QmcConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQoeReference

`func (o *QmcConfigInfo) GetQoeReference() string`

GetQoeReference returns the QoeReference field if non-nil, zero value otherwise.

### GetQoeReferenceOk

`func (o *QmcConfigInfo) GetQoeReferenceOk() (*string, bool)`

GetQoeReferenceOk returns a tuple with the QoeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoeReference

`func (o *QmcConfigInfo) SetQoeReference(v string)`

SetQoeReference sets QoeReference field to given value.


### GetServiceType

`func (o *QmcConfigInfo) GetServiceType() QoeServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *QmcConfigInfo) GetServiceTypeOk() (*QoeServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *QmcConfigInfo) SetServiceType(v QoeServiceType)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *QmcConfigInfo) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSliceScope

`func (o *QmcConfigInfo) GetSliceScope() []Snssai`

GetSliceScope returns the SliceScope field if non-nil, zero value otherwise.

### GetSliceScopeOk

`func (o *QmcConfigInfo) GetSliceScopeOk() (*[]Snssai, bool)`

GetSliceScopeOk returns a tuple with the SliceScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceScope

`func (o *QmcConfigInfo) SetSliceScope(v []Snssai)`

SetSliceScope sets SliceScope field to given value.

### HasSliceScope

`func (o *QmcConfigInfo) HasSliceScope() bool`

HasSliceScope returns a boolean if a field has been set.

### GetAreaScope

`func (o *QmcConfigInfo) GetAreaScope() QmcAreaScope`

GetAreaScope returns the AreaScope field if non-nil, zero value otherwise.

### GetAreaScopeOk

`func (o *QmcConfigInfo) GetAreaScopeOk() (*QmcAreaScope, bool)`

GetAreaScopeOk returns a tuple with the AreaScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaScope

`func (o *QmcConfigInfo) SetAreaScope(v QmcAreaScope)`

SetAreaScope sets AreaScope field to given value.

### HasAreaScope

`func (o *QmcConfigInfo) HasAreaScope() bool`

HasAreaScope returns a boolean if a field has been set.

### GetQoeCollectionEntityAddress

`func (o *QmcConfigInfo) GetQoeCollectionEntityAddress() IpAddr`

GetQoeCollectionEntityAddress returns the QoeCollectionEntityAddress field if non-nil, zero value otherwise.

### GetQoeCollectionEntityAddressOk

`func (o *QmcConfigInfo) GetQoeCollectionEntityAddressOk() (*IpAddr, bool)`

GetQoeCollectionEntityAddressOk returns a tuple with the QoeCollectionEntityAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoeCollectionEntityAddress

`func (o *QmcConfigInfo) SetQoeCollectionEntityAddress(v IpAddr)`

SetQoeCollectionEntityAddress sets QoeCollectionEntityAddress field to given value.

### HasQoeCollectionEntityAddress

`func (o *QmcConfigInfo) HasQoeCollectionEntityAddress() bool`

HasQoeCollectionEntityAddress returns a boolean if a field has been set.

### SetQoeCollectionEntityAddressNil

`func (o *QmcConfigInfo) SetQoeCollectionEntityAddressNil(b bool)`

 SetQoeCollectionEntityAddressNil sets the value for QoeCollectionEntityAddress to be an explicit nil

### UnsetQoeCollectionEntityAddress
`func (o *QmcConfigInfo) UnsetQoeCollectionEntityAddress()`

UnsetQoeCollectionEntityAddress ensures that no value is present for QoeCollectionEntityAddress, not even an explicit nil
### GetQoeTarget

`func (o *QmcConfigInfo) GetQoeTarget() QoeTarget`

GetQoeTarget returns the QoeTarget field if non-nil, zero value otherwise.

### GetQoeTargetOk

`func (o *QmcConfigInfo) GetQoeTargetOk() (*QoeTarget, bool)`

GetQoeTargetOk returns a tuple with the QoeTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoeTarget

`func (o *QmcConfigInfo) SetQoeTarget(v QoeTarget)`

SetQoeTarget sets QoeTarget field to given value.

### HasQoeTarget

`func (o *QmcConfigInfo) HasQoeTarget() bool`

HasQoeTarget returns a boolean if a field has been set.

### GetMdtAlignmentInfo

`func (o *QmcConfigInfo) GetMdtAlignmentInfo() interface{}`

GetMdtAlignmentInfo returns the MdtAlignmentInfo field if non-nil, zero value otherwise.

### GetMdtAlignmentInfoOk

`func (o *QmcConfigInfo) GetMdtAlignmentInfoOk() (*interface{}, bool)`

GetMdtAlignmentInfoOk returns a tuple with the MdtAlignmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdtAlignmentInfo

`func (o *QmcConfigInfo) SetMdtAlignmentInfo(v interface{})`

SetMdtAlignmentInfo sets MdtAlignmentInfo field to given value.

### HasMdtAlignmentInfo

`func (o *QmcConfigInfo) HasMdtAlignmentInfo() bool`

HasMdtAlignmentInfo returns a boolean if a field has been set.

### SetMdtAlignmentInfoNil

`func (o *QmcConfigInfo) SetMdtAlignmentInfoNil(b bool)`

 SetMdtAlignmentInfoNil sets the value for MdtAlignmentInfo to be an explicit nil

### UnsetMdtAlignmentInfo
`func (o *QmcConfigInfo) UnsetMdtAlignmentInfo()`

UnsetMdtAlignmentInfo ensures that no value is present for MdtAlignmentInfo, not even an explicit nil
### GetAvailableRanVisibleQoeMetrics

`func (o *QmcConfigInfo) GetAvailableRanVisibleQoeMetrics() []AvailableRanVisibleQoeMetric`

GetAvailableRanVisibleQoeMetrics returns the AvailableRanVisibleQoeMetrics field if non-nil, zero value otherwise.

### GetAvailableRanVisibleQoeMetricsOk

`func (o *QmcConfigInfo) GetAvailableRanVisibleQoeMetricsOk() (*[]AvailableRanVisibleQoeMetric, bool)`

GetAvailableRanVisibleQoeMetricsOk returns a tuple with the AvailableRanVisibleQoeMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableRanVisibleQoeMetrics

`func (o *QmcConfigInfo) SetAvailableRanVisibleQoeMetrics(v []AvailableRanVisibleQoeMetric)`

SetAvailableRanVisibleQoeMetrics sets AvailableRanVisibleQoeMetrics field to given value.

### HasAvailableRanVisibleQoeMetrics

`func (o *QmcConfigInfo) HasAvailableRanVisibleQoeMetrics() bool`

HasAvailableRanVisibleQoeMetrics returns a boolean if a field has been set.

### GetContainerForAppLayerMeasConfig

`func (o *QmcConfigInfo) GetContainerForAppLayerMeasConfig() string`

GetContainerForAppLayerMeasConfig returns the ContainerForAppLayerMeasConfig field if non-nil, zero value otherwise.

### GetContainerForAppLayerMeasConfigOk

`func (o *QmcConfigInfo) GetContainerForAppLayerMeasConfigOk() (*string, bool)`

GetContainerForAppLayerMeasConfigOk returns a tuple with the ContainerForAppLayerMeasConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerForAppLayerMeasConfig

`func (o *QmcConfigInfo) SetContainerForAppLayerMeasConfig(v string)`

SetContainerForAppLayerMeasConfig sets ContainerForAppLayerMeasConfig field to given value.

### HasContainerForAppLayerMeasConfig

`func (o *QmcConfigInfo) HasContainerForAppLayerMeasConfig() bool`

HasContainerForAppLayerMeasConfig returns a boolean if a field has been set.

### GetMbsCommunicationServiceType

`func (o *QmcConfigInfo) GetMbsCommunicationServiceType() MbsServiceType`

GetMbsCommunicationServiceType returns the MbsCommunicationServiceType field if non-nil, zero value otherwise.

### GetMbsCommunicationServiceTypeOk

`func (o *QmcConfigInfo) GetMbsCommunicationServiceTypeOk() (*MbsServiceType, bool)`

GetMbsCommunicationServiceTypeOk returns a tuple with the MbsCommunicationServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsCommunicationServiceType

`func (o *QmcConfigInfo) SetMbsCommunicationServiceType(v MbsServiceType)`

SetMbsCommunicationServiceType sets MbsCommunicationServiceType field to given value.

### HasMbsCommunicationServiceType

`func (o *QmcConfigInfo) HasMbsCommunicationServiceType() bool`

HasMbsCommunicationServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


