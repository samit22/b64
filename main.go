/*
Copyright © 2023 Samit <info@samitghimire.com.np>
*/
package main

import (
	_ "embed"

	"github.com/samit22/b64/cmd"
)

//go:embed .version
var version []byte

func main() {
	cmd.Version = string(version)
	cmd.Execute()
}
