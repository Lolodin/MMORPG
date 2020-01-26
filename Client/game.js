//import {MainScene} from "./mainScene.js";
import {PlayerScene} from "./PlayerScene.js";


let config = {
    /*eslint no-undef:0*/
    type: Phaser.AUTO,
    width: 1370,
    height:800,
    disableContextMenu: true,
    background: 'black',
    physics: {
        default: 'arcade',
        arcadePhysics: {
            overlapBias: 1
        }
    },
    scene:[PlayerScene],
    pixelArt: false,
    roundPixels: false,
    antialias: false

}
/*eslint no-undef:0*/
/*eslint no-unused-vars:0*/
let game = new Phaser.Game(config);
