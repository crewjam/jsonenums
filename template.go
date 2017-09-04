// Copyright 2017 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

// Added as a .go file to avoid embedding issues of the template.

package main

import "text/template"

var generatedTmpl = template.Must(template.New("generated").Parse(`
// generated by jsonenums {{.Command}}; DO NOT EDIT

package {{.PackageName}}

import (
    "encoding/json"
    "fmt"
)

{{range $typename, $values := .TypesAndValues}}

var (
    _{{$typename}}NameToValue = map[string]{{$typename}} {
        {{range $values}}"{{.}}": {{.}},
        {{end}}
    }

    _{{$typename}}ValueToName = map[{{$typename}}]string {
        {{range $values}}{{.}}: "{{.}}",
        {{end}}
    }
)

// Parse{{$typename}} returns a {{$typename}} given it's string 
// representation, or an error if s is not a valid value of 
// {{$typename}}.
func Parse{{$typename}}(s string) ({{$typename}}, error) {
    v, ok := _{{$typename}}NameToValue[s]
    if ok {
        return v, nil
    }
    var zeroValue {{$typename}}
    return zeroValue, fmt.Errorf("invalid {{$typename}}: %d", s)
}

// String is generated so {{$typename}} satisfies fmt.Stringer.
func (r {{$typename}}) String() string {
    s, ok := _{{$typename}}ValueToName[r]
    if ok {
        return s
    }
    return fmt.Sprintf("{{$typename}}(%d)", r)
}


// MarshalText is generated so {{$typename}} satisfies encoding.TextMarshaler.
func (r {{$typename}}) MarshalText() ([]byte, error) {
    s, ok := _{{$typename}}ValueToName[r]
    if !ok {
        return nil, fmt.Errorf("invalid {{$typename}}: %d", r)
    }
    return []byte(s), nil
}

// UnmarshalText is generated so {{$typename}} satisfies encoding.TextUnmarshaler.
func (r *{{$typename}}) UnmarshalText(data []byte) error {
    v, ok := _{{$typename}}NameToValue[string(data)]
    if !ok {
        return fmt.Errorf("invalid {{$typename}} %q", string(data))
    }
    *r = v
    return nil
}

// MarshalJSON is generated so {{$typename}} satisfies json.Marshaler.
func (r {{$typename}}) MarshalJSON() ([]byte, error) {
    s, ok := _{{$typename}}ValueToName[r]
    if !ok {
        return nil, fmt.Errorf("invalid {{$typename}}: %d", r)
    }
    return json.Marshal(s)
}

// UnmarshalJSON is generated so {{$typename}} satisfies json.Unmarshaler.
func (r *{{$typename}}) UnmarshalJSON(data []byte) error {
    var s string
    if err := json.Unmarshal(data, &s); err != nil {
        return fmt.Errorf("{{$typename}} should be a string, got %s", data)
    }
    v, ok := _{{$typename}}NameToValue[s]
    if !ok {
        return fmt.Errorf("invalid {{$typename}} %q", s)
    }
    *r = v
    return nil
}

{{end}}
`))
