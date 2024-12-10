package main

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	lua "github.com/yuin/gopher-lua"
	"image"
	_ "image/png"
	"io/fs"
	"lua-test/objects"
	"lua-test/tiles"
	"os/exec"
)

const screenWidth = 240
const screenHeight = 224
const screenScale = 3

const tileSize = 16

//go:embed assets/*
var assets embed.FS

func luaPrint(L *lua.LState) int {
	top := L.GetTop()
	for i := 1; i <= top; i++ {
		fmt.Print(L.ToStringMeta(L.Get(i)).String())
		if i != top {
			fmt.Print("\t")
		}
	}
	fmt.Println("")
	return 0
}

func luaBase64Encode(L *lua.LState) int {
	str := L.ToString(1)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("python", "commands/base64encode.py")
	cmd.Stdin = bytes.NewBufferString(str)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	print(stderr.String())
	if err == nil {
		L.Push(lua.LString(stdout.String()))
		return 1
	}
	return 0
}

func luaBase64Decode(L *lua.LState) int {
	str := L.ToString(1)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("python", "commands/base64decode.py")
	cmd.Stdin = bytes.NewBufferString(str)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	print(stderr.String())
	if err == nil {
		L.Push(lua.LString(stdout.String()))
		return 1
	}
	return 0
}

func loadImage(path string) *ebiten.Image {
	f, err := assets.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(f fs.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func init() {}

func main() {
	L := lua.NewState()
	defer L.Close()

	for _, pair := range []struct {
		n string
		f lua.LGFunction
	}{
		{lua.LoadLibName, lua.OpenPackage}, // Must be first
		{lua.TabLibName, lua.OpenTable},
		{lua.StringLibName, lua.OpenString},
	} {
		if err := L.CallByParam(lua.P{
			Fn:      L.NewFunction(pair.f),
			NRet:    0,
			Protect: true,
		}, lua.LString(pair.n)); err != nil {
			panic(err)
		}
	}

	L.SetGlobal("print", L.NewFunction(luaPrint))

	tilesets := map[string]*tiles.Tileset{
		"main": &tiles.Tileset{
			Image: loadImage("assets/tilesets/main.png"),
			Size:  tileSize,
			Aliases: map[string]tiles.Vector{
				"none":   {0, 0},
				"wall":   {1, 0},
				"player": {2, 0},
				"enemy":  {3, 0},
				"coin":   {4, 0},
			},
		},
	}

	tilemap := &tiles.Tilemap{Tileset: tilesets["main"]}
	tilemap.SetTiles(15, 14, [][]string{
		{"wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall"},
		{"wall", "none", "none", "none", "none", "none", "none", "none", "none", "none", "none", "none", "none", "none", "wall"},
		{"wall", "none", "wall", "wall", "wall", "wall", "wall", "wall", "none", "wall", "wall", "wall", "wall", "none", "wall"},
		{"wall", "none", "wall", "none", "none", "none", "none", "none", "none", "none", "none", "none", "wall", "none", "wall"},
		{"wall", "none", "wall", "none", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "none", "wall", "none", "wall"},
		{"wall", "none", "wall", "none", "none", "none", "wall", "none", "wall", "none", "wall", "none", "wall", "none", "wall"},
		{"wall", "none", "wall", "none", "wall", "none", "wall", "none", "wall", "none", "wall", "none", "wall", "none", "wall"},
		{"wall", "none", "wall", "none", "wall", "none", "wall", "none", "none", "none", "wall", "none", "wall", "none", "wall"},
		{"wall", "none", "wall", "none", "wall", "none", "wall", "none", "wall", "none", "none", "none", "wall", "none", "wall"},
		{"wall", "none", "wall", "none", "wall", "none", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "none", "wall"},
		{"wall", "none", "wall", "none", "wall", "none", "none", "none", "none", "none", "none", "none", "wall", "none", "wall"},
		{"wall", "none", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "none", "wall"},
		{"wall", "none", "none", "none", "none", "none", "none", "none", "none", "none", "none", "none", "none", "none", "wall"},
		{"wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall"},
	})

	instances := 1
	g := &Game{
		tilemap:      tilemap,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}

	coinLocations := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for y, row := range coinLocations {
		for x, coin := range row {
			if coin == 1 {
				g.objects = append(g.objects, &objects.Coin{
					Position:   tiles.Vector{X: x, Y: y},
					Tileset:    *tilesets["main"],
					InstanceId: instances,
					Alive:      true,
				})
				instances++
			}
		}
	}

	g.objects = append(g.objects,
		&objects.Enemy{
			Position:   tiles.Vector{X: 1, Y: 1},
			Direction:  objects.West,
			Tilemap:    tilemap,
			InstanceId: instances,
		},
		&objects.Enemy{
			Position:   tiles.Vector{X: 3, Y: 3},
			Direction:  objects.North,
			Tilemap:    tilemap,
			InstanceId: instances + 1,
		},
		&objects.Enemy{
			Position:   tiles.Vector{X: 9, Y: 5},
			Direction:  objects.North,
			Tilemap:    tilemap,
			InstanceId: instances + 2,
		},
	)
	instances += 3

	player := &objects.Player{
		Position:   tiles.Vector{X: 13, Y: 1},
		Direction:  objects.None,
		Tilemap:    tilemap,
		InstanceId: instances,
		L:          L,
		Alive:      true,
	}
	g.objects = append(g.objects, player)
	instances++

	directionsTable := L.NewTable()
	directionsTable.RawSetString("NONE", lua.LNumber(objects.None))
	directionsTable.RawSetString("EAST", lua.LNumber(objects.East))
	directionsTable.RawSetString("NORTH", lua.LNumber(objects.North))
	directionsTable.RawSetString("WEST", lua.LNumber(objects.West))
	directionsTable.RawSetString("SOUTH", lua.LNumber(objects.South))
	L.SetGlobal("DIRECTION", directionsTable)

	L.SetGlobal("length", L.NewFunction(func(L *lua.LState) int {
		table := L.ToTable(1)
		L.Push(lua.LNumber(table.Len()))
		return 1
	}))

	L.SetGlobal("getInstancePosition", L.NewFunction(func(L *lua.LState) int {
		instanceId := L.ToInt(1)
		for _, object := range g.objects {
			if object.GetInstanceId() == instanceId {
				position := object.(objects.Object).GetPosition()
				L.Push(lua.LNumber(position.X))
				L.Push(lua.LNumber(position.Y))
				return 2
			}
		}
		return 0
	}))

	L.SetGlobal("getInstanceIds", L.NewFunction(func(L *lua.LState) int {
		ids := &lua.LTable{}
		for _, object := range g.objects {
			ids.Append(lua.LNumber(object.GetInstanceId()))
		}
		L.Push(ids)
		return 1
	}))

	L.SetGlobal("getInstanceType", L.NewFunction(func(L *lua.LState) int {
		instanceId := L.ToInt(1)
		for _, object := range g.objects {
			if object.GetInstanceId() == instanceId {
				switch object.(type) {
				case *objects.Player:
					L.Push(lua.LString("player"))
				case *objects.Enemy:
					L.Push(lua.LString("enemy"))
				case *objects.Coin:
					L.Push(lua.LString("coin"))
				}
				return 1
			}
		}
		return 0
	}))

	L.SetGlobal("getDirection", L.NewFunction(func(L *lua.LState) int {
		instanceId := L.ToInt(1)
		for _, object := range g.objects {
			if object.GetInstanceId() == instanceId {
				L.Push(lua.LNumber(object.GetDirection()))
				return 1
			}
		}
		return 0
	}))

	L.SetGlobal("setPlayerDirection", L.NewFunction(func(L *lua.LState) int {
		direction := L.ToInt(1)
		player.Direction = objects.Direction(direction)
		return 0
	}))

	L.SetGlobal("setTPS", L.NewFunction(func(L *lua.LState) int {
		tps := L.ToInt(1)
		if tps < 1 {
			tps = 1
		} else if tps > 30 {
			tps = 60
		}
		ebiten.SetTPS(tps)
		return 0
	}))

	L.SetGlobal("isWall", L.NewFunction(func(L *lua.LState) int {
		x := L.ToInt(1)
		y := L.ToInt(2)
		if x < 0 || x >= 15 || y < 0 || y >= 14 {
			L.Push(lua.LBool(true))
			return 1
		}
		tile := tilemap.GetTile(x, y)
		L.Push(lua.LBool(tile == "wall"))
		return 1
	}))

	ebiten.SetWindowSize(screenWidth*screenScale, screenHeight*screenScale)
	ebiten.SetWindowTitle("Lua Game")
	ebiten.SetTPS(2)

	if err := L.DoFile("scripts/init.lua"); err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
