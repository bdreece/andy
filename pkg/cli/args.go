package cli

import (
	"flag"
	"fmt"
	"os"
)

const (
	licenseText string = "andy v0.1.0\n" +
		"Copyright (C) 2024 - Brian Reece\n" +
		"This program comes with ABSOLUTELY NO WARRANTY\n"
)

type Args struct {
	ConfigPath string
}

func Parse() (args Args) {
	flag.Usage = help
	flag.StringVar(&args.ConfigPath, "c", "/etc/andy/config.yml", "config path")

	flag.Parse()

	return
}

func help() {
	fmt.Fprintf(os.Stderr, "%s\n", licenseText)
	flag.PrintDefaults()
}
