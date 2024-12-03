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

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StickTableEntry Stick Table Entry
//
// # One entry in stick table
//
// swagger:model stick_table_entry
type StickTableEntry struct {

	// bytes in cnt
	BytesInCnt *int64 `json:"bytes_in_cnt,omitempty"`

	// bytes in rate
	BytesInRate *int64 `json:"bytes_in_rate,omitempty"`

	// bytes out cnt
	BytesOutCnt *int64 `json:"bytes_out_cnt,omitempty"`

	// bytes out rate
	BytesOutRate *int64 `json:"bytes_out_rate,omitempty"`

	// conn cnt
	ConnCnt *int64 `json:"conn_cnt,omitempty"`

	// conn cur
	ConnCur *int64 `json:"conn_cur,omitempty"`

	// conn rate
	ConnRate *int64 `json:"conn_rate,omitempty"`

	// exp
	Exp *int64 `json:"exp,omitempty"`

	// glitch cnt
	GlitchCnt *int64 `json:"glitch_cnt,omitempty"`

	// glitch rate
	GlitchRate *int64 `json:"glitch_rate,omitempty"`

	// gpc0
	Gpc0 *int64 `json:"gpc0,omitempty"`

	// gpc0 rate
	Gpc0Rate *int64 `json:"gpc0_rate,omitempty"`

	// gpc1
	Gpc1 *int64 `json:"gpc1,omitempty"`

	// gpc1 rate
	Gpc1Rate *int64 `json:"gpc1_rate,omitempty"`

	// gpt0
	Gpt0 *int64 `json:"gpt0,omitempty"`

	// http err cnt
	HTTPErrCnt *int64 `json:"http_err_cnt,omitempty"`

	// http err rate
	HTTPErrRate *int64 `json:"http_err_rate,omitempty"`

	// http req cnt
	HTTPReqCnt *int64 `json:"http_req_cnt,omitempty"`

	// http req rate
	HTTPReqRate *int64 `json:"http_req_rate,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// server id
	ServerID *int64 `json:"server_id,omitempty"`

	// sess cnt
	SessCnt *int64 `json:"sess_cnt,omitempty"`

	// sess rate
	SessRate *int64 `json:"sess_rate,omitempty"`

	// use
	Use bool `json:"use,omitempty"`
}

// Validate validates this stick table entry
func (m *StickTableEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stick table entry based on context it is used
func (m *StickTableEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StickTableEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StickTableEntry) UnmarshalBinary(b []byte) error {
	var res StickTableEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
