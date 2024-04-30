package enums

type SmallLargeChoice string

const (
    Small SmallLargeChoice = "SMALL"
    Large SmallLargeChoice = "LARGE"
)

type GameResult string

const (
    W  GameResult = "W"
    L GameResult = "L"
)

type AuditOperation string

const (
	Select AuditOperation = "SELECT"
	Insert AuditOperation = "INSERT"
	Update AuditOperation = "UPDATE"
)
