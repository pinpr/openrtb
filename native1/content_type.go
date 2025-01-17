package native1

import (
	"database/sql/driver"

	"github.com/pinpr/backend-shared-kit/sql"
	"github.com/pinpr/backend-shared-kit/vld"
)

type (
	ContentType  string
	ContentTypes []ContentType
)

const (
	ContentTypeText     ContentType = "text"
	ContentTypeImage    ContentType = "image"
	ContentTypeVideo    ContentType = "video"
	ContentTypeSubtitle ContentType = "subtitle"
	ContentTypeRating   ContentType = "rating"
	ContentTypeIcon     ContentType = "icon"
)

var (
	allContentTypes = []ContentType{
		ContentTypeText,
		ContentTypeImage,
		ContentTypeVideo,
		ContentTypeSubtitle,
		ContentTypeRating,
		ContentTypeIcon,
	}
)

func (e ContentType) IsValid() error {
	return vld.EnumIsValid(&e, allContentTypes)
}

func (c *ContentTypes) Scan(src interface{}) error {
	return sql.Scan(c, src)
}

func (c ContentTypes) Value() (driver.Value, error) {
	return sql.Value(c)
}
