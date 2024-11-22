package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/*
func makeGrid(x int, y int, val int) [][]int { //tworzy macierz o rozmiarach x y wypełniając ją wartościa val
	matrix := make([][]int, y)

	for v := range matrix {
		matrix[v] = make([]int, x)
	}

	for yv := range matrix {
		for i := 0; i < x; i++ {
			matrix[yv][i] = val
		}
	}

	return matrix
}
*/

type wall struct {
	vector rl.Vector3
	width  float32
	height float32
	length float32
}

func wall_gen_rl(grid [][]int) []wall {
	rows := len(grid)    //Y
	cols := len(grid[0]) //X

	walls := []wall{}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {

			if grid[y][x]&1 != 0 {
				walls = append(walls, wall{
					vector: rl.NewVector3(float32(x)+0.5, 0, float32(y)),
					width:  1,
					height: 1,
					length: 0.1,
				})
			}

			if grid[y][x]&2 != 0 {
				walls = append(walls, wall{
					vector: rl.NewVector3(float32(x)+1, 0, float32(y)+0.5),
					width:  0.1,
					height: 1,
					length: 1,
				})
			}

			if grid[y][x]&4 != 0 {
				walls = append(walls, wall{
					vector: rl.NewVector3(float32(x)+0.5, 0, float32(y)+1),
					width:  1,
					height: 1,
					length: 0.1,
				})

			}

			if grid[y][x]&8 != 0 {
				walls = append(walls, wall{
					vector: rl.NewVector3(float32(x), 0, float32(y)+0.5),
					width:  0.1,
					height: 1,
					length: 1,
				})
			}

		}
	}
	return walls
}

func main() {

	//pole := makeGrid(10, 10, 15)
	//stos := stack{
	//list:    [][]int{},
	//visited: [][]int{},
	//}

	dd := dfs{
		grid:      [][]int{},
		start_pos: []int{},
	}

	grid := [][]int{
		{9, 5, 5, 1, 3, 11},
		{8, 5, 3, 10, 8, 6},
		{8, 1, 4, 4, 0, 3},
		{8, 2, 11, 11, 8, 2},
		{8, 4, 2, 8, 4, 2},
		{12, 7, 12, 6, 13, 6},
	}

	dd.gridInit(51, 51, 15)
	dd.startInit(0, 0)
	dd.createMaze()
	//dd.grid = grid
	fmt.Println(dd.grid)
	//dd.drawMaze(dd.grid)

	ast := astar{
		grid:       grid,
		start_pos:  [2]int{0, 0},
		finish_pos: [2]int{50, 50},
		open_list: stack{
			list: [][]int{},
		},
		closed_list: make(map[[2]int]bool),
		parent:      make(map[[2]int][2]int),
		g_cost:      make(map[[2]int]int),
		path:        [][2]int{},
	}
	fmt.Println(ast.manhEstimate([]int{2, 0}))
	ast.a_star_solving(dd.grid)
	//ast.displaySolution(dd.grid)
	//fmt.Println(findNeighbors(1, 1, 3, 3))
	//fmt.Println(findNeighbors(2, 1, 3, 3))
	//fmt.Println(findNeighbors(1, 2, 3, 3))

	//ast.a_star_solving(dd.grid)

	const screenW = int32(1280)
	const screenH = int32(720)

	rl.InitWindow(screenW, screenH, "Maze")
	rec := rl.NewRectangle(0, 0, float32(screenW), float32(screenH))

	stage := 0

	current_input := 1
	inputX := ""
	inputY := ""

	for !rl.WindowShouldClose() {
		if stage == 0 { //--------------------------------------------Okno startowe
			rl.BeginDrawing()

			rl.ClearBackground(rl.Black)

			txt := "Kamil Wlodarczyk"
			txtlen := rl.MeasureText(txt, 50)
			rl.DrawText(txt, screenW/2-txtlen/2-3, screenH/2-150+3, 50, rl.Magenta)
			rl.DrawText(txt, screenW/2-txtlen/2-1, screenH/2-150+1, 50, rl.Black)
			rl.DrawText(txt, screenW/2-txtlen/2, screenH/2-150, 50, rl.White)
			txt = "labirynt z A*"
			txtlen = rl.MeasureText(txt, 50)
			rl.DrawText(txt, screenW/2-txtlen/2-3, screenH/2-90+3, 50, rl.Magenta)
			rl.DrawText(txt, screenW/2-txtlen/2-1, screenH/2-90+1, 50, rl.Black)
			rl.DrawText(txt, screenW/2-txtlen/2, screenH/2-90, 50, rl.White)
			txt = "Press Enter to progress"
			txtlen = rl.MeasureText(txt, 30)
			rl.DrawText(txt, screenW/2-txtlen/2-3, screenH/2+3, 30, rl.Magenta)
			rl.DrawText(txt, screenW/2-txtlen/2-1, screenH/2+1, 30, rl.Black)
			rl.DrawText(txt, screenW/2-txtlen/2, screenH/2, 30, rl.White)

			rl.EndDrawing()

			if rl.IsKeyDown(rl.KeyEnter) {
				stage = 1 // Przejdź do drugiego etapu
				time.Sleep(500 * time.Millisecond)
			}
		} else if stage == 1 { //--------------------------------------------Wpisywanie wartości labiryntu

			rl.BeginDrawing()

			if rl.IsKeyPressed(rl.KeyTab) {
				if current_input == 1 {
					current_input = 2
				} else {
					current_input = 1
				}
			}

			if rl.IsKeyPressed(rl.KeyBackspace) {
				if current_input == 1 && len(inputX) > 0 {
					inputX = inputX[:len(inputX)-1]
				} else if current_input == 2 && len(inputY) > 0 {
					inputY = inputY[:len(inputY)-1]
				}
			}

			key := rl.GetCharPressed()

			for key > 0 {
				if key >= '0' && key <= '9' {
					if current_input == 1 {
						inputX += string(key)
					} else if current_input == 2 {
						inputY += string(key)
					}
				}
				key = rl.GetCharPressed()
			}

			rl.ClearBackground(rl.Black)
			//rec := rl.NewRectangle(0, 0, float32(screenW), float32(screenH))
			rl.DrawRectangleRec(rec, rl.DarkPurple)

			txt := "Podaj rozmiar labiryntu"
			txtlen := rl.MeasureText(txt, 50)
			rl.DrawText(txt, screenW/2-txtlen/2-1, screenH/2-150+1, 50, rl.Black)
			rl.DrawText(txt, screenW/2-txtlen/2, screenH/2-150, 50, rl.White)

			txt = "TAB aby przelaczyc miedzy wartosciami"
			txtlen = rl.MeasureText(txt, 50)
			rl.DrawText(txt, screenW/2-txtlen/2-1+300, screenH/2-90+1, 20, rl.Black)
			rl.DrawText(txt, screenW/2-txtlen/2+300, screenH/2-90, 20, rl.White)

			//pole X
			koord1X := screenW/2 - 450
			Koord1Y := int32(400)
			//pole Y
			koord2X := screenW/2 + 200
			Koord2Y := int32(400)

			if current_input == 1 {
				rl.DrawRectangle(koord1X-5, Koord1Y-5, 310, 60, rl.Red)
			} else if current_input == 2 {
				rl.DrawRectangle(koord2X-5, Koord2Y-5, 310, 60, rl.Red)
			}

			rl.DrawText("X:", koord1X-70, Koord1Y, 55, rl.Black)
			rl.DrawRectangle(koord1X, Koord1Y, 300, 50, rl.LightGray)
			rl.DrawText(inputX, koord1X+10, Koord1Y+10, 30, rl.Black)

			rl.DrawText("Y:", koord2X-70, Koord2Y, 55, rl.Black)
			rl.DrawRectangle(koord2X, Koord2Y, 300, 50, rl.LightGray)
			rl.DrawText(inputY, koord2X+10, Koord2Y+10, 30, rl.Black)

			txt = "Press Enter to progress"
			txtlen = rl.MeasureText(txt, 30)
			rl.DrawText(txt, screenW/2-txtlen/2-1, screenH/2+160+1, 30, rl.Black)
			rl.DrawText(txt, screenW/2-txtlen/2, screenH/2+160, 30, rl.White)

			rl.EndDrawing()
			if rl.IsKeyDown(rl.KeyEnter) {
				stage = 2
			}
		} else if stage == 2 { //--------------------------------------------ekran 3D
			dupa := 0
			camera := rl.Camera3D{}
			camera.Position = rl.NewVector3(10.0, 10.0, 10.0)
			camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
			camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
			camera.Fovy = 45.0
			camera.Projection = rl.CameraPerspective

			//cubePosition := rl.NewVector3(0.0, 0.0, 0.0)

			rl.SetTargetFPS(60)
			centerX := int(screenW / 2)
			centerY := int(screenH / 2)
			rl.HideCursor()

			rl.SetTargetFPS(60)

			lastTime := time.Now()
			index := 0

			walls := wall_gen_rl(dd.grid)

			for !rl.WindowShouldClose() {

				rl.UpdateCamera(&camera, rl.CameraFree)

				rl.DrawFPS(1200, 10)

				rl.SetMousePosition(centerX, centerY)

				if time.Since(lastTime).Seconds() >= 0.000001 && index < len(ast.visited_list) { //rysowanie nowego pola co określoną ilość sekund
					lastTime = time.Now()
					index++
				}

				if rl.IsKeyDown(rl.KeyZ) {
					camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
				}

				rl.BeginDrawing()

				rl.ClearBackground(rl.RayWhite)

				rl.BeginMode3D(camera)

				for _, v_wall := range walls {
					rl.DrawCube(v_wall.vector, v_wall.width, v_wall.height, v_wall.length, rl.Black)
				}

				//for i := 0; i < index; i++ {
				//rl.DrawCube(rl.NewVector3(float32(ast.visited_list[i][0])+0.5, 0.0, float32(ast.visited_list[i][1])+0.5), 1, 0.05, 1, rl.Red)
				//}

				for _, vn := range ast.visited_list {
					rl.DrawCube(rl.NewVector3(float32(vn[0])+0.5, 0.1, float32(vn[1])+0.5), 1, 0.05, 1, rl.Red)
				}

				for _, vn := range ast.path {
					rl.DrawCube(rl.NewVector3(float32(vn[0])+0.5, 0.2, float32(vn[1])+0.5), 1, 0.1, 1, rl.Yellow)
				}

				if index == len(ast.visited_list) {
					for _, vn := range ast.path {
						rl.DrawCube(rl.NewVector3(float32(vn[0])+0.5, 0.2, float32(vn[1])+0.5), 1, 0.1, 1, rl.Yellow)
					}
				}

				//rl.DrawCube(rl.NewVector3(3+0.5, 0.5, 3), 0.1, 1, 1, rl.Black)

				//rl.DrawCube(cubePosition, 2.0, 2.0, 2.0, rl.Red)
				//rl.DrawCube(rl.NewVector3(20.0, 0.0, 20.0), 2.0, 2.0, 2.0, rl.Red)
				//rl.DrawCubeWires(cubePosition, 2.0, 2.0, 2.0, rl.Maroon)

				rl.DrawGrid(10, 1.0)

				rl.EndMode3D()

				rl.DrawRectangle(10, 10, 320, 133, rl.Fade(rl.SkyBlue, 0.5))
				rl.DrawRectangleLines(10, 10, 320, 133, rl.Blue)

				rl.DrawText("Free camera default controls:", 20, 20, 10, rl.Black)
				rl.DrawText("- Mouse Wheel to Zoom in-out", 40, 40, 10, rl.DarkGray)
				rl.DrawText("- Mouse Wheel Pressed to Pan", 40, 60, 10, rl.DarkGray)
				rl.DrawText("- Z to zoom to (0, 0, 0)", 40, 120, 10, rl.DarkGray)

				rl.EndDrawing()

				fmt.Println("index: ", index)
				fmt.Println("len: ", len(ast.visited_list))
				fmt.Println("-----------------------------------------")

				if dupa == 0 {
					dd.drawMaze(dd.grid)
					dupa = 1
				}

			}
		}

	}
	fmt.Println("X:", inputX)
	fmt.Println("Y:", inputY)

	/*


	 */

	rl.CloseWindow()

}
