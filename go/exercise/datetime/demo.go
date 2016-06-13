package main
import(
	"time"
	"fmt"
	"strconv"
)

func main(){
	test, err := time.Parse(time.UnixDate, "1443544106.731")
	fmt.Println(strconv.ParseFloat("1443544106.731", 64))
	// test, err := time.Parse("20060102150405", "20150930003325")
	fmt.Println(err)
	fmt.Println(test)
}
