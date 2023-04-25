printf "temp:$(openssl passwd temp)\n" > .htpasswd

# Change permissions on filebeat config
sudo chown root ./remote_files/filebeat.yml 
sudo chmod go-w ./remote_files/filebeat.yml