#  X/leave

## Abstract

This document specifies the leave management system module.

This leave module is responsible for the students to apply the leaves and admin to accept the leaves.
It helps the admin to accept the leaves and query all the leaves.
It also helps the students to do the transaction and query the transaction


## Client

## CLI

A user can interact with the `lms` using CLI.

## Query
The `query` commands allows users to  `lms` module.
```
lmsd query lms --help
```

### admin

The `admin` command helps the query the admin registered.

```
lmsd query lms admin 
```
Example:

```
lmsd query lms admin 
```

Example Output:

admins:
- address: cosmos12f70wvzjtwlyrralxat3xqc86fz7rjvuk3jrvx
  name: vishal
  signer: cosmos12f70wvzjtwlyrralxat3xqc86fz7rjvuk3jrvx



### list-all-leaves

The `list-all-leaves` command helps the admin to query the all the leaves.


```
lmsd query lms list-all-leaves
```

Example:
```
lmsd query lms list-all-leaves

```
Example Output:

leaves:
    Id: b172118
  from: "2025-03-01T00:00:00Z"
  leaveid: "1"
  reason: fever
  signer: cosmos12f70wvzjtwlyrralxat3xqc86fz7rjvuk3jrvx
  status: STATUS_UNDEFINED
  to: "2027-03-01T00:00:00Z"



### list-all-students 

The `list-all-students `command helps the admin to get all the students details present in the store.

```
lmsd query lms listallstudents
```

Example:
```
lmsd query lms listallstudents

```
Example Output:

students:
- address: cosmos19hgll56xm2c5pj4y5rra6ms67e30dxvwg72xkw
  id: b172118
  name: student1




### studentleavestatus 

The `studentleavestatus` command helps the user or student to get the status of the applied leave.

```
lmsd query lms studentleavestatus [studentid] [leaveid]
```

Example:
```
lmsd query lms studentleavestatus b172118 1

```
Example Output:

Before the accepting the leave
status:
  Id: b172118
  from: "2025-03-01T00:00:00Z"
  leaveid: "1"
  reason: fever
  signer: cosmos12f70wvzjtwlyrralxat3xqc86fz7rjvuk3jrvx
  status: STATUS_UNDEFINED
  to: "2027-03-01T00:00:00Z"

After accepting the leave


status:
  Id: b172118
  from: "2025-03-01T00:00:00Z"
  leaveid: "1"
  reason: fever
  signer: cosmos12f70wvzjtwlyrralxat3xqc86fz7rjvuk3jrvx
  status: STATUS_ACCEPTED
  to: "2027-03-01T00:00:00Z"




## Transactions 

The `tx` command allows user to interact with `lms` module.
```
lmsd tx lms --help

```
 ### registeradmin
 
 The `registeradmin ` allows the user to register as a admin and make other transactions like adding students and accepting leaves.
 
 
 
```
lmsd tx lms registeradmin [adminname] --from validator-key --chain-id testnet
```

Example:
```
lmsd tx lms registeradmin vishal --from validator-key --chain-id testnet 

```




### addstudent 
The `addstudent` command allows admin to add the students.

```
lmsd tx lms addstudent [studentname] [studentid] [studentaddress] --from validator-key --chain-id testnet
```


Example:
```
lmsd tx lms addstudent student1 b172118 bec32address --from validator-key --chain-id testnet 

```



### applyleave

The `applyleave` command allows the user to apply leaves.

```
lmsd tx lms applyleave [studentid] [reason] [from] [to] --from validator-key --chain-id testnet
```


Example:
```
lmsd tx lms applyleave b172118 fever 2023-mar-1 2023-mar-4 --from validator-key --chain-id testnet

```
Example output:
 it is event triggered when leave is applied.Leaveid is used to check the status of the leave.

```
raw_log: '[{"msg_index":0,"events":[{"type":"message","attributes":[{"key":"action","value":"/cosmos.lms.v1beta1.ApplyLeaveRequest"},{"key":"leaveid:","value":"1"}]}]}]'
```





### acceptleave

The` acceptleave` command allows the admin to accept the leaves based on the studentid and leaveid.

```
lmsd tx lms acceptleave [studentid] [leaveid] --from validator-key --chain-id testnet
```


Example:
```
lmsd tx lms applyleave b172118 1 --from validator-key --chain-id testnet

```

