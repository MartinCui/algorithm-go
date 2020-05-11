package linklistloopdetection

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	rq := require.New(t)
	rq.Nil(nil, findLoopEntrance(nil))
	rq.Nil(nil, findLoopEntrance(&node{key: 0}))
	rq.Nil(nil, findLoopEntrance(&node{key: 0, next: &node{key: 1}}))
	rq.Nil(nil, findLoopEntrance(&node{key: 0, next: &node{key: 1, next: &node{key: 2, next: &node{key: 3}}}}))
	rq.Equal(2, findLoopEntrance(createLoop(2, 4)).key)
	rq.Equal(2, findLoopEntrance(createLoop(2, 15)).key)
	rq.Equal(15, findLoopEntrance(createLoop(15, 2)).key)
	rq.Equal(159, findLoopEntrance(createLoop(159, 11)).key)
	rq.Equal(11, findLoopEntrance(createLoop(11, 159)).key)
}

func TestLoopCreation(t *testing.T) {
	m := 3
	n := 5
	loop := createLoop(m, n)

	nodeInLoop := loop
	var sb strings.Builder
	var loopEntrance *node
	for i := 0; i < m+n-1; i++ {
		sb.WriteString(fmt.Sprintf("%d -> ", nodeInLoop.key))
		if i == m {
			loopEntrance = nodeInLoop
		}
		nodeInLoop = nodeInLoop.next
	}
	lastNode := nodeInLoop
	sb.WriteString(fmt.Sprintf("%d", lastNode.key))
	t.Log(sb.String())
	sb.Reset()

	rq := require.New(t)
	rq.Equal(lastNode.next, loopEntrance)
	for i := 0; i < m; i++ {
		sb.WriteString("     ")
	}
	sb.WriteString("ˆ")
	for i := 0; i < n-2; i++ {
		sb.WriteString("_____")
	}
	sb.WriteString("____|")

	t.Log(sb.String())
}

func createLoop(beforeLoopSize int, loopSize int) *node {
	loopEntrance := &node{key: beforeLoopSize}
	nodeInLoop := loopEntrance
	for i := 1; i < loopSize-1; i++ {
		nodeInLoop.next = &node{key: beforeLoopSize + i}
		nodeInLoop = nodeInLoop.next
	}
	nodeInLoop.next = &node{key: beforeLoopSize + loopSize - 1, next: loopEntrance}

	nodeInLoop = loopEntrance
	for i := beforeLoopSize - 1; i >= 0; i-- {
		nodeInLoop = &node{key: i, next: nodeInLoop}
	}

	return nodeInLoop
}
