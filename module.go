package dlmanager

import (
	"github.com/admpub/nging/v4/application/library/module"

	"github.com/nging-plugins/dlmanager/pkg/handler"
)

const ID = `download`

var Module = module.Module{
	Startup: ID,
	TemplatePath: map[string]string{
		ID: `dlmanager/template/backend`,
	},
	AssetsPath: []string{},
	Navigate:   RegisterNavigate,
	Route:      handler.RegisterRoute,
}
