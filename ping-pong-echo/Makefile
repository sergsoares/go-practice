create:
	@doctl compute droplet create golang-simple-app --image ubuntu-20-04-x64  --region nyc3 --size s-1vcpu-1gb --ssh-keys 43:7d:f6:a5:2e:15:78:4e:58:8a:f8:1a:ae:47:bf:5f --tag-names go-simple-app,lab --wait

stop:
	@doctl compute droplet-action shutdown $(doctl compute droplet list --format ID --no-header --tag-name go-simple-app)

destroy:
	@doctl compute droplet delete $(doctl compute droplet list --format ID --no-header --tag-name go-simple-app) 

build: 
	@echo building

watch: 
	@air

deploy: 
	@scp webapp "root@159.65.249.81:/root"
	@ssh root@159.65.249.81:/root "./webapp"

service.install:
	@scp webapp.service "root@159.65.249.81:/etc/systemd/system/appgo.service"
	@ssh root@159.65.249.81:/root "systemctl daemon-reload"
	@ssh root@159.65.249.81:/root "service appgo enable"
	@ssh root@159.65.249.81:/root "service appgo start"
	@ssh root@159.65.249.81:/root "service appgo status"
	@ssh root@159.65.249.81:/root "mkdir /var/log/appgoservice"	
	@scp webapp.service "root@159.65.249.81:/var/log/appgoservice"
	@ssh root@159.65.249.81:/root "systemctl restart rsyslog.service"	

service.restart:
	@ssh root@159.65.249.81:/root "service appgo status"

logs.journal:
	@ssh root@159.65.249.81:/root "journalctl -b | grep appgo"

logs.tail:
	@ssh root@159.65.249.81:/root "watch tail -n 15 /var/log/appgoservice/output.log"
