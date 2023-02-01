package inventory

type Getter interface {
	GetAll() []InvItem
}
type Adder interface {
	Add(invItem InvItem)
}
type Deleter interface {
	Remove(invItem InvItem)
}
type renamer interface {
	rename(invItem InvItem, name string)
}
type relocater interface {
	relocate(invItem InvItem, location string)
}

type InvItem struct {
	Name     string `json:"Item Name"`
	Location string `json:"Location"`
}

type Container struct {
	Name       string `json:"Cont Name"`
	InvItems   []InvItem
	Containers []Container
}

func New() *Container {
	return &Container{
		InvItems: []InvItem{},
	}
}

func (r *Container) Add(invItem InvItem) {
	r.InvItems = append(r.InvItems, invItem)
}

func (r *Container) GetAll() []InvItem {
	return r.InvItems
}

func (r *Container) Remove(invItem InvItem) {
	// find way to delete item
	r.InvItems = delete(r.InvItems, invItem)

}

func (r *Container) rename(invItem InvItem, name string) {
	// find way to rename item
	invItem.Name = name
}

func (r *Container) relocate(invItem InvItem, location string) {
	// find way to relocate item
	invItem.Location = location
}


