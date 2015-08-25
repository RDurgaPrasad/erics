package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Holds the configurable properties of a vagrant vm
type vm struct {
	Name string
	Box  string
	Ip   string
	Role string
}

func main() {
	var box string
	var iptmpl string
	var ipstart int
	var env string
	flag.StringVar(&box, "b", "ubuntu/trusty64", "c.vm.box setting")
	flag.StringVar(&env, "e", "dev", "Environment name. This is the name of folder/environment that holds the inventory and group_vars")
	flag.StringVar(&iptmpl, "t", "192.168.33.%d", "Sprintf template for c.vm.network ip setting. %d will be replaced by the value of -s option")
	flag.IntVar(&ipstart, "s", 2, "Start value of %d in -t template")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Printf("No machines specified.\n\n")
		help()
		os.Exit(1)
	}

	args := flag.Args()

	// parse the arguments and setup vm configurations
	var vmConfigs []vm
	hasRoles := false
	for _, arg := range args {
		var name string
		var role string
		// check for vmname:user.ansiblerole
		if strings.Contains(arg, ":") {
			split := strings.Split(arg, ":")
			name = split[0]
			role = split[1]
			hasRoles = true
		} else {
			name = arg
		}
		vmConfigs = append(vmConfigs, vm{name, box, fmt.Sprintf(iptmpl, ipstart), role})
		ipstart++
	}

	// setup directory structure.
	// See http://toja.io/using-host-and-group-vars-files-in-ansible/
	//
	// Vagrantfile
	// ansible.cfg
	// playbook.yml
    // required-roles.txt
	// env
	// |- inventory
	// |- group_vars
	//    |- {{Name}}.yml
	//
	// This way we can create multiple environments (dev, test, prod) with different configurations.
	// All we need todo is copy the env folder, rename it and adjust the variables
	// We can then run a specific environment with ansible-playbook -i env playbook.yml 

	makeDir(env)
	makeDir(env + "/group_vars")

	// write the templates
	vagrantTmpl := makeTmpl("Vagrantfile", vagrantfile)
	writeTmpl(vagrantTmpl, vmConfigs)
	inventoryTmpl := makeTmpl(env + "/inventory", inventory)
	writeTmpl(inventoryTmpl, vmConfigs)
	ansibleCfgTmpl := makeTmpl("ansible.cfg", ansibleCfg)
	writeTmpl(ansibleCfgTmpl, vmConfigs)
	playbookTmpl := makeTmpl("playbook.yml", playbook)
	writeTmpl(playbookTmpl, vmConfigs)
	
	// output the roles that might need to be installed
	if hasRoles {
		rolesTmpl := makeTmpl("required-roles.txt", requiredRoles)
		writeTmpl(rolesTmpl, vmConfigs)
		fmt.Printf("To install all roles use ansible-galaxy install -i -r required-roles.txt")
	}

	// write group_vars file for each vm
	for _, vmc := range vmConfigs {
		groupVarsTmpl := makeTmpl(env + "/group_vars/" + vmc.Name + ".yml", groupVars)
		writeTmpl(groupVarsTmpl, vmConfigs)
	}
}

func makeTmpl(name string, source string) *template.Template {
	tmpl, err := template.New(name).Parse(source)
	checkError(err)
	return tmpl
}

func makeDir(dir string) {
	file, err := os.Open(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(dir, 0777);
			checkError(err)
		} else {
			checkError(err)
		}
	} else {
		defer file.Close()
	}

}

func writeTmpl(tmpl *template.Template, vmConfigs []vm) {
	file, err := os.Create(tmpl.Name())
	checkError(err)
	defer file.Close()
	err = tmpl.Execute(file, vmConfigs)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
