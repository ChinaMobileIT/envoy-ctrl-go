package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful"

	"github.com/ChinaMobileIT/envoy-ctrl-go/internal/example"
	"github.com/ChinaMobileIT/envoy-ctrl-go/internal/example/cache"
)

func Start() {
	container := restful.NewContainer()
	ws := new(restful.WebService)
	ws.Route(ws.GET("/reset").To(reset).Produces(restful.MIME_JSON))
	ws.Route(ws.GET("/snapshot").To(snapshot).Produces(restful.MIME_JSON))
	ws.Route(ws.GET("/add/{upstreamHost}/{upstreamPort}/{listenerPort}").To(add).Produces(restful.MIME_JSON))
	container.Add(ws)
	log.Fatal(http.ListenAndServe(":9999", container))
}

// GET /reset
func reset(request *restful.Request, response *restful.Response) {
	snapshotCache := cache.GetCache()
	snapshot := example.Reset()
	if err := snapshotCache.SetSnapshot("test-id", snapshot); err != nil {
		response.WriteAsJson(fmt.Sprintf("snapshot error %q for %+v", err, snapshot))
	}
	response.WriteAsJson(snapshot)
}

// GET /snapshot
func snapshot(request *restful.Request, response *restful.Response) {
	snapshotCache := cache.GetCache()
	snapshot, err := snapshotCache.GetSnapshot("test-id")
	if err != nil {
		response.WriteEntity(err.Error())
		return
	}
	response.WriteAsJson(snapshot)
}

// GET /add
func add(request *restful.Request, response *restful.Response) {
	upstreamHost := request.PathParameter("upstreamHost")
	upstreamPort, _ := strconv.Atoi(request.PathParameter("upstreamPort"))
	listenerPort, _ := strconv.Atoi(request.PathParameter("listenerPort"))

	snapshotCache := cache.GetCache()
	snapshot := example.GenerateSnapshot(upstreamHost, uint32(upstreamPort), uint32(listenerPort))
	if err := snapshotCache.SetSnapshot("test-id", snapshot); err != nil {
		response.WriteAsJson(fmt.Sprintf("snapshot error %q for %+v", err, snapshot))
	}

	response.WriteAsJson(snapshot)
}
