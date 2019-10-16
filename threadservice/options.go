package threadservice

import (
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/textileio/go-textile-core/crypto"
	"github.com/textileio/go-textile-core/thread"
)

// AddOpt is an instance helper for creating add options.
var AddOpt AddOption

// AddOption is used to create AddSettings.
type AddOption func(*AddSettings)

// Thread sets the target thread for an add operation.
func (AddOption) Thread(val thread.ID) AddOption {
	return func(settings *AddSettings) {
		settings.Thread = val
	}
}

// Time sets wall-clock time associated for an add operation.
func (AddOption) Time(val time.Time) AddOption {
	return func(settings *AddSettings) {
		settings.Time = val
	}
}

// Key sets the read encryption key used for an add operation.
// If both KeyLog and this option are absent, the target log's read key is used.
// This option takes precedence over KeyLogID.
func (AddOption) Key(val crypto.EncryptionKey) AddOption {
	return func(settings *AddSettings) {
		settings.Key = val
	}
}

// KeyLog sets the read encryption key used for an add operation
// to the read key of the given log ID.
// If both Key and this option are absent, the target log's read key is used.
func (AddOption) KeyLog(val peer.ID) AddOption {
	return func(settings *AddSettings) {
		settings.KeyLog = val
	}
}

// Addrs holds additional addresses to notify with an add operation.
func (AddOption) Addrs(val []ma.Multiaddr) AddOption {
	return func(settings *AddSettings) {
		settings.Addrs = val
	}
}

// AddSettings holds values used for an add operation.
type AddSettings struct {
	Thread thread.ID
	Time   time.Time
	Key    crypto.EncryptionKey
	KeyLog peer.ID
	Addrs  []ma.Multiaddr
}

// AddOptions returns add settings from options.
func AddOptions(opts ...AddOption) *AddSettings {
	options := &AddSettings{
		Thread: thread.NewIDV1(thread.AccessControlled, 16),
		Time:   time.Now(),
	}

	for _, opt := range opts {
		opt(options)
	}
	return options
}

// PutOpt is an instance helper for creating put options.
var PutOpt PutOption

// PutOption is used to create PutSettings.
type PutOption func(*PutSettings)

// Thread sets the target thread for a put operation.
func (PutOption) Thread(val thread.ID) PutOption {
	return func(settings *PutSettings) {
		settings.Thread = val
	}
}

// Log sets the target log for a put operation.
func (PutOption) Log(val peer.ID) PutOption {
	return func(settings *PutSettings) {
		settings.Log = val
	}
}

// PutSettings holds values used for a put operation.
type PutSettings struct {
	Thread thread.ID
	Log    peer.ID
}

// PutOptions returns put settings from options.
func PutOptions(opts ...PutOption) *PutSettings {
	options := &PutSettings{
		Thread: thread.NewIDV1(thread.AccessControlled, 16),
	}

	for _, opt := range opts {
		opt(options)
	}
	return options
}

// ListenOpt is an instance helper for creating listen options.
var ListenOpt ListenOption

// ListenOption is used to create ListenSettings.
type ListenOption func(*ListenSettings)

// Thread restricts the listener to the given thread.
// Use this option multiple times to build up a list of threads
// to listen to.
func (ListenOption) Thread(val thread.ID) ListenOption {
	return func(settings *ListenSettings) {
		settings.Threads = append(settings.Threads, val)
	}
}

// ListenSettings holds values used for a listen operation.
type ListenSettings struct {
	Threads []thread.ID
}

// ListenOptions returns listen settings from options.
func ListenOptions(opts ...ListenOption) *ListenSettings {
	options := &ListenSettings{}

	for _, opt := range opts {
		opt(options)
	}
	return options
}
