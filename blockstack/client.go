package blockstack

import (
	"fmt"
	"log"

	"github.com/kolo/xmlrpc"
)

// Client is the exportable object that the RPC methods are defined on
type Client struct {
	node *xmlrpc.Client
}

// NewClient creates a new instance of the blockstack-core rpc client
func NewClient(conf ServerConfig) *Client {
	addr := fmt.Sprintf("%s://%s/RPC2", conf.Scheme, conf.Address)
	if conf.Port != "" {
		addr = fmt.Sprintf("%s://%s:%v/RPC2", conf.Scheme, conf.Address, conf.Port)
	}
	client, err := xmlrpc.NewClient(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		node: client,
	}
}

// ServerConfig is connection details for an indivdual blockstack-core node
type ServerConfig struct {
	Address string
	Port    string
	Scheme  string
}

// ServerConfigs is a type to hold multiple ServerConfig
type ServerConfigs []ServerConfig

// RPCError wraps errors from the RPC calls
type RPCError struct {
	Err       string   `json:"error"`
	Traceback []string `json:"traceback"`
}

func (err RPCError) Error() string {
	return err.Err
}

// TestMethod calls an RPC method with the given args
func (bsk *Client) TestMethod(methodName string, args []interface{}) string {
	var result string
	err := bsk.node.Call(methodName, args, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
