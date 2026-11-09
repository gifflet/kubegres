/*
Copyright 2023 Reactive Tech Limited.
"Reactive Tech Limited" is a company located in England, United Kingdom.
https://www.reactive-tech.io

Lead Developer: Alex Arica

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

package statefulset

import (
	"reactive-tech.io/kubegres/internal/controller/states/statefulset"
)

// nextInstanceIndex returns the instance index to use for a new StatefulSet.
//
// The last created index stored in the Kubegres status can lag behind the cluster when the status update which
// followed a deployment was rejected by Kubernetes (optimistic lock). In that case, the deployed StatefulSets are
// the source of truth: the new index must also be greater than every deployed one, otherwise the creation keeps
// failing with "already exists" (see issues #98 and #123).
func nextInstanceIndex(lastCreatedInstanceIndex int32, deployedStatefulSets []statefulset.StatefulSetWrapper) int32 {
	instanceIndex := lastCreatedInstanceIndex
	for _, deployedStatefulSet := range deployedStatefulSets {
		if deployedStatefulSet.InstanceIndex > instanceIndex {
			instanceIndex = deployedStatefulSet.InstanceIndex
		}
	}
	return instanceIndex + 1
}
