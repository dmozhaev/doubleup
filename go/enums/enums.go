package enums

type SmallLargeChoice string

const (
    Small SmallLargeChoice = "SMALL"
    Large SmallLargeChoice = "LARGE"
)

type GameResult string

const (
    Win  GameResult = "WIN"
    Lose GameResult = "LOSE"
    Draw GameResult = "DRAW"
)

type AuditOperation string

const (
	Create AuditOperation = "CREATE"
	Update AuditOperation = "UPDATE"
	Delete AuditOperation = "DELETE"
)
