package testimpl

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getCloudWatchClient(t *testing.T, region string) *cloudwatch.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	require.NoError(t, err, "unable to load AWS config")
	return cloudwatch.NewFromConfig(cfg)
}

func normalizeDashboardBody(t *testing.T, body string) string {
	var parsed any
	require.NoError(t, json.Unmarshal([]byte(body), &parsed), "dashboard body should be valid JSON")
	normalized, err := json.Marshal(parsed)
	require.NoError(t, err, "dashboard body should marshal to JSON")
	return string(normalized)
}

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	t.Run("VerifyTerraformOutputs", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		id := terraform.Output(t, opts, "id")
		name := terraform.Output(t, opts, "name")
		arn := terraform.Output(t, opts, "arn")

		assert.Equal(t, name, id, "id should equal name for CloudWatch dashboard")
		assert.Regexp(t, `^arn:aws:cloudwatch:.*:dashboard/`, arn, "ARN should match CloudWatch dashboard format")
		assert.Contains(t, arn, name, "ARN should contain dashboard name")
	})

	t.Run("VerifyDashboardViaAPI", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		dashboardName := terraform.Output(t, opts, "name")
		region := terraform.Output(t, opts, "region")
		expectedBody := normalizeDashboardBody(t, terraform.Output(t, opts, "dashboard_body"))

		client := getCloudWatchClient(t, region)

		output, err := client.GetDashboard(context.TODO(), &cloudwatch.GetDashboardInput{
			DashboardName: aws.String(dashboardName),
		})
		require.NoError(t, err, "GetDashboard should succeed")
		require.NotNil(t, output, "GetDashboard output should not be nil")

		assert.Equal(t, dashboardName, aws.ToString(output.DashboardName), "dashboard name should match")
		assert.Equal(t, expectedBody, normalizeDashboardBody(t, aws.ToString(output.DashboardBody)), "dashboard body should match")
	})

	t.Run("PutDashboardAndVerifyUpdate", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		dashboardName := terraform.Output(t, opts, "name")
		region := terraform.Output(t, opts, "region")
		updatedBody := `{"widgets":[{"type":"text","x":0,"y":0,"width":24,"height":1,"properties":{"markdown":"Updated by functional test"}}]}`

		client := getCloudWatchClient(t, region)

		_, err := client.PutDashboard(context.TODO(), &cloudwatch.PutDashboardInput{
			DashboardName: aws.String(dashboardName),
			DashboardBody: aws.String(updatedBody),
		})
		require.NoError(t, err, "PutDashboard should succeed")

		output, err := client.GetDashboard(context.TODO(), &cloudwatch.GetDashboardInput{
			DashboardName: aws.String(dashboardName),
		})
		require.NoError(t, err, "GetDashboard after PutDashboard should succeed")
		assert.Equal(t, updatedBody, normalizeDashboardBody(t, aws.ToString(output.DashboardBody)), "dashboard body should reflect PutDashboard update")
	})
}

func TestComposableCompleteReadOnly(t *testing.T, ctx types.TestContext) {
	t.Run("VerifyTerraformOutputs", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		id := terraform.Output(t, opts, "id")
		name := terraform.Output(t, opts, "name")

		assert.Equal(t, name, id, "id should equal name for CloudWatch dashboard")
	})

	t.Run("VerifyDashboardExistsViaAPI", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		dashboardName := terraform.Output(t, opts, "name")
		region := terraform.Output(t, opts, "region")
		expectedBody := normalizeDashboardBody(t, terraform.Output(t, opts, "dashboard_body"))

		client := getCloudWatchClient(t, region)

		output, err := client.GetDashboard(context.TODO(), &cloudwatch.GetDashboardInput{
			DashboardName: aws.String(dashboardName),
		})
		require.NoError(t, err, "GetDashboard should succeed")
		require.NotNil(t, output, "GetDashboard output should not be nil")

		assert.Equal(t, dashboardName, aws.ToString(output.DashboardName), "dashboard name should match")
		assert.Equal(t, expectedBody, normalizeDashboardBody(t, aws.ToString(output.DashboardBody)), "dashboard body should match")
	})
}
