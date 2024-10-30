package main

import (
	"fmt"
	"os"
    "sort"
    "strings"
)

func GetParams() (string, string) {
    if len(os.Args) >= 3 {
        return os.Args[1], os.Args[2]
    } else if len(os.Args) == 2 {
        return os.Args[1], "output.txt"
    } else {
        return "input.txt", "output.txt"
    }
}

func GetFile(input string) (string, error) {
    data, err := os.ReadFile(input)
    if err != nil {
        return "", err
    }
    if len(data) == 0 {
        return "", fmt.Errorf("empty file")
    }
    return string(data), nil
}

func FindUnique(data string) []string {
    mp := make(map[string]int)
    for _, val := range strings.Split(data, "\n"){
        mp[val] += 1
    }
    ans := make([]string, 0)
    for key, val := range mp {
        if val == 1 {
            ans = append(ans, key)
        }
    }
    return ans
}

func SortedStrings(data []string) []string{
    sort.Strings(data)
    return data
}

func GiveFile(data []string, output string) error {
    file, err := os.Create(output)
    if err != nil {
        return err
    }
    defer file.Close()
    for _, val := range data{
        file.WriteString(fmt.Sprintf("%v - %v bytes\n", strings.ToUpper(val), len(strings.ToUpper(val))))
    }
    return nil
}

func main(){
    // Params from command line, default val: "input.txt", "output.txt"
    input, output := GetParams()
    
    if data, err := GetFile(input); err != nil {
        fmt.Println("reading unsuccessful: ", err)
    } else {
        strs := SortedStrings(FindUnique(data))
        if err := GiveFile(strs, output); err != nil {
            fmt.Println("writing unsuccessful: ", err)
        } else {
            fmt.Println("programm completed successfully")
        }
    }
}
