package controllers

import (
	"GoogleFirebase/models"
	"GoogleFirebase/utilities"
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

func MapCollectionUrl(w http.ResponseWriter, r *http.Request) utilities.ResponseJSON {
	returnData := utilities.ResponseJSON{}
	returnData.Code = 400
	returnData.Msg = "Failure"
	returnData.Model = nil

	ctx := context.Background()
	ProjectID := cast.ToString(viper.Get("projectid"))
	client, err := firestore.NewClient(ctx, ProjectID)
	fmt.Println(client, err)

	if err != nil {
		fmt.Println("handle error fireStoreInitializeFunction")

	}

	ID := r.URL.Query().Get("id")
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	collection := r.URL.Query().Get("collection")
	inputobj := make(map[string]interface{})
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &inputobj)

	switch true {
	/* Get all collection json response*/
	case cast.ToString(r.Method) == "GET" && cast.ToString(ID) == "":
		allData := models.GetAll(ctx, client, collection)
		if len(allData) > 0 {
			returnData.Code = 200
			returnData.Msg = "Success"
			returnData.Model = allData
		} else {
			returnData.Code = 401
			returnData.Msg = "Failure:USER GET REQUEST"
		}
		break
	/*Get collection json response By FirebaseID*/
	case cast.ToString(r.Method) == "GET" && cast.ToString(ID) != "":
		allData := models.GetOne(ctx, client, collection, cast.ToString(ID), "")
		if len(allData) > 0 {
			returnData.Code = 200
			returnData.Msg = "Success"
			returnData.Model = allData
		} else {
			returnData.Code = 402
			returnData.Msg = "Failure:USER doesn't exists"
		}
		break
	/*Get collection json response By Key And Value */
	case cast.ToString(r.Method) == "GET" && cast.ToString(key) != "" && cast.ToString(value) != "":
		allData := utilities.GetCollectionFirebaseFunction(collection, cast.ToString(key), cast.ToString(value))
		if len(allData) > 0 {
			returnData.Code = 200
			returnData.Msg = "Success"
			returnData.Model = allData
		} else {
			returnData.Code = 402
			returnData.Msg = "error"
			returnData.Model = nil
		}
		break
		/*Update collection json response same struct of firebase*/
	case cast.ToString(r.Method) == "POST":
		err := models.Update(ctx, client, collection, inputobj)
		if err == nil {
			returnData.Code = 200
			returnData.Msg = "Success"
			returnData.Model = nil
		} else {
			returnData.Code = 403
			returnData.Msg = "Failure:USER  Update error"
		}
		break
	/*Create New collection json response same struct of firebase*/
	case cast.ToString(r.Method) == "PUT":
		err := models.Create(ctx, client, collection, inputobj)
		if err == nil {
			returnData.Code = 200
			returnData.Msg = "Success"
			returnData.Model = nil
		} else {
			returnData.Code = 404
			returnData.Msg = "Failure:User Creation error"
		}

	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure:Authorization error"
		returnData.Model = nil

	}
	return returnData
}
