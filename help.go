package main

import (
	"fmt"
	"flag"
	"os"
)

func help() {
	fmt.Printf("Usage of %s:\n\n", os.Args[0])
	fmt.Printf("\tercis [options] machine1[:ansible role] machine2[:ansible role]...\n\n")
	fmt.Printf("Examples:\n\n")
	fmt.Printf("Setup vagrant VM with basic ansible configurations\n\n")
	fmt.Printf("\terics myvm\n\n")
	fmt.Printf("Setup vagrant VM with ansible galaxy roles\n\n")
	fmt.Printf("\terics redis:azavea.redis web:andyceo.nginx\n\n")
	fmt.Printf("Note: You have to install the ansible roles with\n\n")
	fmt.Printf("\tansible-galaxy install -i -r required-roles.txt\n\n")
	fmt.Printf("Use another OS\n\n")
	fmt.Printf("\terics -b ubuntu/trusty32 myvm\n\n")
	fmt.Printf("Options\n\n")
	flag.PrintDefaults()
}

func init() {
	flag.Usage = help
}