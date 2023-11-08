package data

import (
	"github.com/curio-research/keystone/state"
)

// MAKE SURE TO ADD 'id' FIELD TO ALL SCHEMAS!!!

// randomness variable
type LocalRandSeedSchema struct {
	Id        int `gorm:"primaryKey;autoIncrement:false"`
	RandValue int
}

type Weather int

const (
	Sunny Weather = 1
	Windy Weather = 2
)

type GameSchema struct {
	Id      int `gorm:"primaryKey;autoIncrement:false"`
	GameId  string
	Weather Weather
}

type Terrain bool

const (
	Ground   Terrain = true
	Obstacle Terrain = false
)

type TileSchema struct {
	Id       int       `gorm:"primaryKey;autoIncrement:false"`
	Position state.Pos `gorm:"embedded"`
	Terrain  Terrain
}

type PlayerSchema struct {
	Id              int       `gorm:"primaryKey;autoIncrement:false"`
	Position        state.Pos `gorm:"embedded"`
	Resources       int
	PlayerId        int
	Base64PublicKey string
}

type ProjectileSchema struct {
	Id       int       `gorm:"primaryKey;autoIncrement:false"`
	Position state.Pos `gorm:"embedded"`
}

type AnimalSchema struct {
	Id       int       `gorm:"primaryKey;autoIncrement:false"`
	Position state.Pos `gorm:"embedded"`
}

type ResourceSchema struct {
	Id       int       `gorm:"primaryKey;autoIncrement:false"`
	Position state.Pos `gorm:"embedded"`
	Amount   int
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
	Resource        = state.NewTableAccessor[ResourceSchema]()
)

var SchemaMapping = map[interface{}]*state.TableBaseAccessor[any]{
	&GameSchema{}:          (*state.TableBaseAccessor[any])(Game),
	&LocalRandSeedSchema{}: (*state.TableBaseAccessor[any])(LocalRandomSeed),
	&ProjectileSchema{}:    (*state.TableBaseAccessor[any])(Projectile),
	&TileSchema{}:          (*state.TableBaseAccessor[any])(Tile),
	&PlayerSchema{}:        (*state.TableBaseAccessor[any])(Player),
	&AnimalSchema{}:        (*state.TableBaseAccessor[any])(Animal),
	&ResourceSchema{}:      (*state.TableBaseAccessor[any])(Resource),
}
