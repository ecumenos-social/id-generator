package idgenerator

import (
	"fmt"
)

var (
	TopNodeBits uint8 = 7
	LowNodeBits uint8 = 10

	topNodeMax   int64 = -1 ^ (-1 << TopNodeBits)
	topNodeMask        = topNodeMax << LowNodeBits
	lowNodeMax   int64 = -1 ^ (-1 << LowNodeBits)
	lowNodeMask        = lowNodeMax
	topNodeShift       = LowNodeBits
)

type NodeID struct {
	Top int64
	Low int64
}

func (n *NodeID) Validate() error {
	if n.Top < 0 || n.Top > topNodeMax {
		return fmt.Errorf("invalid top node id (requirements: 0 < top_node_id <= %d, actual: %d)", topNodeMax, n.Top)
	}
	if n.Low < 0 || n.Low > lowNodeMax {
		return fmt.Errorf("invalid low node id (requirements: 0 < low_node_id <= %d, actual: %d)", lowNodeMax, n.Low)
	}

	return nil
}

func (n *NodeID) Int64() int64 {
	return (n.Top)<<topNodeShift | n.Low
}

func ParseInt64ToNodeID(v int64) *NodeID {
	return &NodeID{
		Top: v & topNodeMask >> topNodeShift,
		Low: v & lowNodeMask,
	}
}
