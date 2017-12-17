package draw

import (
	"github.com/YiYeYu/LazySin/tool"
)

//TextureManager is 2d texture manager
type TextureManager struct {
}

//LoadTexture load texture from path
func LoadTexture(path string) (*Texture, error) {
	img, e := tool.LoadImage(path)
	if e != nil {
		return nil, e
	}
	_ = img
	return nil, nil
}
