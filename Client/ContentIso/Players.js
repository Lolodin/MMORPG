export {Players}
class Players extends Phaser.GameObjects.Sprite {
    constructor(scene) {
        super(scene)
        this.activePlayers = [] // Активные игроки на сцене
        this.createSceneAnimation(scene)
        this.scene = scene

    }
    createSceneAnimation(scene) {
        scene.anims.create({
            key: 'PlayerTurn',
            frames: scene.anims.generateFrameNumbers('Player'),
            frameRate: 2,
            repeat: -1
        });
    }
    DrawPlayer(players) {
        for (let i = 0; i<players.length; i++) {
            if (players[i].Name ==  this.scene.ID.Name) {
                let coord =  this.scene.cartesianToIsometric(players[i])
                this.scene.ID.x = coord.x
                this.scene.ID.y = coord.y
            }
            if (!this.activePlayers[players[i].Name] && players[i].Name !=  this.scene.ID.name) {


                let coord =  this.scene.cartesianToIsometric(players[i])
                this.activePlayers[players[i].Name] =  this.scene.add.container(coord.x,coord.y)
                let player =   this.scene.add.sprite(0,-32, "Player")
                player.play("PlayerTurn")
                let Text =  this.scene.add.text(-players[i].Name.length*5,-78,players[i].Name, {fontFamily: 'Arial'})






                this.activePlayers[players[i].Name].add(player)
                this.activePlayers[players[i].Name].add(Text)
            } else if(this.activePlayers[players[i].Name] && players[i].Name !=  this.scene.ID.name) {

                let coord =  this.scene.cartesianToIsometric(players[i])
                this.activePlayers[players[i].Name].setDepth(coord.y+1)
                this.activePlayers[players[i].Name].x =coord.x
                this.activePlayers[players[i].Name].y =coord.y

            }
        }
    }


}