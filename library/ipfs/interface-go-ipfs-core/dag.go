package iface

import (
	ipld "github.com/glvd/link-rest/library/ipfs/go-ipld-format"
)

// APIDagService extends ipld.DAGService
type APIDagService interface {
	ipld.DAGService

	// Pinning returns special NodeAdder which recursively pins added nodes
	Pinning() ipld.NodeAdder
}
