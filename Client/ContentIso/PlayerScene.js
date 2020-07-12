export {PlayerScene}
import {Identification} from "./Identification.js";
import {Players} from "./Players.js";


class PlayerScene extends Phaser.Scene {
    constructor() {
        super("PlayerScene")
        this.ID = {} // Аватар игрока для взаимодействия с сервером
        this.CurrentMap = [] // Текущая отрисованная карта, которая добавлена в группу
        this.LoadChunks = []// загруженные чанки
        this.LoadChunksTree = [] // Деревья
        this.coordinate = 0
        this.CurrentChunk =0
        this.CHUNKIDSIZE = 32
        this.chunkSize =  32 *  32
        this.targetPath = [16,16] // Путь куда должен двигаться персонаж
        //Константы для отправки запроса




            }

    preload(){
        this.load.image('Oak', 'Client/ContentIso/Oak.png');
        this.load.image('Spruce', 'Client/ContentIso/Spruce.png');
        this.load.image('Sand', 'Client/ContentIso/Sand.png');
        this.load.image('Ground', 'Client/ContentIso/sprGrass.png');
        this.load.spritesheet('Water', 'Client/ContentIso/Water.png', {
            frameHeight: 32,
            frameWidth: 64,
        });
        this.load.spritesheet('Player', 'Client/ContentIso/Player.png', {
            frameHeight: 64, frameWidth: 64
        });


        let ident = new Identification(this)
        ident.inServer()
        this.websocket = new WebSocket("ws://localhost:8080/player")
        this.websocket.onopen = (e) => {
            console.log("OPEN", e)
        }






    }
    create() {
        this.anims.create({
            key: 'water',
            frames: this.anims.generateFrameNumbers('Water'),
            frameRate: 5,
            repeat: -1
        });
        this.Players = new Players(this)
        this.input.on('gameobjectup', function (pointer, gameObject){
            gameObject.emit('clicked', gameObject);
        }, this);
        // Получаем текущую карту
        this.GetServerMap(this.ID.x,this.ID.x)
            // Координаты
        let coord = this.isometricTocartesian({x:this.ID.x,y:this.ID.y})
        this.CurrentChunk =  this.getChunkID(coord.x,coord.y)
       // this.Player = this.add.image(this.ID.x,this.ID.y, "Player")
        this.cameras.main.startFollow(this.ID, true)
        this.coordinate = this.getCurrentMap(this.CurrentChunk)
        this.websocket.onmessage = (e)=> {
          //  console.log("on message")
            let players = e.data
            players = JSON.parse(players)
          //  console.log(players)
            this.Players.DrawPlayer(players.players)
        }



    }
    update(time, delta){
        if (this.websocket.readyState === 1) {
            let playerData = {name: this.ID.Name,  x: this.targetPath[0], y: this.targetPath[1]}
            this.websocket.send(JSON.stringify(playerData))
        }
        let coord = this.isometricTocartesian({x:this.ID.x,y:this.ID.y})
        let  nowChunk = this.getChunkID(coord.x, coord.y)
        if (nowChunk[0]!= this.CurrentChunk[0] || nowChunk[1]!=this.CurrentChunk[1]) {
            let newCoordinate = this.getCurrentMap(nowChunk)
            this.CurrentChunk =  nowChunk
            this.clearMap(newCoordinate)
            this.coordinate = newCoordinate
            let cartesianCoord = this.isometricTocartesian({x:this.ID.x, y: this.ID.y})
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
      //  requstMapServer.CurrentMap.forEach((chunk =>this.drawTileChunk(chunk.Map, chunk.ChunkID) ))
        for (let i = 0; i<9;i++) {
            this.drawTileChunk(requstMapServer.CurrentMap[i].Map,  requstMapServer.CurrentMap[i].ChunkID)
            this.drawTree(requstMapServer.CurrentMap[i].Tree, requstMapServer.CurrentMap[i].ChunkID )
        }
    }
    drawTileChunk(chunk, chunkID) {
        // Check chunk is Load
        if (this.LoadChunks[chunkID] == true) {
            return
        }
        // add chunk Group for tiles
        // Load chunk true
        this.CurrentMap[chunkID] = this.add.group()
        this.LoadChunks[chunkID] = true
for (let coordTile in chunk) {
    let tile
    let coordinate = this.cartesianToIsometric(chunk[coordTile])
    if(chunk[coordTile].key == "Water") {
       tile = this.add.sprite(coordinate.x, coordinate.y, chunk[coordTile].key).play('water', true);



    } else {
       tile = this.add.image(coordinate.x, coordinate.y,chunk[coordTile].key )
    }
tile.setDepth(coordinate.y)
tile.setInteractive()
    tile.on('clicked', (tile)=>{
        tile.alpha = 0.5
        setTimeout(()=>tile.alpha =1, 1000)
        let coord = this.isometricTocartesian(tile)
       this.targetPath[0] = coord.x
       this.targetPath[1] = coord.y
        //console.log(this.targetPath)
    }, this)
   // tile.active = false

// add tile in ChunkGroup
    this.CurrentMap[chunkID].add(tile)

}
    }
    drawTree(chunk, chunkID) {
        if (this.LoadChunksTree[chunkID] == true) {
            return
        }
        this.LoadChunksTree[chunkID] = true
        for (let coordTile in chunk) {
            let tree
            let coordinate = this.cartesianToIsometric(chunk[coordTile])
             tree = this.add.image(coordinate.x, coordinate.y -64,chunk[coordTile].tree)
            tree.setDepth(coordinate.y)
            tree.setRotation(chunk[coordTile].age/5)

            //console.log(coordTile, coordinate, "coordTREEE")

            this.CurrentMap[chunkID].add(tree)
            }


// add tile in ChunkGroup


        }

    //cartIso {x: xx y: yy}
    isometricTocartesian(cartIso) {
        let tempISO = new Phaser.Geom.Point((2*cartIso.y + cartIso.x)/2,(2*cartIso.y - cartIso.x)/2 )
        return tempISO
    }
    //cattPt{x: xx y: yy}
    cartesianToIsometric(cartPt){
        let tempPt=new Phaser.Geom.Point(cartPt.x-cartPt.y,(cartPt.x+cartPt.y)/2 );
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
        let tileX = Math.fround(x/this.CHUNKIDSIZE);
        let tileY = Math.fround(y/this.CHUNKIDSIZE);
        let chunkX = null;
        let chunkY = null;
        if (tileX<0)
        {
            chunkX = Math.floor(tileX/this.CHUNKIDSIZE)
        }
        else
        {
            chunkX = Math.ceil(tileX/this.CHUNKIDSIZE);
        }
        if (tileY<0)
        {
            chunkY = Math.floor(tileY/this.CHUNKIDSIZE)
        }
        else
        {
            chunkY = Math.ceil(tileY/this.CHUNKIDSIZE);
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
                delete this.LoadChunks[c]
                delete this.LoadChunksTree[c]
                try {
                    this.CurrentMap[c].clear(true, true)
                    delete this.CurrentMap[c];
                }catch (e) {

                }
            }

        }
    }
}