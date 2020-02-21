package main

import (
	"github.com/anton-johansson/gitlab-ci-kubernetes-namespace-cleaner/pkg/clean"
	"github.com/spf13/cobra"
)

var kubeconfig string

func init() {
	command := &cobra.Command{
		Use:   "clean",
		Short: "Runs the tool",
		Run: func(command *cobra.Command, arguments []string) {
			clean.Clean(kubeconfig)
		},
	}
	command.Flags().StringVarP(&kubeconfig, "kubeconfig", "k", "", "The path to your kubeconfig. If not passed in, in-cluster mode will be used.")
	rootCommand.AddCommand(command)
}
