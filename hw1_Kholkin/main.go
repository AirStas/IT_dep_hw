package main

import (
	"fmt"
	"os"
    "bufio"
    "sort"
    "strings"
    "errors"
)

func MagicStrings(path, input, output string) error {
    mp := make(map[string]int)
    var ans []string
    
    file, err := os.Open(path+input)
    if err != nil{
        return errors.New(fmt.Sprint("error while opening file: ", err))
    }
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        mp[strings.ToLower(scanner.Text())] += 1
    }
    
    file.Close()
    
    for key, val := range mp {
        if val == 1 {
            ans = append(ans, key)
        }
    }
    
    sort.Strings(ans)
    
    file, err = os.Create(path+output)
    if err != nil {
        return errors.New(fmt.Sprint("error while creating file: ", err))
    }
    
    for _, val := range ans{
        file.WriteString(fmt.Sprintf("%v - %v bytes\n", strings.ToUpper(val), len(strings.ToUpper(val))))
    }
    
    file.Close()
    
    return nil
}

func main(){
    //path is a working directory, ex "./"
    path := "./"
    
    //input file in working directory, ex "input.txt"
    input := "input.txt"
    
    //output file in working directory, ex "input.txt"
    output := "output.txt"
    
    if err := MagicStrings(path, input, output); err!=nil{
        fmt.Println("An error occured! Error: ", err)
    } else {
        fmt.Println("Programm completed successfully!")
    }
}
