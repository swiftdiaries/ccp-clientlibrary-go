package create

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/swiftdiaries/ccp-clientlibrary-go/v3/ccp"
	"gopkg.in/yaml.v2"
)

// Cluster is used to create a CCP Cluster with a profile file
func Cluster(filepath string) error {
	// Load CCP creds from environment variables
	username := os.Getenv("CCP_USERNAME")
	password := os.Getenv("CCP_PASSWORD")
	ccpAdmin := os.Getenv("CCP_BASE_URL")

	// Authenticate to CCP Admin
	client := ccp.NewClient(username, password, ccpAdmin)

	err := client.Login(client)
	if err != nil {
		return err
	}

	profileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	var cluster *ccp.Cluster
	err = yaml.Unmarshal(profileBytes, &cluster)
	if err != nil {
		return err
	}
	cluster, err = client.AddCluster(cluster)
	if err != nil {
		return err
	}
	log.Infof("Cluster created with UUID: %v", cluster.UUID)

	return nil
}
