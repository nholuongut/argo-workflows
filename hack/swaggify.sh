#!/usr/bin/env bash
set -eu -o pipefail

# order is important, "REPLACEME" -> "workflow"
cat \
    | sed 's/github.com.nholuongut.argo_workflows.v3.pkg.apis.workflow.v1alpha1./io.nholuongut.REPLACEME.v1alpha1./' \
    | sed 's/github.com.nholuongut.argo_events.pkg.apis.common./io.nholuongut.events.v1alpha1./' \
    | sed 's/github.com.nholuongut.argo_events.pkg.apis.eventsource.v1alpha1./io.nholuongut.events.v1alpha1./' \
    | sed 's/github.com.nholuongut.argo_events.pkg.apis.sensor.v1alpha1./io.nholuongut.events.v1alpha1./' \
    | sed 's/cronworkflow\./io.nholuongut.REPLACEME.v1alpha1./' \
    | sed 's/event\./io.nholuongut.REPLACEME.v1alpha1./' \
    | sed 's/info\./io.nholuongut.REPLACEME.v1alpha1./' \
    | sed 's/workflowarchive\./io.nholuongut.REPLACEME.v1alpha1./' \
    | sed 's/clusterworkflowtemplate\./io.nholuongut.REPLACEME.v1alpha1./' \
    | sed 's/workflowtemplate\./io.nholuongut.REPLACEME.v1alpha1./' \
    | sed 's/workflow\./io.nholuongut.REPLACEME.v1alpha1./' \
    | sed 's/io.nholuongut.REPLACEME.v1alpha1./io.nholuongut.workflow.v1alpha1./' \
    | sed 's/k8s.io./io.k8s./' \
    | sed 's/v1alpha1\.v1alpha1\./v1alpha1\./g'
