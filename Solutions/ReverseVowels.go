package Solutions

import (
	// "fmt"
	"slices"
	"strings"
)

func reverseVowels(s string) string {
    low, high := 0, len(s)-1
    vowels := "aeiouAEIOU"
    runeSlice := []rune(s) // Convert the string to a slice of runes

    for low < high {
        // Move the `low` pointer to the right until it points to a vowel
        for low < high && !strings.Contains(vowels, string(runeSlice[low])) {
            low++
        }
        // Move the `high` pointer to the left until it points to a vowel
        for low < high && !strings.Contains(vowels, string(runeSlice[high])) {
            high--
        }

        if low < high {
            // Swap the vowels at positions pointed by `low` and `high`
            runeSlice[low], runeSlice[high] = runeSlice[high], runeSlice[low]
            low++
            high--
        }
    }
    
    // Convert the rune slice back to a string
    return string(runeSlice)
}

func ReverseVowels(s string) string {
    return reverseVowels(s)
}

// func iReverseVowels(s string) string {
// 	leftString := ""
// 	rightString := ""

// 	leftIndex := 0
// 	rightIndex := len(s) - 1
// 	leftMap := make(map[int]string)
// 	rightMap := make(map[int]string)

//     if len(s) == 1 {
//         return s
//     }

// 	for leftIndex <= rightIndex {
// 		if isVowel(string(s[leftIndex])) {
// 			leftMap[leftIndex] = string(s[leftIndex])
// 		}
//         leftString += string(s[leftIndex])

//         if leftIndex != rightIndex {
//             if isVowel(string(s[rightIndex])) {
//                 rightMap[rightIndex] = string(s[rightIndex])
//             }
//             rightString = string(s[rightIndex]) + rightString
//         }

//         leftString, rightString = swap(leftMap, rightMap, leftString, rightString)
        
// 		leftIndex++
// 		rightIndex--
// 	}
//     leftString, rightString = swap(leftMap, rightMap, leftString, rightString)

//     // fmt.Printf(leftString + rightString + "\n")

// 	return leftString + rightString
// }

func isVowel(value string) bool {
	return slices.Contains([]string{"a", "e", "i", "o", "u"}, strings.ToLower(value))
}

func replaceAtIndex(s string, replacement string, index int) string {
    if index >= 0 {
        runes := []rune(s)
        value := []rune(replacement)
        runes[index] = value[0]
        return string(runes)
    }
    return s
}

// func swap(leftMap map[int]string, rightMap map[int]string, leftString string, rightString string) (string, string) {
//     if len(leftMap) > 0 && len(rightMap) > 0 {
//         left := 0
//         right := 0
//         for rightKey := range rightMap {
//             right = rightKey
//             for leftKey := range leftMap {
//                 left = leftKey
//                 leftString = replaceAtIndex(leftString, rightMap[rightKey], leftKey)
//                 rightString = replaceAtIndex(rightString, leftMap[leftKey], strings.Index(rightString, rightMap[rightKey]))   
//                 break                 
//             }
//             break
//         }
//         // delete items from map
//         delete(leftMap, left)
//         delete(rightMap, right)
//     }

//     // if len(leftMap) > 0 {
//     //     for leftKey := range leftMap {

//     //     }
//     // }

//     return leftString, rightString
// }