/*
 * Referensi :
 * 	- Convert Decimal -> Hex :
 * 		  https://www.rapidtables.com/convert/number/decimal-to-hex.html
 *
 *	- This program (Lissajous generates GIF animations of random Lissajous figures.) :
 *		  https://go.dev/play/p/o2o2lDR3Eg
 */

/*
 * Nama  : Dimas Syauqi Syafa
 * NIM   : 2502004405
 * Kelas : LA20
 */

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"log"
	"net/http"
	"time"
)

const (
	whiteIndex = 0
	blackIndex = 1
)

/*
 * pada parameter dibawah ini, maksud dari code (color.RGBA{0xff, 0xff, 0x00, 0xff})
 * adalah warna kuning, parameter dari color.RGBA sendiri adalah hasil convert
 * dari angka desimal -> hex, lalu masukkan ke parameter
 *
 * sebagai contoh : 0xff jika di convert ke decimal menjadi 255
 * rgb dari kuning sendiri adalah (255, 255, 0)
 *
 * tapi mengapa ada 4 paramter? karena itu RGBA parameter terakhir untuk menentukan
 * transparasi warna, tetapi saya masukkan 255 ke parameter dalam bentuk hex
 * karena saya ingin warna saya tidak transparant.
 */
var palette = []color.Color{color.Black, color.RGBA{0xff, 0xff, 0x00, 0xff}}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {

		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}

		http.HandleFunc("/", handler)

		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 3
		res     = 0.0015
		size    = 640
		nframes = 128
		delay   = 5
	)

	freq  := rand.Float64() * 3.0
	anim  := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img  := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
