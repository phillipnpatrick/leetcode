package Solutions

import (
	"strings"
)

// https://leetcode.com/problems/simplify-path/description/
// You are given an absolute path for a Unix-style file system, which always begins with a slash '/'.
// Your task is to transform this absolute path into its simplified canonical path.
// The rules of a Unix-style file system are as follows:
// 	A single period '.' represents the current directory.
// 	A double period '..' represents the previous/parent directory.
// 	Multiple consecutive slashes such as '//' and '///' are treated as a single slash '/'.
// 	Any sequence of periods that does not match the rules above should be treated as a valid directory or file name. For example, '...' and '....' are valid directory or file names.
// The simplified canonical path should follow these rules:
// 	The path must start with a single slash '/'.
// 	Directories within the path must be separated by exactly one slash '/'.
// 	The path must not end with a slash '/', unless it is the root directory.
// 	The path must not have any single or double periods ('.' and '..') used to denote current or parent directories.
// Return the simplified canonical path.

// Example 1:
// Input: path = "/home/"
// Output: "/home"
// Explanation: The trailing slash should be removed.

// Example 2:
// Input: path = "/home//foo/"
// Output: "/home/foo"
// Explanation: Multiple consecutive slashes are replaced by a single one.

// Example 3:
// Input: path = "/home/user/Documents/../Pictures"
// Output: "/home/user/Pictures"
// Explanation: A double period ".." refers to the directory up a level (the parent directory).

// Example 4:
// Input: path = "/../"
// Output: "/"
// Explanation: Going one level up from the root directory is not possible.

// Example 5:
// Input: path = "/.../a/../b/c/../d/./"
// Output: "/.../b/d"
// Explanation: "..." is a valid name for a directory in this problem.

func simplifyPath(path string) string {
	stack := make([]string, len(path))
	var directory strings.Builder
	thePath := []rune(path)

	for i := 0; i < len(thePath); i++ {
		c := thePath[i]

		if c != '/' {
			directory.WriteRune(c)
		} else {
			directory, stack = updateStack(directory, stack)
		}
	}

	if directory.Len() > 0 {
		_, stack = updateStack(directory, stack)
	}

	s := asString(stack)

	return s
}

func updateStack(directory strings.Builder, stack []string) (strings.Builder, []string) {
	if directory.String() == "." {
		directory.Reset()
	} else if directory.String() == ".." {
		stack = stack[:len(stack)-1]
		directory.Reset()
	} else if directory.Len() > 0 {
		stack = append(stack, directory.String())
		directory.Reset()
	}

	return directory, stack
}

func asString(stack []string) string {
	var sb strings.Builder
	for i := 0; i < len(stack); i++ {
		if stack[i] != "" {
			sb.WriteString("/")
			sb.WriteString(stack[i])
		}
	}

	if sb.Len() == 0 {
		return "/"
	}

	return sb.String()
}
