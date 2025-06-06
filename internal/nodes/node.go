package node

import (
	"fmt"

	"github.com/Prayag2003/kubernetes-simulation/internal/models"
)

var Nodes = []*models.Node{}

func InitNodePool(n int) {
	for i := 0; i < n; i++ {
		node := &models.Node{
			ID:     fmt.Sprintf("node-%d", i),
			Name:   fmt.Sprintf("node-%d", i),
			CPU:    4000, // 4000 millicores = 4 cores
			Memory: 8192, // 8 GB
		}
		Nodes = append(Nodes, node)
	}
}
