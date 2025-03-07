# AzureBlobBrowserDestinationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 
**RegionType** | [**EAzureRegionType**](EAzureRegionType.md) |  | 
**ContainerName** | **string** | Name of the container where you want to store your backup data. | 
**FolderType** | Pointer to [**ECloudBrowserFolderType**](ECloudBrowserFolderType.md) |  | [optional] 

## Methods

### NewAzureBlobBrowserDestinationSpec

`func NewAzureBlobBrowserDestinationSpec(regionType EAzureRegionType, containerName string, ) *AzureBlobBrowserDestinationSpec`

NewAzureBlobBrowserDestinationSpec instantiates a new AzureBlobBrowserDestinationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobBrowserDestinationSpecWithDefaults

`func NewAzureBlobBrowserDestinationSpecWithDefaults() *AzureBlobBrowserDestinationSpec`

NewAzureBlobBrowserDestinationSpecWithDefaults instantiates a new AzureBlobBrowserDestinationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *AzureBlobBrowserDestinationSpec) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AzureBlobBrowserDestinationSpec) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AzureBlobBrowserDestinationSpec) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AzureBlobBrowserDestinationSpec) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetRegionType

`func (o *AzureBlobBrowserDestinationSpec) GetRegionType() EAzureRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AzureBlobBrowserDestinationSpec) GetRegionTypeOk() (*EAzureRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AzureBlobBrowserDestinationSpec) SetRegionType(v EAzureRegionType)`

SetRegionType sets RegionType field to given value.


### GetContainerName

`func (o *AzureBlobBrowserDestinationSpec) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *AzureBlobBrowserDestinationSpec) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *AzureBlobBrowserDestinationSpec) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetFolderType

`func (o *AzureBlobBrowserDestinationSpec) GetFolderType() ECloudBrowserFolderType`

GetFolderType returns the FolderType field if non-nil, zero value otherwise.

### GetFolderTypeOk

`func (o *AzureBlobBrowserDestinationSpec) GetFolderTypeOk() (*ECloudBrowserFolderType, bool)`

GetFolderTypeOk returns a tuple with the FolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderType

`func (o *AzureBlobBrowserDestinationSpec) SetFolderType(v ECloudBrowserFolderType)`

SetFolderType sets FolderType field to given value.

### HasFolderType

`func (o *AzureBlobBrowserDestinationSpec) HasFolderType() bool`

HasFolderType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


