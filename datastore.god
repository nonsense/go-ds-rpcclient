package rpcclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/rpc/jsonrpc"
	"strings"

	"github.com/davecgh/go-spew/spew"
	ds "github.com/ipfs/go-datastore"
	logger "github.com/ipfs/go-log/v2"
)

var log = logger.Logger("rpcclient")

var _ ds.Datastore = (*Datastore)(nil)
var _ ds.Batching = (*Datastore)(nil)

type Datastore struct {
	client   *http.Client
	endpoint string
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

func NewDatastore() *Datastore {
	client := jsonrpc.NewClient("http://localhost:8089/")

	return &Datastore{
		client:   client,
		endpoint: "http://localhost:8089",
	}
}

func (d *Datastore) Close() error {
	if t, ok := d.client.Transport.(*http.Transport); ok {
		t.CloseIdleConnections()
	}
	return nil
}

func (d *Datastore) Get(ctx context.Context, key ds.Key) (value []byte, err error) {
	r := &GetRequest{
		Key: key,
	}

	resp, err := d.client.Call("Get", r)
	if err != nil {
		return nil, err
	}

	spew.Dump(res)

	return resp, nil
}

func (d *Datastore) request(ctx context.Context, method string, path string, body io.Reader, headers ...string) (io.ReadCloser, error) {
	if len(headers)%2 != 0 {
		return nil, fmt.Errorf("headers must be tuples: key1, value1, key2, value2")
	}
	req, err := http.NewRequest(method, c.endpoint+path, body)
	req = req.WithContext(ctx)

	token := strings.TrimSpace(c.cfg.Client.Token)
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	for i := 0; i < len(headers); i = i + 2 {
		req.Header.Add(headers[i], headers[i+1])
	}

	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("unexpected status code received: %s", resp.Status)
	}

	if ct := resp.Header.Get("Content-Type"); ct != "application/json" {
		return nil, fmt.Errorf("unexpected content-type received: %s", ct)
	}

	return resp.Body, nil
}

// Datastore.Read iface

//func (d *Datastore) Get(ctx context.Context, key ds.Key) (value []byte, err error) {
//}

//func (d *Datastore) Has(ctx context.Context, key ds.Key) (bool, error) {
//}

//func (d *Datastore) GetSize(ctx context.Context, key ds.Key) (size int, err error) {
//}

//func (d *Datastore) Query(ctx context.Context, q dsq.Query) (dsq.Results, error) {
//}

//// Datastore.Write iface

//func (d *Datastore) Put(ctx context.Context, key ds.Key, value []byte) error {
//}

//func (d *Datastore) Delete(ctx context.Context, key ds.Key) error {
//}

//// Datastore iface

//func (d *Datastore) Sync(ctx context.Context, prefix ds.Key) error {
//}

//func (d *Datastore) Close() error {
//}

//// Batching iface

//func (d *Datastore) Batch(ctx context.Context) (ds.Batch, error) {
//}

////func (d *Datastore) PutWithTTL(ctx context.Context, key ds.Key, value []byte, ttl time.Duration) error {
////}

////func (d *Datastore) SetTTL(ctx context.Context, key ds.Key, ttl time.Duration) error {
////}

////func (d *Datastore) GetExpiration(ctx context.Context, key ds.Key) (time.Time, error) {
////}

////func (d *Datastore) DiskUsage(ctx context.Context) (uint64, error) {
////}

////func (d *Datastore) CollectGarbage(ctx context.Context) (err error) {
////}
