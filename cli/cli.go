package cli

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)

var (
  _version_   = ""
  _branch_    = ""
  _commitId_  = ""
  _buildTime_ = ""
)

func Start() {
  fmt.Printf("gobuild-git Version: %s, Branch: %s, Build: %s, Build time: %s\n",
    _version_, _branch_, _commitId_, _buildTime_)
  if err := Run(os.Args[1:]); err != nil {
    fmt.Fprintf(os.Stderr, "Failed running %q\n", os.Args[1])
    os.Exit(1)
  }
}

var mainCmd = &cobra.Command{
  Use:          "gobuild-git",
  Short:        "gobuild-git",
  SilenceUsage: true,
}

func init() {
  cobra.EnableCommandSorting = false

  mainCmd.AddCommand(
    startCmd,
  )
}

func Run(args []string) error {
  mainCmd.SetArgs(args)
  return mainCmd.Execute()
}
