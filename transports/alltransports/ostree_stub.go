// +build containers_image_ostree_stub !linux !cgo

package alltransports

import "github.com/containers/image/transports"

func init() {
	transports.Register(transports.NewStubTransport("ostree"))
}
