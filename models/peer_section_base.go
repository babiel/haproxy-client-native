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
	"github.com/go-openapi/validate"
)

// PeerSectionBase Peer Section Base
//
// HAProxy peer_section configuration
//
// swagger:model peer_section_base
type PeerSectionBase struct {

	// default bind
	DefaultBind *DefaultBind `json:"default_bind,omitempty"`

	// default server
	DefaultServer *DefaultServer `json:"default_server,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_.:]+$
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9-_.:]+$`
	Name string `json:"name"`

	// In some configurations, one would like to distribute the stick-table contents
	// to some peers in place of sending all the stick-table contents to each peer
	// declared in the "peers" section. In such cases, "shards" specifies the
	// number of peer involved in this stick-table contents distribution.
	Shards int64 `json:"shards,omitempty"`
}

// Validate validates this peer section base
func (m *PeerSectionBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultBind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeerSectionBase) validateDefaultBind(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultBind) { // not required
		return nil
	}

	if m.DefaultBind != nil {
		if err := m.DefaultBind.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_bind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_bind")
			}
			return err
		}
	}

	return nil
}

func (m *PeerSectionBase) validateDefaultServer(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultServer) { // not required
		return nil
	}

	if m.DefaultServer != nil {
		if err := m.DefaultServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_server")
			}
			return err
		}
	}

	return nil
}

func (m *PeerSectionBase) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this peer section base based on the context it is used
func (m *PeerSectionBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultBind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeerSectionBase) contextValidateDefaultBind(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultBind != nil {
		if err := m.DefaultBind.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_bind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_bind")
			}
			return err
		}
	}

	return nil
}

func (m *PeerSectionBase) contextValidateDefaultServer(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultServer != nil {
		if err := m.DefaultServer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeerSectionBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeerSectionBase) UnmarshalBinary(b []byte) error {
	var res PeerSectionBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}