package snake

type plane struct {
    height int
    width  int
}

func newPlane(h, w int) *plane {
    pln := &plane{
        height: h,
        width:  w,
    }
    return pln
}