package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/airware/vili/kube/extensions/v1beta1"
)

// LogMessage is a wrapper for log messages in the db
type LogMessage struct {
	Time    time.Time `json:"time"`
	Message string    `json:"msg"`
	Level   string    `json:"level"`
}

func getPortFromDeployment(deployment *v1beta1.Deployment) (int32, error) {
	containers := deployment.Spec.Template.Spec.Containers
	if len(containers) == 0 {
		return 0, fmt.Errorf("no containers in controller")
	}
	ports := containers[0].Ports
	if len(ports) == 0 {
		return 0, fmt.Errorf("no ports in controller")
	}
	return ports[0].ContainerPort, nil
}

func getImageTagFromDeployment(deployment *v1beta1.Deployment) (string, error) {
	containers := deployment.Spec.Template.Spec.Containers
	if len(containers) == 0 {
		return "", fmt.Errorf("no containers in deployment")
	}
	image := containers[0].Image
	imageSplit := strings.Split(image, ":")
	if len(imageSplit) != 2 {
		return "", fmt.Errorf("invalid image: %s", image)
	}
	return imageSplit[1], nil
}
