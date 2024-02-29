// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 Isovalent Inc.

package metricsmd

import (
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
)

type LabelValues struct {
	Label  string
	Values []string
}

type LabelOverrides struct {
	Metric    string
	Overrides []LabelValues
}

type Config struct {
	CobraAnnotations map[string]string
	Targets          map[string]string // cli argument -> docs header
	LabelOverrides   []LabelOverrides
	InitMetrics      func(target string, reg *prometheus.Registry, log *slog.Logger) error
}