// Code generated by github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli, DO NOT EDIT.

package webapicap

import (
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/workflows/sdk"
)

func (cfg TriggerConfig) New(w *sdk.WorkflowSpecFactory) TriggerRequestPayloadCap {
	ref := "trigger"
	def := sdk.StepDefinition{
		ID: "web-api-trigger@1.0.0", Ref: ref,
		Inputs: sdk.StepInputs{},
		Config: map[string]any{
			"allowedSenders": cfg.AllowedSenders,
			"allowedTopics":  cfg.AllowedTopics,
			"rateLimiter":    cfg.RateLimiter,
			"requiredParams": cfg.RequiredParams,
		},
		CapabilityType: capabilities.CapabilityTypeTrigger,
	}

	step := sdk.Step[TriggerRequestPayload]{Definition: def}
	return TriggerRequestPayloadCapFromStep(w, step)
}

type RateLimiterConfigCap interface {
	sdk.CapDefinition[RateLimiterConfig]
	GlobalBurst() sdk.CapDefinition[int64]
	GlobalRPS() sdk.CapDefinition[float64]
	PerSenderBurst() sdk.CapDefinition[int64]
	PerSenderRPS() sdk.CapDefinition[float64]
	private()
}

// RateLimiterConfigCapFromStep should only be called from generated code to assure type safety
func RateLimiterConfigCapFromStep(w *sdk.WorkflowSpecFactory, step sdk.Step[RateLimiterConfig]) RateLimiterConfigCap {
	raw := step.AddTo(w)
	return &rateLimiterConfig{CapDefinition: raw}
}

type rateLimiterConfig struct {
	sdk.CapDefinition[RateLimiterConfig]
}

func (*rateLimiterConfig) private() {}
func (c *rateLimiterConfig) GlobalBurst() sdk.CapDefinition[int64] {
	return sdk.AccessField[RateLimiterConfig, int64](c.CapDefinition, "globalBurst")
}
func (c *rateLimiterConfig) GlobalRPS() sdk.CapDefinition[float64] {
	return sdk.AccessField[RateLimiterConfig, float64](c.CapDefinition, "globalRPS")
}
func (c *rateLimiterConfig) PerSenderBurst() sdk.CapDefinition[int64] {
	return sdk.AccessField[RateLimiterConfig, int64](c.CapDefinition, "perSenderBurst")
}
func (c *rateLimiterConfig) PerSenderRPS() sdk.CapDefinition[float64] {
	return sdk.AccessField[RateLimiterConfig, float64](c.CapDefinition, "perSenderRPS")
}

func NewRateLimiterConfigFromFields(
	globalBurst sdk.CapDefinition[int64],
	globalRPS sdk.CapDefinition[float64],
	perSenderBurst sdk.CapDefinition[int64],
	perSenderRPS sdk.CapDefinition[float64]) RateLimiterConfigCap {
	return &simpleRateLimiterConfig{
		CapDefinition: sdk.ComponentCapDefinition[RateLimiterConfig]{
			"globalBurst":    globalBurst.Ref(),
			"globalRPS":      globalRPS.Ref(),
			"perSenderBurst": perSenderBurst.Ref(),
			"perSenderRPS":   perSenderRPS.Ref(),
		},
		globalBurst:    globalBurst,
		globalRPS:      globalRPS,
		perSenderBurst: perSenderBurst,
		perSenderRPS:   perSenderRPS,
	}
}

type simpleRateLimiterConfig struct {
	sdk.CapDefinition[RateLimiterConfig]
	globalBurst    sdk.CapDefinition[int64]
	globalRPS      sdk.CapDefinition[float64]
	perSenderBurst sdk.CapDefinition[int64]
	perSenderRPS   sdk.CapDefinition[float64]
}

func (c *simpleRateLimiterConfig) GlobalBurst() sdk.CapDefinition[int64] {
	return c.globalBurst
}
func (c *simpleRateLimiterConfig) GlobalRPS() sdk.CapDefinition[float64] {
	return c.globalRPS
}
func (c *simpleRateLimiterConfig) PerSenderBurst() sdk.CapDefinition[int64] {
	return c.perSenderBurst
}
func (c *simpleRateLimiterConfig) PerSenderRPS() sdk.CapDefinition[float64] {
	return c.perSenderRPS
}

func (c *simpleRateLimiterConfig) private() {}

type TriggerRequestPayloadCap interface {
	sdk.CapDefinition[TriggerRequestPayload]
	Params() TriggerRequestPayloadParamsCap
	Timestamp() sdk.CapDefinition[int64]
	Topics() sdk.CapDefinition[[]string]
	TriggerEventId() sdk.CapDefinition[string]
	TriggerId() sdk.CapDefinition[string]
	private()
}

// TriggerRequestPayloadCapFromStep should only be called from generated code to assure type safety
func TriggerRequestPayloadCapFromStep(w *sdk.WorkflowSpecFactory, step sdk.Step[TriggerRequestPayload]) TriggerRequestPayloadCap {
	raw := step.AddTo(w)
	return &triggerRequestPayload{CapDefinition: raw}
}

type triggerRequestPayload struct {
	sdk.CapDefinition[TriggerRequestPayload]
}

func (*triggerRequestPayload) private() {}
func (c *triggerRequestPayload) Params() TriggerRequestPayloadParamsCap {
	return TriggerRequestPayloadParamsCap(sdk.AccessField[TriggerRequestPayload, TriggerRequestPayloadParams](c.CapDefinition, "params"))
}
func (c *triggerRequestPayload) Timestamp() sdk.CapDefinition[int64] {
	return sdk.AccessField[TriggerRequestPayload, int64](c.CapDefinition, "timestamp")
}
func (c *triggerRequestPayload) Topics() sdk.CapDefinition[[]string] {
	return sdk.AccessField[TriggerRequestPayload, []string](c.CapDefinition, "topics")
}
func (c *triggerRequestPayload) TriggerEventId() sdk.CapDefinition[string] {
	return sdk.AccessField[TriggerRequestPayload, string](c.CapDefinition, "trigger_event_id")
}
func (c *triggerRequestPayload) TriggerId() sdk.CapDefinition[string] {
	return sdk.AccessField[TriggerRequestPayload, string](c.CapDefinition, "trigger_id")
}

func NewTriggerRequestPayloadFromFields(
	params TriggerRequestPayloadParamsCap,
	timestamp sdk.CapDefinition[int64],
	topics sdk.CapDefinition[[]string],
	triggerEventId sdk.CapDefinition[string],
	triggerId sdk.CapDefinition[string]) TriggerRequestPayloadCap {
	return &simpleTriggerRequestPayload{
		CapDefinition: sdk.ComponentCapDefinition[TriggerRequestPayload]{
			"params":           params.Ref(),
			"timestamp":        timestamp.Ref(),
			"topics":           topics.Ref(),
			"trigger_event_id": triggerEventId.Ref(),
			"trigger_id":       triggerId.Ref(),
		},
		params:         params,
		timestamp:      timestamp,
		topics:         topics,
		triggerEventId: triggerEventId,
		triggerId:      triggerId,
	}
}

type simpleTriggerRequestPayload struct {
	sdk.CapDefinition[TriggerRequestPayload]
	params         TriggerRequestPayloadParamsCap
	timestamp      sdk.CapDefinition[int64]
	topics         sdk.CapDefinition[[]string]
	triggerEventId sdk.CapDefinition[string]
	triggerId      sdk.CapDefinition[string]
}

func (c *simpleTriggerRequestPayload) Params() TriggerRequestPayloadParamsCap {
	return c.params
}
func (c *simpleTriggerRequestPayload) Timestamp() sdk.CapDefinition[int64] {
	return c.timestamp
}
func (c *simpleTriggerRequestPayload) Topics() sdk.CapDefinition[[]string] {
	return c.topics
}
func (c *simpleTriggerRequestPayload) TriggerEventId() sdk.CapDefinition[string] {
	return c.triggerEventId
}
func (c *simpleTriggerRequestPayload) TriggerId() sdk.CapDefinition[string] {
	return c.triggerId
}

func (c *simpleTriggerRequestPayload) private() {}

type TriggerRequestPayloadParamsCap sdk.CapDefinition[TriggerRequestPayloadParams]
