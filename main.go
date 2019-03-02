package main

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
        if err != nil {
            break;
        }
        println(line);
    }
}
