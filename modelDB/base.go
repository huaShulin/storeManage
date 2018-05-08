package modelDB

const (
	PAGE = " ORDER BY ID LIMIT ?,? "
)

type Result struct {
	Number 			int		`gorm:"column:NUMBER" json:"NUMBER"`
}
