export {Identification}
class Identification {
    constructor(scene) {
        this.Scene = scene
    }
// Идентификация на сервере, получение данных персонажа
  async  inServer() {
        let user = prompt("Введите имя вашего персонажа", "Anon")
        let pass = prompt("Введите пароль", "123456")
        this.Scene.ID.Name = user
        this.Scene.ID.Pass = pass

        let data = {name: this.Scene.ID.Name, password: this.Scene.ID.Pass}
        let res = await fetch("/init", {
            method: "POST",
            body: JSON.stringify(data)
        })
        res = await res.json()
      console.log(res)
        if (res.error == "null") {
            this.Scene.ID.x = res.x
            this.Scene.ID.y = res.y
        } else {
            alert("Error server")
        }
    }



}