package cli

import (
  "github.com/spf13/pflag"
)

type flagInfoString struct {
  Name        string
  Shorthand   string
  Description string
  Default     string
}

type flagInfoBool struct {
  Name        string
  Shorthand   string
  Description string
  Default     bool
}

var (
  FlagMainFileName = flagInfoString{
    Name:        "main",
    Shorthand:   "m",
    Description: "[required]main entrance file name",
  }
  FlagOutputFile = flagInfoString{
    Name:        "output",
    Shorthand:   "o",
    Description: "output name",
    Default:     "",
  }
  FlagGO111MODULE = flagInfoString{
    Name:        "go111",
    Shorthand:   "g",
    Description: "if found file 'go.mod',this variable will be set 'on' automatic",
    Default:     "off",
  }
  FlagImportPath = flagInfoString{
    Name:        "import-path",
    Shorthand:   "i",
    Description: "variables import path",
    Default:     "main",
  }
  FlagVersionName = flagInfoString{
    Name:        "version",
    Shorthand:   "v",
    Description: "variable name of version",
    Default:     "_version_",
  }

  FlagBranchName = flagInfoString{
    Name:        "branch",
    Shorthand:   "b",
    Description: "variable name of branch",
    Default:     "_branch_",
  }
  FlagCommitIdName = flagInfoString{
    Name:        "commit",
    Shorthand:   "c",
    Description: "variable name of commitId",
    Default:     "_commitId_",
  }
  FlagBuildTimeName = flagInfoString{
    Name:        "build-time",
    Shorthand:   "t",
    Description: "variable name of buildTime",
    Default:     "_buildTime_",
  }
)

func stringFlag(f *pflag.FlagSet, valPtr *string, flagInfo flagInfoString) {
  f.StringVarP(valPtr,
    flagInfo.Name,
    flagInfo.Shorthand,
    flagInfo.Default,
    flagInfo.Description)
}

func boolFlag(f *pflag.FlagSet, valPtr *bool, flagInfo flagInfoBool) {
  f.BoolVarP(valPtr,
    flagInfo.Name,
    flagInfo.Shorthand,
    flagInfo.Default,
    flagInfo.Description)
}
