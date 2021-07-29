package cloudinit

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateDockerComposeGivesMockedValues(t *testing.T) {
	f, _ := os.Open("mocked_cloud_init.yml")
	content, _ := ioutil.ReadAll(f)

	want := string(content)
	got := GenerateDockerCompose()
	if want != got {
		t.Error("Want : ", want, "Got: ", got)
	}
}
