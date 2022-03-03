/*
Copyright 2020 The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/container-object-storage-interface-api/apis/objectstorage.k8s.io/v1alpha1"
)

// BucketClaimLister helps list BucketClaims.
type BucketClaimLister interface {
	// List lists all BucketClaims in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BucketClaim, err error)
	// BucketClaims returns an object that can list and get BucketClaims.
	BucketClaims(namespace string) BucketClaimNamespaceLister
	BucketClaimListerExpansion
}

// bucketClaimLister implements the BucketClaimLister interface.
type bucketClaimLister struct {
	indexer cache.Indexer
}

// NewBucketClaimLister returns a new BucketClaimLister.
func NewBucketClaimLister(indexer cache.Indexer) BucketClaimLister {
	return &bucketClaimLister{indexer: indexer}
}

// List lists all BucketClaims in the indexer.
func (s *bucketClaimLister) List(selector labels.Selector) (ret []*v1alpha1.BucketClaim, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketClaim))
	})
	return ret, err
}

// BucketClaims returns an object that can list and get BucketClaims.
func (s *bucketClaimLister) BucketClaims(namespace string) BucketClaimNamespaceLister {
	return bucketClaimNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BucketClaimNamespaceLister helps list and get BucketClaims.
type BucketClaimNamespaceLister interface {
	// List lists all BucketClaims in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.BucketClaim, err error)
	// Get retrieves the BucketClaim from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.BucketClaim, error)
	BucketClaimNamespaceListerExpansion
}

// bucketClaimNamespaceLister implements the BucketClaimNamespaceLister
// interface.
type bucketClaimNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BucketClaims in the indexer for a given namespace.
func (s bucketClaimNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BucketClaim, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketClaim))
	})
	return ret, err
}

// Get retrieves the BucketClaim from the indexer for a given namespace and name.
func (s bucketClaimNamespaceLister) Get(name string) (*v1alpha1.BucketClaim, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bucketclaim"), name)
	}
	return obj.(*v1alpha1.BucketClaim), nil
}
