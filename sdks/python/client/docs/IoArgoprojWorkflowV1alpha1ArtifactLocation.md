# IonholuongutWorkflowV1alpha1ArtifactLocation

ArtifactLocation describes a location for a single or multiple artifacts. It is used as single artifact in the context of inputs/outputs (e.g. outputs.artifacts.artname). It is also used to describe the location of multiple artifacts such as the archive location of a single workflow step, which the executor will use as a default location to store its files.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**archive_logs** | **bool** | ArchiveLogs indicates if the container logs should be archived | [optional] 
**artifactory** | [**IonholuongutWorkflowV1alpha1ArtifactoryArtifact**](IonholuongutWorkflowV1alpha1ArtifactoryArtifact.md) |  | [optional] 
**gcs** | [**IonholuongutWorkflowV1alpha1GCSArtifact**](IonholuongutWorkflowV1alpha1GCSArtifact.md) |  | [optional] 
**git** | [**IonholuongutWorkflowV1alpha1GitArtifact**](IonholuongutWorkflowV1alpha1GitArtifact.md) |  | [optional] 
**hdfs** | [**IonholuongutWorkflowV1alpha1HDFSArtifact**](IonholuongutWorkflowV1alpha1HDFSArtifact.md) |  | [optional] 
**http** | [**IonholuongutWorkflowV1alpha1HTTPArtifact**](IonholuongutWorkflowV1alpha1HTTPArtifact.md) |  | [optional] 
**oss** | [**IonholuongutWorkflowV1alpha1OSSArtifact**](IonholuongutWorkflowV1alpha1OSSArtifact.md) |  | [optional] 
**raw** | [**IonholuongutWorkflowV1alpha1RawArtifact**](IonholuongutWorkflowV1alpha1RawArtifact.md) |  | [optional] 
**s3** | [**IonholuongutWorkflowV1alpha1S3Artifact**](IonholuongutWorkflowV1alpha1S3Artifact.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


