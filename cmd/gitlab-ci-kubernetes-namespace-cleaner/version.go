package main

import (
	"fmt"
	"github.com/anton-johansson/gitlab-ci-kubernetes-namespace-cleaner/pkg/version"
	"github.com/spf13/cobra"
)

var short bool

func init() {
	command := &cobra.Command{
		Use:   "version",
		Short: "Prints the version of the tool",
		Run: func(command *cobra.Command, arguments []string) {
			info := version.GetVersionInfo()
			if short {
				fmt.Println(info.Version)
			} else {
				fmt.Println("gitlab-ci-kubernetes-namespace-cleaner " + info.Version)
				fmt.Println("go version: " + info.GoVersion)
				fmt.Println("commit: " + info.Commit)
				fmt.Println("build date: " + info.BuildDate)
				fmt.Println("operating system: " + info.OperatingSystem + "/" + info.Architechture)
			}
		},
	}
	command.Flags().BoolVarP(&short, "short", "s", false, "Whether or not to output the actual version only")
	rootCommand.AddCommand(command)
}
