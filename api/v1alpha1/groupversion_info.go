/*
Copyright 2024.

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

// Package v1alpha1 contains API Schema definitions for the argoproj.io v1alpha1 API group
// +kubebuilder:object:generate=true
// +groupName=argoproj.io
package v1alpha1

import (
	application "github.com/kkb0318/argo-cd-api/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: application.Group, Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion                   = schema.GroupVersion{Group: application.Group, Version: "v1alpha1"}
	ApplicationSchemaGroupVersionKind    = schema.GroupVersionKind{Group: application.Group, Version: "v1alpha1", Kind: application.ApplicationKind}
	AppProjectSchemaGroupVersionKind     = schema.GroupVersionKind{Group: application.Group, Version: "v1alpha1", Kind: application.AppProjectKind}
	ApplicationSetSchemaGroupVersionKind = schema.GroupVersionKind{Group: application.Group, Version: "v1alpha1", Kind: application.ApplicationSetKind}
)

// Resource takes an unqualified resource and returns a Group-qualified GroupResource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// addKnownTypes adds the set of types defined in this package to the supplied scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Application{},
		&ApplicationList{},
		// &AppProject{},
		// &AppProjectList{},
		// &ApplicationSet{},
		// &ApplicationSetList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
