package main

import (
	"fmt"
	"time"

	"github.com/Prayag2003/kubernetes-simulation/internal/kubeapi"
)

func main() {
	kube := kubeapi.NewKubeAPI()

	id := kube.CreatePod("nginx")
	time.Sleep(4 * time.Second)

	kube.DeletePod(id)
	time.Sleep(2 * time.Second)

	fmt.Println("[Main] simulation successfull.")
}
