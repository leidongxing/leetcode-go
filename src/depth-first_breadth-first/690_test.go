package depth_first_breadth_first

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

//员工的重要性
func getImportance(employees []*Employee, id int) int {
	m := map[int]*Employee{}
	for _, employee := range employees {
		m[employee.Id] = employee
	}
	return dfsFind(id, m)

}
func dfsFind(id int, m map[int]*Employee) int {
	if m[id] == nil {
		return 0
	} else {
		sum := m[id].Importance
		for _, subId := range m[id].Subordinates {
			sum += dfsFind(subId, m)
		}
		return sum
	}
}
