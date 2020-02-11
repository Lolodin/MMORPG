//import {MainScene} from "./mainScene.js";
import {PlayerScene} from "./PlayerScene.js";


let config = {
    /*eslint no-undef:0*/
    scale: {
        mode: Phaser.Scale.FIT,
        width: 1536,
        height: 869
    },
    disableContextMenu: true,
    background: 'black',
    physics: {
        default: 'arcade',
        arcadePhysics: {
            overlapBias: 1
        }
    },
    scene:[PlayerScene],
    pixelArt: true,
    roundPixels: true,
    antialias: true

}
/*eslint no-undef:0*/
/*eslint no-unused-vars:0*/
let game = new Phaser.Game(config);
