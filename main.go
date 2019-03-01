package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    const prompt = "gosh$ ";
    var scanner = bufio.NewScanner(os.Stdin);

    for {
        fmt.Print(prompt);
        if scanner.Scan() {
            fmt.Println(scanner.Text());
        } else {
            break;
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err);
    }
}

