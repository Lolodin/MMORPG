package WorldMap

type Player struct {
	Name string
	X int
	Y int
	WalkPath

}
type Players struct {
	P []Player `json:"players"`
}
func NewPlayer(n string, id int) Player {
	p:= Player{}
	p.X = 0
	p.Y = 0
	p.Name = n

	return p

}