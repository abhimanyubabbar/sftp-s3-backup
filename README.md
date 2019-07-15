# sftp-s3-backup
A simple utility to backup the files on sftp to s3


## Usage:
It exposes a simple utility to backup the files onto S3 from the sftp. The usage pattern is as follows:

```
package main

import (
  "log",
  backup "github.com/babbarshaer/sftp-s3-backup"
)

func main() {

  config := backup.Config{
    User:              "user",
    Address:           "address",
    Port:              22,
    PublicKeyLocation: "/home/user/.ssh/id_rsa.pub",
  }

  client := backup.Client(config)
  err := client.Init()
  if err != nil {
    log.Fatalf("Unable to initialize the client: %s", err.Error())
  }

  defer client.Close()

  err = client.Backup(
    "/data/user/uploads",
    "aws.bucket",
    backup.DefaultPathTransformer)

  // If we don't want the default path transformer
  // we can override it with our own implementation.

  if err != nil {
    log.Fatalf("Unable to backup, err: %s", err.Error())
  }
}
```


### AWS Variables
In order to backup the files onto S3, we need to access the AWS variables which would allow us to login to the platform and upload the information into corresponding bucket. More information could be read at [AWS Env](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html)
