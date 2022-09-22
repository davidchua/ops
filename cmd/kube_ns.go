package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/manifoldco/promptui"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	kubensCmd = &cobra.Command{
		Use:   "ns",
		Short: "Manipulate nameservers!",
		Long:  "Quickly switch nameservers with a few keystrokes",
		Run: func(cmd *cobra.Command, args []string) {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				log.Fatalf("error retrieving user's home dir: %#v", err)
			}

			kubeconfig := filepath.Join(homeDir, ".kube", "config")

			config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
			if err != nil {
				panic(err.Error())
			}

			// create the clientset
			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				panic(err.Error())
			}

			l, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
			if err != nil {
				log.Fatal(err)
			}

			var selection []string

			for _, file := range l.Items {
				selection = append(selection, file.Name)
			}

			prompt := promptui.Select{
				Label: "Select Namespace",
				Items: selection,
				Size:  30,
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			// this assumes we're using the default path for our kubeconfig
			pathOptions := clientcmd.NewDefaultPathOptions()

			startingConfig, err := pathOptions.GetStartingConfig()
			if err != nil {
				log.Fatal(err)
			}

			name := startingConfig.CurrentContext

			startingStanza, exists := startingConfig.Contexts[name]
			if !exists {
				startingStanza = clientcmdapi.NewContext()
			}

			startingStanza.Namespace = result

			if err := clientcmd.ModifyConfig(pathOptions, *startingConfig, true); err != nil {
				log.Fatalf("error modifying config: %#v", err)
			}
		},
	}
)

func init() {

	kubeCmd.AddCommand(kubensCmd)

}
