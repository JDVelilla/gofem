// Copyright 2016 The Gofem Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/cpmech/gofem/out"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/plt"
)

func main() {

	// filename
	filename, fnkey := io.ArgToFilename(0, "d2-simple-flux", ".sim", true)

	// start analysis process
	out.Extrap = []string{"nwlx", "nwly"}
	out.Start(filename, 0, 0)

	// define entities
	out.Define("top-middle", out.At{5, 3})
	out.Define("section-A", out.N{-1})
	out.Define("section-B", out.Along{{0, 0}, {10, 0}})

	// load results
	out.LoadResults(nil)

	// compute water discharge along section-A
	nwlx_TM := out.GetRes("nwlx", "top-middle", -1)
	Q := out.Integrate("nwlx", "section-A", "y", -1)
	io.PfYel("Q = %g m³/s [answer: (0.0002999972723204199) 0.0003]\n", Q)

	// plot
	kt := len(out.Times) - 1
	out.Splot("pl-y", "")
	out.Plot("pl", "y", "section-A", &plt.A{L: "t=0"}, 0)
	out.Plot("pl", "y", "section-A", &plt.A{L: io.Sf("t=%g", out.Times[kt])}, kt)
	out.Splot("x-pl", "")
	out.Plot("x", "pl", "section-B", &plt.A{L: "t=0"}, 0)
	out.Plot("x", "pl", "section-B", &plt.A{L: io.Sf("t=%g", out.Times[kt])}, kt)
	out.Splot("nwlx", "")
	out.Plot("t", nwlx_TM, "top-middle", &plt.A{}, -1)
	out.Csplot.Ylbl = "$n_{\\ell}\\cdot w_{\\ell x}$"

	// save
	out.Draw("/tmp", "seep_simple_flux_"+fnkey+".png", -1, -1, false, func(id string) {
		if id == "x-pl" {
			plt.Plot([]float64{0, 10}, []float64{10, 9}, &plt.A{C: "k", Ls: "--"})
		}
	})
}
