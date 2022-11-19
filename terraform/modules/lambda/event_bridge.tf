resource "aws_lambda_permission" "test_lambda_permission" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.test_lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.test_event_bridge.arn
}

resource "aws_cloudwatch_event_rule" "test_event_bridge" {
  name                = "test_event_bridge"
  schedule_expression = "cron(0 9 ? * MON-FRI *)"
}

resource "aws_cloudwatch_event_target" "test_event_bridge_target" {
  rule = aws_cloudwatch_event_rule.test_event_bridge.name
  arn  = aws_lambda_function.test_lambda.arn
}