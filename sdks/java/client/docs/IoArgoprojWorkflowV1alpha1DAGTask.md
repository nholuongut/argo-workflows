

# IonholuongutWorkflowV1alpha1DAGTask

DAGTask represents a node in the graph during DAG execution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**arguments** | [**IonholuongutWorkflowV1alpha1Arguments**](IonholuongutWorkflowV1alpha1Arguments.md) |  |  [optional]
**continueOn** | [**IonholuongutWorkflowV1alpha1ContinueOn**](IonholuongutWorkflowV1alpha1ContinueOn.md) |  |  [optional]
**dependencies** | **List&lt;String&gt;** | Dependencies are name of other targets which this depends on |  [optional]
**depends** | **String** | Depends are name of other targets which this depends on |  [optional]
**hooks** | [**Map&lt;String, IonholuongutWorkflowV1alpha1LifecycleHook&gt;**](IonholuongutWorkflowV1alpha1LifecycleHook.md) | Hooks hold the lifecycle hook which is invoked at lifecycle of task, irrespective of the success, failure, or error status of the primary task |  [optional]
**inline** | [**IonholuongutWorkflowV1alpha1Template**](IonholuongutWorkflowV1alpha1Template.md) |  |  [optional]
**name** | **String** | Name is the name of the target | 
**onExit** | **String** | OnExit is a template reference which is invoked at the end of the template, irrespective of the success, failure, or error of the primary template. DEPRECATED: Use Hooks[exit].Template instead. |  [optional]
**template** | **String** | Name of template to execute |  [optional]
**templateRef** | [**IonholuongutWorkflowV1alpha1TemplateRef**](IonholuongutWorkflowV1alpha1TemplateRef.md) |  |  [optional]
**when** | **String** | When is an expression in which the task should conditionally execute |  [optional]
**withItems** | **List&lt;Object&gt;** | WithItems expands a task into multiple parallel tasks from the items in the list |  [optional]
**withParam** | **String** | WithParam expands a task into multiple parallel tasks from the value in the parameter, which is expected to be a JSON list. |  [optional]
**withSequence** | [**IonholuongutWorkflowV1alpha1Sequence**](IonholuongutWorkflowV1alpha1Sequence.md) |  |  [optional]



