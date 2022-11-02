package ebitenSubimageTest

import (
	"image"
	"image/color"
	"testing"

	"github.com/amanitaverna/frostutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

const FixedWidth = 640
const FixedHeight = 360

func TestMain(m *testing.M) {
	frostutil.OnTestMain(m)
}

func Test_EbitenSetAndAt(t *testing.T) {
	frostutil.QueueLayoutTest(t, func(t *testing.T, outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
		return FixedWidth, FixedHeight
	})
	// We perform this test in both Update and Draw to make sure that any fix works for both.
	frostutil.QueueUpdateTest(t, test_EbitenUpdateSetAndAt)
	frostutil.QueueDrawTest(t, test_EbitenDrawSetAndAt)
}

func test_EbitenUpdateSetAndAt(t *testing.T) {
	test_EbitenSetAndAtImpl(t, "Update")
}

func test_EbitenDrawSetAndAt(t *testing.T, screen *ebiten.Image) {
	test_EbitenSetAndAtImpl(t, "Draw")
}
func test_EbitenSetAndAtImpl(t *testing.T, whichTest string) {
	ass := assert.New(t)
	rect := image.Rect(0, 0, 1024, 1024)
	tCol := color.RGBAModel.Convert(image.Transparent.C)
	blkCol := color.RGBAModel.Convert(image.Black.C)
	img := image.NewRGBA(rect)
	subWidth := 32
	subHeight := 32

	// Initialize img to all black
	for idx := 0; idx < len(img.Pix); idx += 4 {
		img.Pix[idx] = 0
		img.Pix[idx+1] = 0
		img.Pix[idx+2] = 0
		img.Pix[idx+3] = 0xff
	}

	eImg := ebiten.NewImageFromImage(img)
	// verify that eImg is all black
	ass.True(isAllColor(eImg, blkCol), "[%v] eImg isn't all black after being created from an all-black image!", whichTest)
	// test for black pixels and change them all to transparent
	minY := rect.Min.Y
	maxY := rect.Max.Y
	minX := rect.Min.X
	maxX := rect.Max.X
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			if blkCol == eImg.At(x, y) {
				eImg.Set(x, y, image.Transparent.C)
			}
		}
	}
	// verify that eImg is now all transparent
	ass.True(isAllColor(eImg, tCol), "[%v] eImg isn't all transparent after being repainted from black to transparent!", whichTest)

	// Create sub-images
	for subY := 0; subY < rect.Max.Y; subY += subHeight {
		for subX := 0; subX < rect.Max.X; subX += subWidth {
			subRect := image.Rect(subX, subY, subX+subWidth, subY+subHeight)
			subEImg := eImg.SubImage(subRect).(*ebiten.Image)
			// verify that the sub-image is all transparent
			isT := isAllColor(subEImg, tCol)
			isBlk := isAllColor(subEImg, blkCol)
			ass.True(isT, "[%v] Sub-image starting at <%v, %v> isn't all transparent!", whichTest, subX, subY)
			ass.True(isBlk, "[%v] Sub-image starting at <%v, %v> is all black, somehow!", whichTest, subX, subY)
		}
	}
	// verify that eImg is still all transparent
	ass.True(isAllColor(eImg, tCol), "[%v] eImg isn't all transparent after calling SubImage!", whichTest)
}

func isAllColor(eImg *ebiten.Image, col color.Color) bool {
	minY := eImg.Bounds().Min.Y
	maxY := eImg.Bounds().Max.Y
	minX := eImg.Bounds().Min.X
	maxX := eImg.Bounds().Max.X
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			if col != eImg.At(x, y) {
				return false
			}
		}
	}
	return true
}
