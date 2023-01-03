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

// Site Site
//
// Site configuration. Sites are considered as one service and all farms connected to that service.
// Farms are connected to service using use-backend and default_backend directives. Sites let you
// configure simple HAProxy configurations, for more advanced options use /haproxy/configuration
// endpoints.
//
// Example: {"farms":[{"balance":{"algorithm":"roundrobin"},"mode":"http","name":"www_backend","servers":[{"address":"127.0.1.1","name":"www_server","port":4567},{"address":"127.0.1.2","name":"www_server_new","port":4567}],"use_as":"default"}],"name":"test_site","service":{"http_connection_mode":"httpclose","maxconn":2000,"mode":"http"}}
//
// swagger:model site
type Site struct {

	// farms
	Farms []*SiteFarm `json:"farms,omitempty"`

	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_.:]+$
	Name string `json:"name"`

	// service
	Service *SiteService `json:"service,omitempty"`
}

// Validate validates this site
func (m *Site) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFarms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Site) validateFarms(formats strfmt.Registry) error {
	if swag.IsZero(m.Farms) { // not required
		return nil
	}

	for i := 0; i < len(m.Farms); i++ {
		if swag.IsZero(m.Farms[i]) { // not required
			continue
		}

		if m.Farms[i] != nil {
			if err := m.Farms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("farms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("farms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Site) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this site based on the context it is used
func (m *Site) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFarms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Site) contextValidateFarms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Farms); i++ {

		if m.Farms[i] != nil {
			if err := m.Farms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("farms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("farms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Site) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

	if m.Service != nil {
		if err := m.Service.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Site) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Site) UnmarshalBinary(b []byte) error {
	var res Site
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SiteFarm site farm
//
// swagger:model SiteFarm
type SiteFarm struct {

	// balance
	Balance *Balance `json:"balance,omitempty"`

	// cond
	// Enum: [if unless]
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// forwardfor
	Forwardfor *Forwardfor `json:"forwardfor,omitempty"`

	// mode
	// Enum: [http tcp]
	Mode string `json:"mode,omitempty"`

	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_.:]+$
	Name string `json:"name"`

	// servers
	Servers []*Server `json:"servers,omitempty"`

	// use as
	// Required: true
	// Enum: [default conditional]
	UseAs string `json:"use_as"`
}

// Validate validates this site farm
func (m *SiteFarm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForwardfor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseAs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteFarm) validateBalance(formats strfmt.Registry) error {
	if swag.IsZero(m.Balance) { // not required
		return nil
	}

	if m.Balance != nil {
		if err := m.Balance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("balance")
			}
			return err
		}
	}

	return nil
}

var siteFarmTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteFarmTypeCondPropEnum = append(siteFarmTypeCondPropEnum, v)
	}
}

const (

	// SiteFarmCondIf captures enum value "if"
	SiteFarmCondIf string = "if"

	// SiteFarmCondUnless captures enum value "unless"
	SiteFarmCondUnless string = "unless"
)

// prop value enum
func (m *SiteFarm) validateCondEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteFarmTypeCondPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SiteFarm) validateCond(formats strfmt.Registry) error {
	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

func (m *SiteFarm) validateForwardfor(formats strfmt.Registry) error {
	if swag.IsZero(m.Forwardfor) { // not required
		return nil
	}

	if m.Forwardfor != nil {
		if err := m.Forwardfor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forwardfor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("forwardfor")
			}
			return err
		}
	}

	return nil
}

var siteFarmTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","tcp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteFarmTypeModePropEnum = append(siteFarmTypeModePropEnum, v)
	}
}

const (

	// SiteFarmModeHTTP captures enum value "http"
	SiteFarmModeHTTP string = "http"

	// SiteFarmModeTCP captures enum value "tcp"
	SiteFarmModeTCP string = "tcp"
)

// prop value enum
func (m *SiteFarm) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteFarmTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SiteFarm) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *SiteFarm) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

func (m *SiteFarm) validateServers(formats strfmt.Registry) error {
	if swag.IsZero(m.Servers) { // not required
		return nil
	}

	for i := 0; i < len(m.Servers); i++ {
		if swag.IsZero(m.Servers[i]) { // not required
			continue
		}

		if m.Servers[i] != nil {
			if err := m.Servers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var siteFarmTypeUseAsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","conditional"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteFarmTypeUseAsPropEnum = append(siteFarmTypeUseAsPropEnum, v)
	}
}

const (

	// SiteFarmUseAsDefault captures enum value "default"
	SiteFarmUseAsDefault string = "default"

	// SiteFarmUseAsConditional captures enum value "conditional"
	SiteFarmUseAsConditional string = "conditional"
)

// prop value enum
func (m *SiteFarm) validateUseAsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteFarmTypeUseAsPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SiteFarm) validateUseAs(formats strfmt.Registry) error {

	if err := validate.RequiredString("use_as", "body", m.UseAs); err != nil {
		return err
	}

	// value enum
	if err := m.validateUseAsEnum("use_as", "body", m.UseAs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this site farm based on the context it is used
func (m *SiteFarm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBalance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateForwardfor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteFarm) contextValidateBalance(ctx context.Context, formats strfmt.Registry) error {

	if m.Balance != nil {
		if err := m.Balance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("balance")
			}
			return err
		}
	}

	return nil
}

func (m *SiteFarm) contextValidateForwardfor(ctx context.Context, formats strfmt.Registry) error {

	if m.Forwardfor != nil {
		if err := m.Forwardfor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forwardfor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("forwardfor")
			}
			return err
		}
	}

	return nil
}

func (m *SiteFarm) contextValidateServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Servers); i++ {

		if m.Servers[i] != nil {
			if err := m.Servers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteFarm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteFarm) UnmarshalBinary(b []byte) error {
	var res SiteFarm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SiteService site service
//
// swagger:model SiteService
type SiteService struct {

	// http connection mode
	// Enum: [http-tunnel httpclose forced-close http-server-close http-keep-alive]
	HTTPConnectionMode string `json:"http_connection_mode,omitempty"`

	// listeners
	Listeners []*Bind `json:"listeners,omitempty"`

	// maxconn
	Maxconn *int64 `json:"maxconn,omitempty"`

	// mode
	// Enum: [http tcp]
	Mode string `json:"mode,omitempty"`
}

// Validate validates this site service
func (m *SiteService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPConnectionMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListeners(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var siteServiceTypeHTTPConnectionModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http-tunnel","httpclose","forced-close","http-server-close","http-keep-alive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteServiceTypeHTTPConnectionModePropEnum = append(siteServiceTypeHTTPConnectionModePropEnum, v)
	}
}

const (

	// SiteServiceHTTPConnectionModeHTTPDashTunnel captures enum value "http-tunnel"
	SiteServiceHTTPConnectionModeHTTPDashTunnel string = "http-tunnel"

	// SiteServiceHTTPConnectionModeHttpclose captures enum value "httpclose"
	SiteServiceHTTPConnectionModeHttpclose string = "httpclose"

	// SiteServiceHTTPConnectionModeForcedDashClose captures enum value "forced-close"
	SiteServiceHTTPConnectionModeForcedDashClose string = "forced-close"

	// SiteServiceHTTPConnectionModeHTTPDashServerDashClose captures enum value "http-server-close"
	SiteServiceHTTPConnectionModeHTTPDashServerDashClose string = "http-server-close"

	// SiteServiceHTTPConnectionModeHTTPDashKeepDashAlive captures enum value "http-keep-alive"
	SiteServiceHTTPConnectionModeHTTPDashKeepDashAlive string = "http-keep-alive"
)

// prop value enum
func (m *SiteService) validateHTTPConnectionModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteServiceTypeHTTPConnectionModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SiteService) validateHTTPConnectionMode(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPConnectionMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPConnectionModeEnum("service"+"."+"http_connection_mode", "body", m.HTTPConnectionMode); err != nil {
		return err
	}

	return nil
}

func (m *SiteService) validateListeners(formats strfmt.Registry) error {
	if swag.IsZero(m.Listeners) { // not required
		return nil
	}

	for i := 0; i < len(m.Listeners); i++ {
		if swag.IsZero(m.Listeners[i]) { // not required
			continue
		}

		if m.Listeners[i] != nil {
			if err := m.Listeners[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service" + "." + "listeners" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service" + "." + "listeners" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var siteServiceTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","tcp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteServiceTypeModePropEnum = append(siteServiceTypeModePropEnum, v)
	}
}

const (

	// SiteServiceModeHTTP captures enum value "http"
	SiteServiceModeHTTP string = "http"

	// SiteServiceModeTCP captures enum value "tcp"
	SiteServiceModeTCP string = "tcp"
)

// prop value enum
func (m *SiteService) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteServiceTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SiteService) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("service"+"."+"mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this site service based on the context it is used
func (m *SiteService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateListeners(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteService) contextValidateListeners(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Listeners); i++ {

		if m.Listeners[i] != nil {
			if err := m.Listeners[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service" + "." + "listeners" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service" + "." + "listeners" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteService) UnmarshalBinary(b []byte) error {
	var res SiteService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
