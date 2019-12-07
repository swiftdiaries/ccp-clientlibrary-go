package delete

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/swiftdiaries/ccp-clientlibrary-go/v3/ccp"
)

// Cluster deletes the CCP cluster with the name given by the user
func Cluster(name string) error {
	// Load CCP creds from environment variables
	username := os.Getenv("CCP_USERNAME")
	password := os.Getenv("CCP_PASSWORD")
	ccpAdmin := os.Getenv("CCP_URL")

	// Authenticate to CCP Admin
	client := ccp.NewClient(username, password, ccpAdmin)

	err := client.Login(client)
	if err != nil {
		return err
	}

	var cluster *ccp.Cluster
	cluster, err = client.GetCluster(name)
	if err != nil {
		return err
	}
	log.Infof("Cluster has UUID: %v", cluster.UUID)

	err = client.DeleteCluster(*cluster.UUID)
	if err != nil {
		return err
	}
	return nil
}
