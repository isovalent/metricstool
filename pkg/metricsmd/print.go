// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 Isovalent Inc.

package metricsmd

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"

	"golang.org/x/exp/maps"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewCmd creates a Cobra command for generating metrics reference docs.
// It's intended to be used as an add-on to applications that have a Cobra CLI
// and expose Prometheus metrics.
func NewCmd(vp *viper.Viper, log *slog.Logger, config *Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "metrics-docs",
		Hidden:      true,
		Annotations: config.CobraAnnotations,
		ValidArgs:   maps.Keys(config.Targets),
		Args:        cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			if vp != nil {
				flags := cmd.Flags()
				if err := vp.BindPFlags(flags); err != nil {
					return err
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			target := args[0]
			reg := prometheus.NewRegistry()

			err := config.InitMetrics(target, reg, log)
			if err != nil {
				return err
			}

			var b bytes.Buffer
			if config.AutogeneratedComment {
				// Comment to inform people this file is autogenerated.
				b.WriteString("<!-- This file is autogenerated via https://github.com/isovalent/metricstool -->\n\n")
			}
			// Document title
			h := "#"
			for i := 0; i < config.HeadingLevel; i++ {
				h += "#"
			}
			b.WriteString(fmt.Sprintf("%s %s Metrics\n\n", h, config.Targets[target]))
			// Generate metrics reference
			err = Generate(reg, &b, config)
			if err != nil {
				return err
			}
			io.Copy(cmd.OutOrStdout(), &b)

			return nil
		},
	}
	return cmd
}
