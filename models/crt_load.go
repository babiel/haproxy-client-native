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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CrtLoad Certificate load action
//
// # Loads a certificate from a store with options
//
// swagger:model crt_load
type CrtLoad struct {

	// Certificate alias
	Alias string `json:"alias,omitempty"`

	// Certificate filename
	// Required: true
	Certificate string `json:"certificate"`

	// OCSP issuer filename
	Issuer string `json:"issuer,omitempty"`

	// Private key filename
	Key string `json:"key,omitempty"`

	// OCSP response filename
	Ocsp string `json:"ocsp,omitempty"`

	// Automatic OCSP response update
	// Enum: ["enabled","disabled"]
	// +kubebuilder:validation:Enum=enabled;disabled;
	OcspUpdate string `json:"ocsp_update,omitempty"`

	// Signed Certificate Timestamp List filename
	Sctl string `json:"sctl,omitempty"`
}

// Validate validates this crt load
func (m *CrtLoad) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOcspUpdate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrtLoad) validateCertificate(formats strfmt.Registry) error {

	if err := validate.RequiredString("certificate", "body", m.Certificate); err != nil {
		return err
	}

	return nil
}

var crtLoadTypeOcspUpdatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		crtLoadTypeOcspUpdatePropEnum = append(crtLoadTypeOcspUpdatePropEnum, v)
	}
}

const (

	// CrtLoadOcspUpdateEnabled captures enum value "enabled"
	CrtLoadOcspUpdateEnabled string = "enabled"

	// CrtLoadOcspUpdateDisabled captures enum value "disabled"
	CrtLoadOcspUpdateDisabled string = "disabled"
)

// prop value enum
func (m *CrtLoad) validateOcspUpdateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, crtLoadTypeOcspUpdatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CrtLoad) validateOcspUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.OcspUpdate) { // not required
		return nil
	}

	// value enum
	if err := m.validateOcspUpdateEnum("ocsp_update", "body", m.OcspUpdate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this crt load based on context it is used
func (m *CrtLoad) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CrtLoad) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrtLoad) UnmarshalBinary(b []byte) error {
	var res CrtLoad
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
