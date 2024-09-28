// Copyright 2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package info provides plugin information.
package info // import "buf.build/go/bufplugin/info"

import (
	"errors"
	"net/url"
)

// Info is information about a plugin.
type Info interface {
	// URL returns the URL for a plugin.
	//
	// Optional.
	//
	// This typically is the source control repository that contains the plugin's implementation.
	URL() *url.URL
	// License returns the license of the plugin.
	//
	// Optional.
	License() License
	// Doc returns the documentation of the plugin.
	//
	// Optional.
	Doc() Doc

	isInfo()
}

// InfoForSpec returns a new Info for the given Spec.
func InfoForSpec(spec *Spec) (Info, error) {
	if err := ValidateSpec(spec); err != nil {
		return nil, err
	}
	return nil, errors.New("TODO")
}

// *** PRIVATE ***

type info struct {
	url     *url.URL
	license License
	doc     Doc
}

func newInfo(
	url *url.URL,
	license License,
	doc Doc,
) *info {
	return &info{
		url:     url,
		license: license,
		doc:     doc,
	}
}

func (i *info) URL() *url.URL {
	return i.url
}

func (i *info) License() License {
	return i.license
}

func (i *info) Doc() Doc {
	return i.doc
}

func (*info) isInfo() {}
