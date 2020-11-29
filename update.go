package sqlow

import "fmt"

type UpdateData struct {
	Table   TableData
	OldName string
}

//ToUpdate is used to change (update) the contents of the Table. Arguments: OldName If not, enter the current name
func (tab TableData) ToUpdate(oldName string) *UpdateData {
	return &UpdateData{Table: tab, OldName: oldName}
}

func (upd UpdateData) Build() (string, error) {
	existsTable := []string{}
	if database == nil {
		return "", &Error{Msg: "Currently the Database has not been created yet. Create it with New (*sql.DB, Name)."}
	}
	result, err := database.Database.Query(fmt.Sprintf("SHOW TABLES LIKE \"%v\";"), upd.OldName)
	if err != nil {
		return "", err
	}
	if !result.Next() {
		return "", &Error{Msg: "The specified table did not exist."}
	}
	result, err = database.Database.Query(fmt.Sprintf("SHOW COLUMNS FROM `%v`;"), upd.OldName)
	if err != nil {
		return "", err
	}
	for result.Next() {
		var name string
		var del string
		result.Scan(&name, &del, &del, &del, &del, &del)
		existsTable = append(existsTable, name)
	}

}
