package web

import "embed"

//go:embed index.html
var IndexHtml embed.FS

//go:embed assets
var Dist embed.FS
