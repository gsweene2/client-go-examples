package main

import (
	"log"
	"testing"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fake "k8s.io/client-go/kubernetes/fake"
)

func TestGetListOfPodNames(t *testing.T) {
	// Arrange
	var mockPods = []apiv1.Pod{{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod1",
			Namespace: apiv1.NamespaceDefault,
			Labels: map[string]string{
				"A": "B",
				"C": "D",
			},
		},
	}, {
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod2",
			Namespace: apiv1.NamespaceDefault,
			Labels: map[string]string{
				"E": "F",
				"G": "H",
			},
		},
	}}
	var mockPodList = &apiv1.PodList{Items: mockPods}
	kubeClient := fake.NewSimpleClientset(mockPodList)

	// Act
	result := GetListOfPodNames(kubeClient)

	// Assert
	log.Println("Output: ")
	log.Println(result)
}
