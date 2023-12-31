package constants

// ------------------------
// game constants
// ------------------------

var (
	// game tick rate (milliseconds)
	TickRate int = 10

	// animials
	AnimalCreationRate     = 500
	AnimalUpdateRate       = 100
	MaxAnimals         int = 10

	// rate to save world to database (seconds)
	// 0 means never save
	SaveStateToDatabaseRate int = 0

	// time in milliseconds per turn timer count down
	WeatherChangeIntervalMs = 20_000

	// total players count
	TotalPlayersCount = 5

	// milliseconds per small tile troop movement
	SmallTileMoveTimeMs int = 40

	// milliseconds after attack that the troop will be idle before movement
	AttackIdleTimeMs int = 2000

	// in large tiles
	WorldWidth  int = 10
	WorldHeight int = 10

	// entity for global randomness seed
	RandomnessEntity int = 6969
	GameEntity       int = 200

	// fractions to use when generating map
	FreeTilesRatio    = 0.9
	AnimalUpdateRatio = 0.2

	// speed in terms of ticks between each movement (lower the number, the faster the speed!)
	BulletSpeed = 5

	PlayerGold = 10
	AnimalGold = 5
)
