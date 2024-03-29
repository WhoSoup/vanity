package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

// comp compares the input against two separate bytes
func comp(s, l1, l2 byte) bool {
	return s == l1 || s == l2
}

var now = time.Now()
var i uint64
var targets []string

func toUpper(b byte) byte {
	if 'a' <= b && b <= 'z' {
		b -= 'a' - 'A'
	}
	return b
}

// goroutine loop
func find(sleep uint) {
	for {
		publicRaw, private, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			fmt.Println(err)
			continue
		}

		public := base58.Encode(publicRaw)

		for _, target := range targets {
			matched := true
			for i := range target {
				if !comp(public[i], target[i], toUpper(target[i])) {
					matched = false
					break
				}
			}
			if matched {
				fmt.Printf("%s %s\n", public, hex.EncodeToString(private))
			}
		}

		atomic.AddUint64(&i, 1)
		if sleep > 0 {
			time.Sleep(time.Microsecond * time.Duration(sleep))
		}
	}
}

// Start the goroutines using specified settings
func Start(sleep, conc uint) {
	fmt.Fprintln(os.Stderr, "Starting...")

	if conc > 1 {
		for i := uint(0); i < conc-1; i++ {
			go find(sleep)
		}
	}
	find(sleep)
}
