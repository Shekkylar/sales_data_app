package routes

func Init() {

	r := NewRouter()
	port := "localhost:7071"
	r.Run(port)
}
