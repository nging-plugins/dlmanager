package dlmanager

import (
	"github.com/coscms/webcore/library/module"
	"github.com/nging-plugins/dlmanager/application/handler"
)

var LeftNavigate = handler.LeftNavigate

func RegisterNavigate(nc module.Navigate) {
	nc.Backend().AddLeftItems(-1, LeftNavigate)
}
