resource_names_map = {
  "cloudwatch_dashboard" = {
    name       = "cwdashboard1"
    max_length = 255
  }
}

logical_product_family  = "launch"
logical_product_service = "cloudwatch"
class_env               = "dev"
instance_env            = 1
instance_resource       = 1

dashboard_body = <<-EOF
{
  "widgets": [
    {
      "type": "text",
      "x": 0,
      "y": 0,
      "width": 24,
      "height": 1,
      "properties": {
        "markdown": "Example CloudWatch dashboard for testing"
      }
    }
  ]
}
EOF
