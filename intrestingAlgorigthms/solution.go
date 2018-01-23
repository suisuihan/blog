package main

import (
	"fmt"
	"sort"
)


func Palindrome(s string)bool{
	n := len(s)
	if n <= 2{
	    return true		
	}
	var palindrome func (string, int) bool
	palindrome = func(s string, count int)bool{
		n := len(s)
		if n <= 2{
	    	return true		
		}
		start, end := 0, n-1
		for start < end && count <= 1{
			if s[start] == s[end]{
				start++
				end--
			}else{
				if count >= 1{
					return false
				}
				count++
				return palindrome(s[start:end], count) || palindrome(s[start+1:end+1], count)
			}
		}
		return true
	}
	return palindrome(s, 0)
}

func UpgradeLevel(baseValue int, monsters []int)int{
	n := len(monsters)
	if n == 0{
		return baseValue
	}

	var gcd func(int, int)int
	gcd = func(a int, b int)int{
		if b == 0{
			return a
		}else{
			return gcd(b, a%b)
		}
	}
	ret := baseValue
	for i:=0; i<n; i++{
		if ret >= monsters[i]{
			ret += monsters[i]
		}else{
			ret += gcd(ret, monsters[i])
		}
	}
	return ret
}

func LongestSubList(nums []int)int{
	n := len(nums)
	if n <=1 {
		return n
	}
	sortedNums := make([]int, n)
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	//此时问题转换为求nums和sortedNums最长公共子序列问题
	common :=make([][]int, n+1)
	for i:=0; i<n+1; i++{
		common[i] = make([]int, n+1)
	}
	for i:=0; i<n+1; i++{
		common[0][i] = 0
		common[i][0] = 0
	}
	for i:=1; i<n+1;i++{
		for j:=1; j<n+1; j++{		
			if nums[i-1] == sortedNums[j-1]{
				common[i][j] = common[i-1][j-1] + 1
			}else{
				common[i][j] = max(common[i-1][j], common[i][j-1])
			}
		}
	}
	// fmt.Println(common)
	return common[n][n]
}

func max(a int, b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}


func LongestSubList1(nums []int)int{
	n := len(nums)
	if n <=1 {
		return n
	}
	common := make([]int, n)
	common[0] = 1
	for i:=1; i<n; i++{
		max := 0
		for j:=i-1; j>=0; j--{
			if nums[j] <= nums[i] && max < common[j]{
				max = common[j]
			}
		}
		common[i] = max + 1
	}
	ret := common[0]
	for i:=1; i<n; i++{
		if common[i] > ret{
			ret = common[i]
		}
	}
	fmt.Println(common)
	return ret
}


func FullCharaters(s string)string{
    used := make(map[rune]bool)
    for i, j:='A', 'a'; i <= 'Z' && j <= 'z';{
        used[i] = false
        used[j] = false
        i++
        j++
    }
    var ret string
    for _, c := range s{
        if !used[c]{
			ret += string(c)
			used[c] = true
        }
    }
    return ret
}

func leastBeauties(stationBeauties []int, reachStations [][]int)int{
    n := len(stationBeauties)
    if n == 0{
        return 0 
    }else if n == 1{
        return stationBeauties[0]
    }else if n == 2{
        return stationBeauties[0] + stationBeauties[1]
    }

    beauties := make([]int, n)
    beauties[0] = stationBeauties[0]
    for i:=1; i < n; i++{
        minBeauties := beauties[reachStations[i][0]]
        for j:=1; j < len(reachStations[i]); j++{
            if minBeauties > beauties[reachStations[i][j]]{
                minBeauties = beauties[reachStations[i][j]]
            }
        }
        beauties[i] = stationBeauties[i]+minBeauties
    }
    return beauties[n-1]
}

func fullCombine(s string)[]string{
	charaters := make([]rune, )
}

func main(){
	// tests := []struct{
	// 	input string
	// 	expect bool
	// 	want bool
	// }{
	// 	{"coco", true, true},
	// 	{"coc", true, true},
	// 	{"cocod", false, true},
	// 	{"cocodd", false, true},
	// }
	// for _, test := range tests{
	// 	if got := Palindrome(test.input); (got == test.expect) != test.want{
	// 		fmt.Printf("Palindrome(%s)=%t, except=%t, test failed\n", test.input, got, test.expect)
	// 	}
	// }

	// tests := []struct{
	// 	baseValue int
	// 	monsters []int
	// 	expect int
	// 	want bool
	// }{
	// 	{50, []int{50, 105, 200}, 1101, true},
	// 	{20, []int{30, 20, 15, 40, 100}, 205, true},
	// }
	// for _, test := range tests{
	// 	if got := UpgradeLevel(test.baseValue, test.monsters); (got == test.expect) != test.want{
	// 		fmt.Printf("UpgradeLevel(%d, %v)=%d, except=%d, test failed\n", test.baseValue, test.monsters, got, test.expect)
	// 	}
	// }

	// tests := []struct{
	// 	input []int
	// 	expect int
	// 	want bool
	// }{
	// 	{[]int{89, 256, 78, 1, 46, 78, 8}, 3, true},
	// 	{[]int{6, 4, 8, 2, 17}, 3, true},
	// }
	// for _, test := range tests{
	// 	if got := LongestSubList1(test.input); (got == test.expect) != test.want{
	// 		fmt.Printf("LongestSubList1(%v)=%d, except=%d, test failed\n", test.input, got, test.expect)
	// 	}
	// }

	tests := []struct{
		stationBeauties []int
		reachStations [][]int
		expect int
		want bool
	}{
		{[]int{0,1,1,3,6}, [][]int{[]int{0}, []int{0}, []int{1}, []int{0}, []int{2,3}}, 8, true},
	}
	for _, test := range tests{
		if got := leastBeauties(test.stationBeauties, test.reachStations); (got == test.expect) != test.want{
			fmt.Printf("leastBeauties(%v, %v)=%d, except=%d, test failed\n", test.stationBeauties, test.reachStations, got, test.expect)
		}else{
			fmt.Println(got)
		}
	}

	// fmt.Println(FullCharaters("abcqweracb"))
}