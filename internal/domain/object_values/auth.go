package objectValues

type Auth struct {
	Data struct {
		Nombres   string `json:"nombres"`
		ID        uint64 `json:"id"`
		Apellidos string `json:"apellidos"`
		Correo    string `json:"correo"`
		Exp       int    `json:"exp"`
	} `json:"data"`
	Valid bool `json:"valid"`
}
