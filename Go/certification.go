package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the 'ModifyString' function below and add imports if needed.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING str as a parameter.
 */

func ModifyString(str string) string {
    // Trim leading and trailing spaces
    str = strings.TrimSpace(str)

    // Remove digits from the string
    result := ""
    for _, char := range str {
        if !('0' <= char && char <= '9') {
            result += string(char)
        }
    }

    // Reverse the resulting string
    reversed := ""
    for i := len(result) - 1; i >= 0; i-- {
        reversed += string(result[i])
    }

    return reversed
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16*1024*1024)

    str := readLine(reader)

    result := ModifyString(str)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

