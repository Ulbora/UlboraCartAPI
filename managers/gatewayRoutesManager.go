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
package managers

import (
	ch "UlboraApiGateway/cache"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
)

//GetGatewayRoutes route
func (gw *GatewayRoutes) GetGatewayRoutes(getActive bool, routeName string) *GatewayRouteURL {
	var rtnVal GatewayRouteURL
	var rtn = make([]GatewayRouteURL, 0)
	// check cache for saved value---------
	var cp ch.CProxy
	cp.Host = gw.GwCacheHost
	var cid = strconv.FormatInt(gw.ClientID, 10)
	var key = cid + ":" + gw.Route
	//fmt.Print("Key Used for cache: ")
	//fmt.Println(key)
	res := cp.Get(key)
	if res.Success == true {
		rJSON, err := b64.StdEncoding.DecodeString(res.Value)
		if err != nil {
			fmt.Println(err)
		} else {
			err := json.Unmarshal([]byte(rJSON), &rtn)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("Found Gateway route in cache.")
	} else {
		//read db
		fmt.Println("Routes not found in cache, reading db.")
		var a []interface{}
		a = append(a, gw.Route, gw.ClientID)
		rowsPtr := gw.GwDB.DbConfig.GetRouteNameURLList(a...)
		if rowsPtr != nil {
			foundRows := rowsPtr.Rows
			for r := range foundRows {
				foundRow := foundRows[r]
				rowContent := parseGatewayRoutesRow(&foundRow)
				rtn = append(rtn, *rowContent)
			}
			// add to cache now-----
			aJSON, err := json.Marshal(rtn)
			if err != nil {
				fmt.Println(err)
			} else {
				cval := b64.StdEncoding.EncodeToString([]byte(aJSON))
				var i ch.Item
				i.Key = key
				i.Value = cval
				res := cp.Set(&i)
				if res.Success != true {
					fmt.Println("Routes not cached from db.")
				}
			}
		}
	}
	fmt.Println("Routes: ")
	fmt.Println(rtn)
	if len(rtn) > 0 && getActive == true {
		for r := range rtn {
			if rtn[r].Active == true {
				rtnVal = rtn[r]
				break
			}
		}
	} else if len(rtn) > 0 {
		for r := range rtn {
			if rtn[r].Name == routeName {
				rtnVal = rtn[r]
				break
			}
		}
	}
	return &rtnVal
}

func parseGatewayRoutesRow(foundRow *[]string) *GatewayRouteURL {
	var rtn GatewayRouteURL
	if len(*foundRow) > 0 {
		rtn.Route = (*foundRow)[0]
		rtn.Name = (*foundRow)[1]
		rtn.URL = (*foundRow)[2]
		active, err := strconv.ParseBool((*foundRow)[3])
		if err != nil {
			fmt.Print(err)
			rtn.Active = false
		} else {
			rtn.Active = active
		}
	}
	return &rtn
}