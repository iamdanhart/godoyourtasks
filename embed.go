//go:build !dev

package godoyourtasks

import (
	"embed"
	"io/fs"
)

//go:embed client
var embeddedFiles embed.FS

// Use the contents of client/ as the filesystem for proper resolution
var ClientFiles, _ = fs.Sub(embeddedFiles, "client")
