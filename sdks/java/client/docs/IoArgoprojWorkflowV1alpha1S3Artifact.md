

# IonholuongutWorkflowV1alpha1S3Artifact

S3Artifact is the location of an S3 artifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**accessKeySecret** | [**io.kubernetes.client.openapi.models.V1SecretKeySelector**](io.kubernetes.client.openapi.models.V1SecretKeySelector.md) |  |  [optional]
**bucket** | **String** | Bucket is the name of the bucket |  [optional]
**createBucketIfNotPresent** | [**IonholuongutWorkflowV1alpha1CreateS3BucketOptions**](IonholuongutWorkflowV1alpha1CreateS3BucketOptions.md) |  |  [optional]
**encryptionOptions** | [**IonholuongutWorkflowV1alpha1S3EncryptionOptions**](IonholuongutWorkflowV1alpha1S3EncryptionOptions.md) |  |  [optional]
**endpoint** | **String** | Endpoint is the hostname of the bucket endpoint |  [optional]
**insecure** | **Boolean** | Insecure will connect to the service with TLS |  [optional]
**key** | **String** | Key is the key in the bucket where the artifact resides |  [optional]
**region** | **String** | Region contains the optional bucket region |  [optional]
**roleARN** | **String** | RoleARN is the Amazon Resource Name (ARN) of the role to assume. |  [optional]
**secretKeySecret** | [**io.kubernetes.client.openapi.models.V1SecretKeySelector**](io.kubernetes.client.openapi.models.V1SecretKeySelector.md) |  |  [optional]
**useSDKCreds** | **Boolean** | UseSDKCreds tells the driver to figure out credentials based on sdk defaults. |  [optional]



