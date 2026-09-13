// Copyright 2026 The go-ethereum Authors
// This file is part of the go-ethereum library and is licensed under the
// GNU Lesser General Public License, version 3 or later.

package nat

import (
	"fmt"
	"net"
	"testing"
	"time"

	stunV3 "github.com/pion/stun/v3"
)

func TestSTUNv3Loopback(t *testing.T) {
	for _, includeAddress := range []bool{true, false} {
		t.Run(fmt.Sprintf("mapped-address-%t", includeAddress), func(t *testing.T) {
			server, err := net.ListenPacket("udp4", "127.0.0.1:0")
			if err != nil {
				t.Fatal(err)
			}
			defer server.Close()
			if err := server.SetDeadline(time.Now().Add(3 * time.Second)); err != nil {
				t.Fatal(err)
			}
			expected := net.IPv4(203, 0, 113, 7)
			done := make(chan error, 1)
			go func() {
				buf := make([]byte, 1500)
				n, peer, err := server.ReadFrom(buf)
				if err != nil {
					done <- err
					return
				}
				request := new(stunV3.Message)
				request.Raw = buf[:n]
				if err = request.Decode(); err != nil {
					done <- err
					return
				}
				if request.Type != stunV3.BindingRequest {
					done <- fmt.Errorf("unexpected request %v", request.Type)
					return
				}
				setters := []stunV3.Setter{stunV3.NewTransactionIDSetter(request.TransactionID), stunV3.BindingSuccess}
				if includeAddress {
					setters = append(setters, &stunV3.XORMappedAddress{IP: expected, Port: 3456})
				}
				response, err := stunV3.Build(setters...)
				if err == nil {
					_, err = server.WriteTo(response.Raw, peer)
				}
				done <- err
			}()
			got, err := new(stun).externalIP(server.LocalAddr().String())
			if includeAddress {
				if err != nil {
					t.Fatal(err)
				}
				if !got.Equal(expected) {
					t.Fatalf("mapped address %v, expected synthetic %v", got, expected)
				}
			} else if err == nil {
				t.Fatal("missing XOR-MAPPED-ADDRESS was accepted")
			}
			select {
			case err := <-done:
				if err != nil {
					t.Fatal(err)
				}
			case <-time.After(4 * time.Second):
				t.Fatal("synthetic responder did not finish")
			}
		})
	}
}
