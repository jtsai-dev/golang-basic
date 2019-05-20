package main

import "fmt"

func main() {
	// map 使用哈希表，key 必须可以比较相等
	// 除了 slice、map、function 的内建类型都可以作为key
	// struct 也可以作为 key
	maps := create()
	maps = add(maps, "one", 1)
	maps = add(maps, "two", 2)
	foreach(maps)
	fmt.Println(get(maps, "two"))
	fmt.Println(get(maps, "three")) // 获取不存在key的value
	remove(maps, "two")
	foreach(maps)
}

func create() map[string]int {
	return make(map[string]int)
}

func add(maps map[string]int, key string, value int) map[string]int{
	maps[key] = value
	return maps
}

func foreach(maps map[string]int){
	// map 中无法保证顺序
	for k, v := range maps {
		fmt.Println(k, v)
	}
}

func remove(maps map[string]int, key string){
	delete(maps, key)
}

func get(maps map[string]int, key string) int{
	return maps[key]
}