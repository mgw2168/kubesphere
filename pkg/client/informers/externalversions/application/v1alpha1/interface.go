/*
Copyright 2020 The KubeSphere Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubesphere.io/kubesphere/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// HelmApplications returns a HelmApplicationInformer.
	HelmApplications() HelmApplicationInformer
	// HelmApplicationVersions returns a HelmApplicationVersionInformer.
	HelmApplicationVersions() HelmApplicationVersionInformer
	// HelmCategories returns a HelmCategoryInformer.
	HelmCategories() HelmCategoryInformer
	// HelmReleases returns a HelmReleaseInformer.
	HelmReleases() HelmReleaseInformer
	// HelmRepos returns a HelmRepoInformer.
	HelmRepos() HelmRepoInformer
	// Manifests returns a ManifestInformer.
	Manifests() ManifestInformer
	// OperatorApplications returns a OperatorApplicationInformer.
	OperatorApplications() OperatorApplicationInformer
	// OperatorApplicationVersions returns a OperatorApplicationVersionInformer.
	OperatorApplicationVersions() OperatorApplicationVersionInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// HelmApplications returns a HelmApplicationInformer.
func (v *version) HelmApplications() HelmApplicationInformer {
	return &helmApplicationInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// HelmApplicationVersions returns a HelmApplicationVersionInformer.
func (v *version) HelmApplicationVersions() HelmApplicationVersionInformer {
	return &helmApplicationVersionInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// HelmCategories returns a HelmCategoryInformer.
func (v *version) HelmCategories() HelmCategoryInformer {
	return &helmCategoryInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// HelmReleases returns a HelmReleaseInformer.
func (v *version) HelmReleases() HelmReleaseInformer {
	return &helmReleaseInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// HelmRepos returns a HelmRepoInformer.
func (v *version) HelmRepos() HelmRepoInformer {
	return &helmRepoInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Manifests returns a ManifestInformer.
func (v *version) Manifests() ManifestInformer {
	return &manifestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OperatorApplications returns a OperatorApplicationInformer.
func (v *version) OperatorApplications() OperatorApplicationInformer {
	return &operatorApplicationInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OperatorApplicationVersions returns a OperatorApplicationVersionInformer.
func (v *version) OperatorApplicationVersions() OperatorApplicationVersionInformer {
	return &operatorApplicationVersionInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
