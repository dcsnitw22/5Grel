# Modify200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**ImplicitUnsubscribe** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**CallbackReference** | **string** | String providing an URI formatted according to RFC 3986. | 
**AmfServiceName** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**MonitoredResourceUris** | **[]string** |  | 
**SingleNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**ImmediateReport** | Pointer to **bool** |  | [optional] [default to false]
**Report** | [**[]ReportItem**](ReportItem.md) | The execution report contains an array of report items. Each report item indicates one  failed modification.  | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 
**NfChangeFilter** | Pointer to **bool** |  | [optional] [default to false]
**UniqueSubscription** | Pointer to **bool** |  | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 
**UeConSmfDataSubFilter** | Pointer to [**UeContextInSmfDataSubFilter**](UeContextInSmfDataSubFilter.md) |  | [optional] 
**AdjacentPlmns** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**DisasterRoamingInd** | Pointer to **bool** |  | [optional] [default to false]
**DataRestorationCallbackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**UdrRestartInd** | Pointer to **bool** |  | [optional] [default to false]
**ExpectedUeBehaviourThresholds** | Pointer to [**map[string]ExpectedUeBehaviourThreshold**](ExpectedUeBehaviourThreshold.md) | A map(list of key-value pairs) where a valid JSON pointer serves as key of expectedUeBehaviourThresholds | [optional] 

## Methods

### NewModify200Response

`func NewModify200Response(nfInstanceId string, callbackReference string, monitoredResourceUris []string, report []ReportItem, ) *Modify200Response`

NewModify200Response instantiates a new Modify200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModify200ResponseWithDefaults

`func NewModify200ResponseWithDefaults() *Modify200Response`

NewModify200ResponseWithDefaults instantiates a new Modify200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *Modify200Response) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *Modify200Response) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *Modify200Response) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetImplicitUnsubscribe

`func (o *Modify200Response) GetImplicitUnsubscribe() bool`

GetImplicitUnsubscribe returns the ImplicitUnsubscribe field if non-nil, zero value otherwise.

### GetImplicitUnsubscribeOk

`func (o *Modify200Response) GetImplicitUnsubscribeOk() (*bool, bool)`

GetImplicitUnsubscribeOk returns a tuple with the ImplicitUnsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitUnsubscribe

`func (o *Modify200Response) SetImplicitUnsubscribe(v bool)`

SetImplicitUnsubscribe sets ImplicitUnsubscribe field to given value.

### HasImplicitUnsubscribe

`func (o *Modify200Response) HasImplicitUnsubscribe() bool`

HasImplicitUnsubscribe returns a boolean if a field has been set.

### GetExpires

`func (o *Modify200Response) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Modify200Response) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Modify200Response) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *Modify200Response) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetCallbackReference

`func (o *Modify200Response) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *Modify200Response) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *Modify200Response) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.


### GetAmfServiceName

`func (o *Modify200Response) GetAmfServiceName() ServiceName`

GetAmfServiceName returns the AmfServiceName field if non-nil, zero value otherwise.

### GetAmfServiceNameOk

`func (o *Modify200Response) GetAmfServiceNameOk() (*ServiceName, bool)`

GetAmfServiceNameOk returns a tuple with the AmfServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfServiceName

`func (o *Modify200Response) SetAmfServiceName(v ServiceName)`

SetAmfServiceName sets AmfServiceName field to given value.

### HasAmfServiceName

`func (o *Modify200Response) HasAmfServiceName() bool`

HasAmfServiceName returns a boolean if a field has been set.

### GetMonitoredResourceUris

`func (o *Modify200Response) GetMonitoredResourceUris() []string`

GetMonitoredResourceUris returns the MonitoredResourceUris field if non-nil, zero value otherwise.

### GetMonitoredResourceUrisOk

`func (o *Modify200Response) GetMonitoredResourceUrisOk() (*[]string, bool)`

GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredResourceUris

`func (o *Modify200Response) SetMonitoredResourceUris(v []string)`

SetMonitoredResourceUris sets MonitoredResourceUris field to given value.


### GetSingleNssai

`func (o *Modify200Response) GetSingleNssai() Snssai`

GetSingleNssai returns the SingleNssai field if non-nil, zero value otherwise.

### GetSingleNssaiOk

`func (o *Modify200Response) GetSingleNssaiOk() (*Snssai, bool)`

GetSingleNssaiOk returns a tuple with the SingleNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssai

`func (o *Modify200Response) SetSingleNssai(v Snssai)`

SetSingleNssai sets SingleNssai field to given value.

### HasSingleNssai

`func (o *Modify200Response) HasSingleNssai() bool`

HasSingleNssai returns a boolean if a field has been set.

### GetDnn

`func (o *Modify200Response) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *Modify200Response) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *Modify200Response) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *Modify200Response) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *Modify200Response) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Modify200Response) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Modify200Response) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *Modify200Response) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetPlmnId

`func (o *Modify200Response) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *Modify200Response) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *Modify200Response) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *Modify200Response) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetImmediateReport

`func (o *Modify200Response) GetImmediateReport() bool`

GetImmediateReport returns the ImmediateReport field if non-nil, zero value otherwise.

### GetImmediateReportOk

`func (o *Modify200Response) GetImmediateReportOk() (*bool, bool)`

GetImmediateReportOk returns a tuple with the ImmediateReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateReport

`func (o *Modify200Response) SetImmediateReport(v bool)`

SetImmediateReport sets ImmediateReport field to given value.

### HasImmediateReport

`func (o *Modify200Response) HasImmediateReport() bool`

HasImmediateReport returns a boolean if a field has been set.

### GetReport

`func (o *Modify200Response) GetReport() []ReportItem`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *Modify200Response) GetReportOk() (*[]ReportItem, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *Modify200Response) SetReport(v []ReportItem)`

SetReport sets Report field to given value.


### GetSupportedFeatures

`func (o *Modify200Response) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *Modify200Response) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *Modify200Response) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *Modify200Response) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetContextInfo

`func (o *Modify200Response) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *Modify200Response) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *Modify200Response) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *Modify200Response) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.

### GetNfChangeFilter

`func (o *Modify200Response) GetNfChangeFilter() bool`

GetNfChangeFilter returns the NfChangeFilter field if non-nil, zero value otherwise.

### GetNfChangeFilterOk

`func (o *Modify200Response) GetNfChangeFilterOk() (*bool, bool)`

GetNfChangeFilterOk returns a tuple with the NfChangeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfChangeFilter

`func (o *Modify200Response) SetNfChangeFilter(v bool)`

SetNfChangeFilter sets NfChangeFilter field to given value.

### HasNfChangeFilter

`func (o *Modify200Response) HasNfChangeFilter() bool`

HasNfChangeFilter returns a boolean if a field has been set.

### GetUniqueSubscription

`func (o *Modify200Response) GetUniqueSubscription() bool`

GetUniqueSubscription returns the UniqueSubscription field if non-nil, zero value otherwise.

### GetUniqueSubscriptionOk

`func (o *Modify200Response) GetUniqueSubscriptionOk() (*bool, bool)`

GetUniqueSubscriptionOk returns a tuple with the UniqueSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueSubscription

`func (o *Modify200Response) SetUniqueSubscription(v bool)`

SetUniqueSubscription sets UniqueSubscription field to given value.

### HasUniqueSubscription

`func (o *Modify200Response) HasUniqueSubscription() bool`

HasUniqueSubscription returns a boolean if a field has been set.

### GetResetIds

`func (o *Modify200Response) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *Modify200Response) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *Modify200Response) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *Modify200Response) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.

### GetUeConSmfDataSubFilter

`func (o *Modify200Response) GetUeConSmfDataSubFilter() UeContextInSmfDataSubFilter`

GetUeConSmfDataSubFilter returns the UeConSmfDataSubFilter field if non-nil, zero value otherwise.

### GetUeConSmfDataSubFilterOk

`func (o *Modify200Response) GetUeConSmfDataSubFilterOk() (*UeContextInSmfDataSubFilter, bool)`

GetUeConSmfDataSubFilterOk returns a tuple with the UeConSmfDataSubFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeConSmfDataSubFilter

`func (o *Modify200Response) SetUeConSmfDataSubFilter(v UeContextInSmfDataSubFilter)`

SetUeConSmfDataSubFilter sets UeConSmfDataSubFilter field to given value.

### HasUeConSmfDataSubFilter

`func (o *Modify200Response) HasUeConSmfDataSubFilter() bool`

HasUeConSmfDataSubFilter returns a boolean if a field has been set.

### GetAdjacentPlmns

`func (o *Modify200Response) GetAdjacentPlmns() []PlmnId`

GetAdjacentPlmns returns the AdjacentPlmns field if non-nil, zero value otherwise.

### GetAdjacentPlmnsOk

`func (o *Modify200Response) GetAdjacentPlmnsOk() (*[]PlmnId, bool)`

GetAdjacentPlmnsOk returns a tuple with the AdjacentPlmns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjacentPlmns

`func (o *Modify200Response) SetAdjacentPlmns(v []PlmnId)`

SetAdjacentPlmns sets AdjacentPlmns field to given value.

### HasAdjacentPlmns

`func (o *Modify200Response) HasAdjacentPlmns() bool`

HasAdjacentPlmns returns a boolean if a field has been set.

### GetDisasterRoamingInd

`func (o *Modify200Response) GetDisasterRoamingInd() bool`

GetDisasterRoamingInd returns the DisasterRoamingInd field if non-nil, zero value otherwise.

### GetDisasterRoamingIndOk

`func (o *Modify200Response) GetDisasterRoamingIndOk() (*bool, bool)`

GetDisasterRoamingIndOk returns a tuple with the DisasterRoamingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisasterRoamingInd

`func (o *Modify200Response) SetDisasterRoamingInd(v bool)`

SetDisasterRoamingInd sets DisasterRoamingInd field to given value.

### HasDisasterRoamingInd

`func (o *Modify200Response) HasDisasterRoamingInd() bool`

HasDisasterRoamingInd returns a boolean if a field has been set.

### GetDataRestorationCallbackUri

`func (o *Modify200Response) GetDataRestorationCallbackUri() string`

GetDataRestorationCallbackUri returns the DataRestorationCallbackUri field if non-nil, zero value otherwise.

### GetDataRestorationCallbackUriOk

`func (o *Modify200Response) GetDataRestorationCallbackUriOk() (*string, bool)`

GetDataRestorationCallbackUriOk returns a tuple with the DataRestorationCallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRestorationCallbackUri

`func (o *Modify200Response) SetDataRestorationCallbackUri(v string)`

SetDataRestorationCallbackUri sets DataRestorationCallbackUri field to given value.

### HasDataRestorationCallbackUri

`func (o *Modify200Response) HasDataRestorationCallbackUri() bool`

HasDataRestorationCallbackUri returns a boolean if a field has been set.

### GetUdrRestartInd

`func (o *Modify200Response) GetUdrRestartInd() bool`

GetUdrRestartInd returns the UdrRestartInd field if non-nil, zero value otherwise.

### GetUdrRestartIndOk

`func (o *Modify200Response) GetUdrRestartIndOk() (*bool, bool)`

GetUdrRestartIndOk returns a tuple with the UdrRestartInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrRestartInd

`func (o *Modify200Response) SetUdrRestartInd(v bool)`

SetUdrRestartInd sets UdrRestartInd field to given value.

### HasUdrRestartInd

`func (o *Modify200Response) HasUdrRestartInd() bool`

HasUdrRestartInd returns a boolean if a field has been set.

### GetExpectedUeBehaviourThresholds

`func (o *Modify200Response) GetExpectedUeBehaviourThresholds() map[string]ExpectedUeBehaviourThreshold`

GetExpectedUeBehaviourThresholds returns the ExpectedUeBehaviourThresholds field if non-nil, zero value otherwise.

### GetExpectedUeBehaviourThresholdsOk

`func (o *Modify200Response) GetExpectedUeBehaviourThresholdsOk() (*map[string]ExpectedUeBehaviourThreshold, bool)`

GetExpectedUeBehaviourThresholdsOk returns a tuple with the ExpectedUeBehaviourThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUeBehaviourThresholds

`func (o *Modify200Response) SetExpectedUeBehaviourThresholds(v map[string]ExpectedUeBehaviourThreshold)`

SetExpectedUeBehaviourThresholds sets ExpectedUeBehaviourThresholds field to given value.

### HasExpectedUeBehaviourThresholds

`func (o *Modify200Response) HasExpectedUeBehaviourThresholds() bool`

HasExpectedUeBehaviourThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


