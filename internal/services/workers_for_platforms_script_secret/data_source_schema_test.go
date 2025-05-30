// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_script_secret_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_for_platforms_script_secret"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestWorkersForPlatformsScriptSecretDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workers_for_platforms_script_secret.WorkersForPlatformsScriptSecretDataSourceModel)(nil)
	schema := workers_for_platforms_script_secret.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
