//import {MainScene} from "./mainScene.js";
import {PlayerScene} from "./PlayerScene.js";


let config = {
    /*eslint no-undef:0*/
    type: Phaser.AUTO,
    width: 1278,
    height:800,
    background: 'black',
    physics: {
        default: 'arcade',
        arcadePhysics: {
            Gravity: {x: 0, y: 0}
        }
    },
    scene:[PlayerScene],
    pixelArt: true,
    roundPixels: true

}
/*eslint no-undef:0*/
/*eslint no-unused-vars:0*/
let game = new Phaser.Game(config);
