package cli

import (
  "bytes"
  "errors"
  "fmt"
  "github.com/spf13/cobra"
  "log"
  "os"
  "os/exec"
  "runtime"
  "strings"
  "time"
)

var (
  mainFileName  string
  outputFile    string
  GO111MODULE   string
  importPath    string
  versionName   string
  branchName    string
  commitIdName  string
  buildTimeName string
)

var startCmd = &cobra.Command{
  Use:     "start",
  Short:   "start an instance",
  Long:    "",
  Example: "",
  RunE:    runStart,
}

func exists(path string) bool {
  _, err := os.Stat(path)
  if err != nil {
    if os.IsExist(err) {
      return true
    }
    return false
  }
  return true
}

func splitVersion(v string) [3][]byte {
  var x [3][]byte
  s := strings.Split(v, ".")
  sLen := len(s)
  if sLen > 3 {
    s = s[:3]
  }
  for k, v := range s {
    if v != "0" {
      v = strings.TrimLeft(v, "0")
    }
    x[k] = []byte(v)
  }
  return x
}

func VersionCompare(v1, v2 string) (result int) {
  x1 := splitVersion(v1)
  x2 := splitVersion(v2)

  for k := range x1 {
    x1Len := len(x1[k])
    x2Len := len(x2[k])
    if x1Len > x2Len {
      return 1
    } else if x1Len < x2Len {
      return -1
    }

    result = bytes.Compare(x1[k], x2[k])
    if result != 0 {
      break
    }
  }
  return
}

func init() {
  flags := startCmd.Flags()

  stringFlag(flags, &mainFileName, FlagMainFileName)
  stringFlag(flags, &outputFile, FlagOutputFile)
  stringFlag(flags, &GO111MODULE, FlagGO111MODULE)
  stringFlag(flags, &importPath, FlagImportPath)
  stringFlag(flags, &versionName, FlagVersionName)
  stringFlag(flags, &branchName, FlagBranchName)
  stringFlag(flags, &commitIdName, FlagCommitIdName)
  stringFlag(flags, &buildTimeName, FlagBuildTimeName)
}

func runStart(cmd *cobra.Command, args []string) error {
  if runtime.GOOS == "windows" {
    return errors.New("Unsupport windows")
  }
  startTime := time.Now()
  if mainFileName == "" {
    return cmd.Usage()
  }

  if exists("./go.mod") {
    GO111MODULE = "on"
  }

  var e *exec.Cmd
  var out []byte
  var err error
  e = exec.Command("go", "version")
  out, err = e.CombinedOutput()
  if err != nil {
    log.Fatal(err)
  }
  versionSplit := strings.Split(string(out), " ")
  if len(versionSplit) < 3 || !strings.HasPrefix(versionSplit[2], "go") {
    log.Fatal("unknown go version")
  }
  separator := "="
  goVersion := strings.TrimLeft(versionSplit[2], "go")
  if VersionCompare(goVersion, "1.5.0") <= 0 {
    separator = " "
  }

  values := make(map[string]string)

  if versionName != "" {
    e = exec.Command("git", "describe", "--abbrev=0", "--tags")
    out, _ = e.CombinedOutput()
    version := string(out)
    if strings.HasPrefix(version, "fatal") {
      version = "unknown"
    }
    values[versionName] = version
  }

  if branchName != "" {
    e = exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
    out, _ = e.CombinedOutput()
    branch := string(out)
    if strings.HasPrefix(branch, "fatal") {
      branch = "unknown"
    }
    values[branchName] = branch
  }
  if commitIdName != "" {
    e = exec.Command("git", "log", `--pretty=format:"%h"`, "-1")
    out, _ = e.CombinedOutput()
    commitId := string(out)
    if strings.HasPrefix(commitId, "fatal") {
      commitId = "unknown"
    }
    values[commitIdName] = strings.Trim(commitId, `"`)
  }
  if buildTimeName != "" {
    values[buildTimeName] = time.Now().Format(time.RFC3339)
  }
  valueNames := []string{versionName, branchName, commitIdName, buildTimeName}
  ldFlagsArray := []string{}
  for _, v := range valueNames {
    if v == "" {
      continue
    }
    ldFlagsArray = append(ldFlagsArray, fmt.Sprintf("-X %s.%s%s%s", importPath, v, separator, values[v]))
  }
  output := ""
  if outputFile != "" {
    output = fmt.Sprintf("-o %s", outputFile)
  }

  e = exec.Command("sh", "-c", fmt.Sprintf(`GO111MODULE=%s go build %s -ldflags "%s" %s`, GO111MODULE, output,
    strings.Join(ldFlagsArray,
      " "), mainFileName))
  //fmt.Println(e.String())
  err = e.Run()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("build success! use: %s\n", time.Now().Sub(startTime).String())
  return nil
}
