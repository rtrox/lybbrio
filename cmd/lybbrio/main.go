package main

import "lybbrio/internal/commands"

var (
	name      = "lybbrio"
	version   = "x.x.x"
	buildTime = "x"
	revision  = "x"
)

// @contact.name Maintainer
// @contact.url https://lybbr.io

// @license.name AGPLv3
// @license.url <github license link>
func main() {

	err := commands.Execute(commands.AppInfo{
		Name:      name,
		Version:   version,
		Revision:  revision,
		BuildTime: buildTime,
	})
	if err != nil {
		panic(err)
	}
}
