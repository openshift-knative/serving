/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "k8s.io/api/flowcontrol/v1beta1"
	flowcontrolv1beta1 "k8s.io/client-go/applyconfigurations/flowcontrol/v1beta1"
	gentype "k8s.io/client-go/gentype"
	typedflowcontrolv1beta1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
)

// fakePriorityLevelConfigurations implements PriorityLevelConfigurationInterface
type fakePriorityLevelConfigurations struct {
	*gentype.FakeClientWithListAndApply[*v1beta1.PriorityLevelConfiguration, *v1beta1.PriorityLevelConfigurationList, *flowcontrolv1beta1.PriorityLevelConfigurationApplyConfiguration]
	Fake *FakeFlowcontrolV1beta1
}

func newFakePriorityLevelConfigurations(fake *FakeFlowcontrolV1beta1) typedflowcontrolv1beta1.PriorityLevelConfigurationInterface {
	return &fakePriorityLevelConfigurations{
		gentype.NewFakeClientWithListAndApply[*v1beta1.PriorityLevelConfiguration, *v1beta1.PriorityLevelConfigurationList, *flowcontrolv1beta1.PriorityLevelConfigurationApplyConfiguration](
			fake.Fake,
			"",
			v1beta1.SchemeGroupVersion.WithResource("prioritylevelconfigurations"),
			v1beta1.SchemeGroupVersion.WithKind("PriorityLevelConfiguration"),
			func() *v1beta1.PriorityLevelConfiguration { return &v1beta1.PriorityLevelConfiguration{} },
			func() *v1beta1.PriorityLevelConfigurationList { return &v1beta1.PriorityLevelConfigurationList{} },
			func(dst, src *v1beta1.PriorityLevelConfigurationList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.PriorityLevelConfigurationList) []*v1beta1.PriorityLevelConfiguration {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1beta1.PriorityLevelConfigurationList, items []*v1beta1.PriorityLevelConfiguration) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
