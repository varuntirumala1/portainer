package edge

import portainer "github.com/portainer/portainer/api"

// LoadEdgeJobs registers all edge jobs inside corresponding endpoint tunnel
func LoadEdgeJobs(dataStore portainer.DataStore, reverseTunnelService portainer.ReverseTunnelService) error {
	edgeJobs, err := dataStore.EdgeJob().EdgeJobs()
	if err != nil {
		return err
	}

	for _, edgeJob := range edgeJobs {
		for endpointID := range edgeJob.Endpoints {
			reverseTunnelService.AddEdgeJob(endpointID, &edgeJob)
		}
	}

	return nil
}
