package helper

import (
	"math/rand"

	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/state"
)

// =========================
// Algebra
// =========================

// min of 2 ints
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntSliceToInt64Slice(s []int) []int64 {
	s64 := make([]int64, len(s))

	for i, v := range s {
		s64[i] = int64(v)
	}

	return s64
}

func Int64SliceToIntSlice(s []int64) []int {
	slice := make([]int, len(s))

	for i, v := range s {
		slice[i] = int(v)
	}

	return slice
}

func IndexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func Abs(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func HashIntsByXor(a, b int) int {
	return a ^ b
}

// ---------------------------------------
// Randomness
// ---------------------------------------

func GetWorldRandomnessEntity(w state.IWorld) (int, data.LocalRandSeedSchema) {
	randomnessEntity := data.LocalRandomSeed.Get(w, constants.RandomnessEntity)

	return constants.RandomnessEntity, randomnessEntity
}

// function used by systems to get deterministic randomness
// start and end are inclusive
func GetRandomness(w state.IWorld, start int, end int) int {
	seed := data.LocalRandomSeed.Get(w, constants.RandomnessEntity).RandValue

	// generate the next deterministic random seed back into the world
	nextSeed := nextDeterministicRandom(seed)
	data.LocalRandomSeed.Set(w, constants.RandomnessEntity, data.LocalRandSeedSchema{
		RandValue: nextSeed,
	})

	// using this new seed, generate a random number from this range
	rng := rand.New(rand.NewSource(int64(nextSeed)))

	return rng.Intn(end-start+1) + start
}

func nextDeterministicRandom(seed int) int {
	// create a new rand.Rand instance with the provided seed
	rng := rand.New(rand.NewSource(int64(seed)))

	// generate the next random integer. can change this to Intn for a specific range.
	nextRandom := rng.Int()

	return nextRandom
}
