// Client is a simple client for testing out the Alpaca and Polygon APIs.
//
// This client relies on a TOML config file located at ~/.alpaca.toml
// that looks like this:
//
// [alpaca]
// apiKeyID="abc...xyz"
// apiSecretKey="abc...xyz"
package client

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gmlewis/alpaca-trade-api-go/alpaca"
	"github.com/gmlewis/alpaca-trade-api-go/common"
	"github.com/gmlewis/alpaca-trade-api-go/polygon"
	"github.com/mitchellh/go-homedir"
	"github.com/pelletier/go-toml"
)

const (
	baseURL          = "https://paper-api.alpaca.markets"
	settingsFilename = ".alpaca.toml"
)

// Client contains an Alpaca v2 and Polygon v2 API client.
// Client implements the API interface.
type Client struct {
	AClient *alpaca.Client
	PClient *polygon.Client
	AStream *alpaca.Stream
	PStream *polygon.Stream
}

// New returns a new Client.
//
// Config values are read from ${HOME}/.alpaca.toml
// See above.
func New() (*Client, error) {
	homedir, err := homedir.Dir()
	if err != nil {
		return nil, fmt.Errorf("homedir: %v", err)
	}
	filename := filepath.Join(homedir, settingsFilename)
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: %v", err)
	}
	settings, err := toml.Load(string(buf))
	if err != nil {
		return nil, fmt.Errorf("toml.Load: %v", err)
	}

	apiKeyID := settings.Get("alpaca.apiKeyID").(string)
	apiSecretKey := settings.Get("alpaca.apiSecretKey").(string)

	if common.Credentials().ID == "" {
		os.Setenv(common.EnvApiKeyID, apiKeyID)
	}
	if common.Credentials().Secret == "" {
		os.Setenv(common.EnvApiSecretKey, apiSecretKey)
	}
	alpaca.SetBaseUrl(baseURL)

	aClient := alpaca.NewClient(common.Credentials())
	pClient := polygon.NewClient(common.Credentials())

	aStream := alpaca.GetStream()
	pStream := polygon.GetStream()

	return &Client{
		AClient: aClient,
		PClient: pClient,
		AStream: aStream,
		PStream: pStream,
	}, nil
}

// Close closes the stream connection.
func (c *Client) Close() error {
	if c.AStream != nil {
		if err := c.AStream.Close(); err != nil {
			return err
		}
		c.AStream = nil
	}

	if c.PStream != nil {
		if err := c.PStream.Close(); err != nil {
			return err
		}
		c.PStream = nil
	}

	return nil
}
