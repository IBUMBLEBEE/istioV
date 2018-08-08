package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// const (
// 	k8sApiUrl = "http://118.31.57.58:8080/api/v1"
// )

type PodInfo struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		SelfLink        string `json:"selfLink"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			Namespace         string    `json:"namespace"`
			SelfLink          string    `json:"selfLink"`
			UID               string    `json:"uid"`
			ResourceVersion   string    `json:"resourceVersion"`
			Generation        int       `json:"generation"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Labels            struct {
				App     string `json:"app"`
				Version string `json:"version"`
			} `json:"labels"`
		} `json:"metadata"`
		Spec struct {
			Replicas int `json:"replicas"`
			Selector struct {
				App        string `json:"app"`
				Deployment string `json:"deployment"`
			} `json:"selector"`
			Template struct {
				Metadata struct {
					CreationTimestamp interface{} `json:"creationTimestamp"`
					Labels            struct {
						App        string `json:"app"`
						Deployment string `json:"deployment"`
					} `json:"labels"`
				} `json:"metadata"`
				Spec struct {
					Volumes []struct {
						Name     string `json:"name"`
						HostPath struct {
							Path string `json:"path"`
						} `json:"hostPath"`
					} `json:"volumes"`
					Containers []struct {
						Name  string `json:"name"`
						Image string `json:"image"`
						Ports []struct {
							ContainerPort int    `json:"containerPort"`
							Protocol      string `json:"protocol"`
						} `json:"ports"`
						Env []struct {
							Name  string `json:"name"`
							Value string `json:"value"`
						} `json:"env"`
						Resources struct {
							Limits struct {
								CPU    string `json:"cpu"`
								Memory string `json:"memory"`
							} `json:"limits"`
							Requests struct {
								CPU    string `json:"cpu"`
								Memory string `json:"memory"`
							} `json:"requests"`
						} `json:"resources"`
						VolumeMounts []struct {
							Name      string `json:"name"`
							MountPath string `json:"mountPath"`
						} `json:"volumeMounts"`
						TerminationMessagePath string `json:"terminationMessagePath"`
						ImagePullPolicy        string `json:"imagePullPolicy"`
					} `json:"containers"`
					RestartPolicy                 string `json:"restartPolicy"`
					TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
					DNSPolicy                     string `json:"dnsPolicy"`
					SecurityContext               struct {
					} `json:"securityContext"`
					ImagePullSecrets []struct {
						Name string `json:"name"`
					} `json:"imagePullSecrets"`
				} `json:"spec"`
			} `json:"template"`
		} `json:"spec"`
		Status struct {
			Replicas             int `json:"replicas"`
			FullyLabeledReplicas int `json:"fullyLabeledReplicas"`
			ObservedGeneration   int `json:"observedGeneration"`
		} `json:"status"`
	} `json:"items"`
}

var s []string

func pod(c *gin.Context) {
	fmt.Println("pod list function")
	resp, err := http.Get("http://118.31.57.58:8080/api/v1/namespaces/default/replicationcontrollers?namespace=kube-system")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	podinfo := PodInfo{}
	json.Unmarshal([]byte(body), &podinfo)
	for _, item := range podinfo.Items {
		rcname := item.Metadata.Name
		s = append(s, rcname)
	}
	fmt.Println(len(s), s)
}

func main() {
	router := gin.Default()

	podlist := router.Group("/pod")
	podlist.GET("/list", pod)

	router.Run(":9090")
}
