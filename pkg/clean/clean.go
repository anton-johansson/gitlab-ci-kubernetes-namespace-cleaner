package clean

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"regexp"
)

func config(kubeconfig string) *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return config
}

func Clean(kubeconfig string) {
	config := config(kubeconfig)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, namespace := range namespaces.Items {
		var namespaceName string = namespace.ObjectMeta.Name

		matched, err := regexp.MatchString("^gitlab-ci-test-.+", namespaceName)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if !matched {
			continue
		}

		fmt.Println("Processing " + namespaceName)
		pods, err := clientset.CoreV1().Pods(namespaceName).List(metav1.ListOptions{})
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		var numberOfPods int = len(pods.Items)
		if numberOfPods > 0 {
			fmt.Println("Namespace '" + namespaceName + "' has pods, skipping it")
			continue
		}

		fmt.Println("Namespace '" + namespaceName + "' is empty, removing it")
		err = clientset.CoreV1().Namespaces().Delete(namespaceName, &metav1.DeleteOptions{})
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println("Namespace '" + namespaceName + "' was removed")
	}

	fmt.Println("Done")
}
