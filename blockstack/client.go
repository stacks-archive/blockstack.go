package blockstack

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/kolo/xmlrpc"
)

// Client is the exportable object that the RPC methods are defined on
type Client struct {
	node    *xmlrpc.Client
	config  ServerConfig
	getInfo GetInfoResult
}

// Clients is a collection of clients
type Clients []*Client

// MUST POPULATE GETINFO FOR CLIENTS BEFORE RUNNING
func (c Clients) consensusPeers() ([]*Client, []Error) {
	var out []*Client
	var outErr []Error
	var hash string
	numSame := 0
	hashes := make(map[string]int)
	for _, client := range c {
		hashes[client.getInfo.Consensus]++
	}

	// If there is 1 consensus hash between the clients, return them all
	if len(hashes) == 1 {
		return c, []Error{nil}
	}

	// Find the most common ConsensusHash
	for key := range hashes {
		if hashes[key] > numSame {
			numSame = hashes[key]
			hash = key
		}
	}

	// Sort the clients and errors
	for _, client := range c {
		if client.getInfo.Consensus != hash {
			outErr = append(outErr, ClientRegistrationError{URL: client.config.String(), Err: "Client failed consensus check"})
		} else {
			out = append(out, client)
		}
	}

	return out, outErr
}

// register takes and array of Client and makes
// sure they are all valid servers and in consensus
func (c Clients) register() ([]*Client, []Error) {
	var out []*Client
	var getInfo []*Client
	var getInfoErrs []Error

	// Run GetInfo call for each client and save any errors
	for _, client := range c {
		res, err := client.GetInfo()
		// If there is no result to check against, and no error on call
		// And not indexing, the result is the right one.
		if err == nil && !res.Indexing {
			client.getInfo = res
			getInfo = append(getInfo, client)
			// If client is indexing return the error
		} else if err == nil && res.Indexing {
			getInfoErrs = append(getInfoErrs, ClientRegistrationError{URL: client.config.String(), Err: "Client still indexing"})
			// If the error is not nil, return the error
		} else if err != nil {
			getInfoErrs = append(getInfoErrs, ClientRegistrationError{URL: client.config.String(), Err: err.Error()})
		}
	}

	// Pick out the consensusPeers
	out, outErr := Clients(getInfo).consensusPeers()

	// Append the errors
	for _, err := range getInfoErrs {
		outErr = append(outErr, err)
	}

	// Return the good ones and errors
	return out, outErr
}

// ValidClients takes an array of strings and returns a *Client for
// Each valid URL that is in consensus with the others
func ValidClients(urls []string) ([]*Client, []Error) {
	var clients []*Client
	var urlErrs []Error

	// Parse each url and return error if not there
	for _, uri := range urls {
		purl, err := url.Parse(uri)
		if err != nil {
			urlErrs = append(urlErrs, ClientRegistrationError{URL: uri, Err: "Failed to parse URL"})
		} else {
			clients = append(clients, NewClient(ServerConfig{Address: purl.Hostname(), Port: purl.Port(), Scheme: purl.Scheme}))
		}
	}

	// register the clients that pass
	out, outErr := Clients(clients).register()

	// Return all errors
	for _, err := range urlErrs {
		outErr = append(outErr, err)
	}

	return out, outErr
}

// NewClient creates a new instance of the blockstack-core rpc client
func NewClient(conf ServerConfig) *Client {
	client, err := xmlrpc.NewClient(conf.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		node:   client,
		config: conf,
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

// ClientRegistrationError represents an error resulting from a failed RPC call
type ClientRegistrationError struct {
	URL string `json:"url"`
	Err string `json:"error"`
}

// Error satisfies the error interface
func (err ClientRegistrationError) Error() string {
	return err.Err
}

// JSON allows for easy Marshal
func (err ClientRegistrationError) JSON() string {
	byt, e := json.Marshal(err)
	if e != nil {
		log.Fatal(e)
	}
	return string(byt)
}

// PrettyJSON allows for easy Marshal
func (err ClientRegistrationError) PrettyJSON() string {
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
