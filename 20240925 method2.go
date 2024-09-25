package main

import "fmt"

//给定几种不同颜色、体积的盒子，输出体积最大的盒子颜色，最后将他们全部改为黑色

const (
	WHITE = iota
	RED
	YELLOW
	GREEN
	BLUE
	BLACK
)

// 定义盒子的各项参数
type Box struct {
	width, height, depth float64
	color                Color
}

type Color int

type Boxlist []Box

func (b *Box) Setcolor(c Color) {
	b.color = c
}

func (b Boxlist) Paintitallblack() {
	for i, _ := range b {
		b[i].Setcolor(BLACK)
	}
}
func (b Box) volume() float64 {
	return b.height * b.width * b.depth
}

func (c Color) Colorstring() string {
	var a string = "haha"
	switch c {
	case 0:
		a = "WHITE"
	case 1:
		a = "RED"
	case 2:
		a = "YELLOW"
	case 3:
		a = "GREEN"
	case 4:
		a = "BLUE"
	case 5:
		a = "BLACK"
	}
	return a
}

//func (c color)Colorstring string{strings := []string{"","","","","",""} return strings[c]}

func (b Boxlist) biggestone() (Color, float64) {
	a := float64(0)
	c := Color(WHITE)
	for _, i := range b {
		if bv := i.volume(); bv > a {
			a = bv
			c = i.color
		}
	}
	return c, a
}

func main() {
	realboxes := Boxlist{
		Box{3, 8, 3, BLUE},
		Box{6, 2, 4, WHITE},
		Box{6, 2, 8, BLACK},
	}
	bc1, bv1 := realboxes.biggestone()
	rbc1 := bc1.Colorstring()
	fmt.Println("The color of the biggest one is", rbc1)
	fmt.Println("The valume of the biggest one is", bv1)
	fmt.Println("Paint it all black ！")
	realboxes.Paintitallblack()
}
