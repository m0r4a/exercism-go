package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var birdCount int

	for i := 0; i < len(birdsPerDay); i++ {
		birdCount += birdsPerDay[i]
	}

	return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var birdCount int
	// x = 1: 0 - 6
	// x = 2: 7 - 13
	// x = 3: 14 - 20

	for i := (week - 1) * 7; i < (7 * week); i++ {
		birdCount += birdsPerDay[i]
	}

	return birdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}

	return birdsPerDay
}
