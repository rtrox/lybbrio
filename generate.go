// ./generate.go
package lybbrio

import _ "lybbrio/internal/ent/runtime"

//go:generate go run -mod=mod ./internal/ent/entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen
