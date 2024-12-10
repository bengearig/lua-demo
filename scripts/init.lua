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

function getIdsByType(type)
    ids = getInstanceIds()
    tIds = {}
    for i = 1, length(ids) do
        if getInstanceType(ids[i]) == type then
            table.insert(tIds, ids[i])
        end
    end
    return tIds
end

setPlayerDirection(DIRECTION["WEST"])
step = 0

