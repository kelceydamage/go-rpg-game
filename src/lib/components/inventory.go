package components

type Container interface {
	StoreItem(item Item)
	GetItem(index uint16) *Item
}

type Inventory struct {
	Storage map[uint16]*Item
	MaxSize uint16
}

func NewInventory(maxSize uint16) *Inventory {
	return &Inventory{
		Storage: make(map[uint16]*Item),
		MaxSize: maxSize,
	}
}

func (i *Inventory) StoreItem(item Item) {
	i.Storage[i.GetNextInventoryID()] = &item
}

func (i *Inventory) GetNextInventoryID() uint16 {
	return uint16(len(i.Storage))
}

func (i *Inventory) GetItem(index uint16) *Item {
	return i.Storage[index]
}

type Item interface {
	GetIsFlaggedForRemoval() bool
	SetFlaggedForRemoval(value bool)
	GetIsInContainer() bool
	SetIsInContainer(value bool)
	GetTypeID() uint16
	SetTypeID(typeID uint16)
}

type BasicItem struct {
	isFlaggedForRemoval bool
	isInContainer       bool
	typeID              uint16
}

func (b *BasicItem) GetIsFlaggedForRemoval() bool {
	return b.isFlaggedForRemoval
}

func (b *BasicItem) SetFlaggedForRemoval(value bool) {
	b.isFlaggedForRemoval = value
}

func (b *BasicItem) GetTypeID() uint16 {
	return b.typeID
}

func (b *BasicItem) SetTypeID(typeID uint16) {
	b.typeID = typeID
}

func (b *BasicItem) SetIsInContainer(value bool) {
	b.isInContainer = true
}

func (b *BasicItem) GetIsInContainer() bool {
	return b.isInContainer
}
