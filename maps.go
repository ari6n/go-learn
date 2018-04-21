package main
import "fmt"

type Vertex struct {
	Lat, Long float64
}

var (
	m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	m1 map[string]Vertex
)

var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)

	m1 = make(map[string]Vertex)
	m1["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m1["Bell Labs"])

	fmt.Println(m2)
}
