This package contains a test which:
1. Creates a new RGBA image
2. Paints it black
3. Creates an *ebiten.Image from it, using ebite.NewImageFromImage
4. Verifies that the *ebiten.Image is all black (#000000ff)
5. Paints it transparent (#00000000)
6. Verifies that it is all transparent
7. Creates many subimages from it, and checks each one as it is created to see whether they are all black or all transparent, reporting it as a failure if they are all black or if they are not all transparent.

It uses At to check the pixel color and Set to set it at every step, except for when it is creating the initial image (at that one time, it directly modifies its Pix buffer instead).

Some observations:
- This test fails on Ebitengine 2.4.0 - 2.4.9 but passes on 2.3.*. Specifically, the checks of the subimages' pixels show them to be black when they should be fully transparent. 
- If the At and Set code are replaced with code that uses ReadPixels and WritePixels, the test passes.
- Without the subimage calls, At and Set appear to work fine.