package beatnik

// Default MIDI note mapping.
var midiDrums = map[string]byte{
	"K": 36, //Kick

	"SS": 37, // Snare sidestick
	"S":  38, // Snare

	"HC": 42, // Hi-hat closed
	"HO": 46, // Hi-hat open
	"HP": 44, // Hi-hat pedal

	"C1": 49, // Crash 1
	"C2": 57, // Crash 2
	"C3": 52, // Chinese cymbal
	"C4": 55, // Splash cymbal

	"R1": 51, // Ride 1
	"R2": 59, // Ride 2
	"RB": 53, // Ride 2

	"T1": 50, // Tom 1
	"T2": 48, // Tom 2
	"T3": 47, // Tom 3
	"T4": 45, // Tom 4
	"T5": 43, // Tom 5
	"T6": 41, // Tom 6
}

// EZdrummer 2 note mapping.
var ezDrummer = map[string]byte{
	"K": 36, // Kick

	"S":  38, // Snare
	"SR": 40, // Snare rimshot
	"SS": 37, // Snare sidestick

	"HC":  22, // Hi-hat closed (edge)
	"HCT": 42, // Hi-hat closed (tip)
	"HT":  62, // Hi-hat tight (edge)
	"HTT": 63, // Hi-hat tight (tip)
	"HO1": 24, // Hi-hat open 1
	"HO2": 25, // Hi-hat open 2
	"HO3": 26, // Hi-hat open 3
	"HO4": 60, // Hi-hat open 4
	"HO5": 17, // Hi-hat open 5
	"HP":  21, // Hi-hat pedal (closed)
	"HPO": 23, // Hi-hat pedal (open)
	"HS":  65, // Hi-hat seq hits

	"C1":  55, // Crash 1
	"C1M": 56, // Crash 1 muted
	"C2":  49, // Crash 2
	"C2M": 50, // Crash 2 muted
	"C3":  57, // Crash 3
	"C3M": 58, // Crash 3 muted
	"C4":  52, // Crash 4
	"C4M": 54, // Crash 4 muted

	"R":  59, // Ride
	"RB": 53, // Ride bell
	"RW": 51, // Ride bow
	"RM": 83, // Ride muted

	"T1":  48, // Tom 1
	"T1R": 82, // Tom 1 rimshot
	"T2":  47, // Tom 2
	"T2R": 80, // Tom 2 rimshot
	"T3":  45, // Tom 3
	"T3R": 78, // Tom 3 rimshot
	"T4":  43, // Tom 4
	"T4R": 75, // Tom 4 rimshot
	"T5":  41, // Tom 5
	"T5R": 73, // Tom 5 rimshot
}
