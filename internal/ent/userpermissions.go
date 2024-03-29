// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/user"
	"lybbrio/internal/ent/userpermissions"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserPermissions is the model entity for the UserPermissions schema.
type UserPermissions struct {
	config `json:"-"`
	// ID of the ent.
	ID ksuid.ID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID ksuid.ID `json:"user_id,omitempty"`
	// Admin holds the value of the "Admin" field.
	Admin bool `json:"Admin,omitempty"`
	// CanCreatePublic holds the value of the "CanCreatePublic" field.
	CanCreatePublic bool `json:"CanCreatePublic,omitempty"`
	// CanEdit holds the value of the "CanEdit" field.
	CanEdit bool `json:"CanEdit,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserPermissionsQuery when eager-loading is set.
	Edges        UserPermissionsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserPermissionsEdges holds the relations/edges for other nodes in the graph.
type UserPermissionsEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserPermissionsEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserPermissions) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userpermissions.FieldAdmin, userpermissions.FieldCanCreatePublic, userpermissions.FieldCanEdit:
			values[i] = new(sql.NullBool)
		case userpermissions.FieldID, userpermissions.FieldUserID:
			values[i] = new(sql.NullString)
		case userpermissions.FieldCreateTime, userpermissions.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserPermissions fields.
func (up *UserPermissions) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userpermissions.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				up.ID = ksuid.ID(value.String)
			}
		case userpermissions.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				up.CreateTime = value.Time
			}
		case userpermissions.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				up.UpdateTime = value.Time
			}
		case userpermissions.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				up.UserID = ksuid.ID(value.String)
			}
		case userpermissions.FieldAdmin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field Admin", values[i])
			} else if value.Valid {
				up.Admin = value.Bool
			}
		case userpermissions.FieldCanCreatePublic:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field CanCreatePublic", values[i])
			} else if value.Valid {
				up.CanCreatePublic = value.Bool
			}
		case userpermissions.FieldCanEdit:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field CanEdit", values[i])
			} else if value.Valid {
				up.CanEdit = value.Bool
			}
		default:
			up.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserPermissions.
// This includes values selected through modifiers, order, etc.
func (up *UserPermissions) Value(name string) (ent.Value, error) {
	return up.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserPermissions entity.
func (up *UserPermissions) QueryUser() *UserQuery {
	return NewUserPermissionsClient(up.config).QueryUser(up)
}

// Update returns a builder for updating this UserPermissions.
// Note that you need to call UserPermissions.Unwrap() before calling this method if this UserPermissions
// was returned from a transaction, and the transaction was committed or rolled back.
func (up *UserPermissions) Update() *UserPermissionsUpdateOne {
	return NewUserPermissionsClient(up.config).UpdateOne(up)
}

// Unwrap unwraps the UserPermissions entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (up *UserPermissions) Unwrap() *UserPermissions {
	_tx, ok := up.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserPermissions is not a transactional entity")
	}
	up.config.driver = _tx.drv
	return up
}

// String implements the fmt.Stringer.
func (up *UserPermissions) String() string {
	var builder strings.Builder
	builder.WriteString("UserPermissions(")
	builder.WriteString(fmt.Sprintf("id=%v, ", up.ID))
	builder.WriteString("create_time=")
	builder.WriteString(up.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(up.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", up.UserID))
	builder.WriteString(", ")
	builder.WriteString("Admin=")
	builder.WriteString(fmt.Sprintf("%v", up.Admin))
	builder.WriteString(", ")
	builder.WriteString("CanCreatePublic=")
	builder.WriteString(fmt.Sprintf("%v", up.CanCreatePublic))
	builder.WriteString(", ")
	builder.WriteString("CanEdit=")
	builder.WriteString(fmt.Sprintf("%v", up.CanEdit))
	builder.WriteByte(')')
	return builder.String()
}

// UserPermissionsSlice is a parsable slice of UserPermissions.
type UserPermissionsSlice []*UserPermissions
