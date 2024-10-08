package patterns

var Glider = [][2]int{
	{0, 1},
	{1, 2},
	{2, 0},
	{2, 1},
	{2, 2},
}

var LWSS = [][2]int{
	{10, 11}, {10, 12}, {10, 13},
	{11, 10}, {11, 14},
	{12, 14},
	{13, 10}, {13, 13},

	{10, 21}, {10, 22}, {10, 23},
	{11, 20}, {11, 24},
	{12, 24},
	{13, 20}, {13, 23},
}

var Star = [][2]int{
	{15, 30}, {15, 31}, {15, 32},
	{16, 30}, {16, 32},
	{17, 30}, {17, 32},
}
