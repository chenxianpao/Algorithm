package main

import "fmt"

func defineList() {
	var balance [10]float32
	var balance1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance1)
	var balance2 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance2)
	balance[4] = 50.0
	var salary float32 = balance[9]
	fmt.Println(salary)
}

func defineSlice() {
	// make([]T, length, capacity)
	var slice1 []int = make([]int, 10)
	fmt.Println(slice1)
	var slice2 []int = make([]int, 6, 12)
	fmt.Println(slice2[7:8])
	s := []int{1, 2, 3}
	fmt.Println(s)
	// s := arr[startIndex:endIndex]
	// s := arr[startIndex:]
	// s := arr[:endIndex]
}

func defineMap() {
	// var map_variable map[int]int
	// map_variable1 := make(map[int]int)
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}

func main() {
	defineMap()
}
