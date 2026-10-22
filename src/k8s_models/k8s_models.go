package k8s_models

import (
  "log"
  "reflect"

  appsv1 "k8s.io/api/apps/v1"
  corev1 "k8s.io/api/core/v1"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/apimachinery/pkg/util/intstr"
)

var (
  K8sDeployment *appsv1.Deployment
)

func InitDeployment () *appsv1.Deployment {
//  namespace := "default"
  image := "dnsobc/api-gateway:250320" // Change this to your desired image
  deploymentName := "image-pull-test"

  // Define deployment spec
  deployment := &appsv1.Deployment{
    ObjectMeta: metav1.ObjectMeta{
      Name: deploymentName,
    },
    Spec: appsv1.DeploymentSpec{
      Replicas: int32Ptr(1),
      Selector: &metav1.LabelSelector{
        MatchLabels: map[string]string{"app": deploymentName},
      },
      Template: corev1.PodTemplateSpec{
        ObjectMeta: metav1.ObjectMeta{
          Labels: map[string]string{"app": deploymentName},
        },
        Spec: corev1.PodSpec{
          Containers: []corev1.Container{
            {
              Name:  "test-container",
              Image: image,
              //Command: []string{"sh", "-c", "echo Start; sleep 10"}, // Simulate startup
              ReadinessProbe: &corev1.Probe{
                PeriodSeconds: 1,
                InitialDelaySeconds: 2,
                ProbeHandler: corev1.ProbeHandler{
                  HTTPGet: &corev1.HTTPGetAction{
                    Path:   "/actuator/health",
                    Port:   intstr.FromInt(8080),
                    Scheme: corev1.URISchemeHTTP,
                  },
                },
              },
            },
          },
        },
      },
    },
  }
  log.Println(reflect.TypeOf(deployment))
//  log.Println(deployment)

  return deployment
}

func int32Ptr(i int32) *int32 { return &i }
