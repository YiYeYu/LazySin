package draw

import (
	"github.com/go-gl/gl/v2.1/gl"
)

//Render is a render
type Render struct {
}

//Flush call gl flush
func (r *Render) Flush() {
	gl.Flush()
}

//SetClearColor change gl clear color
func (r *Render) SetClearColor(c Color) {
	gl.ClearColorIuiEXT(uint32(c.R), uint32(c.G), uint32(c.B), uint32(c.A))
}

//SetColor change gl color
func (r *Render) SetColor(c Color) {
	gl.Color4ub(c.R, c.G, c.B, c.A)
}

//PutVertex draw gl vertex
func (r *Render) PutVertex(v Vertex) {
	gl.Vertex3i(v.X, v.Y, v.Z)
}

//PutPairCV draw PairCV
func (r *Render) PutPairCV(p PairCV) {
	gl.Color4ub(p.R, p.G, p.B, p.A)
	gl.Vertex3i(p.X, p.Y, p.Z)
}

//PutVertexs draw gl vertex
func (r *Render) PutVertexs(vs []Vertex) {
	for _, v := range vs {
		gl.Vertex3i(v.X, v.Y, v.Z)
	}
}

//DrawSprite draw a sprite
func (r *Render) DrawSprite(s Sprite) {
	r.drawSprite(s)
}

//DrawSprites draw a sprite array
func (r *Render) DrawSprites(ss []Sprite) {
	for _, s := range ss {
		r.drawSprite(s)
	}
}

//

func (r *Render) drawSprite(s Sprite) {
	for _, v := range s.Vertexs {
		r.PutPairCV(v)
	}
}
