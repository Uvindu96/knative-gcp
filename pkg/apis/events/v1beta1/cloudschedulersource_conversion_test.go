/*
Copyright 2020 Google LLC

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

package v1beta1

import (
	"context"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "github.com/google/knative-gcp/pkg/apis/events/v1"
	gcptesting "github.com/google/knative-gcp/pkg/testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"
)

// These variables are used to create a 'complete' version of CloudSchedulerSource where every field is
// filled in.
var (
	// completeCloudSchedulerSource is a CloudSchedulerSource with every field filled in, except TypeMeta.
	// TypeMeta is excluded because conversions do not convert it and this variable was created to
	// test conversions.
	completeCloudSchedulerSource = &CloudSchedulerSource{
		ObjectMeta: gcptesting.CompleteObjectMeta,
		Spec: CloudSchedulerSourceSpec{
			PubSubSpec: gcptesting.CompleteV1beta1PubSubSpec,
			Location:   "location",
			Schedule:   "schedule",
			Data:       "data",
		},
		Status: CloudSchedulerSourceStatus{
			PubSubStatus: gcptesting.CompleteV1beta1PubSubStatus,
			JobName:      "jobName",
		},
	}
)

func TestCloudSchedulerSourceConversionBadType(t *testing.T) {
	good, bad := &CloudSchedulerSource{}, &CloudStorageSource{}

	if err := good.ConvertTo(context.Background(), bad); err == nil {
		t.Errorf("ConvertTo() = %#v, wanted error", bad)
	}

	if err := good.ConvertFrom(context.Background(), bad); err == nil {
		t.Errorf("ConvertFrom() = %#v, wanted error", good)
	}
}

func TestCloudSchedulerSourceConversion(t *testing.T) {
	// Just one for now, just adding the for loop for ease of future changes.
	versions := []apis.Convertible{&v1.CloudSchedulerSource{}}

	tests := []struct {
		name string
		in   *CloudSchedulerSource
	}{{
		name: "min configuration",
		in: &CloudSchedulerSource{
			ObjectMeta: metav1.ObjectMeta{
				Name:       "ps-name",
				Namespace:  "ps-ns",
				Generation: 17,
			},
			Spec: CloudSchedulerSourceSpec{},
		},
	}, {
		name: "full configuration",
		in:   completeCloudSchedulerSource,
	}}
	for _, test := range tests {
		for _, version := range versions {
			t.Run(test.name, func(t *testing.T) {
				ver := version
				// DeepCopy because we will edit it below.
				in := test.in.DeepCopy()
				if err := test.in.ConvertTo(context.Background(), ver); err != nil {
					t.Errorf("ConvertTo() = %v", err)
				}
				got := &CloudSchedulerSource{}
				if err := got.ConvertFrom(context.Background(), ver); err != nil {
					t.Errorf("ConvertFrom() = %v", err)
				}
				// IdentityStatus.ServiceAccountName in v1alpha1 and v1beta1, it doesn't exist in v1.
				// So this is not a round trip, the field will be silently removed.
				in.Status.ServiceAccountName = ""
				ignoreUsername := cmp.AllowUnexported(url.Userinfo{})
				if diff := cmp.Diff(in, got, ignoreUsername); diff != "" {
					t.Errorf("roundtrip (-want, +got) = %v", diff)
				}
			})
		}
	}
}
