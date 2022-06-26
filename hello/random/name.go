// Package to generate random values.
package random

import "math/rand"

var Names = [3]string { "Israel", "Gabriela", "Cristhian" }

// ReverseRunes returns its argument string reversed rune-wise left to right.
func RandomName() string {
    return Names[rand.Intn(len(Names))]
}