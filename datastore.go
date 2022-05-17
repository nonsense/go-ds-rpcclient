package rpcclient

import (
	"context"

	"github.com/ethereum/go-ethereum/rpc"

	ds "github.com/ipfs/go-datastore"
	dsq "github.com/ipfs/go-datastore/query"
	logger "github.com/ipfs/go-log/v2"
)

var log = logger.Logger("rpcclient")

var _ ds.Datastore = (*Datastore)(nil)
var _ ds.Batching = (*Datastore)(nil)

type Datastore struct {
	client *rpc.Client
}

func NewDatastore() (*Datastore, error) {
	client, err := rpc.Dial("http://localhost:8089")
	if err != nil {
		return nil, err
	}

	return &Datastore{
		client: client,
	}, nil
}

// Datastore.Read iface

func (d *Datastore) Get(ctx context.Context, key ds.Key) (value []byte, err error) {
	var resp []byte
	err = d.client.Call(&resp, "rpcdatastore_get", key)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *Datastore) Has(ctx context.Context, key ds.Key) (bool, error) {
	var resp bool
	err := d.client.Call(&resp, "rpcdatastore_has", key)
	if err != nil {
		return false, err
	}

	return resp, nil
}

func (d *Datastore) GetSize(ctx context.Context, key ds.Key) (size int, err error) {
	var resp int
	err = d.client.Call(&resp, "rpcdatastore_getSize", key)
	if err != nil {
		return -1, err
	}

	return resp, nil
}

func (d *Datastore) Query(ctx context.Context, q dsq.Query) (dsq.Results, error) {
	var entries []dsq.Entry
	err := d.client.Call(&entries, "rpcdatastore_query", q)
	if err != nil {
		return nil, err
	}

	results := dsq.ResultsWithEntries(q, entries)

	return results, nil
}

// Datastore.Write iface

func (d *Datastore) Put(ctx context.Context, key ds.Key, value []byte) error {
	return d.client.Call(nil, "rpcdatastore_put", key, value)
}

func (d *Datastore) Delete(ctx context.Context, key ds.Key) error {
	return d.client.Call(nil, "rpcdatastore_delete", key)
}

// Datastore iface

func (d *Datastore) Sync(ctx context.Context, prefix ds.Key) error {
	return d.client.Call(nil, "rpcdatastore_sync", prefix)
}

func (d *Datastore) Close() error {
	return d.client.Call(nil, "rpcdatastore_close")
}

//// Batching iface

func (d *Datastore) Batch(ctx context.Context) (ds.Batch, error) {
	//TODO: not sure if serialisation works for Batch
	var resp ds.Batch
	err := d.client.Call(&resp, "rpcdatastore_batch")
	if err != nil {
		return resp, err
	}

	return resp, nil
}

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
