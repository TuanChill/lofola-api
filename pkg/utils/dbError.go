// utils.go

package utils

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/tuanchill/lofola-api/pkg/response"
)

func HandleMySQLError(err error) string {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		log.Printf("MySQL Error Code: %d", mysqlErr.Number)

		switch mysqlErr.Number {
		case response.SuccessfulCompletion:
			return "Successful completion"
		case response.Warning:
			return "Warning"
		case response.DynamicResultSetsReturned:
			return "Dynamic result sets returned"
		case response.ImplicitZeroBitPadding:
			return "Implicit zero bit padding"
		case response.PrivilegeNotGranted:
			return "Privilege not granted"
		case response.PrivilegeNotRevoked:
			return "Privilege not revoked"
		case response.StringDataRightTruncation:
			return "String data right truncation"
		case response.DeprecatedFeature:
			return "Deprecated feature"
		case response.NoData:
			return "No data"
		// Add more cases as needed
		default:
			return "Unknown MySQL error"
		}
	}
	return "Not a MySQL error"
}
