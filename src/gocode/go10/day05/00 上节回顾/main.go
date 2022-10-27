package main

import "fmt"

func setage(age *int) {
	*age++
}

func main() {

	var age = 22
	setage(&age)
	fmt.Println(age)

	//var names = []string{"yuan", "rain", "alvin"}
	//var ages = []int{32, 33, 45}
	//var info map[string]string

	//china := make([]map[string][]string, 1)
	//china[0] = map[string][]string{"北京": {"海淀"}}
	//fmt.Println(china)

	//var china = make(map[string][]string)
	//china["北京"] = []string{"海淀", "朝阳"}
	//fmt.Println(china)

}
