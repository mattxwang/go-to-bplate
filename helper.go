package main

import "strings"

func intersection(s1, s2 []string) (inter []string) {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	inter = removeDups(inter)
	return
}

func insensitiveIntersection(s1, s2 []string) (inter []string) {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[strings.ToLower(e)] = true
	}
	for _, e := range s2 {
		if hash[strings.ToLower(e)] {
			inter = append(inter, e)
		}
	}
	inter = removeDups(inter)
	return
}

func removeDups(elements []string) (nodups []string) {
	encountered := make(map[string]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}

func splitStringsByComma(input string) []string {
	response := strings.Split(strings.TrimSpace(input), ",")
	if response[0] != "" {
		return response
	}
	return []string{}
}
