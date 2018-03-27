package handlers

/*
 Copyright (C) 2017 Ulbora Labs Inc. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs Inc., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import (
	//env "UlboraApiGateway/environment"
	mgr "UlboraApiGateway/managers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// //HandleGetRouteStatus HandleGetRouteStatus
// func (h Handler) HandleGetRouteStatus(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "GET":
// 		var gwr mgr.GatewayRoutes
// 		gwr.GwDB.DbConfig = h.DbConfig
// 		cid := r.Header.Get("u-client-id")
// 		gwr.ClientID, _ = strconv.ParseInt((cid), 10, 0)
// 		//gwr.APIKey = r.Header.Get("u-api-key")
// 		gwr.GwCacheHost = getCacheHost()
// 		w.Header().Set("Content-Type", "application/json")
// 		vars := mux.Vars(r)
// 		var route string
// 		if vars != nil {
// 			route = vars["route"]
// 		} else {
// 			route = r.URL.Query().Get("route")
// 		}
// 		gwr.Route = route
// 		res := gwr.GetGatewayRouteStatus()
// 		resJSON, err := json.Marshal(res)
// 		fmt.Print("json out: ")
// 		fmt.Println(res)
// 		if err != nil {
// 			log.Println(err.Error())
// 			//http.Error(w, "json output failed", http.StatusInternalServerError)
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprint(w, string(resJSON))
// 	default:
// 		w.WriteHeader(http.StatusNotFound)
// 	}
// }

// //HandleDeleteRouteStatus HandleDeleteRouteStatus
// func (h Handler) HandleDeleteRouteStatus(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "DELETE":
// 		var gwr mgr.GatewayRoutes
// 		gwr.GwDB.DbConfig = h.DbConfig
// 		cid := r.Header.Get("u-client-id")
// 		gwr.ClientID, _ = strconv.ParseInt((cid), 10, 0)
// 		gwr.APIKey = r.Header.Get("u-api-key")
// 		gwr.GwCacheHost = getCacheHost()
// 		w.Header().Set("Content-Type", "application/json")
// 		vars := mux.Vars(r)
// 		var route string
// 		if vars != nil {
// 			route = vars["route"]
// 		} else {
// 			route = r.URL.Query().Get("route")
// 		}
// 		gwr.Route = route
// 		res := gwr.DeleteGatewayRouteStatus()
// 		resJSON, err := json.Marshal(res)
// 		//fmt.Print("json out: ")
// 		//fmt.Println(res)
// 		if err != nil {
// 			log.Println(err.Error())
// 			//http.Error(w, "json output failed", http.StatusInternalServerError)
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprint(w, string(resJSON))
// 	default:
// 		w.WriteHeader(http.StatusNotFound)
// 	}
// }

//HandleGetClusterGwRoutes HandleGetClusterGwRoutes
func (h Handler) HandleGetClusterGwRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var gwr mgr.GatewayRoutes
		gwr.GwDB.DbConfig = h.DbConfig
		gwr.GwCacheHost = getCacheHost()
		cid := r.Header.Get("u-client-id")
		gwr.ClientID, _ = strconv.ParseInt((cid), 10, 0)
		gwr.APIKey = r.Header.Get("u-api-key")

		//gwr.GwCacheHost = env.GetCacheHost()
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		var route string
		if vars != nil {
			route = vars["route"]
		} else {
			route = r.URL.Query().Get("route")
		}
		gwr.Route = route
		res := gwr.GetClusterGwRoutes()
		resJSON, err := json.Marshal(res)
		fmt.Print("json out: ")
		fmt.Println(res)
		if err != nil {
			log.Println(err.Error())
			//http.Error(w, "json output failed", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(resJSON))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
