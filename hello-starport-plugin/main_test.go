package main

import (
	"testing"

	"github.com/motki/cli/text/banner"
)

func Test_Banner(t *testing.T) {
	banner.Printf("HELLO")
}
