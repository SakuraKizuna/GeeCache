package geecache

import pb "seven_days/geeCache/geecache/geecachepb"

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key.
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {

	// Get
	//Get(group string, key string) ([]byte, error)

	Get(int *pb.Request, out *pb.Response) error
}
