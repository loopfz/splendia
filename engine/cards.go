package main

var (
	DeckTier1 = CardDeck{
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  3,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 3,
				RED:   0,
				BLUE:  1,
				GREEN: 0,
				BLACK: 1,
			},
			VP: 0,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   1,
				BLUE:  1,
				GREEN: 2,
				BLACK: 1,
			},
			VP: 0,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   1,
				BLUE:  1,
				GREEN: 1,
				BLACK: 1,
			},
			VP: 0,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  2,
				GREEN: 2,
				BLACK: 1,
			},
			VP: 0,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  0,
				GREEN: 4,
				BLACK: 0,
			},
			VP: 1,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  2,
				GREEN: 0,
				BLACK: 2,
			},
			VP: 0,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   2,
				BLUE:  0,
				GREEN: 0,
				BLACK: 1,
			},
			VP: 0,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  0,
				GREEN: 0,
				BLACK: 3,
			},
			VP: 0,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 1,
				RED:   0,
				BLUE:  0,
				GREEN: 0,
				BLACK: 2,
			},
			VP: 0,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   4,
				BLUE:  0,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 1,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  0,
				GREEN: 2,
				BLACK: 2,
			},
			VP: 0,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 1,
				RED:   2,
				BLUE:  0,
				GREEN: 1,
				BLACK: 1,
			},
			VP: 0,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 1,
				RED:   2,
				BLUE:  0,
				GREEN: 2,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   1,
				BLUE:  1,
				GREEN: 3,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 1,
				RED:   1,
				BLUE:  0,
				GREEN: 1,
				BLACK: 1,
			},
			VP: 0,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 2,
				RED:   0,
				BLUE:  1,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 0,
				RED:   3,
				BLUE:  0,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  0,
				GREEN: 0,
				BLACK: 4,
			},
			VP: 1,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 1,
				RED:   1,
				BLUE:  1,
				GREEN: 0,
				BLACK: 2,
			},
			VP: 0,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 0,
				RED:   2,
				BLUE:  2,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 1,
				RED:   1,
				BLUE:  1,
				GREEN: 0,
				BLACK: 1,
			},
			VP: 0,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 0,
				RED:   2,
				BLUE:  1,
				GREEN: 0,
				BLACK: 2,
			},
			VP: 0,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 1,
				RED:   0,
				BLUE:  3,
				GREEN: 1,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 3,
				RED:   0,
				BLUE:  0,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 1,
				RED:   0,
				BLUE:  1,
				GREEN: 1,
				BLACK: 1,
			},
			VP: 0,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 2,
				RED:   2,
				BLUE:  0,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 1,
				RED:   1,
				BLUE:  0,
				GREEN: 0,
				BLACK: 3,
			},
			VP: 0,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 4,
				RED:   0,
				BLUE:  0,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 1,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 2,
				RED:   0,
				BLUE:  0,
				GREEN: 1,
				BLACK: 2,
			},
			VP: 0,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 2,
				RED:   0,
				BLUE:  1,
				GREEN: 1,
				BLACK: 1,
			},
			VP: 0,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  2,
				GREEN: 1,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 0,
				RED:   1,
				BLUE:  0,
				GREEN: 2,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  4,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 1,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 2,
				RED:   1,
				BLUE:  2,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 0,
				RED:   3,
				BLUE:  0,
				GREEN: 1,
				BLACK: 1,
			},
			VP: 0,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 2,
				RED:   0,
				BLUE:  0,
				GREEN: 2,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 1,
				RED:   1,
				BLUE:  2,
				GREEN: 1,
				BLACK: 0,
			},
			VP: 0,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 1,
				RED:   1,
				BLUE:  1,
				GREEN: 1,
				BLACK: 0,
			},
			VP: 0,
		},
	}

	DeckTier2 = CardDeck{
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   5,
				BLUE:  0,
				GREEN: 0,
				BLACK: 2,
			},
			VP: 2,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   2,
				BLUE:  0,
				GREEN: 3,
				BLACK: 2,
			},
			VP: 1,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 6,
				RED:   0,
				BLUE:  0,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 3,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   4,
				BLUE:  0,
				GREEN: 1,
				BLACK: 2,
			},
			VP: 2,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   5,
				BLUE:  0,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 2,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 0,
				RED:   2,
				BLUE:  3,
				GREEN: 0,
				BLACK: 3,
			},
			VP: 1,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 2,
				RED:   2,
				BLUE:  0,
				GREEN: 0,
				BLACK: 3,
			},
			VP: 1,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 3,
				RED:   0,
				BLUE:  0,
				GREEN: 0,
				BLACK: 5,
			},
			VP: 2,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  0,
				GREEN: 0,
				BLACK: 5,
			},
			VP: 2,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 0,
				RED:   6,
				BLUE:  0,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 3,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 1,
				RED:   0,
				BLUE:  4,
				GREEN: 2,
				BLACK: 0,
			},
			VP: 2,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 4,
				RED:   0,
				BLUE:  2,
				GREEN: 0,
				BLACK: 1,
			},
			VP: 2,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 3,
				RED:   3,
				BLUE:  0,
				GREEN: 2,
				BLACK: 0,
			},
			VP: 1,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  5,
				GREEN: 3,
				BLACK: 0,
			},
			VP: 2,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 2,
				RED:   0,
				BLUE:  3,
				GREEN: 0,
				BLACK: 2,
			},
			VP: 1,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  0,
				GREEN: 6,
				BLACK: 0,
			},
			VP: 3,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  0,
				GREEN: 5,
				BLACK: 0,
			},
			VP: 2,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 0,
				RED:   2,
				BLUE:  1,
				GREEN: 4,
				BLACK: 0,
			},
			VP: 2,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 5,
				RED:   0,
				BLUE:  0,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 2,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 3,
				RED:   0,
				BLUE:  2,
				GREEN: 2,
				BLACK: 0,
			},
			VP: 1,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 0,
				RED:   3,
				BLUE:  0,
				GREEN: 5,
				BLACK: 0,
			},
			VP: 2,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 3,
				RED:   0,
				BLUE:  0,
				GREEN: 3,
				BLACK: 2,
			},
			VP: 1,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  0,
				GREEN: 0,
				BLACK: 6,
			},
			VP: 3,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 2,
				RED:   1,
				BLUE:  0,
				GREEN: 0,
				BLACK: 4,
			},
			VP: 2,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  2,
				GREEN: 3,
				BLACK: 3,
			},
			VP: 1,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   3,
				BLUE:  2,
				GREEN: 2,
				BLACK: 0,
			},
			VP: 1,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 5,
				RED:   0,
				BLUE:  3,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 2,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  6,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 3,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  5,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 2,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  5,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 2,
		},
	}

	DeckTier3 = CardDeck{
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 3,
				RED:   3,
				BLUE:  0,
				GREEN: 3,
				BLACK: 5,
			},
			VP: 3,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 6,
				RED:   0,
				BLUE:  3,
				GREEN: 0,
				BLACK: 3,
			},
			VP: 4,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 7,
				RED:   0,
				BLUE:  0,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 4,
		},
		{
			Color: BLUE,
			Cost: ColorCount{
				WHITE: 7,
				RED:   0,
				BLUE:  3,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 5,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 5,
				RED:   3,
				BLUE:  3,
				GREEN: 0,
				BLACK: 3,
			},
			VP: 3,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 3,
				RED:   0,
				BLUE:  6,
				GREEN: 3,
				BLACK: 0,
			},
			VP: 4,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  7,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 4,
		},
		{
			Color: GREEN,
			Cost: ColorCount{
				WHITE: 3,
				RED:   0,
				BLUE:  7,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 5,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   5,
				BLUE:  3,
				GREEN: 3,
				BLACK: 3,
			},
			VP: 3,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 3,
				RED:   3,
				BLUE:  0,
				GREEN: 0,
				BLACK: 6,
			},
			VP: 4,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  0,
				GREEN: 0,
				BLACK: 7,
			},
			VP: 4,
		},
		{
			Color: WHITE,
			Cost: ColorCount{
				WHITE: 3,
				RED:   0,
				BLUE:  0,
				GREEN: 7,
				BLACK: 0,
			},
			VP: 5,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 3,
				RED:   3,
				BLUE:  3,
				GREEN: 5,
				BLACK: 0,
			},
			VP: 3,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 0,
				RED:   6,
				BLUE:  0,
				GREEN: 3,
				BLACK: 3,
			},
			VP: 4,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 0,
				RED:   7,
				BLUE:  0,
				GREEN: 0,
				BLACK: 0,
			},
			VP: 4,
		},
		{
			Color: BLACK,
			Cost: ColorCount{
				WHITE: 0,
				RED:   7,
				BLUE:  0,
				GREEN: 0,
				BLACK: 3,
			},
			VP: 5,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 3,
				RED:   0,
				BLUE:  5,
				GREEN: 3,
				BLACK: 3,
			},
			VP: 3,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 0,
				RED:   3,
				BLUE:  3,
				GREEN: 6,
				BLACK: 0,
			},
			VP: 4,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 0,
				RED:   0,
				BLUE:  0,
				GREEN: 7,
				BLACK: 0,
			},
			VP: 4,
		},
		{
			Color: RED,
			Cost: ColorCount{
				WHITE: 0,
				RED:   3,
				BLUE:  0,
				GREEN: 7,
				BLACK: 0,
			},
			VP: 5,
		},
	}
)