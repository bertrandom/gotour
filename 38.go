package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    
	a := make([][]uint8, dx)
    for i := 0; i < dx; i++ {
    	a[i] = make([]uint8, dy)
        for j := 0; j < dy; j++ {
        	a[i][j] = (uint8(i) + uint8(j)) / 2 
        }
    }
    
    return a
    
}

func main() {
    pic.Show(Pic)
}