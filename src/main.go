package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    if len(os.Args) < 2 {
        showHelp()
        os.Exit(1)
    }

    command := strings.ToLower(os.Args[1])

    switch command {
    case "codecommit", "cc", "-c":
        fmt.Println("Not implemented yet!")
        os.Exit(0)
    case "codebuild", "cb", "-b":
        fmt.Println("Not implemented yet!")
        os.Exit(0)
    case "help", "h":
        showHelp()
    default:
        fmt.Println("Invalid option:", command)
        os.Exit(1)
    }

}

func showHelp() {
    fmt.Print(`Usage: 
  waws COMMAND [SUBCOMMAND] [OPTIONS|ARGUMENTS]

Commands/subcommands:
    codecommit              cc  -c      CodeCommit operations
    
        list [search]       l   -l      List repositories
        clone REPO_NAME     c   -c      Clone repository

  codebuild                 db  -b      CodeBuild operations

  help                      h   -h      This help
`)
}