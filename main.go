package main

import (
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type wall struct { //ściana labiryntu
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

			if grid[y][x]&1 != 0 { //pierwszy bit
				walls = append(walls, wall{
					vector: rl.NewVector3(float32(x)+0.5, 0, float32(y)),
					width:  1,
					height: 1,
					length: 0.1,
				})
			}

			if grid[y][x]&2 != 0 { //drugi bit
				walls = append(walls, wall{
					vector: rl.NewVector3(float32(x)+1, 0, float32(y)+0.5),
					width:  0.1,
					height: 1,
					length: 1,
				})
			}

			if grid[y][x]&4 != 0 { //trzeci bit
				walls = append(walls, wall{
					vector: rl.NewVector3(float32(x)+0.5, 0, float32(y)+1),
					width:  1,
					height: 1,
					length: 0.1,
				})

			}

			if grid[y][x]&8 != 0 { //czwarty bit
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

	dd := dfs{ //inicjalizacja generowania labiryntu
		grid:      [][]int{},
		start_pos: []int{},
	}

	ast := astar{ //inicjalizacja A*
		grid:       [][]int{},
		start_pos:  [2]int{0, 0},
		finish_pos: [2]int{220, 220},
		open_list: stack{
			list: [][]int{},
		},
		closed_list: make(map[[2]int]bool),
		parent:      make(map[[2]int][2]int),
		g_cost:      make(map[[2]int]int),
		path:        [][2]int{},
	}

	//wielkość ekranu
	const screenW = int32(1280)
	const screenH = int32(720)

	rl.InitWindow(screenW, screenH, "Maze")                          //tytuł okna
	rec := rl.NewRectangle(0, 0, float32(screenW), float32(screenH)) //tło ekranu

	//wartości pomocnicze
	stage := 0
	current_input := 1
	inputX := ""
	inputY := ""
	inputXF := ""
	inputYF := ""
	input_time := ""
	time_val := float64(0)
	val_error := false
	pole := 0

	for !rl.WindowShouldClose() {

		if stage == 0 { //--------------------------------------------Okno startowe (Wzięte z przykładów na githubie biblioteki raylib)
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

			if rl.IsKeyPressed(rl.KeyTab) { //przechodzenie między inputami
				if current_input == 1 {
					current_input = 2
				} else if current_input == 2 {
					current_input = 3
				} else if current_input == 3 {
					current_input = 4
				} else if current_input == 4 {
					current_input = 5
				} else {
					current_input = 1
				}
			}

			if rl.IsKeyPressed(rl.KeyBackspace) { //usuwanie znaków
				if current_input == 1 && len(inputX) > 0 {
					inputX = inputX[:len(inputX)-1]
				} else if current_input == 2 && len(inputY) > 0 {
					inputY = inputY[:len(inputY)-1]
				} else if current_input == 3 && len(inputXF) > 0 {
					inputXF = inputXF[:len(inputXF)-1]
				} else if current_input == 4 && len(inputYF) > 0 {
					inputYF = inputYF[:len(inputYF)-1]
				} else if current_input == 5 && len(input_time) > 0 {
					input_time = input_time[:len(input_time)-1]
				}
			}

			key := rl.GetCharPressed()

			for key > 0 { //wpisywanie liczb
				if (key >= '0' && key <= '9') || key == '.' {
					if current_input == 1 {
						inputX += string(key)
					} else if current_input == 2 {
						inputY += string(key)
					} else if current_input == 3 {
						inputXF += string(key)
					} else if current_input == 4 {
						inputYF += string(key)
					} else if current_input == 5 {
						input_time += string(key)
					}
				}
				key = rl.GetCharPressed()
			}

			rl.ClearBackground(rl.Black)
			rl.DrawRectangleRec(rec, rl.DarkPurple)

			txt := "Podaj rozmiar labiryntu"
			txtlen := rl.MeasureText(txt, 50)
			rl.DrawText(txt, screenW/2-txtlen/2-1, screenH/2-250+1, 50, rl.Black)
			rl.DrawText(txt, screenW/2-txtlen/2, screenH/2-250, 50, rl.White)

			txt = "TAB aby przelaczyc miedzy wartosciami"
			txtlen = rl.MeasureText(txt, 50)
			rl.DrawText(txt, screenW/2-txtlen/2-1+300, screenH/2-190+1, 20, rl.Black)
			rl.DrawText(txt, screenW/2-txtlen/2+300, screenH/2-190, 20, rl.White)

			//pole X
			koord1X := screenW/2 - 450
			Koord1Y := int32(300)
			//pole Y
			koord2X := screenW/2 + 200
			Koord2Y := int32(300)

			//pole X_finish
			koord1XF := screenW/2 - 450
			Koord1YF := int32(400)
			//pole Y_finish
			koord2XF := screenW/2 + 200
			Koord2YF := int32(400)
			//pole czas
			koordCX := screenW/2 - 125
			KoordCY := int32(500)

			if current_input == 1 { //rysowanie czerwonej obwódki wokół aktualnego pola
				rl.DrawRectangle(koord1X-5, Koord1Y-5, 310, 60, rl.Red)
			} else if current_input == 2 {
				rl.DrawRectangle(koord2X-5, Koord2Y-5, 310, 60, rl.Red)
			} else if current_input == 3 {
				rl.DrawRectangle(koord1XF-5, Koord1YF-5, 310, 60, rl.Red)
			} else if current_input == 4 {
				rl.DrawRectangle(koord2XF-5, Koord2YF-5, 310, 60, rl.Red)
			} else if current_input == 5 {
				rl.DrawRectangle(koordCX-5, KoordCY-5, 310, 60, rl.Red)
			}

			rl.DrawText("X:", koord1X-70, Koord1Y, 55, rl.Black)
			rl.DrawRectangle(koord1X, Koord1Y, 300, 50, rl.LightGray)
			rl.DrawText(inputX, koord1X+10, Koord1Y+10, 30, rl.Black)

			rl.DrawText("Y:", koord2X-70, Koord2Y, 55, rl.Black)
			rl.DrawRectangle(koord2X, Koord2Y, 300, 50, rl.LightGray)
			rl.DrawText(inputY, koord2X+10, Koord2Y+10, 30, rl.Black)

			rl.DrawText("Cel_X:", koord1XF-180, Koord1YF, 55, rl.Black)
			rl.DrawRectangle(koord1XF, Koord1YF, 300, 50, rl.LightGray)
			rl.DrawText(inputXF, koord1XF+10, Koord1YF+10, 30, rl.Black)

			rl.DrawText("Cel_Y:", koord2XF-180, Koord2YF, 55, rl.Black)
			rl.DrawRectangle(koord2XF, Koord2YF, 300, 50, rl.LightGray)
			rl.DrawText(inputYF, koord2XF+10, Koord2YF+10, 30, rl.Black)

			rl.DrawText("Time:", koordCX-150, KoordCY, 55, rl.Black)
			rl.DrawRectangle(koordCX, KoordCY, 300, 50, rl.LightGray)
			rl.DrawText(input_time, koordCX+10, KoordCY+10, 30, rl.Black)

			txt = "Press Enter to progress"
			txtlen = rl.MeasureText(txt, 30)
			rl.DrawText(txt, screenW/2-txtlen/2-1, screenH/2+300+1, 30, rl.Black)
			rl.DrawText(txt, screenW/2-txtlen/2, screenH/2+300, 30, rl.White)

			if val_error {
				txt = "Niepoprwane wartosci"
				txtlen = rl.MeasureText(txt, 25)
				rl.DrawText(txt, screenW/2-txtlen/2-1, screenH/2+220+1, 25, rl.Black)
				rl.DrawText(txt, screenW/2-txtlen/2, screenH/2+220, 25, rl.White)
			}

			rl.EndDrawing()
			if rl.IsKeyDown(rl.KeyEnter) { //sprawdzenie warunków do przejścia do następnego etapu

				var err5 error
				x, err1 := strconv.Atoi(inputX)
				y, err2 := strconv.Atoi(inputY)
				x_cel, err3 := strconv.Atoi(inputXF)
				y_cel, err4 := strconv.Atoi(inputYF)
				time_val, err5 = strconv.ParseFloat(input_time, 64)
				pole = x * y

				if err1 == nil && err2 == nil && err3 == nil && err4 == nil && err5 == nil && x > 0 && y > 0 && x_cel > 0 && y_cel > 0 && time_val > 0 && x_cel < x && y_cel < y {
					stage = 2
					ast.finish_pos = [2]int{x_cel, y_cel}
					dd.gridInit(x, y, 15)
					dd.startInit(0, 0)
					dd.createMaze()
					ast.a_star_solving(dd.grid)

				} else {
					val_error = true
				}

			}
		} else if stage == 2 { //--------------------------------------------ekran 3D
			camera := rl.Camera3D{}
			camera.Position = rl.NewVector3(10.0, 10.0, 10.0)
			camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
			camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
			camera.Fovy = 45.0
			camera.Projection = rl.CameraPerspective

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

				if rl.IsKeyDown(rl.KeyZ) {
					camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
				}

				if rl.IsKeyDown(rl.KeyOne) {
					camera.Position = rl.NewVector3(49, 120, 90)
				}
				if rl.IsKeyDown(rl.KeyTwo) {
					camera.Position = rl.NewVector3(120, 170, 120)
				}
				if rl.IsKeyDown(rl.KeyZero) {
					camera.Position = rl.NewVector3(0, 50, 0)
				}
				if rl.IsKeyDown(rl.KeyNine) {
					camera.Position = rl.NewVector3(float32(ast.finish_pos[0]), 50, float32(ast.finish_pos[1]))
				}

				rl.BeginDrawing()

				rl.ClearBackground(rl.RayWhite)

				rl.BeginMode3D(camera)

				for _, v_wall := range walls {
					rl.DrawCube(v_wall.vector, v_wall.width, v_wall.height, v_wall.length, rl.Black)
				}

				if pole > 3000 { //jesli powierzchnia >3000 to niech narysuje całość aby zaoszczędzić czas
					for _, vn := range ast.visited_list {
						rl.DrawCube(rl.NewVector3(float32(vn[0])+0.5, 0.1, float32(vn[1])+0.5), 1, 0.05, 1, rl.Red)
					}

					for _, vn := range ast.path {
						rl.DrawCube(rl.NewVector3(float32(vn[0])+0.5, 0.2, float32(vn[1])+0.5), 1, 0.1, 1, rl.Yellow)
					}
				} else { //rysowanie fragmentami
					for i := 0; i < index; i++ {
						rl.DrawCube(rl.NewVector3(float32(ast.visited_list[i][0])+0.5, 0.0, float32(ast.visited_list[i][1])+0.5), 1, 0.05, 1, rl.Red)
					}
					if index == len(ast.visited_list) {
						for _, vn := range ast.path {
							rl.DrawCube(rl.NewVector3(float32(vn[0])+0.5, 0.1, float32(vn[1])+0.5), 1, 0.1, 1, rl.Yellow)
						}
					}
					if time.Since(lastTime).Seconds() >= time_val && index < len(ast.visited_list) { //rysowanie nowego pola co określoną ilość sekund
						lastTime = time.Now()
						index++
					}
				}

				rl.DrawCube(rl.NewVector3(float32(ast.finish_pos[0])+0.5, 0.3, float32(ast.finish_pos[1])+0.5), 1, 0.05, 1, rl.Green)

				rl.DrawCube(rl.NewVector3(float32(ast.start_pos[0])+0.5, 0.3, float32(ast.start_pos[1])+0.5), 1, 0.05, 1, rl.LightGray)

				rl.EndMode3D()

				rl.DrawRectangle(10, 10, 320, 133, rl.Fade(rl.SkyBlue, 0.5)) //okno z przykładów na githubie biblioteki raylib
				rl.DrawRectangleLines(10, 10, 320, 133, rl.Blue)

				rl.DrawText("Free camera default controls:", 20, 20, 10, rl.Black)
				rl.DrawText("- Mouse Wheel to Zoom in-out", 40, 40, 10, rl.DarkGray)
				rl.DrawText("- Mouse Wheel Pressed to Pan", 40, 60, 10, rl.DarkGray)
				rl.DrawText("- 1,2,9,0 to change camera to preset positions", 40, 80, 10, rl.DarkGray)
				rl.DrawText("- Z to zoom to (0, 0, 0)", 40, 120, 10, rl.DarkGray)

				rl.EndDrawing()

			}
		}

	}
	rl.CloseWindow()

}
