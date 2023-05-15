#!/bin/bash

# load balancer config file
config_file='temp/load_balancer.conf'

# ssh key
key_file='ssh_key/terraform'

# password file
pass_file='.htpasswd'
beat_config='remote_files/filebeat.yml'

# ugly list concatenating of ips from terraform output
rows=$(terraform output -raw minitwit-swarm-leader-ip-address)
rows+=' '
rows+=$(terraform output -json minitwit-swarm-manager-ip-address | jq -r .[])
rows+=' '
rows+=$(terraform output -json minitwit-swarm-worker-ip-address | jq -r .[])
rows+=' '
rows+=$(terraform output -json minitwit-swarm-elastic-ip-address | jq -r .)

# scp the file
for ip in $rows; do
    ssh -o 'StrictHostKeyChecking no' root@$ip -i $key_file "mkdir /loadbalancer"
    ssh -o 'StrictHostKeyChecking no' root@$ip -i $key_file "mkdir /proxy"
    scp -o 'StrictHostKeyChecking no' -i $key_file $config_file root@$ip:/loadbalancer/default.conf
    scp -o 'StrictHostKeyChecking no' -i $key_file $pass_file root@$ip:/proxy/.htpasswd
    scp -o 'StrictHostKeyChecking no' -i $key_file $beat_config root@$ip:/proxy/filebeat.yml
done
