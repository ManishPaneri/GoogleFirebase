package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/spf13/cast"
	"google.golang.org/api/iterator"
)

//Get Firestore All Details Function
func GetAll(ctx context.Context, client *firestore.Client, Collection string) map[int]map[string]interface{} {
	fmt.Println("====Get All Data====")
	returnMap := make(map[int]map[string]interface{})
	iter := client.Collection(Collection).Documents(ctx)
	i := 0
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("Failed to iterate: %v", err)
		}

		returnMap[i] = doc.Data()
		returnMap[i]["FirebaseID"] = cast.ToString(doc.Ref.ID)

		i = i + 1

	}
	return returnMap
}

func GetOne(ctx context.Context, client *firestore.Client, Collection string, FirebaseID string, FieldID string) map[string]interface{} {
	fmt.Println("====Get One Data====", FirebaseID, FieldID)
	returnMap := make(map[string]interface{})
	if FieldID == "" {
		data, err := client.Collection(Collection).Doc(FirebaseID).Get(ctx)
		if err != nil {
			return returnMap
		}
		returnMap = data.Data()

	} else {
		fmt.Println("====Get One Data====", Collection, FirebaseID, FieldID)

		iter := client.Collection(Collection).Where("name", "==", FieldID).Documents(ctx)

		//.Where("name", "==", FieldID).Documents(ctx)
		i := 0
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				fmt.Printf("Failed to iterate: %v", err)
			}

			returnMap = doc.Data()
			returnMap["FirebaseID"] = cast.ToString(doc.Ref.ID)
			i = i + 1
		}
	}

	return returnMap
}

func Update(ctx context.Context, client *firestore.Client, Collection string, InputObj map[string]interface{}) (err error) {
	fmt.Println("====Update Data====")
	fmt.Println("----------Input obj -----------", InputObj)
	if cast.ToString(InputObj["FirebaseID"]) == "" {
		return nil
	}
	_, err = client.Collection(Collection).Doc(cast.ToString(InputObj["FirebaseID"])).Set(ctx, InputObj, firestore.MergeAll)
	if err != nil {
		return err
	}
	return nil
}

func Create(ctx context.Context, client *firestore.Client, Collection string, InputObj map[string]interface{}) (err error) {
	fmt.Println("====Create New Data====")
	_, _, err = client.Collection(Collection).Add(ctx, InputObj)
	if err != nil {
		fmt.Printf("Failed adding : %v", err)
		return err
	}
	return nil

}

