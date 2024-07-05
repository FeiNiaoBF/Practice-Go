package g2cache

// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {
	// Get returns the value for the specified key and the group.
	Get(group string, key string) ([]byte, error)
}

type PeerPicker interface {
	// PickPeer picks a peer according to the input key.
	PickPeer(key string) (peer PeerGetter, ok bool)
}

