package main

import "fmt"

type FullName struct {
    FirstName string
    LastName  string
}

type BirthDate struct {
    Day   int
    Month int
    Year  int
}

type Profile struct {
    FullName
    BirthDate
    NumberOfSiblings byte
    ZodiacSign       rune
}

func main() {
    
    var me = Profile{

        FullName: FullName{
            FirstName: "Mirella",  
            LastName:  "Porciello", 
        },

        BirthDate: BirthDate{
            Day:   12,
            Month: 04,
            Year:  2008, 
        },
		
        NumberOfSiblings: 1,   
        ZodiacSign:       '\u2648', 
    }


    fmt.Println(me)

    fmt.Println("Siblings Before:", me.NumberOfSiblings)
    
	me.NumberOfSiblings++

    fmt.Println("Siblings After:", me.NumberOfSiblings)
}
