/*
Copyright The Karmada Authors.

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

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/apps/v1alpha1"
	scheme "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// WorkloadRebalancersGetter has a method to return a WorkloadRebalancerInterface.
// A group's client should implement this interface.
type WorkloadRebalancersGetter interface {
	WorkloadRebalancers() WorkloadRebalancerInterface
}

// WorkloadRebalancerInterface has methods to work with WorkloadRebalancer resources.
type WorkloadRebalancerInterface interface {
	Create(ctx context.Context, workloadRebalancer *v1alpha1.WorkloadRebalancer, opts v1.CreateOptions) (*v1alpha1.WorkloadRebalancer, error)
	Update(ctx context.Context, workloadRebalancer *v1alpha1.WorkloadRebalancer, opts v1.UpdateOptions) (*v1alpha1.WorkloadRebalancer, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, workloadRebalancer *v1alpha1.WorkloadRebalancer, opts v1.UpdateOptions) (*v1alpha1.WorkloadRebalancer, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.WorkloadRebalancer, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.WorkloadRebalancerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkloadRebalancer, err error)
	WorkloadRebalancerExpansion
}

// workloadRebalancers implements WorkloadRebalancerInterface
type workloadRebalancers struct {
	*gentype.ClientWithList[*v1alpha1.WorkloadRebalancer, *v1alpha1.WorkloadRebalancerList]
}

// newWorkloadRebalancers returns a WorkloadRebalancers
func newWorkloadRebalancers(c *AppsV1alpha1Client) *workloadRebalancers {
	return &workloadRebalancers{
		gentype.NewClientWithList[*v1alpha1.WorkloadRebalancer, *v1alpha1.WorkloadRebalancerList](
			"workloadrebalancers",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.WorkloadRebalancer { return &v1alpha1.WorkloadRebalancer{} },
			func() *v1alpha1.WorkloadRebalancerList { return &v1alpha1.WorkloadRebalancerList{} }),
	}
}
