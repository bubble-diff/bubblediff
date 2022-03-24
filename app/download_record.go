package app

import (
	"context"
	"encoding/json"
	"fmt"

	cosv5 "github.com/tencentyun/cos-go-sdk-v5"

	"github.com/bubble-diff/bubblediff/models"
)

func DownloadRecord(ctx context.Context, taskID int64, recordID string) (record *models.Record, err error) {
	key := fmt.Sprintf("%d/%s", taskID, recordID)
	opt := &cosv5.ObjectGetOptions{ResponseContentType: "application/json"}
	resp, err := cos.Object.Get(ctx, key, opt)
	if err != nil {
		return nil, err
	}

	record = new(models.Record)
	err = json.NewDecoder(resp.Body).Decode(record)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return record, nil
}
