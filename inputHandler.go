package main

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type InputHandler struct {
	initialized    bool
	selectedObject IGameObject
	objectToPlace  IGameObject
}

func (input *InputHandler) InitializeObjectToPlace(object IGameObject) {
	input.objectToPlace = object
	input.initialized = true
}

func (input *InputHandler) HandleInput(
	win *pixelgl.Window,
	cam *pixel.Matrix,
	gameCommands Commands,
	gameObjs *GameObjects,
	gibletAssets ObjectAssets,
	livingAssets ObjectAssets,
	dt float64,
	camSpeed float64,
	camZoom *float64,
	camZoomSpeed float64,
	camPos *pixel.Vec,
	drawHitBox *bool,
) {
	if !input.initialized {
		input.InitializeObjectToPlace(getShallowLivingObject(livingAssets))
	}

	//select giblet
	if win.JustPressed(pixelgl.Key0) {
		input.objectToPlace = getShallowGibletObject(gibletAssets)
	}

	//select living object
	if win.JustPressed(pixelgl.Key1) {
		input.objectToPlace = getShallowLivingObject(livingAssets)
	}

	//place the selected object
	if win.JustPressed(pixelgl.MouseButtonLeft) && !win.Pressed(pixelgl.KeyLeftControl) {
		mouse := cam.Unproject(win.MousePosition())
		// once objectToPlace gets animation information, we can remove the type switch here
		gameCommands[fmt.Sprintf("AddObject: %s", input.objectToPlace.ObjectName())] = gameObjs.AddObject(input.objectToPlace, mouse)
	}

	//handle ctrl functions
	if win.Pressed(pixelgl.KeyLeftControl) {
		win.SetCursorVisible(true)
		if win.JustPressed(pixelgl.MouseButtonRight) {
			mouse := cam.Unproject(win.MousePosition())
			//add a command to commands
			gameCommands[fmt.Sprintf("RemoveObject x:%f, y:%f", mouse.X, mouse.Y)] = gameObjs.RemoveObject(mouse)
		}
		if win.JustPressed(pixelgl.MouseButtonLeft) { //ctrl + left click
			mouse := cam.Unproject(win.MousePosition())
			newSelectedObject, _, hit, err := gameObjs.getSelectedGameObj(mouse)
			if err != nil {
				fmt.Print(err.Error())
			}
			if hit { //hit object
				//unselect last object
				if input.selectedObject != nil {
					input.selectedObject.changeState(idle)
				}

				input.selectedObject = newSelectedObject
				fmt.Println("object id:", input.selectedObject.getID())
				switch input.selectedObject.(type) {
				case *livingObject:
					{
						input.selectedObject.changeState(selected)
					}
				case *GibletObject:
					{

					}
				}
			} else {
				//ctrl + LM click && no object hit
				fmt.Println("ctrl + LM click on empty space")
			}
		}
	}

	//toggle hit box draw
	if win.JustPressed(pixelgl.KeyH) {
		*drawHitBox = !*drawHitBox
	}

	//move camera
	if win.Pressed(pixelgl.KeyA) {
		camPos.X -= camSpeed * dt
	}
	if win.Pressed(pixelgl.KeyD) {
		camPos.X += camSpeed * dt
	}
	if win.Pressed(pixelgl.KeyS) {
		camPos.Y -= camSpeed * dt
	}
	if win.Pressed(pixelgl.KeyW) {
		camPos.Y += camSpeed * dt
	}

	//zoom camera
	*camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

	// //used for framerate test
	// if win.Pressed(pixelgl.MouseButtonLeft) {
	// 	if win.Pressed(pixelgl.KeyLeftShift) {
	// 		mouse := cam.Unproject(win.MousePosition())
	// 		switch objectToPlace.(type) {
	// 		case *livingObject:
	// 			{
	// 				gameObjs = gameObjs.appendLivingObject(livingObjectAssets, mouse)
	// 			}
	// 		case *GibletObject:
	// 			{
	// 				gameObjs = gameObjs.appendGibletObject(gibletObjectAssets, mouse)
	// 			}
	// 		}
	// 	}
	// }

}
