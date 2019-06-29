# `/cmd`
It contains main applications of this projects. The applications are divided into this:


### `backup`
It basically backs up the files directory on the sftp over to S3. In order to run this utility, we need to first set the AWS credentials in the environment variables. Refer this guide for more information [AWS Environ](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html).

```
go run backup/backup.go -u=ubuntu -a=172.13.34.27 -kl=/home/babbar/.ssh/epidemickey.pem -d="/d
ata/backup/distribution/test/uploads" -b="internal.epidemicsound.distribution.sftp.backup"
```


### `move`
It moves the files in one directory over to another. At the time of writing this, it only moves files in the base directory over to the next one.

```
go run move/move.go -user=ubuntu -a=172.34.23.11 -kl=/home/babbar/.ssh/epidemickey.pem -od="
/data/distribution/test/uploads" -nd="/data/backup/distribution/test/uploads"
```

For both commands `help` can be accessed with `-h` option.
