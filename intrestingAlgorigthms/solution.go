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

func main(){
	tests := []struct{
		input string
		expect bool
		want bool
	}{
		{"coco", true, true},
		{"coc", true, true},
		{"cocod", false, true},
		{"cocodd", false, true},
	}
	for _, test := range tests{
		if got := Palindrome(test.input); (got == test.expect) != test.want{
			fmt.Printf("Palindrome(%s)=%t, except=%t, test failed\n", test.input, got, test.expect)
		}
	}
}