/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"context"
	"fmt"
	dynamic_controller "kubedb.dev/provider-aws/cmd/dynamic-controller"
	"os"
	"path/filepath"

	"github.com/crossplane/upjet/pkg/pipeline"

	"kubedb.dev/provider-aws/config"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}
	p, err := config.GetProvider(context.Background(), true)
	pipeline.Run(p, absRootDir)
	dynamic_controller.GenerateController(p, absRootDir)
}
