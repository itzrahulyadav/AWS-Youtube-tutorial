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
2. Create log group [Docs](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/send-cloudtrail-events-to-cloudwatch-logs.html#send-cloudtrail-events-to-cloudwatch-logs-console)
   - Go to the log groups which was created
   - Enable the metric filters
3. Create metric filters using the following pattern
```

{    $.eventName = "StartInstances" || $.eventName = "StopInstances"  }

```

4. Provide metric namespace
5. Provide metric name
6. Give the metric value as 1
7. Click on create
8. Create cloudwatch alarm
9. Select the SNS topic



Refer to the following link for other example [Link](https://repost.aws/knowledge-center/cloudwatch-monitor-cloudtrail-events)
