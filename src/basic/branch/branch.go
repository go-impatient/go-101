package main
import(
	"os"
	"bufio"
	"fmt"
	"strconv"
)

func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func convertToBin(v int) string {
	result := ""
	for ; v > 0; v /= 2  {
		result = strconv.Itoa(v%2) + result
	}
	return result
}

func main() {
	// printFile("abc.txt")
	fmt.Println(convertToBin(3))
}