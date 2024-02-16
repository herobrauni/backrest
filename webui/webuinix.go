//go:build linux || darwin
// +build linux darwin

//go:generate npm install
//go:generate npm run build
//go:generate gzip -r -k dist

package webui

import (
	"embed"
)

//go:embed dist
var content embed.FS
var contentPrefix = "dist"
