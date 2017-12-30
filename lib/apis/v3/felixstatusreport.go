// Copyright (c) 2017 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindFelixStatusReport     = "FelixStatusReport"
	KindFelixStatusReportList = "FelixStatusReportList"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FelixStatusReport contains information about a Felix's status.
type FelixStatusReport struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the FelixStatusReport.
	Spec FelixStatusReportSpec `json:"spec,omitempty"`
}

// FelixStatusReportSpec contains the values describing Felix status.
type FelixStatusReportSpec struct {
	Timestamp     string  `json:"timestamp,omitempty" validate:"omitempty"`
	UptimeSeconds float64 `json:"uptimeSeconds,omitempty" validate:"omitempty"`
	FirstUpdate   bool    `json:"firstUpdate"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FelixStatusReportList contains a list of FelixStatusReport resources.
type FelixStatusReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []FelixStatusReport `json:"items"`
}

// New FelixStatusReport creates a new (zeroed) FelixStatusReport struct with the TypeMetadata
// initialized to the current version.
func NewFelixStatusReport() *FelixStatusReport {
	return &FelixStatusReport{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindFelixStatusReport,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewFelixStatusReportList creates a new 9zeroed) FelixStatusReportList struct with the TypeMetadata
// initialized to the current version.
func NewFelixStatusReportList() *FelixStatusReportList {
	return &FelixStatusReportList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindFelixStatusReportList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
