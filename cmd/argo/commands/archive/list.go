package archive

import (
	"os"
	"sort"

	"github.com/nholuongut/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/nholuongut/argo-workflows/v3/cmd/argo/commands/client"
	workflowarchivepkg "github.com/nholuongut/argo-workflows/v3/pkg/apiclient/workflowarchive"
	wfv1 "github.com/nholuongut/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/nholuongut/argo-workflows/v3/util/printer"
)

func NewListCommand() *cobra.Command {
	var (
		selector  string
		output    string
		chunkSize int64
	)
	command := &cobra.Command{
		Use:   "list",
		Short: "list workflows in the archive",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, apiClient := client.NewAPIClient(cmd.Context())
			serviceClient, err := apiClient.NewArchivedWorkflowServiceClient()
			errors.CheckError(err)
			namespace := client.Namespace()
			listOpts := &metav1.ListOptions{
				FieldSelector: "metadata.namespace=" + namespace,
				LabelSelector: selector,
				Limit:         chunkSize,
			}
			var workflows wfv1.Workflows
			for {
				log.WithField("listOpts", listOpts).Debug()
				resp, err := serviceClient.ListArchivedWorkflows(ctx, &workflowarchivepkg.ListArchivedWorkflowsRequest{ListOptions: listOpts})
				errors.CheckError(err)
				workflows = append(workflows, resp.Items...)
				if resp.Continue == "" {
					break
				}
				listOpts.Continue = resp.Continue
			}
			sort.Sort(workflows)
			err = printer.PrintWorkflows(workflows, os.Stdout, printer.PrintOpts{Output: output, Namespace: true, UID: true})
			errors.CheckError(err)
		},
	}
	command.Flags().StringVarP(&output, "output", "o", "wide", "Output format. One of: json|yaml|wide")
	command.Flags().StringVarP(&selector, "selector", "l", "", "Selector (label query) to filter on, not including uninitialized ones, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	command.Flags().Int64VarP(&chunkSize, "chunk-size", "", 0, "Return large lists in chunks rather than all at once. Pass 0 to disable.")
	return command
}
