# Use metric filters to filter specific logs

1. Create a trail in cloud trail console [Docs](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-create-a-trail-using-the-console-first-time.html#creating-a-trail-in-the-console)
  - Select the s3 bucket
  - select the cloudwatch log group or create a new group
  - The following permissions are needed by the cloudtrails to access cloudwatch
    ```
     {
     "Version": "2012-10-17",
     "Statement": [
    {
      "Sid": "AWSCloudTrailCreateLogStream2014110",
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogStream"
      ],
      "Resource": [
        "arn:aws:logs:ap-south-1:533267257785:log-group:new-cw:log-stream:533267257785_CloudTrail_ap-south-1*"
      ]
    },
    {
      "Sid": "AWSCloudTrailPutLogEvents20141101",
      "Effect": "Allow",
      "Action": [
        "logs:PutLogEvents"
      ],
      "Resource": [
        "arn:aws:logs:ap-south-1:533267257785:log-group:new-cw:log-stream:533267257785_CloudTrail_ap-south-1*"
      ]
    }
     ]
    }

    ```
3. Create log group [Docs](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/send-cloudtrail-events-to-cloudwatch-logs.html#send-cloudtrail-events-to-cloudwatch-logs-console)
4. [Link](https://repost.aws/knowledge-center/cloudwatch-monitor-cloudtrail-events)
5. Create metric filters using the following pattern
```

{ ($.eventName = "StartInstances" || $.eventName = "StopInstances") && ($.responseElements.instancesSet.items[0].instanceId != "") }

```
