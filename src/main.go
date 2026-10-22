package main

import (
  "context"
  "fmt"
  "log"
  "time"

  k8sConfigLoader "pulse/k8s_config_loader"
  k8sModels "pulse/k8s_models"

//  appsv1 "k8s.io/api/apps/v1"
  corev1 "k8s.io/api/core/v1"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//  "k8s.io/apimachinery/pkg/util/intstr"
//  "k8s.io/client-go/kubernetes"
//  "k8s.io/client-go/tools/clientcmd"
)

func main() {
  clientset := k8sConfigLoader.InitK8sClient()

  namespace := "default"
//  image := "dnsobc/api-gateway:250320" // Change this to your desired image
  deploymentName := "image-pull-test"

  deployment := k8sModels.InitDeployment()

	// Record start time
	startTime := time.Now()

	// Create deployment
	_, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create deployment: %v", err)
	}

	var imagePullTime, readinessTime time.Duration

	// Wait for pod readiness
	for {
		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
			LabelSelector: "app=" + deploymentName,
		})
		if err != nil {
			log.Fatalf("Failed to list pods: %v", err)
		}

		for _, pod := range pods.Items {
			for _, status := range pod.Status.ContainerStatuses {
				if status.State.Waiting == nil && status.State.Running != nil {
					imagePullTime = time.Since(startTime)
					fmt.Printf("Image pull time: %v\n", imagePullTime)
				}
			}
			if pod.Status.Conditions != nil {
				for _, cond := range pod.Status.Conditions {
					if cond.Type == corev1.PodReady && cond.Status == corev1.ConditionTrue {
						readinessTime = time.Since(startTime)
						fmt.Printf("Application readiness time: %v\n", readinessTime)
						return
					}
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}

//func int32Ptr(i int32) *int32 { return &i }

