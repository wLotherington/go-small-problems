// kyu5
// https://www.codewars.com/kata/587136ba2eefcb92a9000027/

package kata

import "fmt"

type SnakesLadders struct {
  turn     int
  pos      []int
  snakes   map[int]int
  ladders  map[int]int
  gameOver bool
}

func (sl *SnakesLadders) NewGame() {
  sl.turn = 0
  sl.pos = []int{0, 0}
  sl.gameOver = false
  sl.snakes = map[int]int{
    16: 6,
    46: 25,
    49: 11,
    62: 19,
    64: 60,
    74: 53,
    89: 68,
    92: 88,
    95: 75,
    99: 80,
  }
  sl.ladders = map[int]int{
    2:  38,
    7:  14,
    8:  31,
    15: 26,
    21: 42,
    28: 84,
    36: 44,
    51: 67,
    71: 91,
    78: 98,
    87: 94,
  }
}

func (sl *SnakesLadders) Play(die1, die2 int) string {
  if sl.gameOver {
    return "Game over!"
  }

  p := sl.turn
  dies := die1 + die2

  sl.pos[p] += dies

  if sl.pos[p] == 100 {
    sl.gameOver = true
    return fmt.Sprintf("Player %d Wins!", p+1)
  } else if sl.pos[p] > 100 {
    sl.pos[p] = 200 - sl.pos[p]
  }

  if sl.ladders[sl.pos[p]] > 0 {
    sl.pos[p] = sl.ladders[sl.pos[p]]
  }

  if sl.snakes[sl.pos[p]] > 0 {
    sl.pos[p] = sl.snakes[sl.pos[p]]
  }

  if die1 != die2 {
    sl.turn = 1 - sl.turn
  }

  return fmt.Sprintf("Player %d is on square %d", p+1, sl.pos[p])
}
