package utilities

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
)

type ResponseJSON struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

func GetCollectionFirebaseFunction(collection string, key string, value string) map[string]interface{} {
	fmt.Println("====GetCollectionFirebaseFunction====")
	returnMap := make(map[string]interface{})
	ctx := context.Background()
	ProjectID := cast.ToString(viper.Get("projectid"))
	client, _ := firestore.NewClient(ctx, ProjectID)

	iter := client.Collection(collection).Where(key, "==", value).Documents(ctx)

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
	}

	return returnMap

}
