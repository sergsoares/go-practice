package cloudinit

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

var baseyaml string = `#cloud-config
groups:
  - docker
users:
  - default
  # the docker service account
  - name: docker-service
    groups: docker
package_upgrade: true
packages:
  - apt-transport-https
  - ca-certificates
  - curl
  - gnupg-agent
  - software-properties-common`

var runcmd string = `# power_state:
#   mode: reboot
#   message: Restarting after installing docker & docker-compose
runcmd:
  # install docker following the guide: https://docs.docker.com/install/linux/docker-ce/ubuntu/
  - curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
  - sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
  - sudo apt-get -y update
  - sudo apt-get -y install docker-ce docker-ce-cli containerd.io
  - sudo systemctl enable docker
  # install docker-compose following the guide: https://docs.docker.com/compose/install/
  - sudo curl -L "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
  - sudo chmod +x /usr/local/bin/docker-compose`

func GenerateDockerCompose() string {
	result := fmt.Sprint(baseyaml, "\n", runcmd, `
  - echo dmVyc2lvbjogJzIuMicKc2VydmljZXM6CiAgZXMwMToKICAgIGltYWdlOiBkb2NrZXIuZWxhc3RpYy5jby9lbGFzdGljc2VhcmNoL2VsYXN0aWNzZWFyY2g6Ny4xMy40CiAgICBjb250YWluZXJfbmFtZTogZXMwMQogICAgZW52aXJvbm1lbnQ6CiAgICAgIC0gbm9kZS5uYW1lPWVzMDEKICAgICAgLSBjbHVzdGVyLm5hbWU9ZXMtZG9ja2VyLWNsdXN0ZXIKICAgICAgLSBkaXNjb3Zlcnkuc2VlZF9ob3N0cz1lczAyLGVzMDMKICAgICAgLSBjbHVzdGVyLmluaXRpYWxfbWFzdGVyX25vZGVzPWVzMDEsZXMwMixlczAzCiAgICAgIC0gYm9vdHN0cmFwLm1lbW9yeV9sb2NrPXRydWUKICAgICAgLSAiRVNfSkFWQV9PUFRTPS1YbXM1MTJtIC1YbXg1MTJtIgogICAgdWxpbWl0czoKICAgICAgbWVtbG9jazoKICAgICAgICBzb2Z0OiAtMQogICAgICAgIGhhcmQ6IC0xCiAgICB2b2x1bWVzOgogICAgICAtIGRhdGEwMTovdXNyL3NoYXJlL2VsYXN0aWNzZWFyY2gvZGF0YQogICAgcG9ydHM6CiAgICAgIC0gOTIwMDo5MjAwCiAgICBuZXR3b3JrczoKICAgICAgLSBlbGFzdGljCiAgZXMwMjoKICAgIGltYWdlOiBkb2NrZXIuZWxhc3RpYy5jby9lbGFzdGljc2VhcmNoL2VsYXN0aWNzZWFyY2g6Ny4xMy40CiAgICBjb250YWluZXJfbmFtZTogZXMwMgogICAgZW52aXJvbm1lbnQ6CiAgICAgIC0gbm9kZS5uYW1lPWVzMDIKICAgICAgLSBjbHVzdGVyLm5hbWU9ZXMtZG9ja2VyLWNsdXN0ZXIKICAgICAgLSBkaXNjb3Zlcnkuc2VlZF9ob3N0cz1lczAxLGVzMDMKICAgICAgLSBjbHVzdGVyLmluaXRpYWxfbWFzdGVyX25vZGVzPWVzMDEsZXMwMixlczAzCiAgICAgIC0gYm9vdHN0cmFwLm1lbW9yeV9sb2NrPXRydWUKICAgICAgLSAiRVNfSkFWQV9PUFRTPS1YbXM1MTJtIC1YbXg1MTJtIgogICAgdWxpbWl0czoKICAgICAgbWVtbG9jazoKICAgICAgICBzb2Z0OiAtMQogICAgICAgIGhhcmQ6IC0xCiAgICB2b2x1bWVzOgogICAgICAtIGRhdGEwMjovdXNyL3NoYXJlL2VsYXN0aWNzZWFyY2gvZGF0YQogICAgbmV0d29ya3M6CiAgICAgIC0gZWxhc3RpYwogIGVzMDM6CiAgICBpbWFnZTogZG9ja2VyLmVsYXN0aWMuY28vZWxhc3RpY3NlYXJjaC9lbGFzdGljc2VhcmNoOjcuMTMuNAogICAgY29udGFpbmVyX25hbWU6IGVzMDMKICAgIGVudmlyb25tZW50OgogICAgICAtIG5vZGUubmFtZT1lczAzCiAgICAgIC0gY2x1c3Rlci5uYW1lPWVzLWRvY2tlci1jbHVzdGVyCiAgICAgIC0gZGlzY292ZXJ5LnNlZWRfaG9zdHM9ZXMwMSxlczAyCiAgICAgIC0gY2x1c3Rlci5pbml0aWFsX21hc3Rlcl9ub2Rlcz1lczAxLGVzMDIsZXMwMwogICAgICAtIGJvb3RzdHJhcC5tZW1vcnlfbG9jaz10cnVlCiAgICAgIC0gIkVTX0pBVkFfT1BUUz0tWG1zNTEybSAtWG14NTEybSIKICAgIHVsaW1pdHM6CiAgICAgIG1lbWxvY2s6CiAgICAgICAgc29mdDogLTEKICAgICAgICBoYXJkOiAtMQogICAgdm9sdW1lczoKICAgICAgLSBkYXRhMDM6L3Vzci9zaGFyZS9lbGFzdGljc2VhcmNoL2RhdGEKICAgIG5ldHdvcmtzOgogICAgICAtIGVsYXN0aWMKCnZvbHVtZXM6CiAgZGF0YTAxOgogICAgZHJpdmVyOiBsb2NhbAogIGRhdGEwMjoKICAgIGRyaXZlcjogbG9jYWwKICBkYXRhMDM6CiAgICBkcml2ZXI6IGxvY2FsCgpuZXR3b3JrczoKICBlbGFzdGljOgogICAgZHJpdmVyOiBicmlkZ2U | base64 -d > /root/docker-compose.yml
  - docker-compose -f /root/docker-compose.yml up -d
`)
	return result
}

func GetFileAsBase64(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Errorf("Failure with path: %s", err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Errorf("Failure with file content: %s", err)
	}

	return base64.StdEncoding.EncodeToString(content), nil
}
