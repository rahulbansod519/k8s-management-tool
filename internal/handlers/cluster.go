package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterCluster registers a Kubernetes cluster
func RegisterCluster(c *gin.Context) {
	// TODO: Save cluster info (e.g., in a DB) and issue a token
	c.JSON(http.StatusOK, gin.H{"message": "Cluster registered successfully"})
}

// ClusterStatus checks cluster connection and health
func ClusterStatus(c *gin.Context) {
	// TODO: Implement cluster status logic
	c.JSON(http.StatusOK, gin.H{"status": "Cluster is healthy"})
}
