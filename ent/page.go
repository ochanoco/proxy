// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/tracer-silver-bullet/tracer-silver-bullet/proxy/ent/page"
)

// Page is the model entity for the Page schema.
type Page struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Skip holds the value of the "skip" field.
	Skip          bool `json:"skip,omitempty"`
	project_pages *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Page) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case page.FieldSkip:
			values[i] = new(sql.NullBool)
		case page.FieldID:
			values[i] = new(sql.NullInt64)
		case page.FieldURL:
			values[i] = new(sql.NullString)
		case page.ForeignKeys[0]: // project_pages
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Page", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Page fields.
func (pa *Page) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case page.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case page.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				pa.URL = value.String
			}
		case page.FieldSkip:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field skip", values[i])
			} else if value.Valid {
				pa.Skip = value.Bool
			}
		case page.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field project_pages", value)
			} else if value.Valid {
				pa.project_pages = new(int)
				*pa.project_pages = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Page.
// Note that you need to call Page.Unwrap() before calling this method if this Page
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Page) Update() *PageUpdateOne {
	return (&PageClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the Page entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Page) Unwrap() *Page {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Page is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Page) String() string {
	var builder strings.Builder
	builder.WriteString("Page(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", url=")
	builder.WriteString(pa.URL)
	builder.WriteString(", skip=")
	builder.WriteString(fmt.Sprintf("%v", pa.Skip))
	builder.WriteByte(')')
	return builder.String()
}

// Pages is a parsable slice of Page.
type Pages []*Page

func (pa Pages) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
