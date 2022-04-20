package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	lines := strings.Split(input, "\n")
	adapters := make([]int, len(lines))
	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		adapters[i] = num
	}
	sort.Ints(adapters)

	ones := 1
	threes := 1
	for i := 0; i < len(adapters)-1; i++ {
		if diff := adapters[i+1] - adapters[i]; diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		}
	}

	fmt.Println(ones * threes)
}

func partTwo() {
	lines := strings.Split(input, "\n")
	adapters := make([]int, len(lines))
	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		adapters[i] = num
	}
	sort.Ints(adapters)

	paths := map[int][]int{}
	for i, adapter := range adapters {
		for j := i + 1; j < len(adapters); j++ {
			if adapters[j] > adapter+3 {
				break
			}
			paths[adapter] = append(paths[adapter], adapters[j])
		}
	}

	options := map[int]int{}
	for i := len(adapters) - 1; i >= 0; i-- {
		adapter := adapters[i]
		values, ok := paths[adapter]
		if !ok {
			options[adapter] = 1
		} else {
			res := 0
			for _, v := range values {
				res += options[v]
			}
			options[adapter] = res
		}
	}

	fmt.Println(options[1] + options[2] + options[3])
}

const input = `115
134
121
184
78
84
77
159
133
90
71
185
152
165
39
64
85
50
20
75
2
120
137
164
101
56
153
63
70
10
72
37
86
27
166
186
154
131
1
122
95
14
119
3
99
172
111
142
26
82
8
31
53
28
139
110
138
175
108
145
58
76
7
23
83
49
132
57
40
48
102
11
105
146
149
66
38
155
109
128
181
43
44
94
4
169
89
96
60
69
9
163
116
45
59
15
178
34
114
17
16
79
91
100
162
125
156
65`
