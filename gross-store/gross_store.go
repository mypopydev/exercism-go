package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	stores := map[string]int {
		"quarter_of_a_dozen" : 3,
		"half_of_a_dozen" :	6,
		"dozen" :	12,
		"small_gross" :	120,
		"gross" : 	144,
		"great_gross" :	1728,
	}

	return stores
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
		return false
	}

	bill[item] += units[unit]

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
		return false
	}

	if _, ok := bill[item]; !ok {
		return false
	}

	num := bill[item] - units[unit] 
	if num < 0 {
		return false
	} else if num == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] = num
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if _, ok := bill[item]; !ok {
		return 0, false
	}

	return bill[item], true
}
