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
	AllLeaves  = []byte{0x05}
)

func StudentStoreId(studentid string) []byte {
	key := make([]byte, len(StudentKey)+len(studentid))
	copy(key, StudentKey)
	copy(key[len(StudentKey):], []byte(studentid))
	return key
}
func AdminstoreId(admin string) []byte {
	key := make([]byte, len(AdminKey)+len(admin))
	copy(key, AdminKey)
	copy(key[len(AdminKey):], []byte(admin))
	return key
}

//LeaveStoreId used store the leave request applied by the student
func LeaveStoreId(studentid string) []byte {
	key := make([]byte, len(LeaveKey)+len(studentid))
	copy(key, LeaveKey)
	copy(key[len(LeaveKey):], []byte(studentid))
	return key
}

//LeaveKeyStoreId stores the counter of the respective student
func LeaveKeyStoreId(sid string) []byte {
	key := make([]byte, len(LeaveId)+len(sid))
	copy(key, LeaveId)
	copy(key[len(LeaveId):], []byte(sid))
	// LeaveId = append(LeaveId, []byte(sid))
	return key
}
func AllLeavesStoreId(sid string) []byte {
	key := make([]byte, len(AllLeaves)+len(sid))
	copy(key, AllLeaves)
	copy(key[len(AllLeaves):], []byte(sid))
	return key
}
