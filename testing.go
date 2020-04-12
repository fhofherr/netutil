package netutil

import (
	"testing"
	"time"
)

// GetAddr reads a servers address from addrC.
func GetAddr(t *testing.T, addrC <-chan string, timeout time.Duration) string {
	select {
	case addr := <-addrC:
		return addr
	case <-time.After(timeout):
		t.Fatalf("Time-out after %v", timeout)
	}
	return ""
}

// GetErr reads an error from errC.
func GetErr(t *testing.T, errC <-chan error, timeout time.Duration) error {
	select {
	case err := <-errC:
		return err
	case <-time.After(timeout):
		t.Fatalf("Time-out after %v", timeout)
	}
	return nil
}
