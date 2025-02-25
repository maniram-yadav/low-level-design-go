package iterator

type UserIterator struct {
	index int
	Users []*User
}

func newUserIterator(users []*User) *UserIterator {
	return &UserIterator{index: 0, Users: users}
}

func (ui *UserIterator) HasNext() bool {
	return ui.index < len(ui.Users)
}

func (ui *UserIterator) GetNext() *User {
	if ui.HasNext() {
		user := ui.Users[ui.index]
		ui.index++
		return user
	}
	return nil
}
