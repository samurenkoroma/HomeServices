package types

type PhysicalObjectID string

func (i PhysicalObjectID) String() any {
	return string(i)
}
