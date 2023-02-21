package types

const (
	ModuleName = "leavemanagementsystem"
	StoreKey   = ModuleName
	// RouterKey    = ModuleName
	// QuerierRoute = ModuleName
)

var (
	AdminKey   = []byte{0x01}
	StudentKey = []byte{0x02}
	LeaveKey   = []byte{0x03}
)

func StudentStoreId(studentid string) []byte {
	key := make([]byte, len(StudentKey)+len(studentid))
	copy(key, StudentKey)
	copy(key[len(StudentKey):], studentid)
	return key
}
func AdminstoreId(admin string) []byte {
	key := make([]byte, len(AdminKey)+len(admin))
	copy(key, AdminKey)
	copy(key[len(AdminKey):], []byte(admin))
	return key
	// AdminKey = append(AdminKey, byte(admin))
}
// func leaveId(levid int) []byte {

// }
