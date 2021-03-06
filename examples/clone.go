// Copyright © 2013, 2014, The Go-MERCURY Authors. All rights reserved.
// Use of this source code is governed by a LGPLv2.1
// license that can be found in the LICENSE file.

// +build linux,cgo

package main

import (
	"flag"
	"log"

	"gopkg.in/mercury/go-mercury.v2"
)

var (
	mercurypath string
	name    string
	backend mercury.BackendStore
)

func init() {
	flag.StringVar(&mercurypath, "mercurypath", mercury.DefaultConfigPath(), "Use specified container path")
	flag.StringVar(&name, "name", "rubik", "Name of the original container")
	flag.Var(&backend, "backend", "Backend type to use, possible values are [dir, zfs, btrfs, lvm, aufs, overlayfs, loopback, best]")
	flag.Parse()
}

func main() {
	c, err := mercury.NewContainer(name, mercurypath)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}

	if backend == 0 {
		log.Fatalf("ERROR: %s\n", mercury.ErrUnknownBackendStore)
	}

	log.Printf("Cloning the container using %s backend...\n", backend)
	err = c.Clone(name+"_"+backend.String(), mercury.CloneOptions{
		Backend: backend,
	})
	if err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}
}
