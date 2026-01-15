// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Generate a self-signed X.509 certificate for a TLS server. Outputs to
// 'cert.pem' and 'key.pem' and will overwrite existing files.

// Based on src/crypto/tls/generate_cert.go from the Go SDK
// Modified by the Gopcua Authors for use in creating an OPC-UA compliant client certificate

package main
