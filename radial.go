package main

// RadialAction defines a single choice on the radial menu
type RadialAction struct {
	Name    string
	Execute func() // The logic to run when this circle is left-clicked
}

// Interactable is the semantic interface for anything the player can click on
type Interactable interface {
	GetAvailableActions(g *Game, targetX, targetY int) []RadialAction
}
