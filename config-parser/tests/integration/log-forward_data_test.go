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

package integration_test

const logforward_dgrambind80443 = `
log-forward test
  dgram-bind :80,:443
`
const logforward_dgrambind10001100801000110443 = `
log-forward test
  dgram-bind 10.0.0.1:10080,10.0.0.1:10443
`
const logforward_dgrambind443interfaceeth0 = `
log-forward test
  dgram-bind :443 interface eth0
`
const logforward_dgrambind443interfaceeth1 = `
log-forward test
  dgram-bind :443 interface eth1
`
const logforward_dgrambind443interfacepppoewan = `
log-forward test
  dgram-bind :443 interface pppoe-wan
`
const logforward_dgrambind443namespaceexample = `
log-forward test
  dgram-bind :443 namespace example
`
const logforward_dgrambind443transparent = `
log-forward test
  dgram-bind :443 transparent
`