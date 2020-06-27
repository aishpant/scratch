package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSum2(t *testing.T) {

	root := &TreeNode{Val: 5,
		Left: &TreeNode{Val: 4,
			Left: &TreeNode{Val: 11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{Val: 8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{Val: 4,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1},
			},
		},
	}

	actual := pathSum(root, 22)
	assert.Equal(t, [][]int{[]int{5, 4, 11, 2}, []int{5, 8, 4, 5}}, actual)
}
