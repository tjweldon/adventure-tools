package rotato

import (
	"adventure-tools/utils/numbers"
	"github.com/ungerik/go3d/quaternion"
	"github.com/ungerik/go3d/vec3"
	"image/color"
)

const mag = 1.732050807568877

func RotateHue(rgb *color.RGBA, hueAngle float32) {
	axis := vec3.T{
		1.0 / mag,
		1.0 / mag,
		1.0 / mag,
	}

	rot := quaternion.FromAxisAngle(&axis, hueAngle)

	rgbVec := vec3.T{
		float32(rgb.R),
		float32(rgb.G),
		float32(rgb.B),
	}

	res := rot.RotatedVec3(&rgbVec)
	rgb.R = bound(res[0])
	rgb.G = bound(res[1])
	rgb.B = bound(res[2])
}

func bound(fl float32) uint8 {
	return uint8(numbers.Max(0, numbers.Min(255, int(fl))))
}
