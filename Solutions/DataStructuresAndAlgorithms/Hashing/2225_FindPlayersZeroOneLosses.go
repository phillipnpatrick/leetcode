package Solutions

import "sort"

// https://leetcode.com/problems/find-players-with-zero-or-one-losses/description/
// You are given an integer array matches where matches[i] = [winneri, loseri] indicates that the player winneri defeated player loseri in a match.
// Return a list answer of size 2 where:
// answer[0] is a list of all players that have not lost any matches.
// answer[1] is a list of all players that have lost exactly one match.
// The values in the two lists should be returned in increasing order.

// Note:
// You should only consider the players that have played at least one match.
// The testcases will be generated such that no two matches will have the same outcome.

// Example 1:
// Input: matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]
// Output: [[1,2,10],[4,5,7,8]]
// Explanation:
// Players 1, 2, and 10 have not lost any matches.
// Players 4, 5, 7, and 8 each have lost one match.
// Players 3, 6, and 9 each have lost two matches.
// Thus, answer[0] = [1,2,10] and answer[1] = [4,5,7,8].

// Example 2:
// Input: matches = [[2,3],[1,3],[5,4],[6,4]]
// Output: [[1,2,5,6],[]]
// Explanation:
// Players 1, 2, 5, and 6 have not lost any matches.
// Players 3 and 4 each have lost two matches.
// Thus, answer[0] = [1,2,5,6] and answer[1] = [].

// runtime 62ms | beats 100% | memory 19.22mb beats 90.41%
func findWinners(matches [][]int) [][]int {
	lossesCount := make([]int, 100001)
	ans := make([][]int, 2)
	ans[0] = []int{}
	ans[1] = []int{}

	for i := range lossesCount {
		lossesCount[i] = -1
	}

	for i := range matches {
		winner := matches[i][0]
		loser := matches[i][1]

		if lossesCount[winner] == -1 {
			lossesCount[winner] = 0
		}
		if lossesCount[loser] == -1 {
			lossesCount[loser] = 1
		} else {
			lossesCount[loser]++
		}
	}

	for i := range lossesCount {
		if lossesCount[i] == 0 {
			ans[0] = append(ans[0], i)
		} else if lossesCount[i] == 1 {
			ans[1] = append(ans[1], i)
		}
	}

	return ans
}

// runtime 67ms | beats 97.53% | memory 20.68mb beats 76.71%
func findWinners4(matches [][]int) [][]int {
	lossesCount := make(map[int]int)
	ans := make([][]int, 2)
	ans[0] = []int{}
	ans[1] = []int{}

	for i := 0; i < len(matches); i++ {
		winner := matches[i][0]
		loser := matches[i][1]

		if lossesCount[winner] == 0 {
			lossesCount[winner] = 0
		}

		lossesCount[loser]++
	}

	for player, losses := range lossesCount {
		if losses == 0 {
			ans[0] = append(ans[0], player)
		} else if losses == 1 {
			ans[1] = append(ans[1], player)
		}
	}

	sort.Ints(ans[0])
	sort.Ints(ans[1])

	return ans
}

// runtime 147ms | beats 7.41% | memory 26.44mb beats 17.81%
func findWinners3(matches [][]int) [][]int {
	seen := make(map[int]struct{})
	lossesCount := make(map[int]int)
	ans := make([][]int, 2)
	ans[0] = []int{}
	ans[1] = []int{}

	for i := 0; i < len(matches); i++ {
		winner := matches[i][0]
		loser := matches[i][1]

		seen[winner] = struct{}{}
		seen[loser] = struct{}{}

		lossesCount[loser]++
	}

	for player := range seen {
		if lossesCount[player] == 0 {
			ans[0] = append(ans[0], player)
		} else if lossesCount[player] == 1 {
			ans[1] = append(ans[1], player)
		}
	}

	sort.Ints(ans[0])
	sort.Ints(ans[1])

	return ans
}

// runtime 83ms | beats 74.07% | memory 19.76mb beats 88.36%
func findWinners2(matches [][]int) [][]int {
	zeroLosses, oneLoss, moreLosses := make(map[int]struct{}), make(map[int]struct{}), make(map[int]struct{})
	ans := make([][]int, 2)
	ans[0] = []int{}
	ans[1] = []int{}

	for i := 0; i < len(matches); i++ {
		winner := matches[i][0]

		_, hasOneLoss := oneLoss[winner]
		_, hasMoreLosses := moreLosses[winner]

		if !hasOneLoss && !hasMoreLosses {
			zeroLosses[winner] = struct{}{}
		}

		loser := matches[i][1]
		_, hasZeroLoss := zeroLosses[loser]
		_, hasOneLoss = oneLoss[loser]
		_, hasMoreLosses = moreLosses[loser]

		if hasZeroLoss {
			delete(zeroLosses, loser)
			oneLoss[loser] = struct{}{}
		} else if hasOneLoss {
			delete(oneLoss, loser)
			moreLosses[loser] = struct{}{}
		} else if hasMoreLosses {
			continue
		} else {
			oneLoss[loser] = struct{}{}
		}
	}

	for key := range zeroLosses {
		ans[0] = append(ans[0], key)
	}

	for key := range oneLoss {
		ans[1] = append(ans[1], key)
	}

	sort.Ints(ans[0])
	sort.Ints(ans[1])

	return ans
}

// runtime 95ms | beats 49.38% | memory 23.38mb beats 46.58%
// my solution
func findWinners1(matches [][]int) [][]int {
	ans := make([][]int, 2)
	wins, losses := make(map[int]int), make(map[int]int)

	for i := 0; i < len(matches); i++ {
		wins[matches[i][0]]++
		losses[matches[i][1]]++
	}

	// Check for winners without losses
	ans[0] = []int{}
	for key, _ := range wins {
		_, exists := losses[key]

		if !exists {
			ans[0] = append(ans[0], key)
		}
	}

	// Check for players with one loss
	ans[1] = []int{}
	for key, value := range losses { 
		if value == 1 {
			ans[1] = append(ans[1], key)
		}
	}

	sort.Ints(ans[0])
	sort.Ints(ans[1])

	return ans
}
