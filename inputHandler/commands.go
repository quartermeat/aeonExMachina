package input

import (
	"fmt"
	"sync"

	"github.com/faiface/pixel"
	"github.com/quartermeat/aeonExMachina/assets"
	objects "github.com/quartermeat/aeonExMachina/gameObjects"
)

//Commands is the map of commands to execute
type Commands map[string]ICommand

//ICommand interface is used to execute game commands
type ICommand interface {
	execute(*sync.WaitGroup)
}

//ExecuteCommands executes the queued list of commands
func (commands Commands) ExecuteCommands(waitGroup *sync.WaitGroup) {
	for key, value := range commands {
		fmt.Printf("executing: %s\n", key)
		waitGroup.Add(1)
		go value.execute(waitGroup)
		delete(commands, key)
	}
}

//#region ADD OBJECT COMMAND

type addObjectAtPositionCommand struct {
	gameObjs      *objects.GameObjects
	objectToPlace objects.IGameObject
	position      pixel.Vec
	objectAssets  assets.ObjectAssets
}

func (command *addObjectAtPositionCommand) execute(waitGroup *sync.WaitGroup) {
	switch command.objectToPlace.(type) {
	case *objects.LivingObject:
		{
			*command.gameObjs = command.gameObjs.AppendLivingObject(command.objectAssets, command.position)
		}
	case *objects.GibletObject:
		{
			*command.gameObjs = command.gameObjs.AppendGibletObject(command.objectAssets, command.position)
		}
	}

	waitGroup.Done()
}

//AddObjectAtPosition allows for the addition of a game object
func AddObjectAtPosition(objs *objects.GameObjects, newObject objects.IGameObject, newPosition pixel.Vec) ICommand {
	return &addObjectAtPositionCommand{
		gameObjs:      objs,
		position:      newPosition,
		objectToPlace: newObject,
		objectAssets:  newObject.GetAssets(),
	}
}

//#endregion

//#region REMOVE OBJECT COMMAND

type removeObjectAtPositionCommand struct {
	gameObjs *objects.GameObjects
	position pixel.Vec
}

func (command *removeObjectAtPositionCommand) execute(waitGroup *sync.WaitGroup) {
	_, index, hit, err := command.gameObjs.GetSelectedGameObjAtPosition(command.position)
	if err != nil {
		panic(err)
	}
	if hit {
		*command.gameObjs = command.gameObjs.FastRemoveIndex(index)
	}

	waitGroup.Done()
}

//RemoveObjectAtPosition allows for the removal of a game Object based on Vec location
func RemoveObjectAtPosition(objs *objects.GameObjects, fromPosition pixel.Vec) ICommand {
	return &removeObjectAtPositionCommand{
		gameObjs: objs,
		position: fromPosition,
	}
}

//#endregion

//#region SELECT OBJECT COMMAND
type selectObjectAtPositionCommand struct {
	gameObjs *objects.GameObjects
	position pixel.Vec
}

func (command *selectObjectAtPositionCommand) execute(waitGroup *sync.WaitGroup) {
	selectedObj, _, hit, err := command.gameObjs.GetSelectedGameObjAtPosition(command.position)
	if err != nil {
		panic(err)
	}
	if hit {
		selectedObj.ChangeState(objects.SELECTED_IDLE)
	}

	waitGroup.Done()
}

//SelectObjectAtPosition allows for changing the state of a game object based on Vec location to selected
func SelectObjectAtPosition(objs *objects.GameObjects, fromPosition pixel.Vec) ICommand {
	return &selectObjectAtPositionCommand{
		gameObjs: objs,
		position: fromPosition,
	}
}

//#endregion

//#region MOVE SELECTED OBJECT COMMAND TO POSITION
type moveSelectedObjectToPositionCommand struct {
	gameObjs *objects.GameObjects
	position pixel.Vec
}

func (command *moveSelectedObjectToPositionCommand) execute(waitGroup *sync.WaitGroup) {
	for _, obj := range *command.gameObjs {
		if obj.GetState() == objects.SELECTED_IDLE {
			obj.MoveToPosition(command.position)
		}
	}
	waitGroup.Done()
}

//MoveSelectedObject allows for directing selected objects to move to a position
func MoveSelectedToPositionObject(objs *objects.GameObjects, fromPosition pixel.Vec) ICommand {
	return &moveSelectedObjectToPositionCommand{
		gameObjs: objs,
		position: fromPosition,
	}
}

//endregion
