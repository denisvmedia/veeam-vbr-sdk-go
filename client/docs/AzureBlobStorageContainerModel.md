# AzureBlobStorageContainerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | **string** | Container name. | 
**FolderName** | **string** | Name of the folder that the object storage repository is mapped to. | 
**StorageConsumptionLimit** | Pointer to [**ObjectStorageConsumptionLimitModel**](ObjectStorageConsumptionLimitModel.md) |  | [optional] 

## Methods

### NewAzureBlobStorageContainerModel

`func NewAzureBlobStorageContainerModel(containerName string, folderName string, ) *AzureBlobStorageContainerModel`

NewAzureBlobStorageContainerModel instantiates a new AzureBlobStorageContainerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobStorageContainerModelWithDefaults

`func NewAzureBlobStorageContainerModelWithDefaults() *AzureBlobStorageContainerModel`

NewAzureBlobStorageContainerModelWithDefaults instantiates a new AzureBlobStorageContainerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerName

`func (o *AzureBlobStorageContainerModel) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *AzureBlobStorageContainerModel) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *AzureBlobStorageContainerModel) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetFolderName

`func (o *AzureBlobStorageContainerModel) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *AzureBlobStorageContainerModel) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *AzureBlobStorageContainerModel) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.


### GetStorageConsumptionLimit

`func (o *AzureBlobStorageContainerModel) GetStorageConsumptionLimit() ObjectStorageConsumptionLimitModel`

GetStorageConsumptionLimit returns the StorageConsumptionLimit field if non-nil, zero value otherwise.

### GetStorageConsumptionLimitOk

`func (o *AzureBlobStorageContainerModel) GetStorageConsumptionLimitOk() (*ObjectStorageConsumptionLimitModel, bool)`

GetStorageConsumptionLimitOk returns a tuple with the StorageConsumptionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumptionLimit

`func (o *AzureBlobStorageContainerModel) SetStorageConsumptionLimit(v ObjectStorageConsumptionLimitModel)`

SetStorageConsumptionLimit sets StorageConsumptionLimit field to given value.

### HasStorageConsumptionLimit

`func (o *AzureBlobStorageContainerModel) HasStorageConsumptionLimit() bool`

HasStorageConsumptionLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


