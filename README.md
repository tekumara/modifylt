# modifylt

The default version of a launch template cannot be specified in AWS CloudFormation (see the [docs](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html)), so use modifylt!

modifylt takes command lines flags that mimic `aws ec2 modify-launch-template`.
This makes it a drop-in replacement for aws-cli when you want to keep your deployment containers lean :-)

```
NAME:
   modifylt - modify launch template

USAGE:
   modifylt [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --launch-template-id value  Launch template ID
   --default-version value     Version to set as default
   --help, -h                  show help (default: false)
```

Example:

```
modifylt --launch-template-id lt-1234567890abcedf1 --default-version 2
```

## Alternatives

- terraform
- a cloudformation custom resource
- aws cli
