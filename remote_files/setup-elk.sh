printf "temp:$(openssl passwd -crypt TEMP)\n" > .htpasswd

# Change permissions on filebat config
sudo chown root filebeat.yml 
sudo chmod go-w filebeat.yml