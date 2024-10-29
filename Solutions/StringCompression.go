package Solutions

import (
	// "fmt"
	"strconv"
)

// Given an array of characters chars, compress it using the following algorithm:

// Begin with an empty string s. For each group of consecutive repeating characters in chars:

// If the group's length is 1, append the character to s.
// Otherwise, append the character followed by the group's length.
// The compressed string s should not be returned separately, but instead, be stored in the input character array chars. Note that group lengths that are 10 or longer will be split into multiple characters in chars.

// After you are done modifying the input array, return the new length of the array.

// You must write an algorithm that uses only constant extra space.

// Hint: How do you know if you are at the end of a consecutive group of characters?

// Input: chars = ["a","a","b","b","c","c","c"]
// Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
func compress(chars []byte) int {
	readIndex, writeIndex := 0, 0
    for readIndex < len(chars) {
        currentChar := chars[readIndex]
        currentCount := 0

        // Count the number of consecutive repeating characters
        for readIndex < len(chars) && currentChar == chars[readIndex] {
            readIndex++
            currentCount++
        }

        // Write the compression result to chars
        chars[writeIndex] = currentChar
        writeIndex++

        if currentCount > 1 {
            count := strconv.Itoa(currentCount)
            for _, i := range(count) {
                chars[writeIndex] = byte(i)
                writeIndex++
            }
        }
    }
    return writeIndex
}

func mycompress(chars []byte) (int, string) {
	// fmt.Printf("chars: %v\t%s\n", chars, string(chars))

	indexOfFirstChar := 0
	writeIndex := 0
	for i := 1; i <= len(chars); i++ {
		if i < len(chars) {
			if chars[i] != chars[i-1] {
				writeIndex = updateByteArray(chars, chars[indexOfFirstChar], writeIndex, i - indexOfFirstChar)
				indexOfFirstChar = i
				// fmt.Printf("chars: %v\t%s\n", chars, string(chars))
			}
		} else if len(chars) > 1 {
			writeIndex = updateByteArray(chars, chars[indexOfFirstChar], writeIndex, i - indexOfFirstChar)
			chars = chars[:writeIndex]
			// fmt.Printf("chars: %v\t%s\n", chars, string(chars))
		}
	}
	// fmt.Printf("chars: %v\t%s\n", chars, string(chars))

	return len(chars), string(chars)
}

func updateByteArray(chars []byte, char byte, writeIndex int, count int) int {
	bs := []byte(strconv.Itoa(count))

	chars[writeIndex] = char
	if count > 1 {
		// fmt.Printf("bs: %v (%s)", bs, string(bs))
		
		for i := 0; i < len(bs); i++ {
			chars[writeIndex+i+1] = bs[i]
			// fmt.Printf("\t%s", string(chars))
		}
		// fmt.Printf("\n")
		writeIndex++
	}
	writeIndex += len(bs)

	return writeIndex
}
