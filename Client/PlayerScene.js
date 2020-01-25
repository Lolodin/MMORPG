export {PlayerScene}
class PlayerScene extends Phaser.Scene {
    constructor() {
        super("PlayerScene")
        this.CurrentMap = [] // Текущая отрисованная карта, которая добавлена в группу
        this.LoadChunks = []// загруженные чанки
        this.coordinate = 0
        this.CurrentChunk =0
        this.tileSize = 32
        this.chunkSize =  this.tileSize *  this.tileSize

            }

    preload(){
        this.load.image('Sand', 'Client/ContentIso/Sand.png');
        this.load.image('Ground', 'Client/ContentIso/sprGrass.png');
        this.load.spritesheet('Water', 'Client/ContentIso/Water.png', {
            frameHeight: 32,
            frameWidth: 64,
        });
        this.load.image('Player', 'Client/ContentIso/Player.png');


    }
    create(){
        this.input.on('gameobjectup', function (pointer, gameObject)
        {
            gameObject.emit('clicked', gameObject);
        }, this);
        this.anims.create({
            key: 'water',
            frames: this.anims.generateFrameNumbers('Water'),
            frameRate: 2,
            repeat: -1
        });
        console.log("start")
        this.GetServerMap(0,0)
        this.Player = this.add.image(0,0, "Player" )
        console.log(this.Player, "pl")
        this.Player.setDepth(2)
        this.CurrentChunk =  this.getChunkID(this.Player.x, this.Player.y)
        this.cameras.main.startFollow(this.Player, true)
        this.coordinate = this.getCurrentMap(this.CurrentChunk)

    }
    update(time, delta){
        let cursors = this.input.keyboard.createCursorKeys();
        if (cursors.left.isDown) {
            let cartesianCoord = this.isometricTocartesian({X:this.Player.x, Y: this.Player.y})
            cartesianCoord.x-=5
            let coordinatePlayer = this.cartesianToIsometric({X:cartesianCoord.x, Y: cartesianCoord.y})
            this.Player.x = coordinatePlayer.x
            this.Player.y = coordinatePlayer.y



        }
        if (cursors.right.isDown) {
            let cartesianCoord = this.isometricTocartesian({X:this.Player.x, Y: this.Player.y})
            cartesianCoord.x+=5
            let coordinatePlayer = this.cartesianToIsometric({X:cartesianCoord.x, Y: cartesianCoord.y})
            this.Player.x = coordinatePlayer.x
            this.Player.y = coordinatePlayer.y
        }
        if (cursors.up.isDown) {
            let cartesianCoord = this.isometricTocartesian({X:this.Player.x, Y: this.Player.y})
            cartesianCoord.y-=5
            let coordinatePlayer = this.cartesianToIsometric({X:cartesianCoord.x, Y: cartesianCoord.y})
            this.Player.x = coordinatePlayer.x
            this.Player.y = coordinatePlayer.y
        }
        if (cursors.down.isDown) {
            let cartesianCoord = this.isometricTocartesian({X:this.Player.x, Y: this.Player.y})
            cartesianCoord.y+=5
            let coordinatePlayer = this.cartesianToIsometric({X:cartesianCoord.x, Y: cartesianCoord.y})
            this.Player.x = coordinatePlayer.x
            this.Player.y = coordinatePlayer.y
        }

        let  nowChunk = this.getChunkID(this.Player.x, this.Player.y)
        if (nowChunk[0]!= this.CurrentChunk[0] || nowChunk[1]!=this.CurrentChunk[1]) {
            let newCoordinate = this.getCurrentMap(nowChunk)
            this.CurrentChunk =  this.getChunkID(this.Player.x, this.Player.y)
            this.clearMap(newCoordinate)
            this.coordinate = newCoordinate
            let cartesianCoord = this.isometricTocartesian({X:this.Player.x, Y: this.Player.y})
            this.GetServerMap(cartesianCoord.x, cartesianCoord.y)

        }
    }


//Работа с картой и координатами
    async GetServerMap(X, Y) {
        let Data = {x:X,y:Y, playerID:2}
        let request = await fetch("/map", {
            method: "POST",
            body: JSON.stringify(Data)

        } )
        request = await request.json() // request.CurrentMap[9].Map
     /*
     request = [9]Map, Map = Map["8,8"]{ Grass, X = 8, Y= 8}
      */
      this.drawMapController(request)

    }
    drawMapController(requstMapServer) {
        requstMapServer.CurrentMap.forEach((chunk =>this.drawTileChunk(chunk.Map, chunk.ChunkID) ))
    }
    drawTileChunk(chunk, chunkID) {
        // Check Chunk is Load
        if (this.LoadChunks[chunkID] == true) {
            return
        }
        console.log(this.LoadChunks)
        // add Chunk Group for tiles
        // Load chunk true
        this.CurrentMap[chunkID] = this.add.group()
        this.LoadChunks[chunkID] = true
for (let coordTile in chunk) {
    let tile
    let coordinate = this.cartesianToIsometric(chunk[coordTile])
    if(chunk[coordTile].key == "Water") {
       tile = this.add.sprite(coordinate.x, coordinate.y,chunk[coordTile].key)
       tile.play("water")
    } else {
       tile = this.add.image(coordinate.x, coordinate.y,chunk[coordTile].key )
    }
tile.setInteractive()
    tile.on('clicked', (tile)=>console.log(this.isometricTocartesian({X:tile.x, Y:tile.y})), this)

// add tile in ChunkGroup
    this.CurrentMap[chunkID].add(tile)

}
    }
    isometricTocartesian(cartIso) {
        let tempISO = new Phaser.Geom.Point((2*cartIso.Y + cartIso.X)/2,(2*cartIso.Y - cartIso.X)/2 )
        return tempISO
    }
    cartesianToIsometric(cartPt){
        let tempPt=new Phaser.Geom.Point(cartPt.X-cartPt.Y,(cartPt.X+cartPt.Y)/2 );
        return (tempPt);
    }
    //Возвращает карту чанка
    getCurrentMap(currentChunk) {
        let map = [];
        let coordinateX = currentChunk[0]*this.chunkSize;
        let coordinateY = currentChunk[1]*this.chunkSize;


        map.push(currentChunk);
        let x = coordinateX +this.chunkSize;
        let y = coordinateY +this.chunkSize;
        let xy = this.getChunkID(x,y);

        map.push(xy);

        x = coordinateX + this.chunkSize;
        y = coordinateY;
        xy = this.getChunkID(x,y);

        map.push(xy);
        if (coordinateY<0)
        {
            x = coordinateX + this.chunkSize;
            y = coordinateY - this.chunkSize;
        }
        else {
            x = coordinateX + this.chunkSize;
            y = coordinateY - this.chunkSize - 1;
        }

        xy = this.getChunkID(x,y);
        map.push(xy);
        x = coordinateX;
        y = coordinateY +this.chunkSize;
        xy = this.getChunkID(x,y);

        map.push(xy);
        if (coordinateY<0)
        {
            x = coordinateX;
            y = coordinateY-this.chunkSize;
        }
        else {
            x = coordinateX;
            y = coordinateY - this.chunkSize - 1;
        }

        xy = this.getChunkID(x,y);
        map.push(xy);
        if (coordinateX<0)
        {
            x = coordinateX -this.chunkSize;
            y = coordinateY +this.chunkSize;
        }
        else {
            x = coordinateX - this.chunkSize - 1;
            y = coordinateY + this.chunkSize;
        }

        xy = this.getChunkID(x,y);
        map.push(xy);
        if (coordinateX<0)
        {
            x = coordinateX -this.chunkSize;
            y = coordinateY;
        }
        else {
            x = coordinateX - this.chunkSize - 1;
            y = coordinateY;
        }

        xy = this.getChunkID(x,y);
        map.push(xy);
        if (coordinateX<0 && coordinateY<0)
        {
            x = coordinateX -this.chunkSize;
            y = coordinateY-this.chunkSize;
        }
        else {
            if (coordinateX>0)
            {
                x = coordinateX - this.chunkSize-1;
            }else
            {
                x = coordinateX - this.chunkSize;
            }
            if (coordinateY<0)
            {
                y = coordinateY - this.chunkSize;
            }
            else
            {
                y = coordinateY - this.chunkSize - 1;
            }
        }

        xy = this.getChunkID(x,y);
        map.push(xy);
        return map;

    }
    getChunkID(x, y) {
        let tileX = Math.fround(x/this.tileSize);
        let tileY = Math.fround(y/this.tileSize);
        let chunkX = null;
        let chunkY = null;
        if (tileX<0)
        {
            chunkX = Math.floor(tileX/this.tileSize)
        }
        else
        {
            chunkX = Math.ceil(tileX/this.tileSize);
        }
        if (tileY<0)
        {
            chunkY = Math.floor(tileY/this.tileSize)
        }
        else
        {
            chunkY = Math.ceil(tileY/this.tileSize);
        }
        if (tileX===0)
        {
            chunkX=1;
        }
        if (tileY===0)
        {
            chunkY=1;
        }
        return [chunkX, chunkY];
    }
    //Очистка карты
    clearMap(newCoordinate) {
        for (let i = 0; i<this.coordinate.length;i++) {
            let chunkIsNotExist = true
            newCoordinate.forEach((v) => {

                if (this.coordinate[i][0]==v[0] && this.coordinate[i][1]==v[1]) {

                    chunkIsNotExist = false
                }
            })


            if (chunkIsNotExist) {
                let c = this.coordinate[i][0]+","+this.coordinate[i][1]
                this.LoadChunks[c] = false
                console.log(this.CurrentMap, c, "coordinate group", this.coordinate)

                try {
                    this.CurrentMap[c].clear(true, true)
                }catch (e) {

                }
            }

        }
    }
}