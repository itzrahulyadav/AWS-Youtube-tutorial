# Use metric filters to filter specific logs

1. Create a trail in cloud trail console [Docs](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-create-a-trail-using-the-console-first-time.html#creating-a-trail-in-the-console)
2. Create log group [Docs](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/send-cloudtrail-events-to-cloudwatch-logs.html#send-cloudtrail-events-to-cloudwatch-logs-console)
3. [Link](https://repost.aws/knowledge-center/cloudwatch-monitor-cloudtrail-events)
4. Create metric filters using the following pattern
```

{ ($.eventName = "StartInstances" || $.eventName = "StopInstances") && ($.responseElements.instancesSet.items[0].instanceId != "") }

```
