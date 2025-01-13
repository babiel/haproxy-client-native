// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
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
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatsOptions stats options
//
// swagger:model stats_options
type StatsOptions struct {

	// stats admin
	StatsAdmin bool `json:"stats_admin,omitempty"`

	// stats admin cond
	// Enum: ["if","unless"]
	// +kubebuilder:validation:Enum=if;unless;
	StatsAdminCond string `json:"stats_admin_cond,omitempty"`

	// stats admin cond test
	StatsAdminCondTest string `json:"stats_admin_cond_test,omitempty"`

	// stats auths
	StatsAuths []*StatsAuth `json:"stats_auths,omitempty"`

	// stats enable
	StatsEnable bool `json:"stats_enable,omitempty"`

	// stats hide version
	StatsHideVersion bool `json:"stats_hide_version,omitempty"`

	// stats http requests
	StatsHTTPRequests []*StatsHTTPRequest `json:"stats_http_requests,omitempty"`

	// stats maxconn
	// Minimum: 1
	// +kubebuilder:validation:Minimum=1
	StatsMaxconn int64 `json:"stats_maxconn,omitempty"`

	// stats realm
	StatsRealm bool `json:"stats_realm,omitempty"`

	// stats realm realm
	StatsRealmRealm *string `json:"stats_realm_realm,omitempty"`

	// stats refresh delay
	// Minimum: 0
	// +kubebuilder:validation:Minimum=0
	StatsRefreshDelay *int64 `json:"stats_refresh_delay,omitempty"`

	// stats show desc
	StatsShowDesc *string `json:"stats_show_desc,omitempty"`

	// stats show legends
	StatsShowLegends bool `json:"stats_show_legends,omitempty"`

	// stats show modules
	StatsShowModules bool `json:"stats_show_modules,omitempty"`

	// stats show node name
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	StatsShowNodeName *string `json:"stats_show_node_name"`

	// stats uri prefix
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	StatsURIPrefix string `json:"stats_uri_prefix,omitempty"`
}

// Validate validates this stats options
func (m *StatsOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatsAdminCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatsAuths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatsHTTPRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatsMaxconn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatsRefreshDelay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatsShowNodeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatsURIPrefix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var statsOptionsTypeStatsAdminCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statsOptionsTypeStatsAdminCondPropEnum = append(statsOptionsTypeStatsAdminCondPropEnum, v)
	}
}

const (

	// StatsOptionsStatsAdminCondIf captures enum value "if"
	StatsOptionsStatsAdminCondIf string = "if"

	// StatsOptionsStatsAdminCondUnless captures enum value "unless"
	StatsOptionsStatsAdminCondUnless string = "unless"
)

// prop value enum
func (m *StatsOptions) validateStatsAdminCondEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statsOptionsTypeStatsAdminCondPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatsOptions) validateStatsAdminCond(formats strfmt.Registry) error {
	if swag.IsZero(m.StatsAdminCond) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatsAdminCondEnum("stats_admin_cond", "body", m.StatsAdminCond); err != nil {
		return err
	}

	return nil
}

func (m *StatsOptions) validateStatsAuths(formats strfmt.Registry) error {
	if swag.IsZero(m.StatsAuths) { // not required
		return nil
	}

	for i := 0; i < len(m.StatsAuths); i++ {
		if swag.IsZero(m.StatsAuths[i]) { // not required
			continue
		}

		if m.StatsAuths[i] != nil {
			if err := m.StatsAuths[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stats_auths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stats_auths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StatsOptions) validateStatsHTTPRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.StatsHTTPRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.StatsHTTPRequests); i++ {
		if swag.IsZero(m.StatsHTTPRequests[i]) { // not required
			continue
		}

		if m.StatsHTTPRequests[i] != nil {
			if err := m.StatsHTTPRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stats_http_requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stats_http_requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StatsOptions) validateStatsMaxconn(formats strfmt.Registry) error {
	if swag.IsZero(m.StatsMaxconn) { // not required
		return nil
	}

	if err := validate.MinimumInt("stats_maxconn", "body", m.StatsMaxconn, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *StatsOptions) validateStatsRefreshDelay(formats strfmt.Registry) error {
	if swag.IsZero(m.StatsRefreshDelay) { // not required
		return nil
	}

	if err := validate.MinimumInt("stats_refresh_delay", "body", *m.StatsRefreshDelay, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *StatsOptions) validateStatsShowNodeName(formats strfmt.Registry) error {
	if swag.IsZero(m.StatsShowNodeName) { // not required
		return nil
	}

	if err := validate.Pattern("stats_show_node_name", "body", *m.StatsShowNodeName, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *StatsOptions) validateStatsURIPrefix(formats strfmt.Registry) error {
	if swag.IsZero(m.StatsURIPrefix) { // not required
		return nil
	}

	if err := validate.Pattern("stats_uri_prefix", "body", m.StatsURIPrefix, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this stats options based on the context it is used
func (m *StatsOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatsAuths(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatsHTTPRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatsOptions) contextValidateStatsAuths(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatsAuths); i++ {

		if m.StatsAuths[i] != nil {

			if swag.IsZero(m.StatsAuths[i]) { // not required
				return nil
			}

			if err := m.StatsAuths[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stats_auths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stats_auths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StatsOptions) contextValidateStatsHTTPRequests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatsHTTPRequests); i++ {

		if m.StatsHTTPRequests[i] != nil {

			if swag.IsZero(m.StatsHTTPRequests[i]) { // not required
				return nil
			}

			if err := m.StatsHTTPRequests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stats_http_requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stats_http_requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatsOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatsOptions) UnmarshalBinary(b []byte) error {
	var res StatsOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
