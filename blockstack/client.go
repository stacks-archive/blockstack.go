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
	client, err := xmlrpc.NewClient(conf.String(), nil)
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

func (s ServerConfig) String() string {
	addr := fmt.Sprintf("%s://%s/RPC2", s.Scheme, s.Address)
	if s.Port != "" {
		addr = fmt.Sprintf("%s://%s:%v/RPC2", s.Scheme, s.Address, s.Port)
	}
	return addr
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
