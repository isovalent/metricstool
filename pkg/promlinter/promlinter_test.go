// Copyright 2020 Authors of promlinter
// Copyright 2023 Isovalent Inc.
// SPDX-License-Identifier: Apache-2.0

// Part of this file is copied from https://github.com/yeya24/promlinter/blob/60c138a6e5b7f18dcb76c3944613722dbf84bffc/promlinter_test.go

package promlinter

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestRun(t *testing.T) {
	// The expected linter outputs are read from comments starting with "// want"
	analysistest.Run(t, analysistest.TestData(), Analyzer)
}
