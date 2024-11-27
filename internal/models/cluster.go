package models

type Cluster struct {
	ID          string `bson:"_id,omitempty"`
	Name        string `bson:"name"`
	OwnerID     string `bson:"owner_id"`     // ID of the user who owns the cluster
	ClusterInfo string `bson:"cluster_info"` // Serialized data about the cluster
	Token       string `bson:"token"`        // Authentication token for the cluster
}
