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
	LeaveId    = []byte{0x04}
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
}
func LeaveStoreId(leaveid string) []byte {
	key := make([]byte, len(LeaveKey)+len(leaveid))
	copy(key, LeaveKey)
	copy(key[len(LeaveKey):], []byte(leaveid))
	return key
}

func LeaveKeyStoreId(leave int) []byte {
	LeaveId = append(LeaveId, byte(leave))
	return LeaveId
}
