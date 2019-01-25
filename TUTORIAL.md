# Beatnik Language Tutorial

Example:

```
bpm:120

HC,K. HC.   HC,S. HC.
HC,K. HC,K. HC,S. HC.
HC,K. HC.   HC,S. HC.
HC,K. HC,K. HC,S. HC.
```

Let's break it down.

## Tempo

The first part is the tempo. Syntax is simple: `bpm:120` for 120 BPM.

Varying tempo is currently unsupported.

## Hits

A **hit** is a bunch of drums played at the same time.

A hit contains drums to play, their velocities (volumes) and a duration. Hits are separated by spaces.

Example hit: `HC,K.`

## Drums

`HC` means hi-hat closed, `K` means kick, `S` means snare (see the full list below).

Drums separated by commas are played at the same time. `HC,K` means hi-hat closed played with kick drum.

You can have any number of different drums on the same hit.

## Hit Duration

Duration means the interval between a hit and the next hit after it.

Use `.` and `~` signs to set a hit's duration. Hit duration follows its drums without spaces.

Available durations:

* `nothing` 1/4 bar
* `~~` 1 bar
* `~` 1/2 bar
* `.` 1/8 bar
* `..` 1/16 bar
* `...` 1/32 bar
* `....` 1/64 bar
* `.....` 1/128 bar

Example: `HC,K.` means hi-hat and kick, 1/8 bar.

### Triplets

Adding `>` to the duration will make it a triplet, multiplying the duration by 2/3.

## Velocity

Velocity means the volume of a single drum hit.

Use `+` and `-` signs on each drum to set its velocity.

Available velocities:

* `nothing` Forte (regular)
* `++` Fortississimo
* `+` Fortissimo
* `-` Mezzo forte
* `--` Mezzo piano
* `---` Piano
* `----` Pianissimo
* `-----` Pianississimo

Example: `S+,HC` means snare in fortissimo and hi-hat in forte played at the same time.

## Spacing

Any amount and type of spaces is allowed between hits. That means spaces, new lines, tabs. A single hit (drums+duration) should not have spaces in it.

Take `K K K` for example (3 consecutive kicks with 1/4 bar between them). It is the same as:

    K    K          K

Or:

    K
    K
    K

## Comments

You can write comments with the `#` sign. Meaning everything that follows `#` is ignored on that line.

Example:

```
# Verse
HC,K. HC.   HC,S. HC.
HC,K. HC,K. HC,S. HC.

# Chorus
HC,K. HC.   HC,S. HC.
HC,K. HC.   HC,S. HC.  # Check how this part sits with the bass
```

## Drum Symbols

Currently written for EZDrummer 2. More drum sets will be supported in the future.

|Symbol|Drum|
|--|--|
|`K`|Kick|
|||
|`S`|Snare|
|`SR`|Snare rimshot|
|`SS`|Snare sidestick|
|||
|`HC`|Hi-hat closed (edge)|
|`HCT`|Hi-hat closed (tip)|
|`HT`|Hi-hat tight (edge)|
|`HTT`|Hi-hat tight (tip)|
|`HO1`|Hi-hat open 1|
|`HO2`|Hi-hat open 2|
|`HO3`|Hi-hat open 3|
|`HO4`|Hi-hat open 4|
|`HO5`|Hi-hat open 5|
|`HP`|Hi-hat pedal (closed)|
|`HPO`|Hi-hat pedal (open)|
|`HS`|Hi-hat seq hits|
|||
|`C1`|Crash 1|
|`C1M`|Crash 1 muted|
|`C2`|Crash 2|
|`C2M`|Crash 2 muted|
|`C3`|Crash 3|
|`C3M`|Crash 3 muted|
|`C4`|Crash 4|
|`C4M`|Crash 4 muted|
|||
|`R`|Ride|
|`RB`|Ride bell|
|`RW`|Ride bow|
|`RM`|Ride muted|
|||
|`T1`|Tom 1|
|`T1R`|Tom 1 rimshot|
|`T2`|Tom 2|
|`T2R`|Tom 2 rimshot|
|`T3`|Tom 3|
|`T3R`|Tom 3 rimshot|
|`T4`|Tom 4|
|`T4R`|Tom 4 rimshot|
|`T5`|Tom 5|
|`T5R`|Tom 5 rimshot|
