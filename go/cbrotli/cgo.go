// Copyright 2017 Google Inc. All Rights Reserved.
//
// Distributed under MIT license.
// See file LICENSE for detail or copy at https://opensource.org/licenses/MIT

package cbrotli

// Inform golang build system that it should link brotli libraries.

// NOTE: to enable statically linking, brotlidec and brotlienc should precede brotlicommon

// #cgo LDFLAGS: -lbrotlidec
// #cgo LDFLAGS: -lbrotlienc
// #cgo LDFLAGS: -lbrotlicommon
import "C"
