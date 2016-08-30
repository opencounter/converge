// Copyright © 2016 Asteris, LLC
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

package render

import (
	"errors"
	"fmt"

	"github.com/asteris-llc/converge/graph"
	"github.com/asteris-llc/converge/render/extensions"
)

// ErrUnresolvable is returned by Render if the template string tries to resolve
// unaccesible node properties.
var ErrUnresolvable = errors.New("node is unresolvable")

// Renderer to be passed to preparers, which will render strings
type Renderer struct {
	Graph           func() *graph.Graph
	ID              string
	DotValue        string
	DotValuePresent bool
	Language        *extensions.LanguageExtension
}

// Value of this renderer
func (r *Renderer) Value() (value string, present bool) {
	return r.DotValue, r.DotValuePresent
}

// Render a string with text/template
func (r *Renderer) Render(name, src string) (string, error) {
	r.Language = r.Language.On("param", r.param)
	out, err := r.Language.Render(r.DotValue, name, src)
	if err != nil {
		return "", err
	}
	return out.String(), err
}

func (r *Renderer) param(name string) (string, error) {
	name = "param." + name
	fmt.Println("Getting param for: ", name)
	g := r.Graph()
	fmt.Println("Vertices")
	for _, vertex := range g.Vertices() {
		fmt.Printf("\t%s :: %T\n", vertex, g.Get(vertex))
	}
	val := r.Graph().Get(graph.SiblingID(r.ID, name))
	if val == nil {
		return "", errors.New("param not found")
	}
	return fmt.Sprintf("%+v", val), nil
}
