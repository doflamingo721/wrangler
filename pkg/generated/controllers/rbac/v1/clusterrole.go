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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	"github.com/rancher/wrangler/pkg/generic"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
)

// ClusterRoleController interface for managing ClusterRole resources.
type ClusterRoleController interface {
	generic.ControllerMeta
	ClusterRoleClient

	// OnChange runs the given handler when the controller detects a resource was changed.
	OnChange(ctx context.Context, name string, sync ClusterRoleHandler)

	// OnRemove runs the given handler when the controller detects a resource was changed.
	OnRemove(ctx context.Context, name string, sync ClusterRoleHandler)

	// Enqueue adds the resource with the given name to the worker queue of the controller.
	Enqueue(name string)

	// EnqueueAfter runs Enqueue after the provided duration.
	EnqueueAfter(name string, duration time.Duration)

	// Cache returns a cache for the resource type T.
	Cache() ClusterRoleCache
}

// ClusterRoleClient interface for managing ClusterRole resources in Kubernetes.
type ClusterRoleClient interface {
	// Create creates a new object and return the newly created Object or an error.
	Create(*v1.ClusterRole) (*v1.ClusterRole, error)

	// Update updates the object and return the newly updated Object or an error.
	Update(*v1.ClusterRole) (*v1.ClusterRole, error)

	// Delete deletes the Object in the given name.
	Delete(name string, options *metav1.DeleteOptions) error

	// Get will attempt to retrieve the resource with the specified name.
	Get(name string, options metav1.GetOptions) (*v1.ClusterRole, error)

	// List will attempt to find multiple resources.
	List(opts metav1.ListOptions) (*v1.ClusterRoleList, error)

	// Watch will start watching resources.
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	// Patch will patch the resource with the matching name.
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterRole, err error)
}

// ClusterRoleCache interface for retrieving ClusterRole resources in memory.
type ClusterRoleCache interface {
	// Get returns the resources with the specified name from the cache.
	Get(name string) (*v1.ClusterRole, error)

	// List will attempt to find resources from the Cache.
	List(selector labels.Selector) ([]*v1.ClusterRole, error)

	// AddIndexer adds  a new Indexer to the cache with the provided name.
	// If you call this after you already have data in the store, the results are undefined.
	AddIndexer(indexName string, indexer ClusterRoleIndexer)

	// GetByIndex returns the stored objects whose set of indexed values
	// for the named index includes the given indexed value.
	GetByIndex(indexName, key string) ([]*v1.ClusterRole, error)
}

// ClusterRoleHandler is function for performing any potential modifications to a ClusterRole resource.
type ClusterRoleHandler func(string, *v1.ClusterRole) (*v1.ClusterRole, error)

// ClusterRoleIndexer computes a set of indexed values for the provided object.
type ClusterRoleIndexer func(obj *v1.ClusterRole) ([]string, error)

// ClusterRoleGenericController wraps wrangler/pkg/generic.NonNamespacedController so that the function definitions adhere to ClusterRoleController interface.
type ClusterRoleGenericController struct {
	generic.NonNamespacedControllerInterface[*v1.ClusterRole, *v1.ClusterRoleList]
}

// OnChange runs the given resource handler when the controller detects a resource was changed.
func (c *ClusterRoleGenericController) OnChange(ctx context.Context, name string, sync ClusterRoleHandler) {
	c.NonNamespacedControllerInterface.OnChange(ctx, name, generic.ObjectHandler[*v1.ClusterRole](sync))
}

// OnRemove runs the given object handler when the controller detects a resource was changed.
func (c *ClusterRoleGenericController) OnRemove(ctx context.Context, name string, sync ClusterRoleHandler) {
	c.NonNamespacedControllerInterface.OnRemove(ctx, name, generic.ObjectHandler[*v1.ClusterRole](sync))
}

// Cache returns a cache of resources in memory.
func (c *ClusterRoleGenericController) Cache() ClusterRoleCache {
	return &ClusterRoleGenericCache{
		c.NonNamespacedControllerInterface.Cache(),
	}
}

// ClusterRoleGenericCache wraps wrangler/pkg/generic.NonNamespacedCache so the function definitions adhere to ClusterRoleCache interface.
type ClusterRoleGenericCache struct {
	generic.NonNamespacedCacheInterface[*v1.ClusterRole]
}

// AddIndexer adds  a new Indexer to the cache with the provided name.
// If you call this after you already have data in the store, the results are undefined.
func (c ClusterRoleGenericCache) AddIndexer(indexName string, indexer ClusterRoleIndexer) {
	c.NonNamespacedCacheInterface.AddIndexer(indexName, generic.Indexer[*v1.ClusterRole](indexer))
}
