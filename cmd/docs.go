package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
)

var docsCmd = &cobra.Command{
	Use:          "docs",
	Short:        "Generate CLI docs in various formats",
	Hidden:       true,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = os.MkdirAll("/tmp/wait4x", os.ModePerm)
		if err != nil {
			return errors.New("unable to create directory")
		}
		//switch docType {
		//case "man":
		//	err = doc.GenManTree(rootCmd, &doc.GenManHeader{Title: "HASURA", Section: "3"}, docDirectory)
		//case "md":
			err = doc.GenMarkdownTree(rootCmd, "/tmp/wait4x")
		//case "rest":
		//	err = doc.GenReSTTree(rootCmd, docDirectory)
		//case "yaml":
		//	err = doc.GenYamlTree(rootCmd, docDirectory)
		//default:
		//	return errors.New("unknown type")
		//}
		if err != nil {
			return errors.New("generating docs failed")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}