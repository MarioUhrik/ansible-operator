/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AnsibleJobSpec defines the desired state of AnsibleJob
type AnsibleJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// When set to true, the job is ran as soon as it's first noticed by the controller.
	// Note that this may happen multiple times over the job's lifetime, e.g. if the controller restarts for any reason.
	// If set to false, the execution of the job is only governed by cronSchedule.
	// Defaults to true.
	RunImmediately bool `json:"runImmediately,omitempty"`
	// Cron expression that determines how often this job runs. E.g. "0 1 * * *" to run every day at 01:00.
	// If left empty, the job doesn't run periodically.
	CronSchedule string `json:"cronSchedule,omitempty"`
}

// AnsibleJobStatus defines the observed state of AnsibleJob
type AnsibleJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AnsibleJob is the Schema for the ansiblejobs API
type AnsibleJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AnsibleJobSpec   `json:"spec,omitempty"`
	Status AnsibleJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AnsibleJobList contains a list of AnsibleJob
type AnsibleJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnsibleJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AnsibleJob{}, &AnsibleJobList{})
}
