package backup

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Config struct {
	User              string
	Address           string
	Port              int
	PublicKeyLocation string
}

func NewClient(config Config) *Client {
	return &Client{
		config: config,
		sftp:   nil,
	}
}

type Client struct {
	config Config
	sftp   *sftp.Client
}

func (c *Client) Init() error {

	config := c.config

	key, err := ioutil.ReadFile(config.PublicKeyLocation)
	if err != nil {
		return err
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return err
	}

	sshConfig := &ssh.ClientConfig{
		User: config.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	address := fmt.Sprintf("%s:%d", config.Address, config.Port)
	conn, err := ssh.Dial("tcp", address, sshConfig)
	if err != nil {
		return err
	}

	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		return err
	}

	c.sftp = sftpClient
	return nil
}

func (c *Client) Backup(directory, bucket string) error {

	files, err := c.sftp.ReadDir(directory)
	if err != nil {
		return err
	}

	for i := 0; i < len(files); i++ {

		if files[i].IsDir() {
			fmt.Printf("Found directory: %s, skipping\n", files[i].Name())
			continue
		}

		path := fmt.Sprintf("%s/%s", directory, files[i].Name())
		f, err := c.sftp.Open(path)

		if err != nil {
			fmt.Println("Unable to open file")
			return err
		}

		b := bytes.NewBuffer([]byte{})
		f.WriteTo(b)

		err = Upload(bucket, files[i].Name(), b)
		if err != nil {
			fmt.Println("Unable to upload the information to AWS, exiting.")
			return err
		}
	}

	return nil
}

func (c *Client) Close() error {
	return c.sftp.Close()
}

type S3PathTransform = func(sftpPath string) string
