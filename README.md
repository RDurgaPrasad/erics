# erics
Provides an opinionated starting point of setting up an ansible + Vagrant development environment. See http://toja.io/using-host-and-group-vars-files-in-ansible/

# Example

Setup a Vagrant VM with ubuntu/trusty64, named redis with the ansible galaxy role azavea.redis
```
erics -b ubuntu/trusty64 redis:azavea.redis
vagrant up
ansible-galaxy install -i -r required-roles.txt
ansible-playbook -i dev playbook.yml
```
You may want to change dev/group_vars/redis.yml to override the defaults provided by azavea.redis.
