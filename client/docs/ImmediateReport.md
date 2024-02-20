# ImmediateReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmData** | Pointer to [**AccessAndMobilitySubscriptionData**](AccessAndMobilitySubscriptionData.md) |  | [optional] 
**SmfSelData** | Pointer to [**SmfSelectionSubscriptionData**](SmfSelectionSubscriptionData.md) |  | [optional] 
**UecAmfData** | Pointer to [**UeContextInAmfData**](UeContextInAmfData.md) |  | [optional] 
**UecSmfData** | Pointer to [**UeContextInSmfData**](UeContextInSmfData.md) |  | [optional] 
**UecSmsfData** | Pointer to [**UeContextInSmsfData**](UeContextInSmsfData.md) |  | [optional] 
**SmsSubsData** | Pointer to [**SmsSubscriptionData**](SmsSubscriptionData.md) |  | [optional] 
**SmData** | Pointer to [**SmSubsData**](SmSubsData.md) |  | [optional] 
**TraceData** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 
**SmsMngData** | Pointer to [**SmsManagementSubscriptionData**](SmsManagementSubscriptionData.md) |  | [optional] 
**LcsPrivacyData** | Pointer to [**LcsPrivacyData**](LcsPrivacyData.md) |  | [optional] 
**LcsMoData** | Pointer to [**LcsMoData**](LcsMoData.md) |  | [optional] 
**LcsSubscriptionData** | Pointer to [**LcsSubscriptionData**](LcsSubscriptionData.md) |  | [optional] 
**V2xData** | Pointer to [**V2xSubscriptionData**](V2xSubscriptionData.md) |  | [optional] 
**LcsBroadcastAssistanceTypesData** | Pointer to [**LcsBroadcastAssistanceTypesData**](LcsBroadcastAssistanceTypesData.md) |  | [optional] 
**ProseData** | Pointer to [**ProseSubscriptionData**](ProseSubscriptionData.md) |  | [optional] 
**MbsData** | Pointer to [**MbsSubscriptionData**](MbsSubscriptionData.md) |  | [optional] 
**UcData** | Pointer to [**UcSubscriptionData**](UcSubscriptionData.md) |  | [optional] 
**A2xData** | Pointer to [**A2xSubscriptionData**](A2xSubscriptionData.md) |  | [optional] 

## Methods

### NewImmediateReport

`func NewImmediateReport() *ImmediateReport`

NewImmediateReport instantiates a new ImmediateReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImmediateReportWithDefaults

`func NewImmediateReportWithDefaults() *ImmediateReport`

NewImmediateReportWithDefaults instantiates a new ImmediateReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmData

`func (o *ImmediateReport) GetAmData() AccessAndMobilitySubscriptionData`

GetAmData returns the AmData field if non-nil, zero value otherwise.

### GetAmDataOk

`func (o *ImmediateReport) GetAmDataOk() (*AccessAndMobilitySubscriptionData, bool)`

GetAmDataOk returns a tuple with the AmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmData

`func (o *ImmediateReport) SetAmData(v AccessAndMobilitySubscriptionData)`

SetAmData sets AmData field to given value.

### HasAmData

`func (o *ImmediateReport) HasAmData() bool`

HasAmData returns a boolean if a field has been set.

### GetSmfSelData

`func (o *ImmediateReport) GetSmfSelData() SmfSelectionSubscriptionData`

GetSmfSelData returns the SmfSelData field if non-nil, zero value otherwise.

### GetSmfSelDataOk

`func (o *ImmediateReport) GetSmfSelDataOk() (*SmfSelectionSubscriptionData, bool)`

GetSmfSelDataOk returns a tuple with the SmfSelData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfSelData

`func (o *ImmediateReport) SetSmfSelData(v SmfSelectionSubscriptionData)`

SetSmfSelData sets SmfSelData field to given value.

### HasSmfSelData

`func (o *ImmediateReport) HasSmfSelData() bool`

HasSmfSelData returns a boolean if a field has been set.

### GetUecAmfData

`func (o *ImmediateReport) GetUecAmfData() UeContextInAmfData`

GetUecAmfData returns the UecAmfData field if non-nil, zero value otherwise.

### GetUecAmfDataOk

`func (o *ImmediateReport) GetUecAmfDataOk() (*UeContextInAmfData, bool)`

GetUecAmfDataOk returns a tuple with the UecAmfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUecAmfData

`func (o *ImmediateReport) SetUecAmfData(v UeContextInAmfData)`

SetUecAmfData sets UecAmfData field to given value.

### HasUecAmfData

`func (o *ImmediateReport) HasUecAmfData() bool`

HasUecAmfData returns a boolean if a field has been set.

### GetUecSmfData

`func (o *ImmediateReport) GetUecSmfData() UeContextInSmfData`

GetUecSmfData returns the UecSmfData field if non-nil, zero value otherwise.

### GetUecSmfDataOk

`func (o *ImmediateReport) GetUecSmfDataOk() (*UeContextInSmfData, bool)`

GetUecSmfDataOk returns a tuple with the UecSmfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUecSmfData

`func (o *ImmediateReport) SetUecSmfData(v UeContextInSmfData)`

SetUecSmfData sets UecSmfData field to given value.

### HasUecSmfData

`func (o *ImmediateReport) HasUecSmfData() bool`

HasUecSmfData returns a boolean if a field has been set.

### GetUecSmsfData

`func (o *ImmediateReport) GetUecSmsfData() UeContextInSmsfData`

GetUecSmsfData returns the UecSmsfData field if non-nil, zero value otherwise.

### GetUecSmsfDataOk

`func (o *ImmediateReport) GetUecSmsfDataOk() (*UeContextInSmsfData, bool)`

GetUecSmsfDataOk returns a tuple with the UecSmsfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUecSmsfData

`func (o *ImmediateReport) SetUecSmsfData(v UeContextInSmsfData)`

SetUecSmsfData sets UecSmsfData field to given value.

### HasUecSmsfData

`func (o *ImmediateReport) HasUecSmsfData() bool`

HasUecSmsfData returns a boolean if a field has been set.

### GetSmsSubsData

`func (o *ImmediateReport) GetSmsSubsData() SmsSubscriptionData`

GetSmsSubsData returns the SmsSubsData field if non-nil, zero value otherwise.

### GetSmsSubsDataOk

`func (o *ImmediateReport) GetSmsSubsDataOk() (*SmsSubscriptionData, bool)`

GetSmsSubsDataOk returns a tuple with the SmsSubsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsSubsData

`func (o *ImmediateReport) SetSmsSubsData(v SmsSubscriptionData)`

SetSmsSubsData sets SmsSubsData field to given value.

### HasSmsSubsData

`func (o *ImmediateReport) HasSmsSubsData() bool`

HasSmsSubsData returns a boolean if a field has been set.

### GetSmData

`func (o *ImmediateReport) GetSmData() SmSubsData`

GetSmData returns the SmData field if non-nil, zero value otherwise.

### GetSmDataOk

`func (o *ImmediateReport) GetSmDataOk() (*SmSubsData, bool)`

GetSmDataOk returns a tuple with the SmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmData

`func (o *ImmediateReport) SetSmData(v SmSubsData)`

SetSmData sets SmData field to given value.

### HasSmData

`func (o *ImmediateReport) HasSmData() bool`

HasSmData returns a boolean if a field has been set.

### GetTraceData

`func (o *ImmediateReport) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *ImmediateReport) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *ImmediateReport) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *ImmediateReport) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *ImmediateReport) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *ImmediateReport) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetSmsMngData

`func (o *ImmediateReport) GetSmsMngData() SmsManagementSubscriptionData`

GetSmsMngData returns the SmsMngData field if non-nil, zero value otherwise.

### GetSmsMngDataOk

`func (o *ImmediateReport) GetSmsMngDataOk() (*SmsManagementSubscriptionData, bool)`

GetSmsMngDataOk returns a tuple with the SmsMngData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMngData

`func (o *ImmediateReport) SetSmsMngData(v SmsManagementSubscriptionData)`

SetSmsMngData sets SmsMngData field to given value.

### HasSmsMngData

`func (o *ImmediateReport) HasSmsMngData() bool`

HasSmsMngData returns a boolean if a field has been set.

### GetLcsPrivacyData

`func (o *ImmediateReport) GetLcsPrivacyData() LcsPrivacyData`

GetLcsPrivacyData returns the LcsPrivacyData field if non-nil, zero value otherwise.

### GetLcsPrivacyDataOk

`func (o *ImmediateReport) GetLcsPrivacyDataOk() (*LcsPrivacyData, bool)`

GetLcsPrivacyDataOk returns a tuple with the LcsPrivacyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsPrivacyData

`func (o *ImmediateReport) SetLcsPrivacyData(v LcsPrivacyData)`

SetLcsPrivacyData sets LcsPrivacyData field to given value.

### HasLcsPrivacyData

`func (o *ImmediateReport) HasLcsPrivacyData() bool`

HasLcsPrivacyData returns a boolean if a field has been set.

### GetLcsMoData

`func (o *ImmediateReport) GetLcsMoData() LcsMoData`

GetLcsMoData returns the LcsMoData field if non-nil, zero value otherwise.

### GetLcsMoDataOk

`func (o *ImmediateReport) GetLcsMoDataOk() (*LcsMoData, bool)`

GetLcsMoDataOk returns a tuple with the LcsMoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsMoData

`func (o *ImmediateReport) SetLcsMoData(v LcsMoData)`

SetLcsMoData sets LcsMoData field to given value.

### HasLcsMoData

`func (o *ImmediateReport) HasLcsMoData() bool`

HasLcsMoData returns a boolean if a field has been set.

### GetLcsSubscriptionData

`func (o *ImmediateReport) GetLcsSubscriptionData() LcsSubscriptionData`

GetLcsSubscriptionData returns the LcsSubscriptionData field if non-nil, zero value otherwise.

### GetLcsSubscriptionDataOk

`func (o *ImmediateReport) GetLcsSubscriptionDataOk() (*LcsSubscriptionData, bool)`

GetLcsSubscriptionDataOk returns a tuple with the LcsSubscriptionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsSubscriptionData

`func (o *ImmediateReport) SetLcsSubscriptionData(v LcsSubscriptionData)`

SetLcsSubscriptionData sets LcsSubscriptionData field to given value.

### HasLcsSubscriptionData

`func (o *ImmediateReport) HasLcsSubscriptionData() bool`

HasLcsSubscriptionData returns a boolean if a field has been set.

### GetV2xData

`func (o *ImmediateReport) GetV2xData() V2xSubscriptionData`

GetV2xData returns the V2xData field if non-nil, zero value otherwise.

### GetV2xDataOk

`func (o *ImmediateReport) GetV2xDataOk() (*V2xSubscriptionData, bool)`

GetV2xDataOk returns a tuple with the V2xData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2xData

`func (o *ImmediateReport) SetV2xData(v V2xSubscriptionData)`

SetV2xData sets V2xData field to given value.

### HasV2xData

`func (o *ImmediateReport) HasV2xData() bool`

HasV2xData returns a boolean if a field has been set.

### GetLcsBroadcastAssistanceTypesData

`func (o *ImmediateReport) GetLcsBroadcastAssistanceTypesData() LcsBroadcastAssistanceTypesData`

GetLcsBroadcastAssistanceTypesData returns the LcsBroadcastAssistanceTypesData field if non-nil, zero value otherwise.

### GetLcsBroadcastAssistanceTypesDataOk

`func (o *ImmediateReport) GetLcsBroadcastAssistanceTypesDataOk() (*LcsBroadcastAssistanceTypesData, bool)`

GetLcsBroadcastAssistanceTypesDataOk returns a tuple with the LcsBroadcastAssistanceTypesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsBroadcastAssistanceTypesData

`func (o *ImmediateReport) SetLcsBroadcastAssistanceTypesData(v LcsBroadcastAssistanceTypesData)`

SetLcsBroadcastAssistanceTypesData sets LcsBroadcastAssistanceTypesData field to given value.

### HasLcsBroadcastAssistanceTypesData

`func (o *ImmediateReport) HasLcsBroadcastAssistanceTypesData() bool`

HasLcsBroadcastAssistanceTypesData returns a boolean if a field has been set.

### GetProseData

`func (o *ImmediateReport) GetProseData() ProseSubscriptionData`

GetProseData returns the ProseData field if non-nil, zero value otherwise.

### GetProseDataOk

`func (o *ImmediateReport) GetProseDataOk() (*ProseSubscriptionData, bool)`

GetProseDataOk returns a tuple with the ProseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseData

`func (o *ImmediateReport) SetProseData(v ProseSubscriptionData)`

SetProseData sets ProseData field to given value.

### HasProseData

`func (o *ImmediateReport) HasProseData() bool`

HasProseData returns a boolean if a field has been set.

### GetMbsData

`func (o *ImmediateReport) GetMbsData() MbsSubscriptionData`

GetMbsData returns the MbsData field if non-nil, zero value otherwise.

### GetMbsDataOk

`func (o *ImmediateReport) GetMbsDataOk() (*MbsSubscriptionData, bool)`

GetMbsDataOk returns a tuple with the MbsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsData

`func (o *ImmediateReport) SetMbsData(v MbsSubscriptionData)`

SetMbsData sets MbsData field to given value.

### HasMbsData

`func (o *ImmediateReport) HasMbsData() bool`

HasMbsData returns a boolean if a field has been set.

### GetUcData

`func (o *ImmediateReport) GetUcData() UcSubscriptionData`

GetUcData returns the UcData field if non-nil, zero value otherwise.

### GetUcDataOk

`func (o *ImmediateReport) GetUcDataOk() (*UcSubscriptionData, bool)`

GetUcDataOk returns a tuple with the UcData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcData

`func (o *ImmediateReport) SetUcData(v UcSubscriptionData)`

SetUcData sets UcData field to given value.

### HasUcData

`func (o *ImmediateReport) HasUcData() bool`

HasUcData returns a boolean if a field has been set.

### GetA2xData

`func (o *ImmediateReport) GetA2xData() A2xSubscriptionData`

GetA2xData returns the A2xData field if non-nil, zero value otherwise.

### GetA2xDataOk

`func (o *ImmediateReport) GetA2xDataOk() (*A2xSubscriptionData, bool)`

GetA2xDataOk returns a tuple with the A2xData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA2xData

`func (o *ImmediateReport) SetA2xData(v A2xSubscriptionData)`

SetA2xData sets A2xData field to given value.

### HasA2xData

`func (o *ImmediateReport) HasA2xData() bool`

HasA2xData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


