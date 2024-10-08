package dlmanager

import (
	"github.com/coscms/webcore/library/module"

	"github.com/nging-plugins/dlmanager/application/handler"
)

const ID = `download`

var Module = module.Module{
	TemplatePath: map[string]string{
		ID: `dlmanager/template/backend`,
	},
	AssetsPath:  []string{`dlmanager/public/assets/backend`},
	Navigate:    RegisterNavigate,
	Route:       handler.RegisterRoute,
	DBSchemaVer: 0.0000,
}
