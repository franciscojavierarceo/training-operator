// Copyright 2021 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
)

func TestValidateV1MpiJobSpec(t *testing.T) {
	testCases := []MPIJobSpec{
		{
			MPIReplicaSpecs: nil,
		},
		{
			MPIReplicaSpecs: map[ReplicaType]*ReplicaSpec{
				MPIJobReplicaTypeLauncher: &ReplicaSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{},
						},
					},
				},
			},
		},
		{
			MPIReplicaSpecs: map[ReplicaType]*ReplicaSpec{
				MPIJobReplicaTypeLauncher: &ReplicaSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								corev1.Container{
									Image: "",
								},
							},
						},
					},
				},
			},
		},
		{
			MPIReplicaSpecs: map[ReplicaType]*ReplicaSpec{
				MPIJobReplicaTypeLauncher: &ReplicaSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								corev1.Container{
									Name:  "",
									Image: "kubeflow/tf-dist-mnist-test:1.0",
								},
							},
						},
					},
				},
			},
		},
		{
			MPIReplicaSpecs: map[ReplicaType]*ReplicaSpec{
				MPIJobReplicaTypeLauncher: &ReplicaSpec{
					Replicas: ptr.To[int32](2),
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								corev1.Container{
									Name:  "tensorflow",
									Image: "kubeflow/tf-dist-mnist-test:1.0",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, c := range testCases {
		err := ValidateV1MpiJobSpec(&c)
		if err == nil {
			t.Error("Failed validate the kubeflowv1.MpiJobSpec")
		}
	}
}
