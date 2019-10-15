package threadservice

import (
	"context"

	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/textileio/go-textile-core/thread"
	tstore "github.com/textileio/go-textile-core/threadstore"
)

// Threadservice is an API for working with threads.
type Threadservice interface {
	// Threadstore persists thread log details.
	tstore.Threadstore

	// Host provides a network identity.
	Host() host.Host

	// DAGService provides a DAG API for reading and writing thread logs.
	DAGService() format.DAGService

	// Add data to a thread. Creates a new thread and own log if they don't exist.
	Add(ctx context.Context, body format.Node, opts ...AddOption) (peer.ID, thread.Record, error)

	// Put an existing node to a log.
	Put(ctx context.Context, node thread.Record, opts ...PutOption) error

	// Pull paginates thread log events.
	Pull(ctx context.Context, id thread.ID, lid peer.ID, offset cid.Cid, opts ...PullOption) ([]thread.Record, error)

	// Logs returns info for each log in the given thread.
	Logs(id thread.ID) ([]thread.LogInfo, error)

	// Delete a thread.
	Delete(ctx context.Context, id thread.ID) error
}
