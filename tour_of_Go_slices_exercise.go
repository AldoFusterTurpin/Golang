//ALDO FUSTER TURPIN

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    ret := make([][]uint8, dy)
	
    for i := range ret {
        ret[i] = make([]uint8, dx)
	assign_value_to_slice(i, ret)
    }
    return ret
}

func assign_value_to_slice(i int, slice_of_slices [][]uint8) {
    for j := range slice_of_slices[i] {
        slice_of_slices[i][j] = uint8(i*j)
    }
}

func main() {
	pic.Show(Pic)
}
