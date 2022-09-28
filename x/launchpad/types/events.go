package types

const (
	TypeProjectCreated      = "project_created"
	TypeProjectDeleted      = "project_deleted"
	TypeWithdrawTokens      = "withdraw_tokens"
	TypeSetProjectVerified  = "set_project_verified"
	AttributeValueCategory  = ModuleName
	AttributeProjectID      = "project_id"
	AttributeProjectAddress = "project_address"
)

const (
	PROJECT_INIT   = uint64(0)
	PROJECT_ACTIVE = uint64(1)
	PROJECT_ENDED  = uint64(2)
)
