package engine

// Compound names prefixes. name + 'ane' is alkane. name + 'yl' is alkyl
var names = []string{
	"meth",
	"eth",
	"prop",
	"but",
	"pent",
	"hex",
	"hept",
	"oct",
	"non",
	"dec",
	"undec",
	"dodec",
}

// Prefixes that notate the occurrence of alkyl group
var prefixes = []string{
	"",
	"di",
	"tri",
	"tetra",
	"penta",
	"hexa",
	"hepta",
	"octa",
	"nona",
	"deca",
	"undeca",
	"dodeca",
}

// Prefixe that notate the occurrence of complex alkyl group
var complexPrefixes = []string{
	"",
	"bis",
	"tris",
	"tetrakis",
	"pentakis",
	"hexakis",
	"heptakis",
	"octakis",
	"nonakis",
	"decakis",
	"undecakis",
	"dodecakis",
}

// Array of positions that constructs chain
type Chain []int