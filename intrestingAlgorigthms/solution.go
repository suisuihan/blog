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

func fullCombination(s string)[]string{
	// fmt.Println(s)
	n := len(s)
	if n <= 1{
		return []string{s}
	}
	chars := make(map[rune]bool)
	for _, v := range s{
		chars[v] = true
	}
	var new_s string
	for i:='A'; i <= 'Z'; i++{
		if chars[i]{
			new_s += string(i)
		}
	}
	
	var combine func(string)[]string
	combine = func(ss string)[]string{
		n := len(ss)
		if n <= 1{
			return []string{ss}
		}else if n==2 {
			s1 := ss
			s2 := string(ss[1])+string(ss[0])
			return []string{s1, s2}
		}else{
			var ret []string
			for i:=0; i<n; i++{
				var t []string
				if i < n-1{
					t = combine(ss[0:i] +ss[i+1:])
				}else{
					t = combine(ss[0:i])
				}
				for j:=0; j < len(t); j++{
					ret = append(ret, string(ss[i])+t[j])
				}
			}
			return ret
		}
	}
	return combine(new_s)
}

func LongestComSubstr(s1 string, s2 string)[]string{
	var subStr []string
	n1, n2 := len(s1), len(s2)
	if n1 <=1 || n2 <=1{
		return subStr
	}
	for i := 0; i < n1-1; i++{
		for j := 0; j < n2-1; j++{
			p1, p2 := i, j
			for p1 < n1 && p2 < n2 && s1[p1] == s2[p2] {
				p1++
				p2++
			}
			if p1 > i{
				subStr = append(subStr, s1[i:p1])
			}
		}
		fmt.Println(i)
	}
	var ret []string
	m := len(subStr)
	if m == 0{
		return ret
	}else{
		ret = append(ret, subStr[0])	
		for i:=1; i < len(subStr); i++{
			if len(subStr[i]) > len(ret[0]){
				ret = ret[0:0]
				ret = append(ret, subStr[i])
			}else if len(subStr[i]) == len(subStr[0]){
				ret = append(ret, subStr[i])
			}
		}
	}
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

// func fullCombine(s string)[]string{
// 	charaters := make([]rune, )
// }

func kmp(a string, b string) bool{
	la := len(a)
	lb := len(b)
	if la < lb{
		return false
	}
	next := make([]int, lb-1)
	for i:=0; i < lb-1; i++{
		next[i] = 0
	}
	// for i:=0; i < lb-1; i++{
	// 	j := 0
	// 	start := j + 1
	// 	k := j + 1
	// 	count := 0
	// 	for k <= i {
	// 		if b[j] == b[k]{
	// 			j++
	// 			k++
	// 			count++
	// 		}else{
	// 			k = start + 1
	// 			start++
	// 			j = 0
	// 			count =0
	// 		}
	// 	}
	// 	if count > 0{
	// 		next[i] += count
	// 	}
	// }

	fmt.Println(b)
	for q, k :=1, 0; q < lb-1; q++{
		for k > 0 && b[q] != b[k]{
			k--
		}
		if b[q] == b[k]{
			k++
		}
		next[q] = k
		fmt.Println(q, k, next)
	}

	fmt.Println(next)

	j := 0
	for i:=0; i < la && j < lb;{
		if a[i] == b[j]{
			i++
			j++
		}else{
			if j == 0{
				i++
			}else{
				j = next[j-1]
			}
		}
	}
	return j == lb
}

func shortestSummary(s string, t string) string{
	sLength, tLength := len(s), len(t)
	if sLength < tLength{
		return s
	}
	start, end := 0, -1
	var absStart, absEnd int
	shortestLen := sLength
	mapT := make(map[uint8]int)
	for i:=0; i < tLength; i++{
		mapT[t[i]]++
	}
	isAllExisted := func(map1 map[uint8]int, map2 map[uint8]int) bool{ 
		for k, v := range map2{
			if map1[k] < v{
				return false
			}
		}
		return true
	}
	mapRet := make(map[uint8]int)
	for {
		for end < sLength && !isAllExisted(mapRet, mapT){
			end++
			for end < sLength{
				if _, ok := mapT[s[end]]; ok{
					mapRet[s[end]]++
					break
				}else{
					end++
				}
			}			
		}
		
		for isAllExisted(mapRet, mapT){
			if end - start < shortestLen{
				shortestLen = end - start
				absStart = start
				absEnd = end
			}
			if _, ok := mapRet[s[start]]; ok{
				mapRet[s[start]]--
			}
			start++
		}
		if end >= sLength{
			break
		}
	}
	return s[absStart:absEnd+1]
}

func Monopoly(n int) int{
	var ret int

	return ret
}

func largestRectangleArea(heights []int)int{
    length := len(heights)
    if length == 0{
        return 0
    }else if length == 1{
        return heights[0]
    }
    maxArea := 0
    for i:=0; i < length; i++{
        l, r :=1, 0
        for j:= i-1; j >= 0; j--{
            if heights[j] < heights[i]{
                break
            }else{
                l++
            }
        }
        
        for j:= i+1; j < length; j++{
            if heights[j] < heights[i]{
                break
            }else{
                r++
            }
        }

        curr := heights[i] * (l + r)
        if curr > maxArea{
            maxArea = curr
        }
    }
    return maxArea
}

func largestRectangleArea1(heights []int)int{
	length := len(heights)
    if length == 0{
        return 0
    }else if length == 1{
        return heights[0]
	}
	heights = append(heights, 0)
	maxArea := 0
	stack := []int{0}
    for i:=1; i < length+1; i++{
		for len(stack) >0 && heights[i] < heights[stack[len(stack)-1]]{
			h := heights[stack[len(stack)-1]]
			var w int
			stack = stack[0:len(stack)-1]
			if len(stack) == 0{
				w = i
			}else{
				w = i - 1 - stack[len(stack)-1]
			}
			if maxArea < w*h{
				maxArea = w*h
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}

type Node struct{
	Value int
	Left *Node
	Right *Node
}

func getPath(head *Node, target int)[]int{
	if head == nil{
		return nil
	}
	if head.Value == target{
		return []int{head.Value}
	}

	var heap []int
	var list []*Node
	list = append(list, head)
	heap = append(heap, head.Value)
	allListNil := true
	for len(list) > 0{
		temp := make([]*Node, len(list))
		copy(temp, list)
		fmt.Print("list is [ ")
		for _, v := range list{
			if v == nil{
				fmt.Print("nil ")
			}else{
				fmt.Print(v.Value, " ")
			}
		}
		fmt.Println("]")
		fmt.Println("heap is ", heap)
		list = list[0:0]
		for len(temp) > 0{
			t := temp[0]
			if t == nil{
				heap = append(heap, 0, 0)
				list = append(list, nil, nil)
				if len(temp) == 1{
					temp = temp[0:0]
				}else{
					temp = temp[1:len(temp)]
				}
				continue
			}
			if t.Left != nil{
				heap = append(heap, t.Left.Value)
				list = append(list, t.Left)
				if t.Left.Value == target{
					list = list[0:0]
					break
				}
				allListNil = false
			}else if t.Left == nil{
				heap = append(heap, 0)
				list = append(list, nil)
			}
			
			if t.Right != nil{
				heap = append(heap, t.Right.Value)
				list = append(list, t.Right)
				if t.Right.Value == target{
					list = list[0:0]
					break
				}
				allListNil = false
			}else if t.Right == nil{
				heap = append(heap, 0)
				list = append(list, nil)
			}
			
			if len(temp) == 1{
				temp = temp[0:0]
			}else{
				temp = temp[1:len(temp)]
			}
		}
		if !allListNil{
			allListNil = true
			continue
		}else{
			break
		}
	}
	if heap[len(heap)-1] != target{
		return nil
	}
	var ret []int
	lenHeap := len(heap)
	index := lenHeap-1
	for index > 0{
		ret = append(ret, heap[index])
		index = (index-1)/2
	}
	ret = append(ret, heap[0])
	for start, end := 0, len(ret)-1; start < end; {
		ret[start], ret[end] = ret[end], ret[start]
		start++
		end--
	}
	return ret
}

func getPath1(head *Node, target int)[]int{
	var traverse []int
	var preTreaverse func(h *Node, t int)
	preTreaverse = func(h *Node, t int){
		if h == nil{
			return
		}
		if h.Value == t{
			traverse = append(traverse, t)
			return
		}
		var stack []*Node
		stack = append(stack, h)
		traverse = append(traverse, h.Value)
		h = h.Left
		for !(len(stack) == 0 && h == nil){
			for h != nil{
				traverse = append(traverse, h.Value)
				if h.Value == t{
					return
				}
				stack = append(stack, h)
				h = h.Left
			}
			h = stack[len(stack)-1].Right
			stack = stack[0:len(stack)-1]
		}
	}
	if head != nil{
		preTreaverse(head, target)
	}
	fmt.Println(traverse)
	if traverse[len(traverse)-1] != target{
		return []int{}
	}
	var ret []int
	ret = append(ret, head.Value)
	for head.Value != target{
		fmt.Println(head.Value)
		if head.Left == nil && head.Right == nil{
			return []int{}
		}
		if head.Left == nil{
			ret = append(ret, head.Right.Value)
			head = head.Right
		}else if head.Right == nil{
			ret = append(ret, head.Left.Value)
			head = head.Left
		}else{
			isRight := false
			for i:=0; i < len(traverse); i++{
				if traverse[i] == head.Right.Value{
					ret = append(ret, head.Right.Value)
					head = head.Right
					isRight = true
					if i < len(traverse) - 2{
						traverse = traverse[i+1:]
					}else{
						traverse = traverse[i+1:]
					}
					break
				}
			}
			if !isRight{
				ret = append(ret, head.Left.Value)
				head = head.Left
			}
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

	// tests := []struct{
	// 	baseValue int
	// 	monsters []int
	// 	expect int
	// 	want bool
	// }{
	// 	{50, []int{50, 105, 200}, 110, true},
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

	// tests := []struct{
	// 	stationBeauties []int
	// 	reachStations [][]int
	// 	expect int
	// 	want bool
	// }{
	// 	{[]int{0,1,1,3,6}, [][]int{[]int{0}, []int{0}, []int{1}, []int{0}, []int{2,3}}, 8, true},
	// }
	// for _, test := range tests{
	// 	if got := leastBeauties(test.stationBeauties, test.reachStations); (got == test.expect) != test.want{
	// 		fmt.Printf("leastBeauties(%v, %v)=%d, except=%d, test failed\n", test.stationBeauties, test.reachStations, got, test.expect)
	// 	}else{
	// 		fmt.Println(got)
	// 	}
	// }

	// fmt.Println(FullCharaters("abcqweracb"))
	
	// tests := []struct{
	// 	a string
	// 	b string
	// 	want bool
	// }{
	// 	{"abcabcdabcefghi", "abcdabc", true},
	// 	{"ddabaad", "aa", true},
	// 	{"dfasdfasdfasdf", "abcabcabdabc", false},
	// }
	// for _, test := range tests{
	// 	if got := kmp(test.a, test.b); got != test.want{
	// 		fmt.Printf("kmp(%v, %v)=%t, want=%t, test failed\n", test.a, test.b, got, test.want)
	// 	}
	// }

	// a := make(map[string]int)
	// a["aaa"] = 1
	// a["bbb"] = 2
	// b := a
	// fmt.Println(b)
	// b["ccc"] = 3
	// b["aaa"] = 4
	// fmt.Println(b)
	// fmt.Println(a)
	// s := "aaaa"
	// fmt.Printf("%T\n", s[0])
	// for _, v:= range s{
	// 	fmt.Printf("%T\n", v)
	// }

	
	// tests := []struct{
	// 	a string
	// 	b string
	// 	except string
	// 	want bool
	// }{
	// 	{"ADOBECODEBANC", "ABC", "BANC", true},
	// }
	// for _, test := range tests{
	// 	if got := shortestSummary(test.a, test.b); (got == test.except) != test.want{
	// 		fmt.Printf("shortestSummary(\"%v\",\"%v\")=%s, want=%t, test failed\n", test.a, test.b, got, test.want)
	// 	}
	// }
	
	// tests := []struct{
	// 	input []int
	// 	except int
	// 	want bool
	// }{
	// 	{[]int{2, 1, 5, 6, 2, 3}, 10, true},
	// }
	// for _, test := range tests{
	// 	if got := largestRectangleArea1(test.input); (got == test.except) != test.want{
	// 		fmt.Printf("largestRectangleArea1(\"%v\")=%d, want=%t, test failed\n", test.input, got, test.want)
	// 	}
	// }

	node1 := &Node{6, nil, nil}
	node2 := &Node{5, node1, nil}
	node3 := &Node{9, nil, nil}
	node4 := &Node{4, nil, node2}
	node5 := &Node{7, node3, nil}
	node6 := &Node{8, nil, nil}
	node7 := &Node{2, node4, nil}
	node8 := &Node{3, node5, node6}
	node9 := &Node{1, node7, node8}
	
	fmt.Println(getPath1(node9, 9))

}