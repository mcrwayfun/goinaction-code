package main

import "fmt"

func main() {
	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
	fmt.Println()
	// 调用函数来移除指定的键
	removeColor(colors, "Coral")
	// 调用函数修改指定的键的值
	updateColor(colors,"ForestGreen","update")
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

// removeColor 将指定映射里的键删除
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}

// delete 修改指定key的值
func updateColor(colors map[string]string, key string, value string) {
	colors[key] = value
}
