package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
    if len(os.Args) < 2 {
        showHelp()
        os.Exit(1)
    }

    command := strings.ToLower(os.Args[1])

    var subcommand string

    switch command {
    case "codecommit", "cc", "-c":
        subcommand = strings.ToLower(os.Args[2]);
        switch subcommand {
        case "list", "l", "-l":
            var query string = "repositories[]"
            if len(os.Args) > 3 {
                query = fmt.Sprintf("repositories[?contains(repositoryName, '%s')]", os.Args[3])
            }
            listRepositories(query)
        case "clone", "c", "-c":
        default:
            fmt.Println("Invalid subcommand:", subcommand)
            os.Exit(1)
        }
        os.Exit(0)
    case "codebuild", "cb", "-b":
        fmt.Println("Not implemented yet!")
        os.Exit(0)
    case "help", "h":
        showHelp()
    default:
        fmt.Println("Invalid command:", command)
        os.Exit(1)
    }

}

func showHelp() {
    fmt.Print(`Usage: 
  waws COMMAND [SUBCOMMAND] [OPTIONS|ARGUMENTS]

Commands/subcommands:
    codecommit              cc  -c      CodeCommit operations
    
        list [QUERY]       l   -l      List repositories
        clone REPO_NAME     c   -c      Clone repository

  codebuild                 db  -b      CodeBuild operations

  help                      h   -h      This help
`)
}

func listRepositories(query string) {
    cmd := exec.Command("aws", "codecommit", "list-repositories", "--output", "text", "--query", query)
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error listing repositories:", err)
        os.Exit(1)
    }
    fmt.Println(string(output))
}