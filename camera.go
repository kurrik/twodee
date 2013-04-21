// Copyright 2013 Arne Roomann-Kurrik
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package twodee

import (
	"github.com/go-gl/gl"
)

type Camera struct {
	view   Rectangle
	focus  Point
	width  float64
	height float64
	zoom   float64
}

func NewCamera(x float64, y float64, w float64, h float64) (c *Camera) {
	c = &Camera{
		width:  w,
		height: h,
		focus:  Pt(x+w/2.0, y+h/2.0),
		zoom:   0,
	}
	c.calcView()
	return
}

func (c *Camera) calcView() {
	var (
		ratio = c.height / c.width
		hw = c.width / 2.0
		hh = hw * ratio
		zw = hw * c.zoom
		zh = zw * ratio
	)
	c.view.Min.X = c.focus.X - hw - zw
	c.view.Min.Y = c.focus.Y - hh - zh
	c.view.Max.X = c.focus.X + hw + zw
	c.view.Max.Y = c.focus.Y + hh + zh
}

func (c *Camera) MatchRatio(width int, height int) {
	ratio := float64(height) / float64(width)
	c.height = c.width * ratio
	c.calcView()
}

func (c *Camera) Top(y float64) {
	var (
		dy = y - c.view.Min.Y
	)
	c.focus.Y += dy
	c.calcView()
}

func (c *Camera) Pan(x float64, y float64) {
	c.focus.X += x
	c.focus.Y += y
	c.calcView()
}

func (c *Camera) Zoom(z float64) {
	c.zoom = z
	c.calcView()
}

func (c *Camera) Bounds() Rectangle {
	return c.view
}

func (c *Camera) SetProjection() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(c.view.Min.X, c.view.Max.X, c.view.Max.Y, c.view.Min.Y, -1, 1)
}
