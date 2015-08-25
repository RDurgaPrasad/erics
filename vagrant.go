package main

// vagrant related templates
var (
	vagrantfile string = `# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  {{range .}}
  config.vm.define "{{.Name}}" do |c|
    c.vm.box = "{{.Box}}"
    c.vm.network "private_network", ip: "{{.Ip}}"
  end
  {{end}}
end`
)
