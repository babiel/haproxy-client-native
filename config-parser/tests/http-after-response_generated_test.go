// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/client-native/v6/config-parser/parsers/http"
)

func TestAfterResponseshttp(t *testing.T) {
	tests := map[string]bool{
		"http-after-response allow":        true,
		"http-after-response allow if acl": true,
		"http-after-response set-header Strict-Transport-Security \"max-age=31536000\"":      true,
		"http-after-response add-header X-Header \"foo=bar\"":                                true,
		"http-after-response add-header X-Header \"foo=bar\" if acl":                         true,
		"http-after-response add-header X-Header \"foo=bar\" unless acl":                     true,
		"http-after-response allow unless acl":                                               true,
		"http-after-response del-header X-Value":                                             true,
		"http-after-response del-header X-Value -m GET":                                      true,
		"http-after-response del-header X-Value -m GET if acl":                               true,
		"http-after-response del-header X-Value -m GET unless acl":                           true,
		"http-after-response replace-header Set-Cookie (C=[^;]*);(.*) \\1;ip=%bi;\\2":        true,
		"http-after-response replace-header Set-Cookie (C=[^;]*);(.*) \\1;ip=%bi;\\2 if acl": true,
		"http-after-response replace-value Cache-control ^public$ private":                   true,
		"http-after-response replace-value Cache-control ^public$ private if acl":            true,
		"http-after-response set-status 431":                                                 true,
		"http-after-response set-status 503 reason \"SlowDown\"":                             true,
		"http-after-response set-status 500 reason \"ServiceUnavailable\" if acl":            true,
		"http-after-response set-status 500 reason \"ServiceUnavailable\" unless acl":        true,
		"http-after-response set-var(sess.last_redir) res.hdr(location)":                     true,
		"http-after-response set-var(sess.last_redir) res.hdr(location) if acl":              true,
		"http-after-response set-var(sess.last_redir) res.hdr(location) unless acl":          true,
		"http-after-response strict-mode on":                                                 true,
		"http-after-response strict-mode off":                                                true,
		"http-after-response unset-var(sess.last_redir)":                                     true,
		"http-after-response unset-var(sess.last_redir) if acl":                              true,
		"http-after-response unset-var(sess.last_redir) unless acl":                          true,
		"http-after-response set-map(map.lst) %[src] %[res.hdr(X-Value)] if value":           true,
		"http-after-response set-map(map.lst) %[src] %[res.hdr(X-Value)]":                    true,
		"http-after-response del-acl(map.lst) [src]":                                         true,
		"http-after-response del-map(map.lst) %[src] if ! value":                             true,
		"http-after-response del-map(map.lst) %[src]":                                        true,
		"http-after-response sc-add-gpc(1,2) 1":                                              true,
		"http-after-response sc-add-gpc(1,2) 1 if is-error":                                  true,
		"http-after-response sc-inc-gpc(1,2)":                                                true,
		"http-after-response sc-inc-gpc(1,2) if is-error":                                    true,
		"http-after-response sc-inc-gpc0(1)":                                                 true,
		"http-after-response sc-inc-gpc0(1) if FALSE":                                        true,
		"http-after-response sc-inc-gpc1(1)":                                                 true,
		"http-after-response sc-inc-gpc1(1) if FALSE":                                        true,
		"http-after-response sc-set-gpt(1,2) 10":                                             true,
		"http-after-response sc-set-gpt0(1) hdr(Host),lower":                                 true,
		"http-after-response sc-set-gpt0(1) 10":                                              true,
		"http-after-response sc-set-gpt0(1) hdr(Host),lower if FALSE":                        true,
		"http-after-response do-log":                                                         true,
		"http-after-response do-log if FALSE":                                                true,
		"http-after-response":                                                                false,
		"http-after-response set-header":                                                     false,
		"http-after-response set-header x-foo":                                               false,
		"http-after-response del-header":                                                     false,
		"http-after-response replace-header Set-Cookie":                                      false,
		"http-after-response replace-header Set-Cookie (C=[^;]*);(.*)":                       false,
		"http-after-response replace-value Cache-control":                                    false,
		"http-after-response replace-value Cache-control ^public$":                           false,
		"http-after-response set-status":                                                     false,
		"http-after-response set-status error":                                               false,
		"http-after-response set-var(sess.last_redir)":                                       false,
		"http-after-response strict-mode":                                                    false,
		"http-after-response strict-mode 1":                                                  false,
		"http-after-response strict-mode 0":                                                  false,
		"http-after-response unset-var()":                                                    false,
		"http-after-response set-map(map.lst) %[src]":                                        false,
		"http-after-response del-acl(map.lst)":                                               false,
		"http-after-response del-map(map.lst)":                                               false,
		"http-after-response sc-add-gpc":                                                     false,
		"http-after-response sc-inc-gpc":                                                     false,
		"http-after-response sc-inc-gpc0":                                                    false,
		"http-after-response sc-inc-gpc1":                                                    false,
		"http-after-response sc-set-gpt0(1)":                                                 false,
		"http-after-response sc-set-gpt0":                                                    false,
		"http-after-response sc-set-gpt0(1) if FALSE":                                        false,
		"---":     false,
		"--- ---": false,
	}
	parser := &http.AfterResponses{}
	for command, shouldPass := range tests {
		t.Run(command, func(t *testing.T) {
			line := strings.TrimSpace(command)
			lines := strings.SplitN(line, "\n", -1)
			var err error
			parser.Init()
			if len(lines) > 1 {
				for _, line = range lines {
					line = strings.TrimSpace(line)
					if err = ProcessLine(line, parser); err != nil {
						break
					}
				}
			} else {
				err = ProcessLine(line, parser)
			}
			if shouldPass {
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				result, err := parser.Result()
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				var returnLine string
				if result[0].Comment == "" {
					returnLine = result[0].Data
				} else {
					returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
				}
				if command != returnLine {
					t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, command))
				}
			} else {
				if err == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
				}
				_, parseErr := parser.Result()
				if parseErr == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
				}
			}
		})
	}
}
