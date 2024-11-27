package repositories

import (
	"context"
	"time"

	"k8s-management-tool/config"
	"k8s-management-tool/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

// RegisterCluster adds a new cluster to the database
func RegisterCluster(cluster models.Cluster) error {
	collection := config.DB.Collection("clusters") // Initialize collection dynamically
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, cluster)
	return err
}

// GetClustersByOwner retrieves clusters for a specific user
func GetClustersByOwner(ownerID string) ([]models.Cluster, error) {
	collection := config.DB.Collection("clusters") // Initialize collection dynamically
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var clusters []models.Cluster
	cursor, err := collection.Find(ctx, bson.M{"owner_id": ownerID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var cluster models.Cluster
		if err := cursor.Decode(&cluster); err != nil {
			return nil, err
		}
		clusters = append(clusters, cluster)
	}
	return clusters, nil
}
