package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/manifoldco/promptui"
)

var (
	kubeCmd = &cobra.Command{
		Use:   "kube",
		Short: "Kube manipulation",
		Long:  "Kube manipulation",
		Run: func(cmd *cobra.Command, args []string) {
			// pull files in ~/.kube directory
			homeDir, err := os.UserHomeDir()
			if err != nil {
				log.Fatalf("error retrieving user's home dir: %#v", err)
			}
			files, err := ioutil.ReadDir(fmt.Sprintf("%s/.kube", homeDir))
			if err != nil {
				log.Fatal(err)
			}

			var selection []string

			for _, file := range files {
				selection = append(selection, fmt.Sprintf("%s/.kube/%s", homeDir, file.Name()))
			}

			prompt := promptui.Select{
				Label: "Select Kubeconfig",
				Items: selection,
				Size:  10,
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			kubeConfigPath := fmt.Sprintf("%s/.kube/config", homeDir)
			os.Remove(kubeConfigPath)
			err = os.Symlink(result, kubeConfigPath)
			if err != nil {
				log.Fatalf("error setting symlink", err)
			}
		},
	}
)

func init() {

	rootCmd.AddCommand(kubeCmd)

}
