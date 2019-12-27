package rectanglefilter

// Rectangle struct
type Rectangle struct {
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	InsertTime string `json:"time"`
}

// AddRectangleSt struct
type AddRectangleSt struct {
	Main  Rectangle   `json:"main"`
	Input []Rectangle `json:"input"`
}

// Init function
func Init() {
	dbInit()
}

func collided(first Rectangle, second Rectangle) bool {
	if first.X+first.Width < second.X {
		return false // "second" is right
	}
	if second.X+second.Width < first.X {
		return false //"second" is left
	}
	if first.Y+first.Height < second.Y {
		return false // "second" is up
	}
	if second.Y+second.Height < first.Y {
		return false // "other" is down
	}
	return true
}

// AddRectangle function
func AddRectangle(ar AddRectangleSt) {
	for _, element := range ar.Input {
		if collided(ar.Main, element) {
			AddRectangleDb(element)
		}
	}
}

// GetAllRectangle function
func GetAllRectangle() []Rectangle {
	return GetAllRectangleDb()
}

// GetUniqueRectangles function
func GetUniqueRectangles() []Rectangle {
	return GetUniqueRectanglesDb()
}
