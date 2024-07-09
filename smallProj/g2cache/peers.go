package g2cache

import pb "g2cache/g2cachepb/g2cachepb"

// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {
	// Get returns the value for the specified key and the group.
	Get(in *pb.Request, out *pb.Response) error
}

type PeerPicker interface {
	// PickPeer picks a peer according to the input key.
	PickPeer(key string) (peer PeerGetter, ok bool)
}
