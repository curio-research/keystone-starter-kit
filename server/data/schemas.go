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

type PlayerSchema struct {
	Id        int       `gorm:"primaryKey"`
	Position  state.Pos `gorm:"embedded"`
	Resources int
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
//      table accessors
// ----------------------------

var (
	Game            = state.NewTableAccessor[GameSchema]()
	Tile            = state.NewTableAccessor[TileSchema]()
	Player          = state.NewTableAccessor[PlayerSchema]()
	LocalRandomSeed = state.NewTableAccessor[LocalRandSeedSchema]()
	Animal          = state.NewTableAccessor[AnimalSchema]()
	Projectile      = state.NewTableAccessor[ProjectileSchema]()
)

var TableSchemasToAccessors = map[interface{}]*state.TableBaseAccessor[any]{
	&GameSchema{}:          (*state.TableBaseAccessor[any])(Game),
	&TileSchema{}:          (*state.TableBaseAccessor[any])(Tile),
	&PlayerSchema{}:        (*state.TableBaseAccessor[any])(Player),
	&LocalRandSeedSchema{}: (*state.TableBaseAccessor[any])(LocalRandomSeed),
	&AnimalSchema{}:        (*state.TableBaseAccessor[any])(Animal),
	&ProjectileSchema{}:    (*state.TableBaseAccessor[any])(Projectile),
}
