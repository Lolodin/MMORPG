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
                this.scene.ID.x = players[i].x
                this.scene.ID.y = players[i].y
            }
            if (!this.activePlayers[players[i].Name] && players[i].Name !=  this.scene.ID.name) {



                this.activePlayers[players[i].Name] =  this.scene.add.container(players[i].x,players[i].y)
                let player =   this.scene.add.image(0,0, "Player")
                let Text =  this.scene.add.text(-players[i].Name.length*5,-32,players[i].Name, {fontFamily: 'Arial'})
                this.activePlayers[players[i].Name].add(player)
                this.activePlayers[players[i].Name].add(Text)
            }
            else if (this.activePlayers[players[i].Name] && players[i].Name !=  this.scene.ID.name) {


                this.activePlayers[players[i].Name].setDepth(players[i].y)
                this.activePlayers[players[i].Name].x =players[i].x
                this.activePlayers[players[i].Name].y =players[i].y

            }

        }

    }


}