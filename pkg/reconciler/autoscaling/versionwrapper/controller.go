package versionwrapper

import (
	"context"

	kubeclient "knative.dev/pkg/client/injection/kube/client"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/injection"
	"knative.dev/serving/pkg/reconciler/autoscaling/hpa"
	hpav2beta2 "knative.dev/serving/pkg/reconciler/autoscaling/hpav2beta2"
)

func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	kc := kubeclient.Get(ctx)
	// Starting from 4.11 (1.24.0) we can use the new API version (autoscaling/v2) of HorizontalPodAutoscaler
	// As we also need to support 4.8 we also need provide the controller using the old API version (autoscaling/v2beta2)
	if err := injection.CheckMinimumVersion(kc.Discovery(), "1.24.0"); err == nil {
		return hpa.NewController(ctx, cmw)
	} else {
		return hpav2beta2.NewController(ctx, cmw)
	}
}
