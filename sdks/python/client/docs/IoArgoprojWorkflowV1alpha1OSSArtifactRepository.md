# IonholuongutWorkflowV1alpha1OSSArtifactRepository

OSSArtifactRepository defines the controller configuration for an OSS artifact repository

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**access_key_secret** | [**SecretKeySelector**](SecretKeySelector.md) |  | [optional] 
**bucket** | **str** | Bucket is the name of the bucket | [optional] 
**create_bucket_if_not_present** | **bool** | CreateBucketIfNotPresent tells the driver to attempt to create the OSS bucket for output artifacts, if it doesn&#39;t exist | [optional] 
**endpoint** | **str** | Endpoint is the hostname of the bucket endpoint | [optional] 
**key_format** | **str** | KeyFormat is defines the format of how to store keys. Can reference workflow variables | [optional] 
**lifecycle_rule** | [**IonholuongutWorkflowV1alpha1OSSLifecycleRule**](IonholuongutWorkflowV1alpha1OSSLifecycleRule.md) |  | [optional] 
**secret_key_secret** | [**SecretKeySelector**](SecretKeySelector.md) |  | [optional] 
**security_token** | **str** | SecurityToken is the user&#39;s temporary security token. For more details, check out: https://www.alibabacloud.com/help/doc-detail/100624.htm | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


