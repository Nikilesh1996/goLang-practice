package main

import (
	"fmt"
)

func getUserIds(userIds []string) map[string]int {
	counts := make(map[string]int)

	for _, userId := range userIds {
		counts[userId]++
	}

	return counts
}

func main() {
	countsMap := getUserIds([]string {"nik", "nik", "gow"})

	for key, value := range countsMap {
		fmt.Println(key, value)
	}
}