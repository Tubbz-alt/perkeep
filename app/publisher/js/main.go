//go:generate ./gensearchtypes.sh

/*
Copyright 2016 The Camlistore Authors.

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

// This package contains source code for gopherjs, to generate javascript code
// that is included in the publisher web UI.
package main

import (
	"fmt"

	"camlistore.org/pkg/blob"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

var jQuery = jquery.NewJQuery

func main() {
	//show jQuery Version on console:
	print("Your current jQuery version is: " + jQuery().Jquery)

	js.Global.Set("RenderMembers", RenderMembers)
	js.Global.Set("RenderFile", StartRenderFile)
}

func host() (string, error) {
	h := js.Global.Get("host")
	if undefOrEmptyString(h) {
		return "", fmt.Errorf("No Host in header")
	}
	return h.String(), nil
}

func scheme() (string, error) {
	s := js.Global.Get("scheme")
	if undefOrEmptyString(s) {
		return "", fmt.Errorf("No Scheme in header")
	}
	return s.String(), nil
}

func subjectBasePath() (string, error) {
	basePath := js.Global.Get("subjectBasePath")
	if undefOrEmptyString(basePath) {
		return "", fmt.Errorf("No SubjectBasePath in header")
	}
	return basePath.String(), nil
}

func subject() (blob.Ref, error) {
	o := js.Global.Get("subject")
	if undefOrEmptyString(o) {
		return blob.Ref{}, fmt.Errorf("No Subject in header")
	}
	sbj := o.String()
	br, ok := blob.Parse(sbj)
	if !ok {
		return blob.Ref{}, fmt.Errorf("invalid blobref %q", sbj)
	}
	return br, nil
}

func publishedRoot() (blob.Ref, error) {
	o := js.Global.Get("publishedRoot")
	if undefOrEmptyString(o) {
		return blob.Ref{}, fmt.Errorf("No PublishedRoot in header")
	}
	root := o.String()
	br, ok := blob.Parse(root)
	if !ok {
		return blob.Ref{}, fmt.Errorf("invalid blobref %q", root)
	}
	return br, nil
}

func pathPrefix() (string, error) {
	prefix := js.Global.Get("pathPrefix")
	if undefOrEmptyString(prefix) {
		return "", fmt.Errorf("No PathPrefix in header")
	}
	return prefix.String(), nil
}

func undefOrEmptyString(o *js.Object) bool {
	return o == js.Undefined || o.String() == ""
}
