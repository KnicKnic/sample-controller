/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/scylladb/go-set/strset"
	"github.com/yyyar/gobetween/config"
	"github.com/yyyar/gobetween/launch"
	"github.com/yyyar/gobetween/manager"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var (
	masterURL      string
	kubeconfig     string
	clusterGroup   string
	clusterNetwork string
	clusterName    string
)

// func NodeIsReady(node v1.a){
// 	isReady := false
// 	for _, condition := range node.
// }

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	// // set up signals so we handle the first shutdown signal gracefully
	// stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	// exampleClient, err := clientset.NewForConfig(cfg)
	// if err != nil {
	// 	klog.Fatalf("Error building example clientset: %s", err.Error())
	// }

	// kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	// exampleInformerFactory := informers.NewSharedInformerFactory(exampleClient, time.Second*30)

	// controller := NewController(kubeClient, exampleClient,
	// 	kubeInformerFactory.Apps().V1().Deployments(),
	// 	exampleInformerFactory.Samplecontroller().V1alpha1().Foos())

	// // notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	// // Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	// kubeInformerFactory.Start(stopCh)
	// exampleInformerFactory.Start(stopCh)

	// if err = controller.Run(2, stopCh); err != nil {
	// 	klog.Fatalf("Error running controller: %s", err.Error())
	// }
	launch.Launch(config.Config{})
	delay := time.NewTicker(2 * time.Second)
	for {
		set := strset.New()

		services, err := kubeClient.CoreV1().Services(v1.NamespaceAll).List(metav1.ListOptions{TypeMeta: metav1.TypeMeta{Kind: "Service"}})
		for _, service := range services.Items {
			if service.Spec.Type == v1.ServiceTypeLoadBalancer {
				// fmt.Printf("external LBs %v", service)

				vipName := service.Namespace + "%" + service.Name
				set.Add(vipName)

				if service.Spec.LoadBalancerIP != "" {
					continue
				}

				local_address, err := ensureVip(vipName, clusterNetwork, clusterGroup, clusterName)
				if err != nil {
					fmt.Println("Error making vip", err)
					continue
				}
				if local_address == "" || local_address == "0.0.0.0" {

					fmt.Printf("Error invalid local_address %#v\n", local_address)
					continue
				}

				service.Spec.LoadBalancerIP = local_address
				_, err = kubeClient.CoreV1().Services(service.Namespace).Update(&service)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		}

		// get list of manager objects and delete any not needed
		names, err := listVipNames(clusterGroup, clusterName)
		if err != nil {
			fmt.Println("Couldn't get list of vips", err)
		} else {
			for _, name := range names {
				if !set.Has(name) {
					manager.Delete(name)
					deleteVip(name, clusterName)
				}
			}
		}

		<-delay.C
	}

}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&clusterGroup, "group", "", "The group to create the cluster in")
	flag.StringVar(&clusterNetwork, "network", "", "The network to create the VIPs on")
	flag.StringVar(&clusterName, "cluster", "localhost", "The cluster name to use, default value is localhost")
}
