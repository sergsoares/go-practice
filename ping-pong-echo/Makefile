create:
	@doctl compute droplet create golang-simple-app --image ubuntu-20-04-x64  --region nyc3 --size s-1vcpu-1gb --ssh-keys 43:7d:f6:a5:2e:15:78:4e:58:8a:f8:1a:ae:47:bf:5f --tag-names go-simple-app,lab --wait

stop:
	@doctl compute droplet-action shutdown $(doctl compute droplet list --format ID --no-header --tag-name go-simple-app)

delete:
	@doctl compute droplet delete $(doctl compute droplet list --format ID --no-header --tag-name go-simple-app) 

build: 
	@echo building

watch: 
	@air

deploy: 
	@scp webapp "root@159.65.249.81:/root"
	@ssh root@159.65.249.81:/root "./webapp"
