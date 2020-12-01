//-----------------------------------------------------------------------------
/*

Text Example

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"

	. "github.com/deadsy/sdfx/sdf"
)

//-----------------------------------------------------------------------------

func main() {

	f, err := LoadFont("cmr10.ttf")
	//f, err := LoadFont("Times_New_Roman.ttf")
	//f, err := LoadFont("wt064.ttf")

	if err != nil {
		log.Fatalf("can't read font file %s", err)
	}

	t := NewText("SDFX!\nHello,\nWorld!")
	//t := NewText("相同的不同")

	s2d, err := TextSDF2(f, t, 10.0)
	if err != nil {
		log.Fatalf("can't generate text sdf2 %s", err)
	}

	err = RenderDXF(s2d, 600, "shape.dxf")
	if err != nil {
		log.Fatalf("dxf render error %s", err)
	}
	err = RenderSVG(s2d, 600, "shape.svg", "fill:none;stroke:black;stroke-width:0.1")
	if err != nil {
		log.Fatalf("svg render error %s", err)
	}
	s3d := ExtrudeRounded3D(s2d, 1.0, 0.2)
	err = RenderSTL(s3d, 600, "shape.stl")
	if err != nil {
		log.Fatalf("stl render error %s", err)
	}
}

//-----------------------------------------------------------------------------
