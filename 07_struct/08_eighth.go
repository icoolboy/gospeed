package main

import (
    "fmt"
)

type Master struct {
    Name      string
    Family    string
    Information
    IsEnable  bool
}

type Information struct {
    Age  uint8
    ExpertInfo
}

type ExpertInfo struct {
    Category        string
    SubCategory     []string
    WorkTimeByYear  uint8
    
}

func main() {
    s := Master{Family : "Akbarimanesh", Name : "Erfan"}
    fmt.Println(s)
    
    s.Name = "Erfan"
    fmt.Println(s.Name)
    fmt.Println(s.Family)
    
    s.Information = Information{Age : 20}
    s.IsEnable = true
    
    s.Information.ExpertInfo.Category = "Programming"
    s.Information.ExpertInfo.SubCategory = []string{"Go", "PHP", "C++", "NodeJs", "Python"}
    //s.Information = Information{ExpertInfo : ExpertInfo{WorkTimeByYear : 7}}
    //if this line not be comment, struct just have one value which is WorkTimeByYear = 7
    fmt.Println(s)
}