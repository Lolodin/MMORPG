export {GameMap}
class GameMap {
    constructor(scene) {
        this.scene = scene
    }
    drawMapController(requstMapServer) {
        for (let i = 0; i<9;i++) {
            this.drawTileChunk(requstMapServer.CurrentMap[i].Map,  requstMapServer.CurrentMap[i].ChunkID)
            this.drawTree(requstMapServer.CurrentMap[i].Tree, requstMapServer.CurrentMap[i].ChunkID )
        }
    }
    drawTileChunk(chunk, chunkID) {
        // Check chunk is Load
        if (this.scene.LoadChunks[chunkID] == true) {
            return
        }
        // add chunk Group for tiles
        // Load chunk true
        this.scene.CurrentMap[chunkID] = this.scene.add.group()
        this.scene.LoadChunks[chunkID] = true
        for (let coordTile in chunk) {
            let tile

            if(chunk[coordTile].key == "Water") {
                tile = this.scene.add.sprite(chunk[coordTile].x, chunk[coordTile].y, chunk[coordTile].key).play('water', true);



            } else {
                tile = this.scene.add.image(chunk[coordTile].x, chunk[coordTile].y,chunk[coordTile].key )
            }
            tile.setDepth(1)
            tile.setInteractive()
            this.scene.CurrentMap[chunkID].add(tile)

        }
        console.log( this.scene.CurrentMap)
    }
    drawTree(chunk, chunkID) {
        if (this.scene.LoadChunksTree[chunkID] == true) {
            return
        }
        this.scene.LoadChunksTree[chunkID] = true
        for (let coordTile in chunk) {
            let tree
            tree = this.scene.add.image(chunk[coordTile].x, chunk[coordTile].y,chunk[coordTile].tree)
            tree.setDepth(chunk[coordTile].y+12)
            tree.setRotation(chunk[coordTile].age/5)
            this.scene.CurrentMap[chunkID].add(tree)
        }


    }
    clearMap(newCoordinate) {

        for (let i = 0; i<this.scene.coordinate.length;i++) {
            let chunkIsNotExist = true
            newCoordinate.forEach((v) => {
                if (this.scene.coordinate[i][0]==v[0] && this.scene.coordinate[i][1]==v[1]) {
                    chunkIsNotExist = false
                }
            })

            if (chunkIsNotExist) {
                let c = this.scene.coordinate[i][0]+","+this.scene.coordinate[i][1]
                delete this.scene.LoadChunks[c]
                delete this.scene.LoadChunksTree[c]
                console.log(this.scene.CurrentMap, this.scene.coordinate, c, "TEST")
                try {
                    this.scene.CurrentMap[c].clear(true, true)
                    delete this.scene.CurrentMap[c];
                } catch (e) {

                    console.log("Error clear map", e)
                }
            }

        }
    }
}