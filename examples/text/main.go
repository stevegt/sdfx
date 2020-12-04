//-----------------------------------------------------------------------------
/*

Text Example

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"

	"github.com/deadsy/sdfx/sdf"
)

//-----------------------------------------------------------------------------

func main() {

	f, err := sdf.LoadFont("cmr10.ttf")
	//f, err := sdf.LoadFont("Times_New_Roman.ttf")
	//f, err := sdf.LoadFont("wt064.ttf")

	if err != nil {
		log.Fatalf("can't read font file %s", err)
	}

	t := sdf.NewText("SDFX!\nHello,\nWorld!")
	//t := sdf.NewText("相同的不同")

	s2d, err := sdf.TextSDF2(f, t, 10.0)
	if err != nil {
		log.Fatalf("can't generate text sdf2 %s", err)
	}

	sdf.RenderDXF(s2d, 600, "shape.dxf")
	sdf.RenderSVG(s2d, 600, "shape.svg", "fill:none;stroke:black;stroke-width:0.1")

	s3d, err := sdf.ExtrudeRounded3D(s2d, 1.0, 0.2)
	if err != nil {
		log.Fatal(err)
	}
	sdf.RenderSTL(s3d, 600, "shape.stl")
}

//-----------------------------------------------------------------------------
