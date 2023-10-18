package constants

// ------------------------
// game constants
// ------------------------

var (
	// game tick rate (milliseconds)
	TickRate int = 100

	// rate to save world to database (seconds)
	// 0 means never save
	SaveStateToDatabaseRate int = 0

	// time in milliseconds per turn timer count down
	WeatherChangeIntervalMs = 20_000

	MaxNPCInWorld int = 100

	// total players count
	TotalPlayersCount = 5

	// milliseconds per small tile troop movement
	SmallTileMoveTimeMs int = 40

	// milliseconds after attack that the troop will be idle before movement
	AttackIdleTimeMs int = 2000

	// in large tiles
	WorldWidth  int = 7
	WorldHeight int = 7

	// entity for global randomness seed
	RandomnessEntity int = 6969
	GameEntity       int = 200
)
