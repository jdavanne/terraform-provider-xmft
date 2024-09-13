package tfhelper

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type onlyFieldsValidator struct {
	Max int
	Min int
}

// Description returns a plain text description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v onlyFieldsValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("string length must be between %d and %d", v.Min, v.Max)
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v onlyFieldsValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("string length must be between `%d` and `%d`", v.Min, v.Max)
}

// Validate runs the main validation logic of the validator, reading configuration data out of `req` and updating `resp` with diagnostics.
func (v onlyFieldsValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	// If the value is unknown or null, there is nothing to validate.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	slog.Info("onlyFieldsValidator", "req", req.Config.Raw.String())
}

func onlyFields(minLength int, maxLength int) onlyFieldsValidator {
	return onlyFieldsValidator{
		Max: maxLength,
		Min: minLength,
	}
}
