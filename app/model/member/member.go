package m_member

type Member struct {
	model.Model
	Uid uint
	Email string
	Level uint
	Role string
	State string
}
