package appstatus

import (
	v1 "github.com/acorn-io/acorn/pkg/apis/internal.acorn.io/v1"
	"github.com/acorn-io/acorn/pkg/ports"
	"github.com/acorn-io/baaah/pkg/router"
	networkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

func (a *appStatusRenderer) readRouter() error {
	// reset state
	a.app.Status.AppStatus.Routers = map[string]v1.RouterStatus{}

	for routerName := range a.app.Status.AppSpec.Routers {
		s := v1.RouterStatus{
			CommonStatus: v1.CommonStatus{
				Defined:      ports.IsLinked(a.app, routerName),
				LinkOverride: ports.LinkService(a.app, routerName),
			},
		}

		ingress := &networkingv1.Ingress{}
		err := a.c.Get(a.ctx, router.Key(a.app.Status.Namespace, routerName), ingress)
		if apierrors.IsNotFound(err) {
			//ignore
		} else if err != nil {
			return err
		} else {
			s.Defined = true
		}

		s.Ready, _ = a.isServiceReady(routerName)
		s.Ready = s.Defined && s.Ready
		a.app.Status.AppStatus.Routers[routerName] = s
	}

	return nil
}