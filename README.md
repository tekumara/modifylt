# modifylt

The default version of a launch template cannot be specified in AWS CloudFormation (see the [docs](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html)), so use modifylt!

modifylt input params mimic `aws ec2 modify-launch-template`, eg:

```
modifylt --launch-template-id lt-1234567890abcedf1 --default-version 2
```
