package main
 
import "fmt"
 
const (
    Aries       = '\u2648' // Widder        March 21-April 19
    Taurus      = '\u2649' // Stier         April 20-May 20
    Gemini      = '\u264a' // Zwillinge     May 21-June 20
    Cancer      = '\u264b' // Krebs         June 21-July 22.
    Leo         = '\u264c' // Löwe          July 23-August 22
    Virgo       = '\u264d' // Jungfrau      August 23-September 22
    Libra       = '\u264e' // Waage         September 23-October 22
    Scorpius    = '\u264f' // Skorpion      October 23-November 21
    Sagittarius = '\u2650' // Schütze       November 22-December 21
    Capricornus = '\u2651' // Steinbock     December 22-January 19
    Aquarius    = '\u2652' // Wassermann    Jan 20-Feb 18
    Pisces      = '\u2653' // Fische        Feb 19-March 20
)
 
func outputWithZodiacSign(p Person) {
    var zodiacSign rune = '?'
 
    if (p.Month == 1 && p.Day >= 20) || (p.Month == 2 && p.Day <= 18) {
        zodiacSign = Aquarius
    } else if (p.Month == 2 && p.Day >= 19) || (p.Month == 3 && p.Day <= 20) {
        zodiacSign = Pisces
    } else if (p.Month == 3 && p.Day >= 21) || (p.Month == 4 && p.Day <= 19) {
        zodiacSign = Aries
    } else if (p.Month == 4 && p.Day >= 20) || (p.Month == 5 && p.Day <= 20) {
        zodiacSign = Taurus
    } else if (p.Month == 5 && p.Day >= 21) || (p.Month == 6 && p.Day <= 20) {
        zodiacSign = Gemini
    } else if (p.Month == 6 && p.Day >= 21) || (p.Month == 7 && p.Day <= 22) {
        zodiacSign = Cancer
    } else if (p.Month == 7 && p.Day >= 23) || (p.Month == 8 && p.Day <= 22) {
        zodiacSign = Leo
    } else if (p.Month == 8 && p.Day >= 23) || (p.Month == 9 && p.Day <= 22) {
        zodiacSign = Virgo
    } else if (p.Month == 9 && p.Day >= 23) || (p.Month == 10 && p.Day <= 22) {
        zodiacSign = Libra
    } else if (p.Month == 10 && p.Day >= 23) || (p.Month == 11 && p.Day <= 21) {
        zodiacSign = Scorpius
    } else if (p.Month == 11 && p.Day >= 22) || (p.Month == 12 && p.Day <= 21) {
        zodiacSign = Sagittarius
    } else if (p.Month == 12 && p.Day >= 22) || (p.Month == 1 && p.Day <= 19) {
        zodiacSign = Capricornus
    }
 
    // TODO: Assign proper value to zodiacSign using if/else branching.
    // NOTE: The runes are defined above as constants.
 
    fmt.Printf("%s %s, born on %02d.%02d.%04d, has the zodiac sign %c.\n",
        p.FirstName, p.LastName, p.Day, p.Month, p.Year, zodiacSign)
}
 
type FullName struct {
    FirstName string
    LastName  string
}
type BirthDate struct {
    Day   byte
    Month byte
    Year  uint16
}
 
type Person struct {
    FullName
    BirthDate
}
 
func main() {
    grace := Person{FullName{"Grace", "Hopper"}, BirthDate{9, 12, 1906}}
    dennis := Person{FullName{"Dennis", "Ritchie"}, BirthDate{9, 9, 1941}}
    rick := Person{FullName{"Rick", "Astley"}, BirthDate{6, 2, 1966}}
    edsger := Person{FullName{"Edsger", "Dijkstra"}, BirthDate{11, 5, 1930}}
    alan := Person{FullName{"Alan", "Turing"}, BirthDate{23, 6, 1912}}
 
    outputWithZodiacSign(grace)
    outputWithZodiacSign(dennis)
    outputWithZodiacSign(rick)
    outputWithZodiacSign(edsger)
    outputWithZodiacSign(alan)
}