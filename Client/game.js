import {MainScene} from "./mainScene.js";


let config = {
    /*eslint no-undef:0*/
    type: Phaser.AUTO,
    width: 1800,
    height:600,
    background: 'black',
    physics: {
        default: 'arcade',
        arcadePhysics: {
            Gravity: {x: 0, y: 0}
        }
    },
    scene:[MainScene],
    pixelArt: true,
    roundPixels: true

}
/*eslint no-undef:0*/
/*eslint no-unused-vars:0*/
let game = new Phaser.Game(config);
console.log(MainScene);