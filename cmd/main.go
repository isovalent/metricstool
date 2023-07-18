// Copyright 2023 Isovalent Inc.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/isovalent/metricstool/pkg/promlinter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(promlinter.Analyzer)
}
