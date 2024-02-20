# SmsfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmsfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**SmsfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 

## Methods

### NewSmsfInfo

`func NewSmsfInfo(smsfInstanceId string, plmnId PlmnId, ) *SmsfInfo`

NewSmsfInfo instantiates a new SmsfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsfInfoWithDefaults

`func NewSmsfInfoWithDefaults() *SmsfInfo`

NewSmsfInfoWithDefaults instantiates a new SmsfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmsfInstanceId

`func (o *SmsfInfo) GetSmsfInstanceId() string`

GetSmsfInstanceId returns the SmsfInstanceId field if non-nil, zero value otherwise.

### GetSmsfInstanceIdOk

`func (o *SmsfInfo) GetSmsfInstanceIdOk() (*string, bool)`

GetSmsfInstanceIdOk returns a tuple with the SmsfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfInstanceId

`func (o *SmsfInfo) SetSmsfInstanceId(v string)`

SetSmsfInstanceId sets SmsfInstanceId field to given value.


### GetPlmnId

`func (o *SmsfInfo) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *SmsfInfo) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *SmsfInfo) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetSmsfSetId

`func (o *SmsfInfo) GetSmsfSetId() string`

GetSmsfSetId returns the SmsfSetId field if non-nil, zero value otherwise.

### GetSmsfSetIdOk

`func (o *SmsfInfo) GetSmsfSetIdOk() (*string, bool)`

GetSmsfSetIdOk returns a tuple with the SmsfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfSetId

`func (o *SmsfInfo) SetSmsfSetId(v string)`

SetSmsfSetId sets SmsfSetId field to given value.

### HasSmsfSetId

`func (o *SmsfInfo) HasSmsfSetId() bool`

HasSmsfSetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


