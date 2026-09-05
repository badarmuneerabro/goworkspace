package main

type Doubler interface{
	Double()
}

type Doubleint int

func (d *DoubleInt) Double(){
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double(){
	for i := range d {
		d[i] = d[i] * 2
	}
}

func DoublerCompare(d1, d2 Doubler) Doubler{
	return d1 == d2
}
func main(){
	var di DoubleInt = 10
	var di2 DoubleInt = 10
	var dis = DoubleIntSlice{1, 2, 3}
	var dis2 = DoubleIntSlice{1, 2, 3}
	
	fmt.Println(DoublerCompare(&di, &di2)) // false because comparing addresses of different pointers.
	fmt.Println(DoubleCompare(&di, dis)) // Why this is false? Aren't both Doubler?
}