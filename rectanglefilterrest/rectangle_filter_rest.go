package main

import (
	"log"
	"net/http"
	"rectanglefilter"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	rectanglefilter.Init()
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/", getAllRectangleRest),
		rest.Post("/", addRectangleRest),
		rest.Get("/unique", getUniqueRectanglesRest),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func getAllRectangleRest(w rest.ResponseWriter, r *rest.Request) {
	rectangles := rectanglefilter.GetAllRectangle()
	w.WriteJson(&rectangles)
}

func addRectangleRest(w rest.ResponseWriter, r *rest.Request) {
	rectangles := rectanglefilter.AddRectangleSt{}
	err := r.DecodeJsonPayload(&rectangles)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rectanglefilter.AddRectangle(rectangles)
	w.WriteJson(&rectangles)
}

func getUniqueRectanglesRest(w rest.ResponseWriter, r *rest.Request) {
	rectangles := rectanglefilter.GetUniqueRectangles()
	if len(rectangles) != 0 {
		w.WriteJson(&rectangles[0])
	}
}
