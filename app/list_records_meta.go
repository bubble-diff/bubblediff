package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	"github.com/bubble-diff/bubblediff/models"
)

func ListRecordsMeta(ctx context.Context, taskid int64, path string) (metas []*models.RecordMeta, err error) {
	key := fmt.Sprintf("task%d_records_meta", taskid)
	values, err := rdb.LRange(ctx, key, 0, -1).Result()
	metas = make([]*models.RecordMeta, 0)
	for _, val := range values {
		meta := new(models.RecordMeta)

		err = json.Unmarshal([]byte(val), meta)
		if err != nil {
			log.Printf("[app.ListRecordsMeta] unmarshal failed, val=%s, %s", val, err)
			return nil, err
		}

		matched := true
		if len(path) != 0{
			matched, err = regexp.MatchString(fmt.Sprintf(`^%s.*`, path), meta.Path)
			if err != nil {
				log.Printf("[app.ListRecordsMeta] match regexp string failed, pattern=^%s.*, path=%s, %s", path, meta.Path, err)
				return nil, err
			}
		}

		if matched {
			metas = append(metas, meta)
		}
	}
	return metas, nil
}
