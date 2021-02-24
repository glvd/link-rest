package iface

import (
	peer "github.com/glvd/link-rest/library/libp2p/go-libp2p-core/peer"
	mbase "github.com/glvd/link-rest/library/multiformats/go-multibase"
)

func FormatKeyID(id peer.ID) string {
	if s, err := peer.ToCid(id).StringOfBase(mbase.Base36); err != nil {
		panic(err)
	} else {
		return s
	}
}

// FormatKey formats the given IPNS key in a canonical way.
func FormatKey(key Key) string {
	return FormatKeyID(key.ID())
}
