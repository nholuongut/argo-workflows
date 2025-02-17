

# IonholuongutWorkflowV1alpha1WorkflowStep

WorkflowStep is a reference to a template to execute in a series of step

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**arguments** | [**IonholuongutWorkflowV1alpha1Arguments**](IonholuongutWorkflowV1alpha1Arguments.md) |  |  [optional]
**continueOn** | [**IonholuongutWorkflowV1alpha1ContinueOn**](IonholuongutWorkflowV1alpha1ContinueOn.md) |  |  [optional]
**hooks** | [**Map&lt;String, IonholuongutWorkflowV1alpha1LifecycleHook&gt;**](IonholuongutWorkflowV1alpha1LifecycleHook.md) | Hooks holds the lifecycle hook which is invoked at lifecycle of step, irrespective of the success, failure, or error status of the primary step |  [optional]
**inline** | [**IonholuongutWorkflowV1alpha1Template**](IonholuongutWorkflowV1alpha1Template.md) |  |  [optional]
**name** | **String** | Name of the step |  [optional]
**onExit** | **String** | OnExit is a template reference which is invoked at the end of the template, irrespective of the success, failure, or error of the primary template. DEPRECATED: Use Hooks[exit].Template instead. |  [optional]
**template** | **String** | Template is the name of the template to execute as the step |  [optional]
**templateRef** | [**IonholuongutWorkflowV1alpha1TemplateRef**](IonholuongutWorkflowV1alpha1TemplateRef.md) |  |  [optional]
**when** | **String** | When is an expression in which the step should conditionally execute |  [optional]
**withItems** | **List&lt;Object&gt;** | WithItems expands a step into multiple parallel steps from the items in the list |  [optional]
**withParam** | **String** | WithParam expands a step into multiple parallel steps from the value in the parameter, which is expected to be a JSON list. |  [optional]
**withSequence** | [**IonholuongutWorkflowV1alpha1Sequence**](IonholuongutWorkflowV1alpha1Sequence.md) |  |  [optional]



