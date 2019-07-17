package main

import (
	"fmt"
	"log"
	"os"

	backup "github.com/babbarshaer/sftp-s3-backup"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "sftp-backup"
	app.Usage = "backing up the files from sftp to s3"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "user, u",
			Usage: "sftp server user",
		},
		cli.StringFlag{
			Name:  "address, a",
			Usage: "Server Address",
		},
		cli.IntFlag{
			Name:  "port, p",
			Value: 22,
			Usage: "server port",
		},
		cli.StringFlag{
			Name:  "keylocation, kl",
			Value: "~/.ssh/id_rsa.pub",
			Usage: "Location for public key",
		},
		cli.StringFlag{
			Name:  "dir, d",
			Usage: "directory to backup",
		},
		cli.StringFlag{
			Name:  "bucket, b",
			Usage: "bucket to backup the data into",
		},
	}

	// Apply the backup action when we need to perform the move
	// operation over to the contents of the directory.
	app.Action = func(c *cli.Context) error {
		return backItUp(c)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func backItUp(c *cli.Context) error {

	config := backup.Config{
		User:              c.String("user"),
		Address:           c.String("address"),
		Port:              c.Int("port"),
		PublicKeyLocation: c.String("keylocation"),
	}

	client := backup.NewClient(config)
	err := client.Init()
	if err != nil {
		return err
	}

	defer client.Close()
	fmt.Println("Successfully connected the client onto the sftp")

	fmt.Printf("Backing up: %s to %s\n", c.String("dir"), c.String("bucket"))
	return client.Backup(
		c.String("dir"),
		c.String("bucket"),
		backup.DefaultS3PathTransformer)
}
