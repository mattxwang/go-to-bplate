package main

func intersection(s1, s2 []string) (inter []string) {
    hash := make(map[string]bool)
    for _, e := range s1 {
        hash[e] = true
    }
    for _, e := range s2 {
        // If elements present in the hashmap then append intersection list.
        if hash[e] {
            inter = append(inter, e)
        }
    }
    //Remove dups from slice.
    inter = removeDups(inter)
    return
}

//Remove dups from slice.
func removeDups(elements []string)(nodups []string) {
    encountered := make(map[string]bool)
    for _, element := range elements {
        if !encountered[element] {
            nodups = append(nodups, element)
            encountered[element] = true
        }
    }
    return
}