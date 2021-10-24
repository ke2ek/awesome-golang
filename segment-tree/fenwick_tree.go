package tree

/*
Fenwick Tree (a.k.a. Binary Indexed Tree) is used to calculate range sum fastly using partial sum.
It saves memory space than a segment tree.
In the Fenwick Tree, the sum of range [i, j] can be calculated by subtracting psum[i-1] from psum[j].
It means to need only one subtree to calculate the sum in a fenwick tree,
whereas each parent node of a segment tree needs two children to calculate the sum of its own range.

Since the size difference between each node and its subtree is twice,
each index in a fenwick tree can be represented as a binary number so that it is easier to find the lower or upper section.

Ref. https://cp-algorithms.com/data_structures/fenwick.html

Given an array, A, where its index starts with 1:

[A]
| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |

[Fenwick Tree]
|-----------------------------8-|
|-------------4-|
|-----2-|       |-----6-|
|-1-|   |-3-|   |-5-|   |-7-|   |

The binary number of 8 is 1000.
The binary number of 4 is 100.
The binary number of 2 is 10.
The binary number of 1 is 1.

The range-length of 8 is 2^3 = 8 since the number of zero in its binary number is 3 (8->1000).
The range-length of 7 is 2^0 = 1 since the number of zero in its binary number is 0 (7->111).
The range-length of 6 is 2^1 = 2 since the number of zero in its binary number is 1 (6->110).
The range-length of 4 is 2^2 = 4 since the number of zero in its binary number is 2 (4->100).

Notice that the range to be added can be found by subtracting the last 1 bit (LSB) of each binary number.

e.g., psum[7] = tree[7] + tree[6] + tree[4] -> psum[111] = tree[111] + tree[110] + tree[100]

In summary,
1. Each index of tree represents the range [0:index].
2. Each value of tree represents the partial sum of [0:index].
3. a lower or upper section of each value can be easily found by manipulating BIT operations.
*/

type FenwickTree struct {
	tree []int
}

func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{tree: make([]int, n+1)}
}

// O(log2(N))
// Return the partial sum of [0:index].
func (this *FenwickTree) Psum(index int) int {
	index++ // 1-indexed
	ret := 0
	for index > 0 {
		ret += this.tree[index]
		index &= index - 1
	}
	return ret
}

// Update when a new value is added or subtracted at the given position.
func (this *FenwickTree) Add(index, diff int) {
	index++
	for index < len(this.tree) {
		this.tree[index] += diff
		index += index & -index // its proof of concept will be explained below.
	}
}

/*
To find the next section that it will reflect the change, it applies BIT manipulation.
For instance, the upper section of 3 is 4, 8.

binary of 3 = 011
binary of -3 = 101
3 & -3 = 011 & 101 = 001
011(3) + 001 = 100(4)

binary of 4 = 0100
binary of -4 = 1011 + 1 = 1100
4 & -4 = 0100 & 1100 = 0100
0100(4) + 0100 = 1000(8)

Also, the upper section of 7 is 8.

binary of 7 = 0111
binary of -7 = 1001
7 & -7 = 0111 & 1001 = 0001
0111(7) + 0001 = 1000(8)
*/
