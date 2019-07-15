package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/sftp"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh"
)

func main() {

	app := cli.NewApp()
	app.Name = "sftp-move"
	app.Usage = "moving the files over sftp from one directory to another"
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
			Name:  "olddir, od",
			Usage: "directory to move the contents from",
		},
		cli.StringFlag{
			Name:  "newdir, nd",
			Usage: "directory to move the contents to",
		},
	}

	// Apply the move action when we need to perform the move
	// operation over to the contents of the directory.
	app.Action = func(c *cli.Context) error {
		return move(c)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func move(c *cli.Context) error {

	client, err := getSFTPClient(
		c.String("user"),
		c.String("address"),
		c.Int("port"),
		c.String("keylocation"),
	)

	if err != nil {
		fmt.Printf("Unable to get sftp client, err: %s\n", err)
		return err
	}
	defer client.Close()
	fmt.Println("Successfully connected to the sftp")

	oldDirectory := c.String("olddir")
	newDirectory := c.String("newdir")

	fmt.Printf("OldDir: %s, NewDir: %s \n", oldDirectory, newDirectory)

	files, err := client.ReadDir(oldDirectory)
	for i := 0; i < len(files); i++ {

		oldPath := fmt.Sprintf("%s/%s", oldDirectory, files[i].Name())
		newPath := fmt.Sprintf("%s/%s", newDirectory, files[i].Name())

		err = client.PosixRename(oldPath, newPath)
		fmt.Printf("Successfully moved, old: %s, new: %s\n", oldPath, newPath)
	}

	if err != nil {
		fmt.Printf("Unable to complete the move of files, error: %s\n", err.Error())
		return err
	}

	fmt.Println("Finished transferring the data between the directories")
	return nil
}

// getSFTPClient fetches the sftp client.
func getSFTPClient(user, address string, port int, location string) (*sftp.Client, error) {
	key, err := ioutil.ReadFile(location)
	if err != nil {
		return nil, err
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		// For now allow the InsecureIgnoreHostKey
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := fmt.Sprintf("%s:%d", address, port)
	conn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, err
	}

	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return sftpClient, nil
}
