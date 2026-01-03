package memlim

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type memlimProcessor struct {
	logger *zap.Logger
	next   consumer.Metrics
}

func (p *memlimProcessor) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (p *memlimProcessor) Start(ctx context.Context, host component.Host) error {
	p.logger.Info("memlim started (pass-through)")
	return nil
}

func (p *memlimProcessor) Shutdown(ctx context.Context) error {
	p.logger.Info("memlim stopped")
	return nil
}

func (p *memlimProcessor) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	p.logger.Info("🔵 ConsumeMetrics called")
	return p.next.ConsumeMetrics(ctx, md)
}
