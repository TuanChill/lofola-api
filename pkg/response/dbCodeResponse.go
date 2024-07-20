package response

const (
	ER_DUP_KEY                            = 1022 // Duplicate key error
	ER_OUTOFMEMORY                        = 1037 // Out of memory error
	ER_CON_COUNT_ERROR                    = 1040 // Too many connections error
	ER_BAD_DB_ERROR                       = 1049 // Unknown database error
	ER_TABLE_EXISTS_ERROR                 = 1050 // Table already exists error
	ER_BAD_TABLE_ERROR                    = 1051 // Unknown table error
	ER_NON_UNIQ_ERROR                     = 1052 // Non-unique column error
	ER_BAD_FIELD_ERROR                    = 1054 // Unknown column error
	ER_WRONG_FIELD_WITH_GROUP             = 1055 // Incorrect column in group by clause error
	ER_DUP_FIELDNAME                      = 1060 // Duplicate column name error
	ER_DUP_KEYNAME                        = 1061 // Duplicate key name error
	ER_DUP_ENTRY                          = 1062 // Duplicate entry error
	ER_PARSE_ERROR                        = 1064 // SQL syntax error
	ER_NO_SUCH_TABLE                      = 1146 // Table doesn't exist error
	ER_CANT_DO_THIS_DURING_AN_TRANSACTION = 1179 // Can't do this operation during a transaction
	ER_LOCK_WAIT_TIMEOUT                  = 1205 // Lock wait timeout exceeded error
	ER_NO_REFERENCED_ROW                  = 1216 // Foreign key constraint fails error
	ER_ROW_IS_REFERENCED                  = 1217 // Foreign key constraint fails error
	ER_DATA_TOO_LONG                      = 1406 // Data too long for column error
	ER_WRONG_VALUE_COUNT_ON_ROW           = 1136 // Column count doesn't match value count error
)
