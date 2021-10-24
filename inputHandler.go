package main

// type inputHandler struct{
// 	selectedObject gameObject
// 	objectToPlace gameObject
// }

// //select giblet
// if win.JustPressed(pixelgl.Key0) {
// 	switch objectToPlace.(type) {
// 	case *gibletObject:
// 		{
// 			//do nothing, already selected
// 		}
// 	case *livingObject:
// 		{
// 			objectToPlace = getShallowGibletObject(gibletAnimKeys, gibletAnims, gibletSheet)
// 		}
// 	}
// }

// //select living object
// if win.JustPressed(pixelgl.Key1) {
// 	switch objectToPlace.(type) {
// 	case *gibletObject:
// 		{
// 			objectToPlace = getShallowLivingObject(pinkAnimKeys, pinkAnims, pinkSheet)
// 		}
// 	case *livingObject:
// 		{
// 			//do nothing, already selected
// 		}
// 	}
// }

// //place the selected object
// if win.JustPressed(pixelgl.MouseButtonLeft) && !win.Pressed(pixelgl.KeyLeftControl) {
// 	mouse := cam.Unproject(win.MousePosition())
// 	// once objectToPlace gets animation information, we can remove the type switch here
// 	gameCommands[fmt.Sprintf("AddObject: %s", objectToPlace.ObjectName())] = gameObjs.AddObject(objectToPlace, mouse)
// }

// //handle ctrl functions
// if win.Pressed(pixelgl.KeyLeftControl) {
// 	win.SetCursorVisible(true)
// 	if win.JustPressed(pixelgl.MouseButtonRight) {
// 		mouse := cam.Unproject(win.MousePosition())
// 		//add a command to commands
// 		gameCommands[fmt.Sprintf("RemoveObject x:%f, y:%f", mouse.X, mouse.Y)] = gameObjs.RemoveObject(mouse)
// 	}
// 	if win.JustPressed(pixelgl.MouseButtonLeft) { //ctrl + left click
// 		mouse := cam.Unproject(win.MousePosition())
// 		newSelectedObject, _, hit, err := gameObjs.getSelectedGameObj(mouse)
// 		if err != nil {
// 			fmt.Print(err.Error())
// 		}
// 		if hit { //hit object
// 			//unselect last object
// 			if selectedObject != nil {
// 				selectedObject.changeState(idle)
// 			}

// 			selectedObject = newSelectedObject
// 			fmt.Println("object id:", selectedObject.getID())
// 			switch selectedObject.(type) {
// 			case *livingObject:
// 				{
// 					selectedObject.changeState(selected)
// 				}
// 			case *gibletObject:
// 				{

// 				}
// 			}
// 		} else {
// 			//ctrl + LM click && no object hit
// 			fmt.Println("ctrl + LM click on empty space")
// 		}
// 	}
// }

// //toggle hit box draw
// if win.JustPressed(pixelgl.KeyH) {
// 	drawHitBox = !drawHitBox
// }

// //move camera
// if win.Pressed(pixelgl.KeyA) {
// 	camPos.X -= camSpeed * dt
// }
// if win.Pressed(pixelgl.KeyD) {
// 	camPos.X += camSpeed * dt
// }
// if win.Pressed(pixelgl.KeyS) {
// 	camPos.Y -= camSpeed * dt
// }
// if win.Pressed(pixelgl.KeyW) {
// 	camPos.Y += camSpeed * dt
// }

// //zoom camera
// camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

// //used for framerate test
// if win.Pressed(pixelgl.MouseButtonLeft) {
// 	if win.Pressed(pixelgl.KeyLeftShift) {
// 		mouse := cam.Unproject(win.MousePosition())
// 		switch objectToPlace.(type) {
// 		case *livingObject:
// 			{
// 				gameObjs = gameObjs.appendLivingObject(pinkAnimKeys, pinkAnims, pinkSheet, mouse)
// 			}
// 		case *gibletObject:
// 			{
// 				gameObjs = gameObjs.appendGibletObject(gibletAnimKeys, gibletAnims, gibletSheet, mouse)
// 			}
// 		}
// 	}
// }
