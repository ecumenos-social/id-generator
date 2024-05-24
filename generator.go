package idgenerator

type Generator interface {
	Generate() ID
}

type generator struct {
	node *Node
}

func New(node *NodeID) (Generator, error) {
	if err := node.Validate(); err != nil {
		return nil, err
	}

	n, err := NewNode(node.Int64())
	if err != nil {
		return nil, err
	}

	return &generator{node: n}, nil
}

func (g *generator) Generate() ID {
	return g.node.Generate()
}
