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

package clientv3

import (
	"context"

	apiv3 "github.com/projectcalico/libcalico-go/lib/apis/v3"
	"github.com/projectcalico/libcalico-go/lib/options"
	validator "github.com/projectcalico/libcalico-go/lib/validator/v3"
	"github.com/projectcalico/libcalico-go/lib/watch"
)

// FelixStatusReportInterface has methods to work with FelixStatusReport resources.
type FelixStatusReportInterface interface {
	Create(ctx context.Context, res *apiv3.FelixStatusReport, opts options.SetOptions) (*apiv3.FelixStatusReport, error)
	Update(ctx context.Context, res *apiv3.FelixStatusReport, opts options.SetOptions) (*apiv3.FelixStatusReport, error)
	Delete(ctx context.Context, name string, opts options.DeleteOptions) (*apiv3.FelixStatusReport, error)
	Get(ctx context.Context, name string, opts options.GetOptions) (*apiv3.FelixStatusReport, error)
	List(ctx context.Context, opts options.ListOptions) (*apiv3.FelixStatusReportList, error)
	Watch(ctx context.Context, opts options.ListOptions) (watch.Interface, error)
}

// felixStatusReports implements FelixStatusReportInterface
type felixStatusReports struct {
	client client
}

// Create takes the representation of a FelixStatusReport and creates it.
// Returns the stored representation of the FelixStatusReport, and an error
// if there is any.
func (r felixStatusReports) Create(ctx context.Context, res *apiv3.FelixStatusReport, opts options.SetOptions) (*apiv3.FelixStatusReport, error) {
	if err := validator.Validate(res); err != nil {
		return nil, err
	}

	out, err := r.client.resources.Create(ctx, opts, apiv3.KindFelixStatusReport, res)
	if out != nil {
		return out.(*apiv3.FelixStatusReport), err
	}
	return nil, err
}

// Update takes the representation of a FelixStatusReport and updates it.
// Returns the stored representation of the FelixStatusReport, and an error
// if there is any.
func (r felixStatusReports) Update(ctx context.Context, res *apiv3.FelixStatusReport, opts options.SetOptions) (*apiv3.FelixStatusReport, error) {
	if err := validator.Validate(res); err != nil {
		return nil, err
	}

	out, err := r.client.resources.Update(ctx, opts, apiv3.KindFelixStatusReport, res)
	if out != nil {
		return out.(*apiv3.FelixStatusReport), err
	}
	return nil, err
}

// Delete takes name of the FelixStatusReport and deletes it. Returns an
// error if one occurs.
func (r felixStatusReports) Delete(ctx context.Context, name string, opts options.DeleteOptions) (*apiv3.FelixStatusReport, error) {
	out, err := r.client.resources.Delete(ctx, opts, apiv3.KindFelixStatusReport, noNamespace, name)
	if out != nil {
		return out.(*apiv3.FelixStatusReport), err
	}
	return nil, err
}

// Get takes name of the FelixStatusReport, and returns the corresponding
// FelixStatusReport object, and an error if there is any.
func (r felixStatusReports) Get(ctx context.Context, name string, opts options.GetOptions) (*apiv3.FelixStatusReport, error) {
	out, err := r.client.resources.Get(ctx, opts, apiv3.KindFelixStatusReport, noNamespace, name)
	if out != nil {
		return out.(*apiv3.FelixStatusReport), err
	}
	return nil, err
}

// List returns the list of FelixStatusReport objects that match the supplied options.
func (r felixStatusReports) List(ctx context.Context, opts options.ListOptions) (*apiv3.FelixStatusReportList, error) {
	res := &apiv3.FelixStatusReportList{}
	if err := r.client.resources.List(ctx, opts, apiv3.KindFelixStatusReport, apiv3.KindFelixStatusReportList, res); err != nil {
		return nil, err
	}
	return res, nil
}

// Watch returns a watch.Interface that watches the FelixStatusReport that
// match the supplied options.
func (r felixStatusReports) Watch(ctx context.Context, opts options.ListOptions) (watch.Interface, error) {
	return r.client.resources.Watch(ctx, opts, apiv3.KindFelixStatusReport, nil)
}
