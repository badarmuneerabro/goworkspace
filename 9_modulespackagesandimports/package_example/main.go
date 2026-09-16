package main
import (

	"fmt"
	"github.com/badarmuneerabro/goworkspace/9_modulespackagesandimports/package_example/do-formatter"
	"github.com/badarmuneerabro/goworkspace/9_modulespackagesandimports/package_example/math"
	
)

func main(){
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}