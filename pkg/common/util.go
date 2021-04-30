/*
 * Copyright 2020 The Multicluster-Scheduler Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"strings"

	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

func SplitLabelsOrAnnotations(m map[string]string) (map[string]string, map[string]string) {
	mc := make(map[string]string)
	other := make(map[string]string)
	for k, v := range m {
		if strings.HasPrefix(k, KeyPrefix) {
			mc[k] = v
		} else {
			other[k] = v
		}
	}
	return mc, other
}

func GetGroupInfoFromPodInfo(info *framework.PodInfo) (string, string, string) {
	groupName, ok := info.Pod.Annotations[AnnotationKeyGroupName]
	if !ok {
		return "", "", ""
	}
	groupCreatedAt, ok := info.Pod.Annotations[AnnotationKeyGroupCreatedAt]
	if !ok {
		return groupName, "", ""
	}
	groupPriority, ok := info.Pod.Annotations[AnnotationKeyGroupPriority]
	if !ok {
		return groupName, groupCreatedAt, ""
	}
	return groupName, groupCreatedAt, groupPriority
}