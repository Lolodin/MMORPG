import {Identification} from "./Identification.js";
import {Players} from "./Players.js";
import {GameMap} from "./Map.js";

export {TopDownClient}

class TopDownClient extends Phaser.Scene{
    constructor() {
        super({key: "SceneMain"})
        this.ID = {} // Аватар игрока для взаимодействия с сервером
        this.CurrentMap = [] // Текущая отрисованная карта, которая добавлена в группу
        this.LoadChunks = []// загруженные чанки
        this.LoadChunksTree = [] // Деревья
        this.coordinate = []
        this.CurrentChunk =0
        this.tileSize = 16
        this.chunkSize =  16 *  16
       // Путь куда должен двигаться персонаж

}


    preload() {
        this.load.spritesheet('Water', 'Client/Content/sprWater.png', {
            frameHeight: 16,
            frameWidth: 16,
        });
        this.load.image('Sand', 'Client/Content/sprSand.png');
        this.load.image('Ground', 'Client/Content/sprGrass.png');
        this.load.image('Player', 'Client/Content/Player.png');
        this.load.image('Oak', 'Client/Content/Oak.png');
        this.load.image('Spruce', 'Client/Content/Spruce.png');

        //Идентификация
        let ident = new Identification(this)
        ident.inServer()

        //Открываем соединение
        this.websocket = new WebSocket("ws://localhost:8080/player")
        this.websocket.onopen = (e) => {
            console.log("OPEN", e)
        }



    }
    create() {
        this.anims.create({
            key: 'water',
            frames: this.anims.generateFrameNumbers('Water'),
            frameRate: 7,
            repeat: -1
        });
        this.Players = new Players(this)
        this.Map = new GameMap(this)
        this.GetServerMap(this.ID.x,this.ID.x)

        this.CurrentChunk =  this.getChunkID(this.ID.x,this.ID.y)
        this.cameras.main.startFollow(this.ID, true)
        this.coordinate = this.getCurrentMap(this.CurrentChunk)
        this.websocket.onmessage = (e)=> {
            //  console.log("on message")
            let players = e.data
            players = JSON.parse(players)
            //  console.log(players)
            this.Players.DrawPlayer(players.players)
        }
        /*
        Рисуем игроков на игровой карте
         */
        this.websocket.onmessage = (e)=> {
            let players = e.data
            players = JSON.parse(players)
            console.log(players)
            this.Players.DrawPlayer(players.players)
        }
        console.log(this.CurrentMap)

    }
    update(time, delta) {
        let cursors = this.input.keyboard.createCursorKeys();
        if (cursors.left.isDown) {
            this.ID.x-=1
        }
        if (cursors.right.isDown) {
            this.ID.x+=1
        }
        if (cursors.up.isDown) {
            this.ID.y-=1
        }
        if (cursors.down.isDown) {
            this.ID.y+=1
        }


        if (this.websocket.readyState === 1) {
            let playerData = {name: this.ID.Name,  x: this.ID.x, y: this.ID.y}
            this.websocket.send(JSON.stringify(playerData))
        }



        let  nowChunk = this.getChunkID(this.ID.x, this.ID.y)
        if (nowChunk[0]!= this.CurrentChunk[0] || nowChunk[1]!=this.CurrentChunk[1]) {
            let newCoordinate = this.getCurrentMap(nowChunk)
            this.CurrentChunk =  nowChunk
            this.Map.clearMap(newCoordinate)
            this.coordinate = newCoordinate
            this.GetServerMap(this.ID.x, this.ID.y)
        }

    }
    // Получаем Игровую карту
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
        this.Map.drawMapController(request)

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
    //Возвращает карту чанка
    getCurrentMap(currentChunk) {
        let map = [];
        let coordinateX = currentChunk[0]* this.chunkSize;
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






}