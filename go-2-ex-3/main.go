package main

import "fmt"

func main() {
	// TODO: create a map called "modules"

    modules := map[int]string{
        104: "Einführung in die Programmierung",
        117: "Datenbanken modellieren und implementieren",
        346: "Cloud-Lösungen konzipieren und realisieren",
    }

    fmt.Println("Modul 104:", modules[104])
    fmt.Println("Modul 117:", modules[117])
    fmt.Println("Modul 346:", modules[346])

    delete(modules, 117)

    modules[200] = "IT-Sicherheit Grundlagen"

    modules[104] = "Programmierung mit Go"

	// TODO: delete one
	// TODO: add one
	// TODO: replace one
    fmt.Println(modules)
}
