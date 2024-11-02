

# IonholuongutWorkflowV1alpha1ArtifactPaths

ArtifactPaths expands a step from a collection of artifacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**archive** | [**IonholuongutWorkflowV1alpha1ArchiveStrategy**](IonholuongutWorkflowV1alpha1ArchiveStrategy.md) |  |  [optional]
**archiveLogs** | **Boolean** | ArchiveLogs indicates if the container logs should be archived |  [optional]
**artifactory** | [**IonholuongutWorkflowV1alpha1ArtifactoryArtifact**](IonholuongutWorkflowV1alpha1ArtifactoryArtifact.md) |  |  [optional]
**from** | **String** | From allows an artifact to reference an artifact from a previous step |  [optional]
**fromExpression** | **String** | FromExpression, if defined, is evaluated to specify the value for the artifact |  [optional]
**gcs** | [**IonholuongutWorkflowV1alpha1GCSArtifact**](IonholuongutWorkflowV1alpha1GCSArtifact.md) |  |  [optional]
**git** | [**IonholuongutWorkflowV1alpha1GitArtifact**](IonholuongutWorkflowV1alpha1GitArtifact.md) |  |  [optional]
**globalName** | **String** | GlobalName exports an output artifact to the global scope, making it available as &#39;{{io.nholuongut.workflow.v1alpha1.outputs.artifacts.XXXX}} and in workflow.status.outputs.artifacts |  [optional]
**hdfs** | [**IonholuongutWorkflowV1alpha1HDFSArtifact**](IonholuongutWorkflowV1alpha1HDFSArtifact.md) |  |  [optional]
**http** | [**IonholuongutWorkflowV1alpha1HTTPArtifact**](IonholuongutWorkflowV1alpha1HTTPArtifact.md) |  |  [optional]
**mode** | **Integer** | mode bits to use on this file, must be a value between 0 and 0777 set when loading input artifacts. |  [optional]
**name** | **String** | name of the artifact. must be unique within a template&#39;s inputs/outputs. | 
**optional** | **Boolean** | Make Artifacts optional, if Artifacts doesn&#39;t generate or exist |  [optional]
**oss** | [**IonholuongutWorkflowV1alpha1OSSArtifact**](IonholuongutWorkflowV1alpha1OSSArtifact.md) |  |  [optional]
**path** | **String** | Path is the container path to the artifact |  [optional]
**raw** | [**IonholuongutWorkflowV1alpha1RawArtifact**](IonholuongutWorkflowV1alpha1RawArtifact.md) |  |  [optional]
**recurseMode** | **Boolean** | If mode is set, apply the permission recursively into the artifact if it is a folder |  [optional]
**s3** | [**IonholuongutWorkflowV1alpha1S3Artifact**](IonholuongutWorkflowV1alpha1S3Artifact.md) |  |  [optional]
**subPath** | **String** | SubPath allows an artifact to be sourced from a subpath within the specified source |  [optional]



