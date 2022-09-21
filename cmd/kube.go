package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"syscall"
	"time"

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

			// sort by last access date!
			sort.Slice(files, func(i, j int) bool {

				a := files[i].Sys().(*syscall.Stat_t).Atim.Sec
				b := files[j].Sys().(*syscall.Stat_t).Atim.Sec

				return a > b
			})

			for _, file := range files {
				selection = append(selection, fmt.Sprintf("%s/.kube/%s", homeDir, file.Name()))
			}

			prompt := promptui.Select{
				Label: "Select Kubeconfig",
				Items: selection,
				Size:  15,
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
				log.Fatalf("error setting symlink %#v", err)
			}
			// Note: Chtimes may have an issue with timedrift on Azure/SMB shares
			// see: https://github.com/golang/go/issues/31880
			err = os.Chtimes(result, time.Now(), time.Time{})
			if err != nil {
				log.Fatalf("error touching %#v", err)
			}
		},
	}
)

func init() {

	rootCmd.AddCommand(kubeCmd)

}
