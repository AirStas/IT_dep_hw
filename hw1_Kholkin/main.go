package main

import (
	"fmt"
	"os"
    "sort"
    "strings"
)

func ReadFile(input string) ([]string, error) {
    data, err := os.ReadFile(input)
    if err != nil {
        return nil, err
    }
    if len(data) == 0 {
        return nil, fmt.Errorf("empty file")
    }
    return strings.Split(string(data), "\n"), nil
}

func FindUniqueStrings(data []string) []string {
    mp := make(map[string]int)
    for _, val := range data{
        mp[val] += 1
    }
    ans := make([]string, 0, len(mp))
    for key, val := range mp {
        if val == 1 {
            ans = append(ans, key)
        }
    }
    return ans
}

func SortedStrings(data []string) []string {
    sort.Strings(data)
    return data
}

func StringsToUpper(data []string) []string {
    for key, val := range data{
        data[key] = strings.ToUpper(val)
    }
    return data
}

func AddStringsSize(data []string) []string {
    for key, val := range data{
        data[key] = fmt.Sprintf("%v - %v bytes\n", val, len(val))
    }
    return data
}

func WriteFile(data []string, output string) error {
    file, err := os.Create(output)
    if err != nil {
        return err
    }
    defer file.Close()
    for _, val := range data{
        file.WriteString(val)
    }
    return nil
}

func main(){
    if len(os.Args) < 3{
        fmt.Printf("got %v argument(s), expected at least 3\n", len(os.Args))
        return
    }
    input, output := os.Args[1], os.Args[2]
    
    if data, err := ReadFile(input); err != nil {
        fmt.Println("reading unsuccessful: ", err)
    } else {
        strs := SortedStrings(FindUniqueStrings(data))
        strs = StringsToUpper(strs)
        strs = AddStringsSize(strs)
    
        if err := WriteFile(strs, output); err != nil {
            fmt.Println("writing unsuccessful: ", err)
            return
        }
        fmt.Println("programm completed successfully")
    }
}
