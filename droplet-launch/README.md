# Create a Droplet for K3S

- Install [doctl](https://github.com/digitalocean/doctl)

```bash
# Auth in Digital Ocean
doctl auth init 

# After Configured launch droplet with cloudinit config
go run main -name=mydroplet
```
