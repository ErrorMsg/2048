package main
import ("fmt"
		"time"
		"math/rand"
	   )

type new_array struct{
    array [4][4]int
}

func UP(ia_up new_array, score int64) (new_array,int64){
	ia_ck := ia_up
    for z:=0;z<4;z++{
        x := 0
        t := 0
        for {
            if ia_up.array[x][z] == 0{
                t++
                for y:=x;y<(4-t);y++{
                    ia_up.array[y][z] = ia_up.array[y+1][z]
                    ia_up.array[y+1][z] = 0
                }
            }else{
                    x++
            }
            if x + t == 3{
                break
            }
        }
    }
    ia_new, new_score := PLUS(ia_up, score)
	if ia_ck == ia_new{
		return ia_up,score
	}else {
    	return ia_new, new_score
	}
}

func LEFT(ia_left new_array, score int64) (new_array,int64){
	var new_right new_array
	for z:=0;z<4;z++{
		for x:=0;x<4;x++{
			new_right.array[x][z] = ia_left.array[3-z][x]
		}
	}
	up_plus,new_score := UP(new_right,score)
	var new_left new_array
	for z:=0;z<4;z++{
		for x:=0;x<4;x++{
			new_left.array[x][z] = up_plus.array[z][3-x]
		}
	}
	return new_left,new_score
}

func RIGHT(ia_right new_array, score int64) (new_array,int64){
	var new_left new_array
	for z:=0;z<4;z++{
		for x:=0;x<4;x++{
			new_left.array[x][z] = ia_right.array[z][3-x]
		}
	}
	up_plus,new_score := UP(new_left,score)
	var new_right new_array
	for z:=0;z<4;z++{
		for x:=0;x<4;x++{
			new_right.array[x][z] = up_plus.array[3-z][x]
		}
	}
	return new_right,new_score
}

func DOWN(ia_down new_array, score int64) (new_array,int64){
	var new_inverse new_array
	for z:=0;z<4;z++{
		for x:=0;x<4;x++{
			new_inverse.array[x][z] = ia_down.array[3-x][z]
		}
	}
	up_plus,new_score := UP(new_inverse,score)
	var new_down new_array
	for z:=0;z<4;z++{
		for x:=0;x<4;x++{
			new_down.array[x][z] = up_plus.array[3-x][z]
		}
	}
	return new_down,new_score
}

func PLUS(ia_plus new_array, score int64) (new_array,int64){
    for z:=0;z<4;z++{
        for x:=0;x<3;x++{
            if ia_plus.array[x][z] == ia_plus.array[x+1][z]{
                ia_plus.array[x][z] = 2*ia_plus.array[x][z]
				ia_plus.array[x+1][z] = 0
				score = score + int64(ia_plus.array[x][z])
				for y:=x+1;y<3;y++{
                    ia_plus.array[y][z] = ia_plus.array[y+1][z]
                    ia_plus.array[y+1][z] = 0
				}
            }
        }
    }
	return ia_plus,score
}

func NEXT(ia_next new_array) new_array{
	fmt.Println("next is: ")
	var positions [] int
	for z:=0;z<4;z++{
		for x:=0;x<4;x++{
			if ia_next.array[x][z] == 0{
				p := x*10 + z
				positions = append(positions, p)
			}
		}
	}
	values := [10] int {2,2,2,4,2,2,2,2,2,2}
	np,nv := PICK(len(positions),len(values))
	addr := positions[np]
	ia_next.array[addr/10][addr%10] = values[nv]
	return ia_next
}

func PTA(ia_pta new_array){
	for i:=0;i<4;i++{
		for j:=0;j<4;j++ {
			tp := ia_pta.array[i][j]
			switch{
				case tp == 0:
					fmt.Printf("|         ")
				case tp > 0 && tp < 10:
					fmt.Printf("|%-8d", ia_pta.array[i][j])
				case tp >= 10 && tp < 100:
					fmt.Printf("|%-7d", ia_pta.array[i][j])
				case tp >= 100 && tp < 1000:
					fmt.Printf("|%-6d", ia_pta.array[i][j])
				case tp >= 1000:
					fmt.Printf("|%-5d", ia_pta.array[i][j])
			}
		}
		fmt.Println("|")
	}
	fmt.Println("--------------------------------")
}

func START(){
	var score int64 = 0
	var ia_start new_array
	for i:=0;i<2;i++{
		x,z := PICK(4,4)
		ia_start.array[x][z] = 2
		time.Sleep(time.Microsecond)
	}
	GAME(ia_start, score)
}

func PICK(i,j int) (int,int){
	var (
		interval int64
		position int
		value int
	)
	interval = int64(rand.Intn(100))
	seed1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	seed2 := rand.New(rand.NewSource(time.Now().UnixNano()+interval))
	position = seed1.Intn(i)
	value = seed2.Intn(j)
	return position,value
}

func CHECK(ia_check new_array) bool{
	var (
		ckp bool
		ckv bool
	)
	ckp = true
	ckv = true
	for z:=0;z<4;z++{
		for x:=0;x<4;x++{
			if ia_check.array[x][z] == 0{
				ckp = false
				break
			}
		}
	}
	for z:=0;z<3;z++{
		for x:=0;x<3;x++{
			if ia_check.array[x][z] == ia_check.array[x+1][z] || ia_check.array[x][z] == ia_check.array[x][z+1]{
				ckv = false
				break
			}
		}
	}
	for z:=0;z<3;z++{
		if ia_check.array[3][z] == ia_check.array[3][z+1]{
			ckv = false
			break
		}
	}
	for x:=0;x<3;x++{
		if ia_check.array[x][3] == ia_check.array[x+1][3]{
			ckv = false
			break
		}
	}
	return ckp&&ckv
}

func GAME(ia_game new_array, score int64){
	fmt.Println("GAME START")
	PTA(ia_game)
	fmt.Println("Press W/A/S/D to move, press R to restart, or press E to quit!")
	for{
		var (
			si string
			fg bool
			rt bool
		)
		fmt.Scanln(&si)
		switch{
			case si == "w" || si == "W":
				a1,s1 := UP(ia_game,score)
				if a1 == ia_game{
					fmt.Println("Please move in another way!")
					fg = CHECK(a1)
				}else {
					fmt.Println("score:", s1)
					ia_game, score = NEXT(a1), s1
					PTA(ia_game)
					fg = CHECK(ia_game)
				}

			case si == "s" || si == "S":
				a2,s2 := DOWN(ia_game,score)
				if a2 == ia_game{
					fmt.Println("Please move in another way!")
					fg = CHECK(a2)
				}else {
					fmt.Println("score:", s2)
					ia_game, score = NEXT(a2), s2
					PTA(ia_game)
					fg = CHECK(ia_game)
				}

			case si == "a" || si == "A":
				a3,s3 := LEFT(ia_game,score)
				if a3 == ia_game{
					fmt.Println("Please move in another way!")
					fg = CHECK(a3)
				}else {
					fmt.Println("score:", s3)
					ia_game, score = NEXT(a3), s3
					PTA(ia_game)
					fg = CHECK(ia_game)
				}

			case si == "d" || si == "D":
				a4,s4 := RIGHT(ia_game,score)
				if a4 == ia_game{
					fmt.Println("Please move in another way!")
					fg = CHECK(a4)
				}else {
					fmt.Println("score:", s4)
					ia_game, score = NEXT(a4), s4
					PTA(ia_game)
					fg = CHECK(ia_game)
				}

			case si == "esc" || si == "e" || si == "E":
				fg = true

			case si == "r" || si == "R":
				rt = true
				fmt.Println("Restarting......")
				START()

			default:
				fmt.Println("Please press correct key!")
				continue
		}
		if fg{
			fmt.Println("GAME OVER")
			fmt.Println("Your score is: " , score)
			break
		}
		if rt{
			break
		}
	}
}


func main() {
	START()
}
 
