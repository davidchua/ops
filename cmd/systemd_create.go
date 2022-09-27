package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/rajatjindal/kubectl-modify-secret/pkg/editor"
)

// look at this to do parsing https://github.com/sergeymakinen/go-systemdconf
// maybne there's no reason to do so.

//https://github.com/rajatjindal/kubectl-modify-secret/blob/master/pkg/cmd/modify_secret.go
var (
	systemPath       string = "/etc/systemd/system"
	systemdCreateCmd        = &cobra.Command{
		Use:   "create",
		Short: "Creates a new systemd service file",
		Long:  "Creates a new systemd service file. Requires superuser access",
		Run: func(cmd *cobra.Command, args []string) {
			var serviceName string
			if len(args) == 0 {
				log.Fatalln("expecting an argument but got 0")
			}

			if len(args) == 1 {
				serviceName = args[0]
			}

			// name, description, dependencies, path, args
			bb, err := tmpl.Open("templates/systemd/create.tmpl")
			if err != nil {
				log.Fatal(err)
			}

			// create a tempfile like in modify-secret, fill the tmp file with the
			// content of the template, and then only if the user saves the tmp file,
			// it gets created in the respective dir.
			// so if the file doesn't get saved, nothing is saved and there's no need
			// for cleanup.

			tmpFile, err := ioutil.TempFile("", serviceName)
			if err != nil {
				log.Fatalf("error creating temp file: %#v", err)
			}

			_, err = io.Copy(tmpFile, bb)
			if err != nil {
				log.Fatalf("error copying file from template: %#v", err)
			}

			// with the file, add the template inside, open it up into editor
			err = editor.Edit(tmpFile.Name())
			if err != nil {
				log.Fatalf("error editing: %#v", err)
			}

			// create a new file

			newFilePath := fmt.Sprintf("%s/%s.service", systemPath, serviceName)
			f, err := os.Create(newFilePath)
			if err != nil {
				log.Fatal(err)
			}

			tmpFile.Seek(0, 0)

			_, err = io.Copy(f, tmpFile)
			if err != nil {
				log.Fatalf("io copy is %#v", err)
			}

			fmt.Printf("💾 File saved at %s\n", newFilePath)

		},
	}
)

func init() {

	// ops systemd create <name> (opens up a vim editor with a template and you fill it
	// up yourself)
	systemdCmd.AddCommand(systemdCreateCmd)

}
