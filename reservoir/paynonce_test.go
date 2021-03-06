// Copyright (c) 2014-2017 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package reservoir_test

import (
	"encoding/json"
	"github.com/bitmark-inc/bitmarkd/block"
	"github.com/bitmark-inc/bitmarkd/chain"
	"github.com/bitmark-inc/bitmarkd/mode"
	"github.com/bitmark-inc/bitmarkd/reservoir"
	"testing"
)

func TestPayNonce(t *testing.T) {

	nonce := reservoir.PayNonce{
		0x2b, 0xa1, 0x54, 0x14, 0x46, 0x74, 0x29, 0x1d,
	}
	expected := `"2ba154144674291d"`

	t.Logf("pay nonce: %#v", nonce)

	buffer, err := json.Marshal(nonce)
	if nil != err {
		t.Fatalf("marshal JSON error: %v", err)
	}

	t.Logf("pay nonce: %s", buffer)

	actual := string(buffer)
	if expected != actual {
		t.Fatalf("pay nonce expected: %#v  actual: %#v", expected, actual)
	}

	var nonce2 reservoir.PayNonce
	err = json.Unmarshal(buffer, &nonce2)
	if nil != err {
		t.Fatalf("unmarshal JSON error: %v", err)
	}

	if nonce != nonce2 {
		t.Fatalf("pay once expected: %#v  actual: %#v", nonce, nonce2)
	}
}

func TestNewPayNonceBitmark(t *testing.T) {

	// dependant on the genesis digest for bitmark
	expected := `"445f81247a6fdecc"`

	setup(t)
	defer teardown(t)

	d, n := block.Get()
	t.Logf("block: %d  %#v", n, d)

	nonce := reservoir.NewPayNonce()
	t.Logf("pay nonce: %#v", nonce)

	buffer, err := json.Marshal(nonce)
	if nil != err {
		t.Fatalf("marshal JSON error: %v", err)
	}

	t.Logf("pay nonce: %s", buffer)

	actual := string(buffer)
	if expected != actual {
		t.Fatalf("pay nonce expected: %#v  actual: %#v", expected, actual)
	}
}

func TestNewPayNonceTesting(t *testing.T) {

	// dependant on the genesis digest for testing
	expected := `"d1cc53a056227402"`

	mode.Initialise(chain.Testing)
	setup(t)
	defer teardown(t)

	d, n := block.Get()
	t.Logf("block: %d  %#v", n, d)

	nonce := reservoir.NewPayNonce()
	t.Logf("pay nonce: %#v", nonce)

	buffer, err := json.Marshal(nonce)
	if nil != err {
		t.Fatalf("marshal JSON error: %v", err)
	}

	actual := string(buffer)
	if expected != actual {
		t.Fatalf("pay nonce expected: %#v  actual: %#v", expected, actual)
	}

}
