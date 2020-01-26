export {Identification}
class Identification {
    constructor(scene) {
        this.Scene = scene
    }

    inServer() {
        let user = prompt("Введите имя вашего персонажа", "Anon")
        this.Scene.ID.Name = user

    }


}