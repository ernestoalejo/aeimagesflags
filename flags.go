package aeimagesflags

import (
	"fmt"
	"image/color"
	"strings"
)

type rotation uint64

const (
	Rotate90  = rotation(90)
	Rotate180 = rotation(180)
	Rotate270 = rotation(270)
)

type Flags struct {
	// Size of the largest dimension
	Size uint64

	// Serve the original image
	Original bool

	Width  uint64
	Height uint64

	// Smart square crop that attempts cropping to faces
	SmartSquareCropFaces bool

	// Alternative smart square crop
	SmartSquareCrop bool

	CircularCrop bool

	SquareCrop bool

	// Crop to the smallest of: Width, Height or Size
	SmallestCrop bool

	// Flip the image on its axis
	VerticalFlip, HorizontalFlip bool

	Rotate rotation

	// Force specific image formats.
	// Forcing PNG, WebP and GIF outputs can work in combination with circular
	// crops for a transparant background. Forcing JPG can be combined with border
	// color to fill in backgrounds in transparent images.
	ForceJPEG, ForcePNG bool
	ForceWebP, ForceGIF bool

	// Generate MP4 from input animated GIF
	MP4 bool

	// Remove any animation of the GIF file
	KillAnimation bool

	// Width of the border, in pixels
	Border uint64

	// Color of the border, black by default (transparent defaults to 255)
	BorderColor color.RGBA

	// Add header to force browser download of the image
	ForceDownload bool

	// Set Cache-Control max-age directive to these number of days
	ExpiresDays uint64

	// Change the quality of the JPEG output
	JPEGQuality uint64

	// Intensity of the blur (max 100)
	Blur uint64

	// Size of the border gradient
	BorderGradientSize uint64

	// Color of the border gradient
	BorderGradientColor color.RGBA
}

// Apply takes the serving URL (generated using appengine/images.ServingURL)
// and apply the flags to it returning the full URL to use.
func Apply(URL string, flags Flags) string {
	if flags.JPEGQuality > 100 {
		panic("jpeg quality cannot be bigger than 100")
	}
	if flags.JPEGQuality > 0 && !flags.ForceJPEG {
		panic("force jpeg format to change the quality of the output")
	}
	if flags.Original && (flags.Size != 0 || flags.Width != 0 || flags.Height != 0) {
		panic("do not ask for original resolution specifying size, width or height")
	}
	if flags.Width > 3000 || flags.Height > 3000 || flags.Size > 3000 {
		panic("size, width and height should be under 3000 pixels")
	}
	if flags.SmartSquareCrop && flags.SmartSquareCropFaces {
		panic("cannot activate two kinds of smart square crop at the same time")
	}
	if active(flags.SmartSquareCrop, flags.SmartSquareCropFaces, flags.SquareCrop, flags.CircularCrop, flags.SmallestCrop) > 1 {
		panic("cannot activate several kinds of crop at the same time")
	}
	if flags.Blur > 100 {
		panic("cannot blur more than 100% percent")
	}
	if flags.BorderGradientColor.A > 0 {
		panic("do not fill border gradient color alpha; it is not used")
	}

	serialized := []string{}

	if flags.Size != 0 {
		serialized = append(serialized, fmt.Sprintf("s%d", flags.Size))
	}
	if flags.Original {
		serialized = append(serialized, "s0")
	}
	if flags.Width != 0 {
		serialized = append(serialized, fmt.Sprintf("w%d", flags.Width))
	}
	if flags.Height != 0 {
		serialized = append(serialized, fmt.Sprintf("h%d", flags.Height))
	}
	if flags.SmartSquareCropFaces {
		serialized = append(serialized, "p")
	}
	if flags.SmartSquareCrop {
		serialized = append(serialized, "pp")
	}
	if flags.CircularCrop {
		serialized = append(serialized, "cc")
	}
	if flags.SmallestCrop {
		serialized = append(serialized, "ci")
	}
	if flags.SquareCrop {
		serialized = append(serialized, "c")
	}
	if flags.VerticalFlip {
		serialized = append(serialized, "fv")
	}
	if flags.HorizontalFlip {
		serialized = append(serialized, "fh")
	}
	if flags.Rotate != 0 {
		serialized = append(serialized, fmt.Sprintf("r%d", flags.Rotate))
	}
	if flags.ForceJPEG {
		serialized = append(serialized, "rj")
	}
	if flags.ForcePNG {
		serialized = append(serialized, "rp")
	}
	if flags.ForceWebP {
		serialized = append(serialized, "rw")
	}
	if flags.ForceGIF {
		serialized = append(serialized, "rg")
	}
	if flags.MP4 {
		serialized = append(serialized, "rh")
	}
	if flags.KillAnimation {
		serialized = append(serialized, "k")
	}
	if flags.Border > 0 {
		serialized = append(serialized, fmt.Sprintf("b%d", flags.Border))
	}
	if flags.BorderColor.A > 0 || flags.BorderColor.R > 0 || flags.BorderColor.G > 0 || flags.BorderColor.B > 0 {
		color := fmt.Sprintf("0x%x%x%x%x", flags.BorderColor.A, flags.BorderColor.R, flags.BorderColor.G, flags.BorderColor.B)
		color = strings.ToLower(color)
		serialized = append(serialized, color)
	}
	if flags.ForceDownload {
		serialized = append(serialized, "d")
	}
	if flags.ExpiresDays > 0 {
		serialized = append(serialized, fmt.Sprintf("e%d", flags.ExpiresDays))
	}
	if flags.JPEGQuality > 0 {
		serialized = append(serialized, fmt.Sprintf("l%d", flags.JPEGQuality))
	}
	if flags.Blur > 0 {
		serialized = append(serialized, fmt.Sprintf("fSoften=1,%d,0:", flags.Blur))
	}
	if flags.BorderGradientSize > 0 {
		color := fmt.Sprintf("%x%x%x%x", flags.BorderGradientColor.R, flags.BorderGradientColor.G, flags.BorderGradientColor.B)
		color = strings.ToLower(color)
		serialized = append(serialized, fmt.Sprintf("fVignette=1,%d,1.4,0,%s", flags.BorderGradientSize, color))
	}

	if len(serialized) == 0 {
		return URL
	}
	return fmt.Sprintf("%s=%s", URL, strings.Join(serialized, "-"))
}

func active(vars ...bool) uint64 {
	var n uint64
	for _, v := range vars {
		if v {
			n++
		}
	}

	return n
}
