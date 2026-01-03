package main

import (
	"log"
	"os"

	"example.com/memlim-lab/processor/memlim"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/provider/fileprovider"
	"go.opentelemetry.io/collector/exporter/debugexporter"
	"go.opentelemetry.io/collector/processor"

	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"
	"go.opentelemetry.io/collector/service/telemetry/otelconftelemetry"

	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/receiver"
)

func main() {
	info := component.BuildInfo{
		Command:     "memlim-lab",
		Description: "A lab for testing memory limits in OpenTelemetry Collector components",
		Version:     "0.0.1",
	}

	settings := otelcol.CollectorSettings{
		BuildInfo: info,

		// Register what this binary can run (components compiled into it).
		Factories: func() (otelcol.Factories, error) {
			var err error
			f := otelcol.Factories{
				Telemetry: otelconftelemetry.NewFactory(),
			}

			f.Receivers, err = otelcol.MakeFactoryMap[receiver.Factory](
				otlpreceiver.NewFactory(),
			)
			if err != nil {
				return otelcol.Factories{}, err
			}

			f.Exporters, err = otelcol.MakeFactoryMap[exporter.Factory](
				debugexporter.NewFactory(),
			)
			if err != nil {
				return otelcol.Factories{}, err
			}

			f.Processors, err = otelcol.MakeFactoryMap[processor.Factory](
				memlim.NewFactory(),
			)
			if err != nil {
				return otelcol.Factories{}, err
			}

			return f, nil
		},

		// Tell the collector how to load config.
		ConfigProviderSettings: otelcol.ConfigProviderSettings{
			ResolverSettings: confmap.ResolverSettings{
				ProviderFactories: []confmap.ProviderFactory{
					fileprovider.NewFactory(),
				},
				DefaultScheme: "file",
			},
		},
	}

	cmd := otelcol.NewCommand(settings)
	if err := cmd.Execute(); err != nil {
		log.Println("Error executing command:", err)
		os.Exit(1)
	}
}
