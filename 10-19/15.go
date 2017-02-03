package main

func Euler15() {

	numOfPaths := GetLatticePathNums(20, 20)

	PrintAnswer(15, numOfPaths)
}

func GetLatticePathNums(height int, width int) int {

	//fmt.Printf("GetLatticePathNums  height: %d ,  width:%d\n", height, width)

	if height == 0 && width == 0 {
		return 0
	} else if height == 1 && width == 0 {
		return 1
	} else if width == 1 && height == 0 {
		return 1
	} else if width == 0 {
		return GetLatticePathNums(height-1, 0)
	} else if height == 0 {
		return GetLatticePathNums(0, width-1)
	} else {
		return GetLatticePathNums(height-1, width) + GetLatticePathNums(height, width-1)
	}
}
