package ragel

//go:generate ragel -Z -G2 match.ragel
//go:generate ragel -Z -G2 index.ragel
//go:generate ragel -Z -G2 uuid_match.ragel
//go:generate ragel -Z -G2 uuid_index.ragel
//go:generate ragel -Z -G2 aa_match.ragel

// Stuff ...
type Stuff struct{}
