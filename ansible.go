package main

// ansible related templates
var (
	inventory string = `{{range .}}[{{.Name}}]
{{.Name}} ansible_ssh_host={{.Ip}} ansible_ssh_port=22 ansible_ssh_user=vagrant ansible_ssh_private_key_file=.vagrant/machines/{{.Name}}/virtualbox/private_key

{{end}}`

	ansibleCfg string = `[defaults]
host_key_checking=False
`
	
	playbook string = `---{{range .}}{{if .Role}}
- hosts: {{.Name}}
  sudo: yes
  roles:
    - { role: {{.Role}} }{{end}}{{end}}
`
	
	requiredRoles string = `{{range .}}{{if .Role}}{{.Role}}
{{end}}{{end}}`

	groupVars string = `---
`
)
