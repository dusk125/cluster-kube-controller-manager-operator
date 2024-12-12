// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	configv1 "github.com/openshift/api/config/v1"
	internal "github.com/openshift/client-go/config/applyconfigurations/internal"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ClusterMonitoringApplyConfiguration represents a declarative configuration of the ClusterMonitoring type for use
// with apply.
type ClusterMonitoringApplyConfiguration struct {
	metav1.TypeMetaApplyConfiguration    `json:",inline"`
	*metav1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                                 *ClusterMonitoringSpecApplyConfiguration `json:"spec,omitempty"`
	Status                               *configv1.ClusterMonitoringStatus        `json:"status,omitempty"`
}

// ClusterMonitoring constructs a declarative configuration of the ClusterMonitoring type for use with
// apply.
func ClusterMonitoring(name string) *ClusterMonitoringApplyConfiguration {
	b := &ClusterMonitoringApplyConfiguration{}
	b.WithName(name)
	b.WithKind("ClusterMonitoring")
	b.WithAPIVersion("config.openshift.io/v1")
	return b
}

// ExtractClusterMonitoring extracts the applied configuration owned by fieldManager from
// clusterMonitoring. If no managedFields are found in clusterMonitoring for fieldManager, a
// ClusterMonitoringApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// clusterMonitoring must be a unmodified ClusterMonitoring API object that was retrieved from the Kubernetes API.
// ExtractClusterMonitoring provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractClusterMonitoring(clusterMonitoring *configv1.ClusterMonitoring, fieldManager string) (*ClusterMonitoringApplyConfiguration, error) {
	return extractClusterMonitoring(clusterMonitoring, fieldManager, "")
}

// ExtractClusterMonitoringStatus is the same as ExtractClusterMonitoring except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractClusterMonitoringStatus(clusterMonitoring *configv1.ClusterMonitoring, fieldManager string) (*ClusterMonitoringApplyConfiguration, error) {
	return extractClusterMonitoring(clusterMonitoring, fieldManager, "status")
}

func extractClusterMonitoring(clusterMonitoring *configv1.ClusterMonitoring, fieldManager string, subresource string) (*ClusterMonitoringApplyConfiguration, error) {
	b := &ClusterMonitoringApplyConfiguration{}
	err := managedfields.ExtractInto(clusterMonitoring, internal.Parser().Type("com.github.openshift.api.config.v1.ClusterMonitoring"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(clusterMonitoring.Name)

	b.WithKind("ClusterMonitoring")
	b.WithAPIVersion("config.openshift.io/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ClusterMonitoringApplyConfiguration) WithKind(value string) *ClusterMonitoringApplyConfiguration {
	b.TypeMetaApplyConfiguration.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ClusterMonitoringApplyConfiguration) WithAPIVersion(value string) *ClusterMonitoringApplyConfiguration {
	b.TypeMetaApplyConfiguration.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ClusterMonitoringApplyConfiguration) WithName(value string) *ClusterMonitoringApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ClusterMonitoringApplyConfiguration) WithGenerateName(value string) *ClusterMonitoringApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ClusterMonitoringApplyConfiguration) WithNamespace(value string) *ClusterMonitoringApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ClusterMonitoringApplyConfiguration) WithUID(value types.UID) *ClusterMonitoringApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ClusterMonitoringApplyConfiguration) WithResourceVersion(value string) *ClusterMonitoringApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ClusterMonitoringApplyConfiguration) WithGeneration(value int64) *ClusterMonitoringApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ClusterMonitoringApplyConfiguration) WithCreationTimestamp(value apismetav1.Time) *ClusterMonitoringApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ClusterMonitoringApplyConfiguration) WithDeletionTimestamp(value apismetav1.Time) *ClusterMonitoringApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ClusterMonitoringApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ClusterMonitoringApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ClusterMonitoringApplyConfiguration) WithLabels(entries map[string]string) *ClusterMonitoringApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Labels == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *ClusterMonitoringApplyConfiguration) WithAnnotations(entries map[string]string) *ClusterMonitoringApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Annotations == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *ClusterMonitoringApplyConfiguration) WithOwnerReferences(values ...*metav1.OwnerReferenceApplyConfiguration) *ClusterMonitoringApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.ObjectMetaApplyConfiguration.OwnerReferences = append(b.ObjectMetaApplyConfiguration.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *ClusterMonitoringApplyConfiguration) WithFinalizers(values ...string) *ClusterMonitoringApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.ObjectMetaApplyConfiguration.Finalizers = append(b.ObjectMetaApplyConfiguration.Finalizers, values[i])
	}
	return b
}

func (b *ClusterMonitoringApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &metav1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *ClusterMonitoringApplyConfiguration) WithSpec(value *ClusterMonitoringSpecApplyConfiguration) *ClusterMonitoringApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *ClusterMonitoringApplyConfiguration) WithStatus(value configv1.ClusterMonitoringStatus) *ClusterMonitoringApplyConfiguration {
	b.Status = &value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *ClusterMonitoringApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Name
}