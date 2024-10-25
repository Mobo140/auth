package сonverter

import (
	"fmt"

	desc "github.com/Mobo140/microservices/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	RoleUser    = 0
	RoleAdmin   = 1
	UnknownRole = -1
)

func mapRoleFromIntToDesc(role int64) (desc.Role, error) {
	switch role {
	case RoleUser:
		return desc.Role_USER, nil
	case RoleAdmin:
		return desc.Role_ADMIN, nil
	default:
		return desc.Role_USER, fmt.Errorf("invalid role value: %d", role)
	}
}

func mapRoleFromDescToInt(role desc.Role) (int64, error) {
	switch role {
	case desc.Role_USER:
		return RoleUser, nil
	case desc.Role_ADMIN:
		return RoleAdmin, nil
	default:
		return UnknownRole, fmt.Errorf("invalid role value: %d", role)
	}
}

func mapNameFromDescToString(name *wrapperspb.StringValue) string {
	return name.GetValue()
}

func mapEmailFromDescToString(email *wrapperspb.StringValue) string {
	return email.GetValue()
}
