// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		name,
		age,
		address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r != nil && r.Name != "" && r.Address != nil && r.Address["street"] != "" {
		return true
	}

	return false
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	if r != nil {
		r.Name = ""
		r.Age = 0
		r.Address = nil
	}
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var i int
	for _, r := range residents {
		if r.HasRequiredInfo() == true {
			i++
		}
	}

	return i
}