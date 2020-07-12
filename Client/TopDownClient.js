
    export {TopDownClient}

class TopDownClient extends Phaser.Scene{
    constructor() {
        super({key: "SceneMain"})
        this.GetServerMap()
        this.CurrentChunk = [1,1]//Начальный чанк
        this.coordinate = this.getCurrentMap(this.CurrentChunk) // Массив Координат чанков на текущей карте
        this.tileSize = 16 // Размер тайтла
        this.chunkSize = 16 * this.tileSize // Размер чанка
        this.LoadChunks = [] // Загруженные чанки
        this.CurrentMap = [] // Текущая отрисованная карта
        this.activePlayers = [] //Активные игроки на карте
        this.Player = {} // Аватар игрока для обработки сервером
        this.map = 0//карта для отрисовки

}


    preload() {
        this.load.spritesheet('Water', 'Client/Content/sprWater.png', {
            frameHeight: 16,
            frameWidth: 16,
        });
        this.load.image('Mount', 'Client/Content/sprSand.png');
        this.load.image('Ground', 'Client/Content/sprGrass.png');
        this.load.image('Player', 'Client/Content/Player.png');
        this.load.image('Oak', 'Client/Content/Oak.png');
        this.load.image('Spruce', 'Client/Content/Spruce.png');
        this.Identification() // Идентификация игрока
            //Соединение с вебсокетом
        this.websocket = new WebSocket("ws://localhost:8080/player")
        this.websocket.onopen = function (e) {

        }



    }
    create() {
        this.anims.create({
            key: 'water',
            frames: this.anims.generateFrameNumbers('Water'),
            frameRate: 7,
            repeat: -1
        });

       this.P = this.add.image(8,8, "Player")
        this.Player.x = 8
        this.Player.y = 8
       this.P.setDepth(2)
       this.cameras.main.startFollow(this.P, true)
        this.GetServerMap(this.P.x, this.P.y)

        /*
        Рисуем игроков на игровой карте
         */
        this.websocket.onmessage = (e)=> {
            let players = e.data
            players = JSON.parse(players)
            console.log(players)
            this.DrawPlayers(players)
        }
        console.log(this.CurrentMap)

    }
    update(time, delta) {
// Блок управления
let cursors = this.input.keyboard.createCursorKeys();
if (cursors.left.isDown) {
    this.Player.x-=1
}
if (cursors.right.isDown) {
    this.Player.x+=1
        }
        if (cursors.up.isDown) {
            this.Player.y-=1
        }
        if (cursors.down.isDown) {
            this.Player.y+=1
        }


      let  nowChunk = this.getChunkID(this.P.x, this.P.y)
        if (nowChunk[0]!= this.CurrentChunk[0] || nowChunk[1]!=this.CurrentChunk[1]) {
            let newCoordinate = this.getCurrentMap(nowChunk)
            this.CurrentChunk = this.getChunkID(this.P.x, this.P.y)
            this.clearMap(newCoordinate)
            this.GetServerMap(this.P.x,this.P.y)
            this.coordinate = newCoordinate
        }
        //Посылаем данные о местоположении игрока на сервер
let playerData = {Name: this.Player.name, X: this.Player.x, Y: this.Player.y}
        JSON.stringify(playerData)
        this.websocket.send(JSON.stringify(playerData))
    }

    // Функции для взаимодействия с сервером
    //Идентификация
    Identification() {
        let pro = prompt("введите имя", "anon")
        this.Player.name = pro
    }
    // Получаем Игровую карту
    async GetServerMap(X, Y) {
        let Data = {x:X,y:Y, playerID:2}
        let request = await fetch("/map", {
            method: "POST",
            body: JSON.stringify(Data)

        } )
        request = request.json()
        request.then((data)=>
        {
            this.map =  data.CurrentMap
            this.DrawMap(this.map)
        })



    }

/*
Функции для рединга игры
 */
// Функция отрисовки карты
    DrawMap(map) {

        map.forEach(m =>this.DrawChunc(m)
        )
    }
    //Отрисовка чанка и игровых объектов на нем
    DrawChunc(chunck) {

        if (this.LoadChunks[chunck.ChunkID] == true) {
            return
        }
        this.CurrentMap[chunck.ChunkID] = this.add.group()
        this.LoadChunks[chunck.ChunkID] = true

        for (let item in chunck.Map) {
            {
                let tile
                let coordinate = this.cartesianToIsometric(chunck.Map[item].X, chunck.Map[item].Y)
                if (chunck.Map[item].key == "Water") {
                    tile = this.add.sprite(coordinate.x, coordinate.y, "Water")
                    tile.play("water")
                } else{
                    tile = this.add.image(coordinate.x, coordinate.y, chunck.Map[item].key)
                }


                this.CurrentMap[chunck.ChunkID].add(tile)
            }

        }
        for (let item in chunck.Tree) {


                 let tileTree
            let coordinate = this.cartesianToIsometric(chunck.Tree[item].X, chunck.Tree[item].Y)
                 tileTree = this.add.image(coordinate.x, coordinate.y, chunck.Tree[item].tree)
            tileTree.setDepth(2)
                 this.CurrentMap[chunck.ChunkID].add(tileTree)


        }
    }
    cartesianToIsometric(cartPt){
        let tempPt=new Phaser.Geom.Point(cartPt.X-cartPt.Y,(cartPt.X+cartPt.Y)/2 );
        return (tempPt);
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

                try {
                    this.CurrentMap[c].clear(true, true)
                }catch (e) {

                }
            }

        }
    }
    //Отрисовка игроков
    DrawPlayers(players) {

        for (let i = 0; i<players.players.length; i++) {
            if (players.players[i].Name == this.Player.name) {
                this.P.x = players.players[i].X
                this.P.y = players.players[i].Y
            }

            if (!this.activePlayers[players.players[i].Name] && players.players[i].Name != this.Player.name) {
                this.activePlayers[players.players[i].Name] = this.add.container(players.players[i].X,players.players[i].Y)
                let player =  this.add.image(0,0, "Player")
                let Text = this.add.text(-players.players[i].Name.length*5,-23,players.players[i].Name)
                this.activePlayers[players.players[i].Name].setDepth(2)
                this.activePlayers[players.players[i].Name].add(player)
                this.activePlayers[players.players[i].Name].add(Text)
            } else if(this.activePlayers[players.players[i].Name] && players.players[i].Name != this.Player.name) {


                this.activePlayers[players.players[i].Name].x =players.players[i].X
                this.activePlayers[players.players[i].Name].y =players.players[i].Y
            }}




    }

    //Вспомогательные функции для работы с картой и координатами
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






}