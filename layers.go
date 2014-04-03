// Copyright 2014 Arne Roomann-Kurrik
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

type Layer interface {
	Render()
	Update()
	Delete()
	HandleMouseEvent(evt *MouseEvent) bool
}

type Layers struct {
	layers []Layer
}

func NewLayers() *Layers {
	return &Layers{}
}

func (l *Layers) Push(layer Layer) {
	l.layers = append(l.layers, layer)
}

func (l *Layers) Pop() (layer Layer) {
	var (
		index = len(l.layers) - 1
	)
	layer = l.layers[index]
	l.layers = l.layers[:index]
	return
}

func (l *Layers) Render() {
	for _, layer := range l.layers {
		layer.Render()
	}
}

func (l *Layers) Update() {
	for _, layer := range l.layers {
		layer.Update()
	}
}

func (l *Layers) Delete() {
	for _, layer := range l.layers {
		layer.Delete()
	}
}

func (l *Layers) HandleMouseEvent(evt *MouseEvent) bool {
	for _, layer := range l.layers {
		if layer.HandleMouseEvent(evt) == false {
			return false
		}
	}
	return true
}