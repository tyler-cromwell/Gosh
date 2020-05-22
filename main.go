package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

import "github.com/chzyer/readline"


func main() {
    const prompt = "gosh$ ";
    rl, err := readline.New(prompt);

    if err != nil {
        panic(err);
    }
    defer rl.Close();

    for {
        line, err := rl.Readline();
        input := strings.Trim(line, " ");

        if err != nil || input == "exit" {
            break;
        } else if input == "" {
            continue;
        } else {
            tokens := strings.Split(input, " ");

            cmd := tokens[0];
            args := tokens[1:];

            if cmd == "cd" {
                if len(args) == 0 {
                    home, _ := os.UserHomeDir();
                    os.Chdir(home);
                } else {
                    os.Chdir(args[0]);
                }
            } else {
                command := exec.Command(cmd, args...);
                stdoutStderr, _ := command.CombinedOutput();
                fmt.Printf("%s", stdoutStderr);
            }
        }
    }
}
