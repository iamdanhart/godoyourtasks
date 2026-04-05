//go:build dev

package godoyourtasks

import "os"

var ClientFiles = os.DirFS("client")