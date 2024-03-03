package generator

import (
	"gRPC/pb"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyBoardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERRITY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(stringSet ...string) string {
	n := len(stringSet)
	if n == 0 {
		return ""
	}

	return stringSet[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}

	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)
}

func randomInt(lo, hi int) int {
	return lo + rand.Intn(hi-lo+1)
}

// rand.Float64 return 0.0 ~ 1.0
func randomFloat64(lo, hi float64) float64 {
	return lo + rand.Float64()*(hi-lo)
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	}

	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	return &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}

}

func randomScreenPanel() pb.Screen_Panel {
	switch rand.Intn(3) {
	case 1:
		return pb.Screen_IPS
	case 2:
		return pb.Screen_OLED
	default:
		return pb.Screen_UNKOWN
	}
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	if brand == "Dell" {
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	}

	return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
}
