import {TopDownClient} from "./TopDownClient.js";
// import {PlayerScene} from "./ContentIso/PlayerScene.js";
import {GameMenu} from "./GameMenu.js";


let config = {
    /*eslint no-undef:0*/
    scale: {
        // mode: Phaser.Scale.ENVELOP,
        // autoCenter: Phaser.Scale.CENTER_BOTH,
        width: 800   ,
        height: 600
    },
    dom: {
        createContainer: true
    },
    disableContextMenu: true,
    background: 'black',
    physics: {
        default: 'arcade',
        arcadePhysics: {
            overlapBias: 1
        }
    },
    scene:[TopDownClient],
    pixelArt: true,
    roundPixels: true,
    antialias: true

}
/*eslint no-undef:0*/
/*eslint no-unused-vars:0*/
let game = new Phaser.Game(config);
