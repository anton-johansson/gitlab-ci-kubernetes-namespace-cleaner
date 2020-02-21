package main

import (
        "fmt"
        "github.com/spf13/cobra"
        "os"
)

var rootCommand = &cobra.Command{
        Use:   "cleaner",
        Short: "A tool for...",
}

func main() {
        if err := rootCommand.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}
