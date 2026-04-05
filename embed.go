//go:build !dev

package godoyourtasks

import (
	"embed"
	"io/fs"
)

//go:embed client
var embeddedFiles embed.FS

var ClientFiles, _ = fs.Sub(embeddedFiles, "client")
