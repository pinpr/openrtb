package adcom1

import (
	"database/sql/driver"
	"github.com/samber/lo"
	"github.com/pinpr/backend-shared-kit/sql"
	"github.com/pinpr/backend-shared-kit/vld"
	"golang.org/x/exp/slices"
)

// APIFramework represents API frameworks either supported by a placement or required by an ad.
type APIFramework int
type APIFrameworkEnum string
type APIFrameworks []APIFrameworkEnum

// API frameworks either supported by a placement or required by an ad.
//
// Values of 500+ hold vendor-specific codes.
const (
	APIVPAID10 APIFrameworkEnum = "VPAID1.0" // 1 VPAID 1.0
	APIVPAID20 APIFrameworkEnum = "VPAID2.0" // 2 VPAID 2.0
	APIMRAID10 APIFrameworkEnum = "MRAID1.0" // 3 MRAID 1.0
	APIORMMA   APIFrameworkEnum = "ORMMA"    // 4 ORMMA
	APIMRAID20 APIFrameworkEnum = "MRAID2.0" // 5 MRAID 2.0
	APIMRAID30 APIFrameworkEnum = "MRAID3.0" // 6 MRAID 3.0
	APIOMID10  APIFrameworkEnum = "OMID1.0"  // 7 OMID 1.0
	APISIMID10 APIFrameworkEnum = "SIMID1.0" // 8 SIMID 1.0
	APISIMID11 APIFrameworkEnum = "SIMID1.1" // 9 SIMID 1.1
)

var allAPIs = []APIFrameworkEnum{
	APIVPAID10,
	APIVPAID20,
	APIMRAID10,
	APIORMMA,
	APIMRAID20,
	APIMRAID30,
	APIOMID10,
	APISIMID10,
	APISIMID11,
}

func (e APIFrameworkEnum) IsValid() error {
	return vld.EnumIsValid(&e, allAPIs)
}

func (e APIFrameworkEnum) ToValue() APIFramework {
	return APIFramework(slices.Index(allAPIs, e) + 1)
}

func (e APIFramework) ToValue() APIFrameworkEnum {
	return allAPIs[e-1]
}

func (a *APIFrameworks) Scan(src interface{}) error {
	return sql.Scan(a, src)
}

func (a APIFrameworks) Value() (driver.Value, error) {
	return sql.Value(a)
}

func (a APIFrameworks) ToValues() []APIFramework {
	return lo.Map(a, func(e APIFrameworkEnum, _ int) APIFramework {
		return e.ToValue()
	})
}
