package v1beta1

import (
	"github.com/appscode/stash/apis/stash/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ofst "kmodules.xyz/offshoot-api/api/v1"
)

const (
	ResourceKindBackupConfigurationTemplate     = "BackupConfigurationTemplate"
	ResourcePluralBackupConfigurationTemplate   = "backupconfigurationtemplates"
	ResourceSingularBackupConfigurationTemplate = "backupconfigurationtemplate"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BackupConfigurationTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupConfigurationTemplateSpec `json:"spec,omitempty"`
}

type BackupConfigurationTemplateSpec struct {
	// RepositorySpec is used to create Repository crd for respective workload
	v1alpha1.RepositorySpec `json:",inline"`
	Schedule                string `json:"schedule,omitempty"`
	// Task specify the Task crd that specifies steps for backup process
	// +optional
	Task TaskRef `json:"task,omitempty"`
	// RetentionPolicy indicates the policy to follow to clean old backup snapshots
	RetentionPolicy v1alpha1.RetentionPolicy `json:"retentionPolicy,omitempty"`
	// RuntimeSettings allow to specify Resources, NodeSelector, Affinity, Toleration, ReadinessProbe etc.
	//+optional
	RuntimeSettings ofst.RuntimeSettings `json:"runtimeSettings,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BackupConfigurationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupConfigurationTemplate `json:"items,omitempty"`
}
