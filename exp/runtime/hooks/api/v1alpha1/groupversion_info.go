/*
Copyright 2022 The Kubernetes Authors.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"

	runtimecatalog "sigs.k8s.io/cluster-api/internal/runtime/catalog"
)

var (
	// GroupVersion is the group version identifying RuntimeHooks defined in this package
	// and their request and response types.
	GroupVersion = schema.GroupVersion{Group: "hooks.runtime.cluster.x-k8s.io", Version: "v1alpha1"}

	// catalogBuilder is used to add RuntimeHooks and their request and response types
	// to a Catalog.
	catalogBuilder = &runtimecatalog.Builder{GroupVersion: GroupVersion}

	// AddToCatalog adds RuntimeHooks defined in this package and their request and
	// response types to a catalog.
	AddToCatalog = catalogBuilder.AddToCatalog
)

func init() {
	// Add Open API definitions for RuntimeHooks request and response types in this package
	// NOTE: the GetOpenAPIDefinitions func is automatically generated by openapi-gen.
	catalogBuilder.RegisterOpenAPIDefinitions(GetOpenAPIDefinitions)
}
