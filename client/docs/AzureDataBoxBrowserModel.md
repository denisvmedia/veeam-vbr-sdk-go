# AzureDataBoxBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server used to connect to the object storage. | [optional] 
**Containers** | Pointer to [**[]AzureDataBoxContainerBrowserModel**](AzureDataBoxContainerBrowserModel.md) | Array of containers that reside in the Azure storage account. | [optional] 

## Methods

### NewAzureDataBoxBrowserModel

`func NewAzureDataBoxBrowserModel() *AzureDataBoxBrowserModel`

NewAzureDataBoxBrowserModel instantiates a new AzureDataBoxBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDataBoxBrowserModelWithDefaults

`func NewAzureDataBoxBrowserModelWithDefaults() *AzureDataBoxBrowserModel`

NewAzureDataBoxBrowserModelWithDefaults instantiates a new AzureDataBoxBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *AzureDataBoxBrowserModel) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AzureDataBoxBrowserModel) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AzureDataBoxBrowserModel) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AzureDataBoxBrowserModel) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetContainers

`func (o *AzureDataBoxBrowserModel) GetContainers() []AzureDataBoxContainerBrowserModel`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *AzureDataBoxBrowserModel) GetContainersOk() (*[]AzureDataBoxContainerBrowserModel, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *AzureDataBoxBrowserModel) SetContainers(v []AzureDataBoxContainerBrowserModel)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *AzureDataBoxBrowserModel) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


