package main

import "strings"

import "github.com/chzyer/readline"


func main() {
    const prompt = "gosh$ ";
    var rl, err = readline.New(prompt);

    if err != nil {
        panic(err);
    }
    defer rl.Close();

    for {
        var line, err = rl.Readline();
        var input = strings.Trim(line, " ");

        if err != nil {
            break;
        } else if input == "exit" {
            break;
        }

        println(input);
    }
}
