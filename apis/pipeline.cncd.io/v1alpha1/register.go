package v1alpha1

import (
	"k8s.io/api/rbac/v1alpha1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	version = "v1alpha1"
)

// Pipeline constants
const (
	PipelineKind       = "Pipeline"
	PipelineName       = "pipeline"
	PipelineNamePlural = "pipelines"
	PipelineScope      = apiextensionsv1beta1.ClusterScoped
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: v1alpha1.GroupName, Version: version}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return VersionKind(kind).GroupKind()
}

// VersionKind takes an unqualified kind and returns back a Group qualified GroupVersionKind
func VersionKind(kind string) schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind(kind)
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Pipeline{},
		&PipelineList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
