//go:build executor
// +build executor

package e2e

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nholuongut/argo-workflows/v3/test/e2e/fixtures"
)

type RunAsNonRootSuite struct {
	fixtures.E2ESuite
}

func (s *RunAsNonRootSuite) TestRunAsNonRootWorkflow() {
	s.Need(fixtures.None(fixtures.Docker))
	s.Given().
		Workflow("@smoke/runasnonroot-workflow.yaml").
		When().
		SubmitWorkflow().
		WaitForWorkflow(fixtures.ToBeSucceeded)
}

func (s *RunAsNonRootSuite) TestRunAsNonRootWithOutputParams() {
	s.Need(fixtures.None(fixtures.Docker, fixtures.K8SAPI, fixtures.Kubelet))
	s.Given().
		Workflow("@smoke/runasnonroot-output-params-pipeline.yaml").
		When().
		SubmitWorkflow().
		WaitForWorkflow(fixtures.ToBeSucceeded)
}

func TestRunAsNonRootSuite(t *testing.T) {
	suite.Run(t, new(RunAsNonRootSuite))
}
