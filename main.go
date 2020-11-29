package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	//MySql Driver
	_ "github.com/go-sql-driver/mysql"
)

var supportSQL = []string{"mysql"}

type sqlData struct {
	SQLType   string
	Database  *sql.DB
	LoginAuth *LoginData
}

// Error is Original Error
type Error struct {
	Msg string
}

// LoginData is the information required to login is saved.
type LoginData struct {
	UserName string
	Password string
	Host     string
	Port     int
	DBname   string
}

//ColumnData is Stores SQL Column data.
type ColumnData struct {
	Name            string
	DataType        DataType
	PrimaryKey      bool
	NotNull         bool
	UniqueIndex     bool
	Unsigned        bool
	ZeroFill        bool
	AutoIncremental bool
	Default         interface{}
	Property        interface{}
}

//TableData is Stores SQL Table data.
type TableData struct {
	Name    string
	Columns []ColumnData
	DBName  string
}

// DataType is SQL DataType
type DataType struct {
	TypeName       string
	Type           string
	UNSIGNED       bool
	UndignedType   string
	ZEROFILL       bool
	MaxLength      int
	DefaultPropaty string
	AutoIncrement  bool
	PrimaryKey     bool
}

func contains(list []string, target string) bool {
	for _, value := range list {
		if value == target {
			return true
		}
	}
	return false
}

var (
	// TINYINT is the one that exists in SQL and can store numbers from -128 to 127. If you add UNSIGNED, you can store up to 255, but you can only use integers.
	TINYINT = DataType{TypeName: "TINYINT", Type: "INT8", UNSIGNED: true, ZEROFILL: true, UndignedType: "UNIT8", AutoIncrement: true, PrimaryKey: true}

	// SMALLINT is the one that exists in SQL and can store numbers from -32768 to 32767. With UNSIGNED, you can store up to 65535, but you can only use integers.
	SMALLINT = DataType{TypeName: "SMALLINT", Type: "INT16", UNSIGNED: true, ZEROFILL: true, UndignedType: "UNIT16", AutoIncrement: true, PrimaryKey: true}

	// MEDIUMINT is the one that exists in SQL and can store numbers from -8388608 to 8388607. With UNSIGNED, you can store up to 16777215, but you can only use integers.
	MEDIUMINT = DataType{TypeName: "MEDIUMINT", Type: "INT32", UNSIGNED: true, ZEROFILL: true, UndignedType: "UNIT32", AutoIncrement: true, PrimaryKey: true}

	// INT is the one that exists in SQL and can store numbers from -2147483648 to 2147483647. With UNSIGNED, you can store up to 4294967295, but you can only use integers.
	INT = DataType{TypeName: "INT", Type: "INT32", UNSIGNED: true, ZEROFILL: true, UndignedType: "UNIT32", AutoIncrement: true, PrimaryKey: true}

	// INTEGER is the one that exists in SQL and can store numbers from -2147483648 to 2147483647. With UNSIGNED, you can store up to 4294967295, but you can only use integers.
	INTEGER = DataType{TypeName: "INTEGER", Type: "INT32", UNSIGNED: true, ZEROFILL: true, UndignedType: "UNIT32", AutoIncrement: true, PrimaryKey: true}

	// BIGINT is the one that exists in SQL and can store numbers from -9223372036854775808 to 9223372036854775807. With UNSIGNED, you can store up to 18446744073709551615, but you can only use integers.
	BIGINT = DataType{TypeName: "BIGINT", Type: "INT64", UNSIGNED: true, ZEROFILL: true, UndignedType: "UNIT64", AutoIncrement: true, PrimaryKey: true}

	// BOOLEAN is a type in SQL that can store True or False.
	BOOLEAN = DataType{TypeName: "BOOLEAN", Type: "BOOL", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: true}

	// BOOL is a type in SQL that can store True or False.
	BOOL = DataType{TypeName: "BIGINT", Type: "BOOL", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: false}

	// BIT is a SQL type that can store bit values.
	BIT = DataType{TypeName: "UNIT", Type: "UNIT", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: true}

	// FLOAT is a type that exists in SQL and can store accurate decimals up to the 7th decimal place.
	FLOAT = DataType{TypeName: "FLOAT", Type: "FLOAT32", UNSIGNED: true, ZEROFILL: true, AutoIncrement: true, PrimaryKey: true}

	// DOUBLE is a type that exists in SQL and can store accurate decimals up to the 7th decimal place.
	DOUBLE = DataType{TypeName: "DOUBLE", Type: "FLOAT64", UNSIGNED: true, ZEROFILL: true, AutoIncrement: true, PrimaryKey: true}

	// DATE is a type that exists in SQL. You can save the year, month, and day.
	DATE = DataType{TypeName: "DATE", Type: "DATE", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: true}

	// DATETIME is a type that exists in SQL. You can save the year, month, day, hour, minute, and second.
	DATETIME = DataType{TypeName: "DATETIME", Type: "DATE", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: true}

	// TIMESTAMP is a type that exists in SQL. You can save the year, month, day, hour, minute, and second. Also, if no value is explicitly assigned, the date and time will be set automatically when the value is changed.
	TIMESTAMP = DataType{TypeName: "TIMESTAMP", Type: "DATE", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: true}

	// TIME is a type that exists in SQL. You can save hours, minutes, and seconds.
	TIME = DataType{TypeName: "TIME", Type: "DATE", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: true}

	// VARCHAR is a SQL type. You can store the specified character string.
	VARCHAR = DataType{TypeName: "VARCHAR", Type: "STRING", UNSIGNED: false, ZEROFILL: false, MaxLength: 65535, DefaultPropaty: "255", AutoIncrement: false, PrimaryKey: false}

	// TEXT is a SQL type. You can store the specified character string.
	TEXT = DataType{TypeName: "TEXT", Type: "STRING", UNSIGNED: false, ZEROFILL: false, MaxLength: 14090025, DefaultPropaty: "255", AutoIncrement: false, PrimaryKey: false}

	// MIDIUMTEXT is a SQL type. You can store the specified character string.
	MIDIUMTEXT = DataType{TypeName: "MIDIUMTEXT", Type: "STRING", UNSIGNED: false, ZEROFILL: false, MaxLength: 3741318945, DefaultPropaty: "255", AutoIncrement: false, PrimaryKey: false}

	// LONGTEXT is a SQL type. You can store the specified character string.
	LONGTEXT = DataType{TypeName: "LONGTEXT", Type: "STRING", UNSIGNED: false, ZEROFILL: false, MaxLength: 4294967295, DefaultPropaty: "255", AutoIncrement: false, PrimaryKey: false}

	// ENUM is a type of SQL. Can store one of the specified string lists.
	ENUM = DataType{TypeName: "TEXT", Type: "LIST", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: false}

	// SET is a type of SQL. Can store multiple in the specified string list.
	SET = DataType{TypeName: "SET", Type: "LIST", UNSIGNED: false, ZEROFILL: false, AutoIncrement: false, PrimaryKey: false}
)

func (err *Error) Error() string {
	return fmt.Sprintf("Does not support the following SQL: %s", err.Msg)
}

func connect(sqlType string, user string, password string, host string, port int, dbname string) (sqlData, error) {
	sqlType = strings.ToUpper(sqlType)
	sqldata := sqlData{}
	if result := contains(supportSQL, sqlType); !result {
		return sqldata, &Error{Msg: sqlType}
	}
	db, err := sql.Open(sqlType, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname))
	if err == nil {
		sqldata.Database = db
		sqldata.SQLType = sqlType
		sqldata.LoginAuth = &LoginData{
			UserName: user,
			Password: password,
			Host:     host,
			Port:     port,
			DBname:   dbname,
		}
		return sqldata, nil
	}
	return sqldata, err
}

func (data sqlData) Ping() error {
	return data.Database.Ping()
}

func (data sqlData) Close() error {
	return data.Database.Close()
}

func (col *ColumnData) setPrimaryKey(check bool) *ColumnData {
	col.PrimaryKey = check
	return col
}

func (col *ColumnData) setNotNull(check bool) *ColumnData {
	col.NotNull = check
	return col
}

func (col *ColumnData) setUniqueIndex(check bool) *ColumnData {
	col.UniqueIndex = check
	return col
}
func (col *ColumnData) setUnsigned(check bool) *ColumnData {
	if col.DataType.UNSIGNED {
		col.Unsigned = check
		return col
	}
	return col
}

func (col *ColumnData) setZeroFill(check bool) *ColumnData {
	if col.DataType.ZEROFILL {
		col.ZeroFill = check
		return col
	}
	return col
}

func (col *ColumnData) setAutoIncrement(check bool) *ColumnData {
	if col.DataType.AutoIncrement {
		col.AutoIncremental = check
		return col
	}
	return col
}

func (col *ColumnData) setDefault(value interface{}) *ColumnData {
	col.Default = value
	return col
}

// Column function is Create a Column. return *ColumnData
func Column(name string, dataType DataType) *ColumnData {
	column := ColumnData{Name: name, DataType: dataType}
	return &column
}

func (tab TableData) build() (string, error) {
	autoIncrement := false
	columns := tab.Columns
	name := tab.Name
	dbname := tab.DBName
	primarys := []string{}
	uniqueIndex := []string{}
	result := fmt.Sprintf("CREATE TABLE '%v'.'%v' (", dbname, name)
	for _, i := range columns {
		if i.AutoIncremental {
			if autoIncrement {
				return "", &Error{Msg: "Up to one AutoIncrement can be set for each table."}
			}
			autoIncrement = true
		}
		if i.PrimaryKey {
			primarys = append(primarys, i.Name)
		}
		if i.UniqueIndex {
			uniqueIndex = append(uniqueIndex, i.Name)
		}
		r, err := i.build()
		if err != nil {
			return "", err
		}
		result += r
	}
	if len(primarys) >= 1 {
		result += ","
		result += fmt.Sprintf("PRIMARY KEY (%v)", toSQLList(primarys))
	}
	if len(uniqueIndex) >= 1 {
		for _, i := range uniqueIndex {
			result += ","
			result += fmt.Sprintf("UNIQUE INDEX `%v_UNIQUE` (`%v` ASC) VISIBLE", i, i)
		}
	}
	return result, nil
}

func (col ColumnData) build() (string, error) {
	result := fmt.Sprintf("'%s' ", col.Name)
	switch col.DataType.TypeName {
	case "ENUM":
		if defa := col.Property; defa != nil {
			if res, ok := defa.([]string); ok {
				result += fmt.Sprintf(" %v", fmt.Sprintf("ENUM(%v)", toSQLList(res)))
			}
		} else {
			return "", &Error{Msg: "The SQL syntax could not be created successfully because the property is not set in the ENUM of the Column."}
		}
	case "SET":
		if prop := col.Property; prop != nil {
			if res, ok := prop.([]string); ok {
				result += fmt.Sprintf(" %v", fmt.Sprintf("SET(%v)", toSQLList(res)))
			}
		} else {
			return "", &Error{Msg: "The SQL syntax could not be created successfully because the property is not set in the SET of the Column."}
		}
	default:
		if prop := col.Property; prop != nil {
			if res, ok := prop.(int); ok {
				result += fmt.Sprintf(" %v(%v)", col.DataType.TypeName, res)
			} else if res, ok := prop.(string); ok {
				result += fmt.Sprintf(" %v('%v')", col.DataType.TypeName, res)
			}
		} else {
			result += col.DataType.TypeName
		}
	}

	dataType := col.DataType.TypeName
	if col.NotNull {
		result += " NOT NULL"
	}
	if col.AutoIncremental {
		if col.DataType.AutoIncrement {
			result += " AUTO_INCREMENT"
		}
	}
	if col.ZeroFill {
		if col.DataType.ZEROFILL {
			result += " ZEROFILL"
		}
	}
	if col.Unsigned {
		if col.DataType.UNSIGNED {
			result += " UNSIGNED"
		}
	}
	if !col.DataType.AutoIncrement {
		if defa := col.Default; defa != nil {
			switch dataType {
			case "TINYINT", "SMALLINT", "MEDIUMINT", "INT", "INTEGER", "BIGINT", "FLOAT", "DOUBLE":
				if col.Unsigned {
					if res, ok := defa.(uint64); ok {
						result += fmt.Sprintf(" %v", res)
					} else {
						return "", &Error{Msg: "The default value could not be successfully converted to a number."}
					}
				} else {
					if res, ok := defa.(int64); ok {
						result += fmt.Sprintf(" %v", res)
					} else {
						return "", &Error{Msg: "The default value could not be successfully converted to a number."}
					}
				}
			case "BOOL", "BOOLEAN":
				if res, ok := defa.(bool); ok {
					result += fmt.Sprintf(" %v", res)
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to Boolean."}
				}
			case "DATE":
				if res, ok := defa.(time.Time); ok {
					result += fmt.Sprintf(" '%v'", toSQLDate(res))
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to time.Time."}
				}
			case "DATETIME":
				if res, ok := defa.(time.Time); ok {
					result += fmt.Sprintf(" '%v'", toSQLDateTime(res))
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to time.Time."}
				}
			case "TIMESTAMP":
				if res, ok := defa.(time.Time); ok {
					result += fmt.Sprintf(" '%v'", toSQLDateTime(res))
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to time.Time."}
				}
			case "TIME":
				if res, ok := defa.(time.Time); ok {
					result += fmt.Sprintf(" '%v'", toSQLTime(res))
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to time.Time."}
				}
			case "TEXT", "VARCHAR", "MIDIUMTEXT", "LONGTEXT", "ENUM":
				if res, ok := defa.(string); ok {
					result += fmt.Sprintf(" '%v'", res)
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to String."}
				}
			case "SET":
				if res, ok := defa.([]string); ok {
					result += fmt.Sprintf(" %v", toSQLList(res))
				} else {
					return "", &Error{Msg: "The default value could not be successfully converted to List."}
				}
			}
		}
	}
	return result, nil
}

func toSQLDate(val time.Time) string {
	return fmt.Sprintf("%v-%v-%v", val.Year(), val.Month(), val.Day())
}

func toSQLDateTime(val time.Time) string {
	return fmt.Sprintf("%v-%v-%v %v:%v:%v", val.Year(), val.Month(), val.Day(), val.Hour(), val.Minute(), val.Second())
}

func toSQLTime(val time.Time) string {
	return fmt.Sprintf("%v:%v:%v", val.Hour(), val.Minute(), val.Second())
}

func toSQLList(val []string) string {
	tmp := []string{}
	for _, value := range val {
		tmp = append(tmp, fmt.Sprintf("'%v'", string(value)))
	}
	return strings.Join(tmp, ",")
}
