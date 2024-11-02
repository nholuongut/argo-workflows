package apiserver

import (
	"github.com/nholuongut/argo-workflows/v3/config"
	"github.com/nholuongut/argo-workflows/v3/server/auth/sso"
)

var emptyConfigFunc = func() interface{} { return &Config{} }

type Config struct {
	config.Config
	// SSO in settings for single-sign on
	SSO sso.Config `json:"sso,omitempty"`
}
