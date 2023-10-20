package data

import (
	"github.com/curio-research/keystone/state"
)

// MAKE SURE TO ADD 'id' FIELD TO ALL SCHEMAS!!!

// randomness variable
type LocalRandSeedSchema struct {
	RandValue int
	Id        int `gorm:"primaryKey"`
}

type Weather int

const (
	Sunny Weather = 1
	Windy Weather = 2
)

type GameSchema struct {
	Id      int `gorm:"primaryKey"`
	Weather Weather
}

type Terrain bool

const (
	Ground   Terrain = true
	Obstacle Terrain = false
)

type TileSchema struct {
	Id       int       `gorm:"primaryKey"`
	Position state.Pos `gorm:"embedded"`
	Terrain  Terrain
}

type ObstacleTileSchema struct {
	Id       int       `gorm:"primaryKey"`
	Position state.Pos `gorm:"embedded"`
}

type PlayerSchema struct {
	Id        int       `gorm:"primaryKey"`
	Position  state.Pos `gorm:"embedded"`
	Resources int
	PlayerId  int
}

type ProjectileSchema struct {
	Id       int       `gorm:"primaryKey"`
	Position state.Pos `gorm:"embedded"`
}

type AnimalSchema struct {
	Id       int       `gorm:"primaryKey"`
	Position state.Pos `gorm:"embedded"`
}

// ----------------------------
//
//	table accessors
//
// ----------------------------

var (
	Game            = state.NewTableAccessor[GameSchema]()
	LocalRandomSeed = state.NewTableAccessor[LocalRandSeedSchema]()
	Projectile      = state.NewTableAccessor[ProjectileSchema]()
	Tile            = state.NewTableAccessor[TileSchema]()
	Player          = state.NewTableAccessor[PlayerSchema]()
	Animal          = state.NewTableAccessor[AnimalSchema]()
)

var TableSchemasToAccessors = map[interface{}]state.ITable{
	&GameSchema{}:          Game,
	&LocalRandSeedSchema{}: LocalRandomSeed,
	&ProjectileSchema{}:    Projectile,
	&TileSchema{}:          Tile,
	&PlayerSchema{}:        Player,
	&AnimalSchema{}:        Animal,
}
