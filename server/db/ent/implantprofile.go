// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/bishopfox/sliver/server/db/ent/implantprofile"
	"github.com/facebookincubator/ent/dialect/sql"
)

// ImplantProfile is the model entity for the ImplantProfile schema.
type ImplantProfile struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ImplantProfile) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ImplantProfile fields.
func (ip *ImplantProfile) assignValues(values ...interface{}) error {
	if m, n := len(values), len(implantprofile.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ip.ID = int(value.Int64)
	values = values[1:]
	return nil
}

// Update returns a builder for updating this ImplantProfile.
// Note that, you need to call ImplantProfile.Unwrap() before calling this method, if this ImplantProfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (ip *ImplantProfile) Update() *ImplantProfileUpdateOne {
	return (&ImplantProfileClient{config: ip.config}).UpdateOne(ip)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ip *ImplantProfile) Unwrap() *ImplantProfile {
	tx, ok := ip.config.driver.(*txDriver)
	if !ok {
		panic("ent: ImplantProfile is not a transactional entity")
	}
	ip.config.driver = tx.drv
	return ip
}

// String implements the fmt.Stringer.
func (ip *ImplantProfile) String() string {
	var builder strings.Builder
	builder.WriteString("ImplantProfile(")
	builder.WriteString(fmt.Sprintf("id=%v", ip.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ImplantProfiles is a parsable slice of ImplantProfile.
type ImplantProfiles []*ImplantProfile

func (ip ImplantProfiles) config(cfg config) {
	for _i := range ip {
		ip[_i].config = cfg
	}
}
