package main

import "fmt"

type Student struct {
    FirstName string
    LastName  string
}

type Class struct {
    Name     string
    Students []Student
}

type Module struct {
    Number int
    Classes []Class
}

func main() {
    
    classA := Class{
        Name: "Klasse S-INA24bL",
        Students: []Student{
            {"Alondra", "Prado"},
            {"Aleksander", "Stranc"},
            {"Jan", "Lehmann"},
        },
    }

    classB := Class{
        Name: "Klasse S-INP24bL ",
        Students: []Student{
            {"Joline", "Steger"},
            {"Valeriia", "Nevalikhina"},
            {"Benjamin", "Greber"},
        },
    }

    modules := map[int]Module{
        104: {Number: 104, Classes: []Class{classA}},
        117: {Number: 117, Classes: []Class{classA, classB}},
        346: {Number: 346, Classes: []Class{classB}},
    }


    fmt.Println("Klassenverwaltung:")
    fmt.Println("Klasse S-INA24bL:", classA)
    fmt.Println("Klasse S-INP24bL:", classB)

    fmt.Println("Module:")
    for number, module := range modules {
        fmt.Println("Modul", number, "wird besucht von:")
        for _, class := range module.Classes {
            fmt.Println
		}
    }
}