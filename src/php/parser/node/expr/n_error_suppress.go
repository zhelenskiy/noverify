package expr

import (
	"github.com/VKCOM/noverify/src/php/parser/freefloating"
	"github.com/VKCOM/noverify/src/php/parser/node"
	"github.com/VKCOM/noverify/src/php/parser/position"
	"github.com/VKCOM/noverify/src/php/parser/walker"
)

// ErrorSuppress node
type ErrorSuppress struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Expr         node.Node
}

// NewErrorSuppress node constructor
func NewErrorSuppress(Expression node.Node) *ErrorSuppress {
	return &ErrorSuppress{
		FreeFloating: nil,
		Expr:         Expression,
	}
}

// SetPosition sets node position
func (n *ErrorSuppress) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ErrorSuppress) GetPosition() *position.Position {
	return n.Position
}

func (n *ErrorSuppress) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ErrorSuppress) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.Expr != nil {
		n.Expr.Walk(v)
	}

	v.LeaveNode(n)
}
