// Copyright (c) 2014-2017 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package zmqutil

import (
	zmq "github.com/pebbe/zmq4"
	"sync"
)

// to ensure only one auth start
var oneTimeAuthStart sync.Once

// initilaise the ZMQ security subsystem
func StartAuthentication() error {

	err := error(nil)
	oneTimeAuthStart.Do(func() {

		// initialise encryption
		zmq.AuthSetVerbose(false)
		//zmq.AuthSetVerbose(true)
		err = zmq.AuthStart()
	})

	return err
}
