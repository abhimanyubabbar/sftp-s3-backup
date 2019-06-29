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
  client.Init()

  err := client.Backup("/data/user/uploads", "aws.bucket")
  if err != nil {
    log.Fatalf("Unable to backup, err: %s", err.Error())
  }
}
```
