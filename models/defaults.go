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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Defaults Defaults with all it's children resources
//
// swagger:model Defaults
type Defaults struct {
	DefaultsBase `json:",inline"`

	// HTTP check list
	HTTPCheckList HTTPChecks `json:"http_check_list,omitempty"`

	// HTTP error rule list
	HTTPErrorRuleList HTTPErrorRules `json:"http_error_rule_list,omitempty"`

	// log target list
	LogTargetList LogTargets `json:"log_target_list,omitempty"`

	// q UI c initial rule list
	QUICInitialRuleList QUICInitialRules `json:"quic_initial_rule_list,omitempty"`

	// TCP check rule list
	TCPCheckRuleList TCPChecks `json:"tcp_check_rule_list,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Defaults) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 DefaultsBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.DefaultsBase = aO0

	// AO1
	var dataAO1 struct {
		HTTPCheckList HTTPChecks `json:"http_check_list,omitempty"`

		HTTPErrorRuleList HTTPErrorRules `json:"http_error_rule_list,omitempty"`

		LogTargetList LogTargets `json:"log_target_list,omitempty"`

		QUICInitialRuleList QUICInitialRules `json:"quic_initial_rule_list,omitempty"`

		TCPCheckRuleList TCPChecks `json:"tcp_check_rule_list,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.HTTPCheckList = dataAO1.HTTPCheckList

	m.HTTPErrorRuleList = dataAO1.HTTPErrorRuleList

	m.LogTargetList = dataAO1.LogTargetList

	m.QUICInitialRuleList = dataAO1.QUICInitialRuleList

	m.TCPCheckRuleList = dataAO1.TCPCheckRuleList

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Defaults) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.DefaultsBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		HTTPCheckList HTTPChecks `json:"http_check_list,omitempty"`

		HTTPErrorRuleList HTTPErrorRules `json:"http_error_rule_list,omitempty"`

		LogTargetList LogTargets `json:"log_target_list,omitempty"`

		QUICInitialRuleList QUICInitialRules `json:"quic_initial_rule_list,omitempty"`

		TCPCheckRuleList TCPChecks `json:"tcp_check_rule_list,omitempty"`
	}

	dataAO1.HTTPCheckList = m.HTTPCheckList

	dataAO1.HTTPErrorRuleList = m.HTTPErrorRuleList

	dataAO1.LogTargetList = m.LogTargetList

	dataAO1.QUICInitialRuleList = m.QUICInitialRuleList

	dataAO1.TCPCheckRuleList = m.TCPCheckRuleList

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this defaults
func (m *Defaults) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DefaultsBase
	if err := m.DefaultsBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPCheckList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPErrorRuleList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogTargetList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQUICInitialRuleList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTCPCheckRuleList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Defaults) validateHTTPCheckList(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPCheckList) { // not required
		return nil
	}

	if err := m.HTTPCheckList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("http_check_list")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("http_check_list")
		}
		return err
	}

	return nil
}

func (m *Defaults) validateHTTPErrorRuleList(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPErrorRuleList) { // not required
		return nil
	}

	if err := m.HTTPErrorRuleList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("http_error_rule_list")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("http_error_rule_list")
		}
		return err
	}

	return nil
}

func (m *Defaults) validateLogTargetList(formats strfmt.Registry) error {

	if swag.IsZero(m.LogTargetList) { // not required
		return nil
	}

	if err := m.LogTargetList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("log_target_list")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("log_target_list")
		}
		return err
	}

	return nil
}

func (m *Defaults) validateQUICInitialRuleList(formats strfmt.Registry) error {

	if swag.IsZero(m.QUICInitialRuleList) { // not required
		return nil
	}

	if err := m.QUICInitialRuleList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("quic_initial_rule_list")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("quic_initial_rule_list")
		}
		return err
	}

	return nil
}

func (m *Defaults) validateTCPCheckRuleList(formats strfmt.Registry) error {

	if swag.IsZero(m.TCPCheckRuleList) { // not required
		return nil
	}

	if err := m.TCPCheckRuleList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tcp_check_rule_list")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("tcp_check_rule_list")
		}
		return err
	}

	return nil
}

// ContextValidate validate this defaults based on the context it is used
func (m *Defaults) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DefaultsBase
	if err := m.DefaultsBase.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPCheckList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPErrorRuleList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogTargetList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQUICInitialRuleList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTCPCheckRuleList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Defaults) contextValidateHTTPCheckList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.HTTPCheckList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("http_check_list")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("http_check_list")
		}
		return err
	}

	return nil
}

func (m *Defaults) contextValidateHTTPErrorRuleList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.HTTPErrorRuleList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("http_error_rule_list")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("http_error_rule_list")
		}
		return err
	}

	return nil
}

func (m *Defaults) contextValidateLogTargetList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LogTargetList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("log_target_list")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("log_target_list")
		}
		return err
	}

	return nil
}

func (m *Defaults) contextValidateQUICInitialRuleList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.QUICInitialRuleList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("quic_initial_rule_list")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("quic_initial_rule_list")
		}
		return err
	}

	return nil
}

func (m *Defaults) contextValidateTCPCheckRuleList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TCPCheckRuleList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tcp_check_rule_list")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("tcp_check_rule_list")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Defaults) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Defaults) UnmarshalBinary(b []byte) error {
	var res Defaults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
