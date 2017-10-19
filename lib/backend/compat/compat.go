// Copyright (c) 2016 Tigera, Inc. All rights reserved.
//
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

package compat

import (
	log "github.com/sirupsen/logrus"

	"github.com/projectcalico/libcalico-go/lib/backend/model"
)

// ToTagsLabelsRules converts a Profile KVPair to separate KVPair types for Keys,
// Labels and Rules. These separate KVPairs are used to write three separate objects
// that make up a single profile.
func ToTagsLabelsRules(d *model.KVPair) (t, l, r *model.KVPair) {
	log.Infof("d.Value is %#v", d.Value)
	p := d.Value.(*model.Profile)
	pk := d.Key.(model.ProfileKey)

	t = &model.KVPair{
		Key:   model.ProfileTagsKey{pk},
		Value: p.Tags,
	}
	l = &model.KVPair{
		Key:   model.ProfileLabelsKey{pk},
		Value: p.Labels,
	}
	r = &model.KVPair{
		Key:   model.ProfileRulesKey{pk},
		Value: &p.Rules,
	}

	return t, l, r
}
