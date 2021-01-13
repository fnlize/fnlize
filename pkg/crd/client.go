/*
Copyright 2016 The Fission Authors.

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

package crd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	genClientset "github.com/fission/fission/pkg/apis/genclient/clientset/versioned"
)

type (
	FissionClient struct {
		genClientset.Interface
	}
)

// GetKubeconfig Get kube config file path from environment var $KUBECONFIG,
// or get $HOME/.kube/config
func GetKubeconfig() (kubeconfig string, err error) {
	kubeconfig = os.Getenv("KUBECONFIG")
	if len(kubeconfig) != 0 {
		return
	}

	var homeDir string
	if homeDir, err = os.UserHomeDir(); err != nil {
		err = fmt.Errorf("cannot get kube config file")
		return
	}
	var fileInfo os.FileInfo
	if fileInfo, err = os.Stat(path.Join(homeDir, ".kube", "config")); err != nil {
		err = fmt.Errorf("cannot get kube config file")
		return
	}
	if fileInfo.IsDir() {
		err = fmt.Errorf("cannot get kube config file")
		return
	}
	kubeconfig = path.Join(homeDir, ".kube", "config")
	return
}

// GetClientset Get a kubernetes client using the kubeconfig file at the
// environment var $KUBECONFIG, or an in-cluster config if that's undefined.
func GetClientset() (config *rest.Config, err error) {
	var kubeConfig string
	if kubeConfig, err = GetKubeconfig(); err != nil {
		fmt.Printf("get kube config file with error: %v", err)
		err = nil // clean errors

		if config, err = rest.InClusterConfig(); err != nil {
			return
		}
	} else {
		if config, err = clientcmd.BuildConfigFromFlags("", kubeConfig); err != nil {
			return
		}
	}
	return
}

// GetKubernetesClient Get a kubernetes client using the kubeconfig file at the
// environment var $KUBECONFIG, or an in-cluster config if that's undefined.
func GetKubernetesClient() (*rest.Config, *kubernetes.Clientset, *apiextensionsclient.Clientset, error) {
	var config *rest.Config
	var err error

	if config, err = GetClientset(); err != nil {
		return nil, nil, nil, err
	}

	// creates the client set
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, err
	}

	apiExtClientset, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, err
	}

	return config, clientset, apiExtClientset, nil
}

func MakeFissionClient() (*FissionClient, *kubernetes.Clientset, *apiextensionsclient.Clientset, error) {
	config, kubeClient, apiExtClient, err := GetKubernetesClient()
	if err != nil {
		return nil, nil, nil, err
	}

	// make a CRD REST client with the config
	crdClient, err := genClientset.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, err
	}

	fc := &FissionClient{
		Interface: crdClient,
	}
	return fc, kubeClient, apiExtClient, nil
}

func (fc *FissionClient) WaitForCRDs() error {
	start := time.Now()
	for {
		fi := fc.CoreV1().Functions(metav1.NamespaceDefault)
		_, err := fi.List(metav1.ListOptions{})
		if err != nil {
			time.Sleep(100 * time.Millisecond)
		} else {
			return nil
		}

		if time.Since(start) > 30*time.Second {
			return errors.New("timeout waiting for CRDs")
		}
	}
}

// GetDynamicClient creates and returns new dynamic client or returns an error
func GetDynamicClient() (dynamic.Interface, error) {
	var config *rest.Config
	var err error

	if config, err = GetClientset(); err != nil {
		return nil, err
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return dynamicClient, nil
}
