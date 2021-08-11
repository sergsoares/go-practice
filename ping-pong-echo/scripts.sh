deploy(){
    DROPLET_IP=$(getDropletIP)

    go build
	scp webapp "root@${DROPLET_IP}:/root"
	ssh root@${DROPLET_IP} "./webapp"
}

getDropletId(){
	DROPLET_ID=$(doctl compute droplet list --format ID --no-header --tag-name go-simple-app)
    echo $DROPLET_ID
}

getDropletIP(){
    DROPLET_IP=$(doctl compute droplet list --format PublicIPv4 --no-header --tag-name go-simple-app)
    echo $DROPLET_IP
}

connect() {
    ssh root@$(getDropletIP)
}

service.copy() {
    DROPLET_IP=$(getDropletIP)

    scp webapp.service "root@${DROPLET_IP}:/etc/systemd/system/appgo.service"
}

service.install() {
    DROPLET_IP=$(getDropletIP)

    ssh "root@${DROPLET_IP}" "systemctl daemon-reload"
    ssh "root@${DROPLET_IP}" "service appgo enable"
    ssh "root@${DROPLET_IP}" "service appgo start"
    ssh "root@${DROPLET_IP}" "service appgo status"
    ssh "root@${DROPLET_IP}" "mkdir /var/log/appgoservice"	
    scp webapp.service "root@${DROPLET_IP}:/var/log/appgoservice"
    ssh "root@${DROPLET_IP}" "systemctl restart rsyslog.service"
}

service.restart() {
	ssh root@$(getDropletIP) "service appgo status"
}

logs.journal(){
	ssh root@$(getDropletIP) "journalctl -b | grep appgo"
}

logs.tail(){
	ssh root@$(getDropletIP) "watch tail -n 15 /var/log/appgoservice/output.log"
}

create() {
	doctl compute droplet create golang-simple-app --image ubuntu-20-04-x64  --region nyc3 --size s-1vcpu-1gb --ssh-keys 43:7d:f6:a5:2e:15:78:4e:58:8a:f8:1a:ae:47:bf:5f --tag-names go-simple-app,lab --wait
}

stop(){
	doctl compute droplet-action shutdown $(getDropletId)
}

destroy(){
	doctl compute droplet delete $(getDropletId) 
}

$1
