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
	v1 "k8s.io/api/authorization/v1"
	gentype "k8s.io/client-go/gentype"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
)

// fakeSubjectAccessReviews implements SubjectAccessReviewInterface
type fakeSubjectAccessReviews struct {
	*gentype.FakeClient[*v1.SubjectAccessReview]
	Fake *FakeAuthorizationV1
}

func newFakeSubjectAccessReviews(fake *FakeAuthorizationV1) authorizationv1.SubjectAccessReviewInterface {
	return &fakeSubjectAccessReviews{
		gentype.NewFakeClient[*v1.SubjectAccessReview](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("subjectaccessreviews"),
			v1.SchemeGroupVersion.WithKind("SubjectAccessReview"),
			func() *v1.SubjectAccessReview { return &v1.SubjectAccessReview{} },
		),
		fake,
	}
}
