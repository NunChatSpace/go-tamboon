package list

type Donater struct {
	Name   string
	Amount float64
}

type Top3Rank []Donater

func (trl Top3Rank) ToString() (result []string) {
	for i := 0; i < len(trl); i++ {
		result = append(result, trl[i].Name)
	}

	return result
}

func (trl Top3Rank) Add(donater Donater) Top3Rank {
	rank := trl.getRank(donater)
	return trl.insert(trl, rank, donater)
}

func (trl Top3Rank) getRank(donater Donater) int {
	// fmt.Printf("%+v\n", trl)
	length := len(trl)
	if length == 0 {
		return 0
	}

	limit := 0
	if length < 3 {
		limit = length
	} else {
		limit = 3
	}

	for i := 0; i < limit; i++ {
		if trl[i].Amount < donater.Amount {
			return i
		}
	}

	return length - 1
}

func (trl Top3Rank) insert(t3r Top3Rank, index int, value Donater) Top3Rank {
	if len(t3r) == index {
		return append(t3r, value)
	}

	t3r = append(t3r[:index+1], t3r[index:]...)
	t3r[index] = value

	return t3r
}
