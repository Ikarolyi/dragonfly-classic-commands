package permissions

import (
	"encoding/json"
	"io/fs"
	"os"
)

type PermissionEntry struct {
	XUID string `json:"xuid"`
	PermissionLevel string `json:"permission"`
}

var db []PermissionEntry
var cache map[string]PermissionLevel

const FILE_PATH = "permissions.json"
const RWRWR_FILEMODE fs.FileMode = 0664

var LevelToSring = map[PermissionLevel]string{
	LEVEL_NORMAL: "member",
	LEVEL_OPERATOR: "operator",
	LEVEL_HOST: "host",
	LEVEL_AUTOMATION: "automation",
	LEVEL_ADMIN: "admin",
}

var StringToLevel = map[string]PermissionLevel{
	"member": LEVEL_NORMAL,
	"operator": LEVEL_OPERATOR,
	"host": LEVEL_HOST,
	"automation": LEVEL_AUTOMATION,
	"admin": LEVEL_ADMIN, 
}

func Init() {
	if _, e := os.Stat(FILE_PATH); e == fs.ErrNotExist{
		os.Create(FILE_PATH)
	}

	file, _ := os.ReadFile(FILE_PATH)
	cache = make(map[string]PermissionLevel)
	db = make([]PermissionEntry, 0)
	json.Unmarshal(file, &db)
}

func Buffer() {
	for xuid, level := range cache {
		Put(xuid, level)
	}
	SavePermissions()
}

func Put(xuid string, level PermissionLevel){
	var index int = -1
	for i, e := range db {
		if e.XUID == xuid{
			index = i
			break
		}
	}
	str_level := LevelToSring[level]

	if index == -1{
		db = append(db, PermissionEntry{xuid, str_level})
	}else{
		(db)[index].PermissionLevel = str_level
	}
}

func Get(xuid string) PermissionLevel{
	var index int = -1
	for i, e := range db {
		if e.XUID == xuid{
			index = i
			break
		}
	}

	if index == -1{
		return LEVEL_NORMAL
	}else{
		return StringToLevel[db[index].PermissionLevel]
	}
}

func SavePermissions(){
	val, _ := json.MarshalIndent(db, "", "  ")
	err := os.WriteFile(FILE_PATH, val, RWRWR_FILEMODE)
	if err == nil{
		println("Saved permissions.json")
	}
}