package draw

//PairCV contains Color & Vertex
type PairCV struct {
	Color
	Vertex
}

//Sprite is a sprite
type Sprite struct {
	Vertexs []PairCV
}
