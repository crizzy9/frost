/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/otiai10/copy"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	source    string
	target    string
	adopt     bool
	overwrite bool
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Creates a symbolic link between two files or directories",
	Long: `Creates a symbolic link between two files or directories.
        Can adopt the target contents to the source or overwrite the target with the source.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("config file: ", viper.ConfigFileUsed())
		// fmt.Printf("Link called with source: %s and target: %s\n", viper.Get("global"), viper.Get("plugins"))
		// fmt.Println(viper.Get("global"))
		// fmt.Println(viper.Get("global.Author"))
		// fmt.Println(viper.Get("global.project-root"))
		// fmt.Println(viper.Get("plugins.1.name"))
		// // viper.Length("plugins")
		// fmt.Println(len(viper.GetStringSlice("plugins")))
		// fmt.Println(len(viper.GetStringMapStringSlice("plugins.1")))

    create()
	},
}

func CopyDirectory(srcDir, dest string) error {
  opt := copy.Options{
      Skip: func(info os.FileInfo, src, dest string) (bool, error) {
      return strings.HasSuffix(src, ".git"), nil
    },
  }

  return copy.Copy(srcDir, dest, opt)
}

func checkExistingLink() bool {
	if _, err := os.Lstat(target); err == nil {
		sr, _ := os.Readlink(target)
		if sr != source {
			fmt.Println("[SKIP] Symlink exists but points to a different source")
			// os.Remove(target)
			return false
		} else {
			fmt.Println("[SKIP] Symlink already exists!")
			return false
		}
	}
  return true
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func create() {
	source = resolveFilePathWithEnvVar(source)
	target = resolveFilePathWithEnvVar(target)

	fmt.Println("Creating symlink: ", source, "->", target)

  fileInfo, err := os.Stat(target)
  shouldCreateLink:= false
  if err != nil {
    fmt.Println("Target not found")
    shouldCreateLink = true
  }

  switch fileInfo.Mode() & os.ModeType {
  case os.ModeSymlink:
    if checkExistingLink() {
      shouldCreateLink = true
    }
  default:
    if adopt {
      if pathExists(source){
        os.RemoveAll(source)
      }
      CopyDirectory(target, source)
    } else if (overwrite) {
      os.RemoveAll(target)
    }
    shouldCreateLink = true
  }

  if shouldCreateLink {
    os.Symlink(source, target)
  }
}

func delete() {
}

func init() {
	rootCmd.AddCommand(linkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	linkCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "Source for the connection")
	source = resolveFilePathWithEnvVar(source)
	linkCmd.PersistentFlags().StringVarP(&target, "target", "t", "", "Target for the connection")
	target = resolveFilePathWithEnvVar(target)

	linkCmd.Flags().BoolVar(&adopt, "adopt", false, "Adopt existing contents from the target to source")
	linkCmd.Flags().BoolVar(&overwrite, "overwrite", false, "Overwrite the target with the source")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
