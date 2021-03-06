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
	mercurypath    string
	template   string
	distro     string
	release    string
	arch       string
	name       string
	verbose    bool
	flush      bool
	validation bool
)

func init() {
	flag.StringVar(&mercurypath, "mercurypath", mercury.DefaultConfigPath(), "Use specified container path")
	flag.StringVar(&template, "template", "download", "Template to use")
	flag.StringVar(&distro, "distro", "ubuntu", "Template to use")
	flag.StringVar(&release, "release", "trusty", "Template to use")
	flag.StringVar(&arch, "arch", "amd64", "Template to use")
	flag.StringVar(&name, "name", "rubik", "Name of the container")
	flag.BoolVar(&verbose, "verbose", false, "Verbose output")
	flag.BoolVar(&flush, "flush", false, "Flush the cache")
	flag.BoolVar(&validation, "validation", false, "GPG validation")
	flag.Parse()
}

func main() {
	c, err := mercury.NewContainer(name, mercurypath)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}

	log.Printf("Creating container...\n")
	if verbose {
		c.SetVerbosity(mercury.Verbose)
	}

	options := mercury.TemplateOptions{
		Template:             template,
		Distro:               distro,
		Release:              release,
		Arch:                 arch,
		FlushCache:           flush,
		DisableGPGValidation: validation,
	}

	if err := c.Create(options); err != nil {
		log.Printf("ERROR: %s\n", err.Error())
	}
}
