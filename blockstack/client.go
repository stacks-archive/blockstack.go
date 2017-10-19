package blockstack

import (
	"encoding/json"
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

// Response is an interface to allow for common methods between responses
type Response interface {
	JSON() string
	PrettyJSON() string
}

// Error models an error coming out of the blockstack lib.
type Error interface {
	Error() string
	JSON() string
	PrettyJSON() string
}

// RPCError wraps errors returned by the blockstack-core node
type RPCError struct {
	Err       string   `json:"error"`
	RPC       string   `json:"rpc_method"`
	Traceback []string `json:"traceback"`
}

// Error satisfies the error interface
func (err RPCError) Error() string {
	return err.Err
}

// JSON allows for easy Marshal
func (err RPCError) JSON() string {
	byt, e := json.Marshal(err)
	if e != nil {
		log.Fatal(e)
	}
	return string(byt)
}

// PrettyJSON allows for easy Marshal
func (err RPCError) PrettyJSON() string {
	byt, e := json.MarshalIndent(err, "", "    ")
	if e != nil {
		log.Fatal(e)
	}
	return string(byt)
}

// CallError represents an error resulting from a failed RPC call
type CallError struct {
	RPC string `json:"rpc_method"`
	Err error  `json:"error"`
}

// Error satisfies the error interface
func (err CallError) Error() string {
	return err.Err.Error()
}

// JSON allows for easy Marshal
func (err CallError) JSON() string {
	byt, e := json.Marshal(err)
	if e != nil {
		log.Fatal(e)
	}
	return string(byt)
}

// PrettyJSON allows for easy Marshal
func (err CallError) PrettyJSON() string {
	byt, e := json.MarshalIndent(err, "", "    ")
	if e != nil {
		log.Fatal(e)
	}
	return string(byt)
}

// JSONUnmarshalError represents an error resulting from a failed RPC call
type JSONUnmarshalError struct {
	RPC string `json:"rpc_method"`
	Err error  `json:"error"`
}

// Error satisfies the error interface
func (err JSONUnmarshalError) Error() string {
	return err.Err.Error()
}

// JSON allows for easy Marshal
func (err JSONUnmarshalError) JSON() string {
	byt, e := json.Marshal(err)
	if e != nil {
		log.Fatal(e)
	}
	return string(byt)
}

// PrettyJSON allows for easy Marshal
func (err JSONUnmarshalError) PrettyJSON() string {
	byt, e := json.MarshalIndent(err, "", "    ")
	if e != nil {
		log.Fatal(e)
	}
	return string(byt)
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
