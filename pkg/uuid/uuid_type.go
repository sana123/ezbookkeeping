package uuid

type UuidType uint8

// Types of uuid
const (
	UUID_TYPE_DEFAULT     UuidType = 0
	UUID_TYPE_USER        UuidType = 1
	UUID_TYPE_ACCOUNT     UuidType = 2
	UUID_TYPE_TRANSACTION UuidType = 3
	UUID_TYPE_CATEGORY    UuidType = 4
	UUID_TYPE_TAG         UuidType = 5
)
