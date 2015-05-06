package main

import "fmt"

func main() {
	//var x [5]int

	//x[4] = 10

	//fmt.Println(x) //[0 0 0 0 10]

	//fmt.Println(x[4]) // 10

	//var x [5]float64

	/* [5] == 0,1,2,3,4 */
	/*x[0] = 20
	x[1] = 21
	x[2] = 45
	x[3] = 59
	x[4] = 71*/
	/*
		also using => x := [5]float64{20, 21, 45, 59, 71}

		x := [5]float64{
			20,
			21,
			45,
			59,
			71,
		}
	*/

	//var total float64 = 0
	
	/*for i := 0; i < 5; i++ {
		total += x[i]
	}

	fmt.Println(total / 5) // 43.2*/

	/*for i := 0; i < len(x); i++ {
		total += x[i]
	}

	//invalid operation: total / 5 (mismatched types float64 and int)
	//fmt.Println(total / len(x)) // use type conversion for fixed

	fmt.Println(total / float64(len(x))) // 43.2 it's work :)*/

	/*

	like c#'s foreach
	for i, val := range x { // i declared and not used
		total += val
	}

	fmt.Println(total / float64(len(x))) */

	/*for _, val := range x {
		total += val
	}

	fmt.Println(total / float64(len(x))) // it's work :)*/

	/* SLICES */

	// var x []float64 slices

	//x := make([]float64, 5)

	// x := make([]float64, 5, 10) // 3rd parameter == slice capacity

	/*
	arr := [5]float64{1,2,3,4,5}

	x := arr[0:5] // [low : high]

	fmt.Println(x) // [1 2 3 4 5]

	*/

	/* SLICE FUNCTIONS */

	/*sl1 := []int{1,2,3}
	sl2 := append(sl1, 4,5)

	fmt.Println(sl1, sl2) // [1 2 3] [1 2 3 4 5]*/

	/*sl1 := []int{1,2,3}
	sl2 := make([]int, 2)

	copy(sl2, sl1) // copy sl1's first two values into sl2. because sl2 length == 2

	fmt.Println(sl2, sl1) // [1 2] [1 2 3]*/

	/* MAPS */

	/*
	var x map[string]int

	x["key"] = 10

	fmt.Println(x)*/

	/*
	x := make(map[string]int)

	x["key"] = 10
	x["key"] = 20
	x["val"] = 30

	//fmt.Println(x) // map[val:20 key:10]

	fmt.Println("X Key: ", x["key"]) // latest key (element)
	fmt.Println("X Val: ", x["val"]) // latest val (element)

	//fmt.Println(x["key"])

	delete(x, "key") // delete key
	
	fmt.Println(x["key"]) // 0
	*/

	// chemical elements example

	/*elements := make(map[string]string)

	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["He"]) // Helium

	name, ok := elements["Fe"]

	fmt.Println(name, ok) // false

	if name, ok := elements["Fe"]; ok {
		fmt.Println(name, ok)
	}
	*/

	/*elements := map[string]string {
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"C": "Carbon",
		"Ne": "Neon",
	}

	fmt.Println(elements["H"]) // Hydrogen
	*/

	elements := map[string]map[string]string{
        "H": map[string]string{
            "name":"Hydrogen", 
            "state":"gas",
        },
        "He": map[string]string{
            "name":"Helium", 
            "state":"gas",
        },
        "Li": map[string]string{
            "name":"Lithium", 
            "state":"solid",
        },
        "Be": map[string]string{
            "name":"Beryllium", 
            "state":"solid",
        },
        "B":  map[string]string{
            "name":"Boron",
            "state":"solid",
        },
        "C":  map[string]string{
            "name":"Carbon",
            "state":"solid",
        },
        "N":  map[string]string{
            "name":"Nitrogen",
            "state":"gas",
        },
        "O":  map[string]string{
            "name":"Oxygen",
            "state":"gas",
        },
        "F":  map[string]string{
            "name":"Fluorine",
            "state":"gas",
        },
        "Ne":  map[string]string{
            "name":"Neon",
            "state":"gas",
        },
    }

    /*if el, ok := elements["Li"]; ok {    
        fmt.Println(el["name"], el["state"]) // Lithium Solid
    }*/

    // fmt.Println(elements["He"]["state"]) // gas
    // fmt.Println(elements["name"]["gas"]) // nil


}
