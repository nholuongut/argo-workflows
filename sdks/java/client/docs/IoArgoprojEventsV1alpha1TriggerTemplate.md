

# IonholuongutEventsV1alpha1TriggerTemplate

TriggerTemplate is the template that describes trigger specification.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**argoWorkflow** | [**IonholuongutEventsV1alpha1ArgoWorkflowTrigger**](IonholuongutEventsV1alpha1ArgoWorkflowTrigger.md) |  |  [optional]
**awsLambda** | [**IonholuongutEventsV1alpha1AWSLambdaTrigger**](IonholuongutEventsV1alpha1AWSLambdaTrigger.md) |  |  [optional]
**azureEventHubs** | [**IonholuongutEventsV1alpha1AzureEventHubsTrigger**](IonholuongutEventsV1alpha1AzureEventHubsTrigger.md) |  |  [optional]
**conditions** | **String** |  |  [optional]
**custom** | [**IonholuongutEventsV1alpha1CustomTrigger**](IonholuongutEventsV1alpha1CustomTrigger.md) |  |  [optional]
**http** | [**IonholuongutEventsV1alpha1HTTPTrigger**](IonholuongutEventsV1alpha1HTTPTrigger.md) |  |  [optional]
**k8s** | [**IonholuongutEventsV1alpha1StandardK8STrigger**](IonholuongutEventsV1alpha1StandardK8STrigger.md) |  |  [optional]
**kafka** | [**IonholuongutEventsV1alpha1KafkaTrigger**](IonholuongutEventsV1alpha1KafkaTrigger.md) |  |  [optional]
**log** | [**IonholuongutEventsV1alpha1LogTrigger**](IonholuongutEventsV1alpha1LogTrigger.md) |  |  [optional]
**name** | **String** | Name is a unique name of the action to take. |  [optional]
**nats** | [**IonholuongutEventsV1alpha1NATSTrigger**](IonholuongutEventsV1alpha1NATSTrigger.md) |  |  [optional]
**openWhisk** | [**IonholuongutEventsV1alpha1OpenWhiskTrigger**](IonholuongutEventsV1alpha1OpenWhiskTrigger.md) |  |  [optional]
**pulsar** | [**IonholuongutEventsV1alpha1PulsarTrigger**](IonholuongutEventsV1alpha1PulsarTrigger.md) |  |  [optional]
**slack** | [**IonholuongutEventsV1alpha1SlackTrigger**](IonholuongutEventsV1alpha1SlackTrigger.md) |  |  [optional]



