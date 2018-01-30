## 问题，有两个字符串 a 和 b，b 是否在 a 中？

## 分析：

很容易想到暴力破解，a 从 i 匹配，匹配上后 i++，继续向后匹配，如果匹配不上，则返回 i+1，继续匹配。时间复杂度为O(n2)。

KMP 算法利用的是当前匹配的前缀，当前已匹配的字符串中，如果前缀n个字符和后缀n个字符相同，那么可以从b[n+1]个字符开始比较，这样无需回溯。这个过程中，已匹配的字符串有(len(b)-1)种情况，需要求出每种情况的n，有数组next[n-1]，即next数组。next指的是下一个匹配b的哪一位。最后一位无需比较。例如，对于“abcdabc” ，next = [0,0,0,0,1,2]，当匹配到b[5] = b != a[i]时，下一个要判断的是a[i] 和 b[next[5-1]]的值，即比较a[i]和b[1]。如果 b[0] 和 a[i] 不想等，则比较 b[0] 和 a[i+1]。这部分代码如下:

~~~
a b a a b a d a b c e f g h i
a b a d a b c

a b a a b a d a b c e f g h i
    a b a d a b c

a b a a b a d a b c e f g h i
      a b a d a b c
~~~


~~~golang
...
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
...
~~~


如何求next?

考虑暴力的情况：对于b[i]，从b[1]开始，如果b[0] == b[1]，判断b[1] 是否等于 b[2]...一直到b[i]。如果不想等，那么从判断b[0] == b[2]开始...，一直到b[i]。如果有n个值相等，则 next[i] = n。
~~~golang
la := len(a)
	lb := len(b)
	if la < lb{
		return false
	}
	next := make([]int, lb-1)
	for i:=0; i < lb-1; i++{
		next[i] = 0
	}
	for i:=0; i < lb-1; i++{
		j := 0
		start := j + 1
		k := j + 1
		count := 0
		for k <= i {
			if b[j] == b[k]{
				j++
				k++
				count++
			}else{
				k = start + 1
				start++
				j = 0
				count =0
			}
		}
		if count > 0{
			next[i] += count
		}
	}
~~~

有没有更好的办法呢？ 设 q 为索引，k 为匹配的长度，q 从1 开始。此时已经有k个匹配了，q=q+1，如果此时b[q] == b[k](k为长度，前缀下一个要匹配的所以是k)，则next[q] = k；如果 q= q+1, b[q] 和 b[k]不匹配呢？求b[q] 和 b[k-1] 的是否匹配，一直到k==0
~~~golang
	for q, k :=1, 0; q < lb-1; q++{
		for k > 0 && b[q] != b[k]{
			k--
		}
		if b[q] == b[k]{
			k++
		}
		next[q] = k
	}
~~~

## 代码:

~~~golang
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

	for q, k :=1, 0; q < lb-1; q++{
		for k > 0 && b[q] != b[k]{
			k--
		}
		if b[q] == b[k]{
			k++
		}
		next[q] = k
	}

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

func main(){
	tests := []struct{
		a string
		b string
		want bool
	}{
		{"abcabcdabcefghi", "abcdabc", true},
		{"ddabaad", "aa", true},
		{"dfasdfasdfasdf", "abcabcabdabc", false},
	}
	for _, test := range tests{
		if got := kmp(test.a, test.b); got != test.want{
			fmt.Printf("kmp(%v, %v)=%t, want=%t, test failed\n", test.a, test.b, got, test.want)
		}
	}
}
~~~