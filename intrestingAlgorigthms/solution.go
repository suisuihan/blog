package main

import (
	"fmt"
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

	tests := []struct{
		baseValue int
		monsters []int
		expect int
		want bool
	}{
		{50, []int{50, 105, 200}, 1101, true},
		{20, []int{30, 20, 15, 40, 100}, 205, true},
	}
	for _, test := range tests{
		if got := UpgradeLevel(test.baseValue, test.monsters); (got == test.expect) != test.want{
			fmt.Printf("UpgradeLevel(%d, %v)=%d, except=%d, test failed\n", test.baseValue, test.monsters, got, test.expect)
		}
	}
}