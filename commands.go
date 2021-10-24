package main

import (
	"fmt"
	"sync"

	"github.com/faiface/pixel"
)

//Commands is the map of commands to execute
type Commands map[string]ICommand

//ICommand interface is used to execute game commands
type ICommand interface {
	execute(*sync.WaitGroup)
}

//concurrently execute queued game commands
func (commands Commands) executeCommands(waitGroup *sync.WaitGroup) {
	for key, value := range commands {
		fmt.Printf("executing: %s\n", key)
		waitGroup.Add(1)
		go value.execute(waitGroup)
		delete(commands, key)
	}
}

type addObjectCommand struct {
	gameObjs      *GameObjects
	objectToPlace IGameObject
	position      pixel.Vec
	objectAssets  ObjectAssets
}

func (command *addObjectCommand) execute(waitGroup *sync.WaitGroup) {
	switch command.objectToPlace.(type) {
	case *livingObject:
		{
			*command.gameObjs = command.gameObjs.appendLivingObject(command.objectAssets, command.position)
		}
	case *GibletObject:
		{
			*command.gameObjs = command.gameObjs.appendGibletObject(command.objectAssets, command.position)
		}
	}

	waitGroup.Done()
}

//AddObject allows for the addition of a game object
func (objects *GameObjects) AddObject(newObject IGameObject, newPosition pixel.Vec) ICommand {
	return &addObjectCommand{
		gameObjs:      objects,
		position:      newPosition,
		objectToPlace: newObject,
		objectAssets:  newObject.GetAssets(),
	}
}

type removeObjectCommand struct {
	gameObjs *GameObjects
	position pixel.Vec
}

func (command *removeObjectCommand) execute(waitGroup *sync.WaitGroup) {
	selectedObj, index, hit, err := command.gameObjs.getSelectedGameObj(command.position)
	if err != nil {
		panic(err)
	}
	if hit {
		fmt.Println("object id:", selectedObj.getID(), " removed")
		*command.gameObjs = command.gameObjs.fastRemoveIndex(index)
	} else {
		fmt.Println("RemoveObjectCommmand: no object selected")
	}
	hit = false

	waitGroup.Done()
}

//RemoveObject allows for the removal of a game Object based on Vec location
func (objects *GameObjects) RemoveObject(fromPosition pixel.Vec) ICommand {
	return &removeObjectCommand{
		gameObjs: objects,
		position: fromPosition,
	}
}
