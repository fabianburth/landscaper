// Copyright 2020 Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
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

package core

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// ConditionStatus is the status of a condition.
type ConditionStatus string

// ConditionType is a string alias.
type ConditionType string

const (
	// ConditionTrue means a resource is in the condition.
	ConditionTrue ConditionStatus = "True"
	// ConditionFalse means a resource is not in the condition.
	ConditionFalse ConditionStatus = "False"
	// ConditionUnknown means Gardener can't decide if a resource is in the condition or not.
	ConditionUnknown ConditionStatus = "Unknown"
	// ConditionProgressing means the condition was seen true, failed but stayed within a predefined failure threshold.
	// In the future, we could add other intermediate conditions, e.g. ConditionDegraded.
	ConditionProgressing ConditionStatus = "Progressing"

	// ConditionCheckError is a constant for a reason in condition.
	ConditionCheckError = "ConditionCheckError"
)

// ErrorCode is a string alias.
type ErrorCode string

const (
	// ErrorUnauthorized indicates that the last error occurred due to invalid credentials.
	ErrorUnauthorized ErrorCode = "ERR_UNAUTHORIZED"
	// ErrorCleanupResources indicates that the last error occurred due to resources are stuck in deletion.
	ErrorCleanupResources ErrorCode = "ERR_CLEANUP"
	// ErrorConfigurationProblem indicates that the last error occurred due a configuration problem.
	ErrorConfigurationProblem ErrorCode = "ERR_CONFIGURATION_PROBLEM"
)

// Condition holds the information about the state of a resource.
type Condition struct {
	// DataType of the Shoot condition.
	Type ConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=ConditionType"`
	// Status of the condition, one of True, False, Unknown.
	Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime" protobuf:"bytes,3,opt,name=lastTransitionTime"`
	// Last time the condition was updated.
	LastUpdateTime metav1.Time `json:"lastUpdateTime" protobuf:"bytes,4,opt,name=lastUpdateTime"`
	// The reason for the condition's last transition.
	Reason string `json:"reason" protobuf:"bytes,5,opt,name=reason"`
	// A human readable message indicating details about the transition.
	Message string `json:"message" protobuf:"bytes,6,opt,name=message"`
	// Well-defined error codes in case the condition reports a problem.
	// +optional
	Codes []ErrorCode `json:"codes,omitempty" protobuf:"bytes,7,rep,name=codes,casttype=ErrorCode"`
}

type Operation string

const (
	ReconcileOperation Operation = "reconcile"
)

// ObjectReference is the reference to a kubernetes object.
type ObjectReference struct {
	// Name is the name of the kubernetes object.
	Name string `json:"name"`

	// Namespace is the namespace of kubernetes object.
	Namespace string `json:"namespace"`
}

// NamespacedName returns the namespaced name for the object reference
func (r *ObjectReference) NamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

// TypedObjectReference is a reference to a typed kubernetes object.
type TypedObjectReference struct {
	// APIGroup is the group for the resource being referenced.
	// If APIGroup is not specified, the specified Kind must be in the core API group.
	// For any other third-party types, APIGroup is required.
	APIGroup string `json:"apiGroup"`
	// Kind is the type of resource being referenced
	Kind string `json:"kind"`

	ObjectReference `json:",inline"`
}

// NamedObjectReference is a named reference to a specific resource.
type NamedObjectReference struct {
	// Name is the unique name of the reference.
	Name string `json:"name"`

	// Reference is the reference to an object.
	Reference ObjectReference `json:"ref"`
}

// VersionedObjectReference is a reference to a object with its last observed resource generation.
// This struct is used by status fields.
type VersionedObjectReference struct {
	ObjectReference `json:",inline"`

	// ObservedGeneration defines the last observed generation of the referenced resource.
	ObservedGeneration int64 `json:"observedGeneration"`
}

// VersionedObjectReference is a named reference to a object with its last observed resource generation.
// This struct is used by status fields.
type VersionedNamedObjectReference struct {
	// Name is the unique name of the reference.
	Name string `json:"name"`

	// Reference is the reference to an object.
	Reference VersionedObjectReference `json:"ref"`
}
