package pb_adapter

import (
	"fmt"

	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

//Sample Db operations for Pocketbase, you can use SQL for queries and also custom pocketbase functions to fetch data too. For more information, see the Extending with GO documentation of Pocketbasse.

func (p *PocketBaseServer) FetchData() error {
	record, err := p.PBServer.Dao().FindRecordById("News", "ltzoxlrmnz2vt74")
	if err != nil {
		fmt.Println("The Error is", err)
		return err
	}
	fmt.Printf("Records for the fetch is %v", record.Get("Summary"))
	return nil
}

func (p *PocketBaseServer) PutData() error {
	collection, err := p.PBServer.Dao().FindCollectionByNameOrId("News")
	if err != nil {
		fmt.Println("This is the error", err)
		return err
	}

	new_record := models.NewRecord(collection)
	form := forms.NewRecordUpsert(p.PBServer, new_record)
	form.LoadData(map[string]any{
		"News":    "test-2",
		"Summary": "test-2 summary",
		"URL":     "This is a summary. Hahah",
	})

	if err := form.Submit(); err != nil {
		fmt.Println("The Error in Put", err)
		return err
	}

	return nil
}
