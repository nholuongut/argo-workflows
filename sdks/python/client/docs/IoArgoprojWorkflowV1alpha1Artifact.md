# IonholuongutWorkflowV1alpha1Artifact

Artifact indicates an artifact to place at a specified path

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | name of the artifact. must be unique within a template&#39;s inputs/outputs. | 
**archive** | [**IonholuongutWorkflowV1alpha1ArchiveStrategy**](IonholuongutWorkflowV1alpha1ArchiveStrategy.md) |  | [optional] 
**archive_logs** | **bool** | ArchiveLogs indicates if the container logs should be archived | [optional] 
**artifactory** | [**IonholuongutWorkflowV1alpha1ArtifactoryArtifact**](IonholuongutWorkflowV1alpha1ArtifactoryArtifact.md) |  | [optional] 
**_from** | **str** | From allows an artifact to reference an artifact from a previous step | [optional] 
**from_expression** | **str** | FromExpression, if defined, is evaluated to specify the value for the artifact | [optional] 
**gcs** | [**IonholuongutWorkflowV1alpha1GCSArtifact**](IonholuongutWorkflowV1alpha1GCSArtifact.md) |  | [optional] 
**git** | [**IonholuongutWorkflowV1alpha1GitArtifact**](IonholuongutWorkflowV1alpha1GitArtifact.md) |  | [optional] 
**global_name** | **str** | GlobalName exports an output artifact to the global scope, making it available as &#39;{{io.nholuongut.workflow.v1alpha1.outputs.artifacts.XXXX}} and in workflow.status.outputs.artifacts | [optional] 
**hdfs** | [**IonholuongutWorkflowV1alpha1HDFSArtifact**](IonholuongutWorkflowV1alpha1HDFSArtifact.md) |  | [optional] 
**http** | [**IonholuongutWorkflowV1alpha1HTTPArtifact**](IonholuongutWorkflowV1alpha1HTTPArtifact.md) |  | [optional] 
**mode** | **int** | mode bits to use on this file, must be a value between 0 and 0777 set when loading input artifacts. | [optional] 
**optional** | **bool** | Make Artifacts optional, if Artifacts doesn&#39;t generate or exist | [optional] 
**oss** | [**IonholuongutWorkflowV1alpha1OSSArtifact**](IonholuongutWorkflowV1alpha1OSSArtifact.md) |  | [optional] 
**path** | **str** | Path is the container path to the artifact | [optional] 
**raw** | [**IonholuongutWorkflowV1alpha1RawArtifact**](IonholuongutWorkflowV1alpha1RawArtifact.md) |  | [optional] 
**recurse_mode** | **bool** | If mode is set, apply the permission recursively into the artifact if it is a folder | [optional] 
**s3** | [**IonholuongutWorkflowV1alpha1S3Artifact**](IonholuongutWorkflowV1alpha1S3Artifact.md) |  | [optional] 
**sub_path** | **str** | SubPath allows an artifact to be sourced from a subpath within the specified source | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


