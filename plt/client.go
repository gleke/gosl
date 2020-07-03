// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import "github.com/cpmech/gosl/chk"

// clientType holds the global variables and manages communication
type clientType struct {
	browser *clientBrowser
	python  *clientPython
}

var client *clientType

// Begin begins sequence of plot commands (e.g. connects to the server)
// name is the name of the plot (required)
func Begin(name string, options ...struct {
	Port   string // server port when using the browser [default = 8081]
	Python bool   // use python's matplotlib instead of plotting server
}) {
	// check
	if name == "" {
		chk.Panic("The name of the plot cannot be empty")
	}

	// get options
	port := "8081"
	python := false
	if len(options) > 0 {
		port = options[0].Port
		python = options[0].Python
		if port == "" {
			port = "8081"
		}
	}

	// handle python case
	if python {
		python := newClientPython()
		client = &clientType{browser: nil, python: python}
		return
	}

	// handle plotting server case
	browser := newClientBrowser(name, port)
	client = &clientType{browser: browser, python: nil}
}