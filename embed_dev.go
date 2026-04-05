//go:build dev

package godoyourtasks

import "os"

// use this locally for hot reload of the frontend
var ClientFiles = os.DirFS("client")
