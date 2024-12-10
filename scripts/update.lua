--[[

BUILT-IN VARIABLES:

DIRECTION = {
    NONE = 0,
    EAST = 1,
    NORTH = 2,
    WEST = 3
    SOUTH = 4
}

BUILT-IN FUNCTIONS:

length(table) --> length: number
getInstanceIds() --> instanceIds: table
getInstancePosition(instanceId) --> x: number, y: number
getInstanceType(instanceId) --> type: string
getDirection() --> direction: number
setPlayerDirection(direction) --> nil
setTPS(tps) --> nil
isWall(x, y) --> isWall: boolean

]]

local pId = getIdsByType("player")[1]
local pX, pY = getInstancePosition(pId)

if step == 0 and pX == 8 then
    setPlayerDirection(DIRECTION.SOUTH)
    step = step + 1
elseif step == 1 and pY == 3 then
    setPlayerDirection(DIRECTION.WEST)
    step = step + 1
elseif step == 2 and pX == 4 then
    setPlayerDirection(DIRECTION.NONE)
    step = step + 1
end
