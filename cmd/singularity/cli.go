// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package main

import (
	"github.com/sylabs/singularity/cmd/singularity/cli"
	"github.com/sylabs/singularity/internal/pkg/buildcfg"
	"github.com/sylabs/singularity/internal/pkg/sylog"
	"github.com/sylabs/singularity/internal/pkg/util/goversion"
	useragent "github.com/sylabs/singularity/pkg/util/user-agent"
)

func main() {
	if err := goversion.Check(); err != nil {
		sylog.Fatalf("%s", err)
	}

	// In cli/singularity.go
	cli.ExecuteSingularity()
}

func init() {
	useragent.InitValue(buildcfg.PACKAGE_NAME, buildcfg.PACKAGE_VERSION)
}
