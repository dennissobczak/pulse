package k8s_config_loader

import (
  "log"
//  "reflect"

  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/tools/clientcmd"
  "k8s.io/client-go/rest"
)

var (
  K8sClient *kubernetes.Clientset
)

func loadNonClusterLocalConfig() *rest.Config {
  // Load Kubernetes configuration from kubeconfig file
  k8sClientConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
  if err != nil {
    log.Fatalf("Failed to load kubeconfig: %v", err)
  }

//  log.Println(reflect.TypeOf(k8sClientConfig))
//  log.Println(k8sClientConfig)
  log.Println("SUCCESS: k8sClientConfig found!")

  return k8sClientConfig
}

func InitK8sClient () *kubernetes.Clientset {
  clientset, err := kubernetes.NewForConfig(loadNonClusterLocalConfig())
  if err != nil {
    log.Fatalf("Failed to create Kubernetes client: %v", err)
  }

//  log.Println(reflect.TypeOf(clientset))
  log.Println("SUCCESS: k8sClient initialized!")

  K8sClient = clientset

  return K8sClient
}

//func loadClusterLocalConfig() {
//}
