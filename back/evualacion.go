package main

import (
	"fmt"
	"sort"
)

func calcularAcciones(datos []DataResponse) ([]DataResponse, error) {
	actionScores := map[string]int{
		"upgraded by":       3,
		"target raised by":  2,
		"initiated by":      1,
		"reiterated by":     0,
		"target lowered by": -2,
		"downgraded by":     -3,
	}

	ratingScores := map[string]int{
		"Strong-Buy":          3,
		"Top Pick":            3,
		"Speculative Buy":     3,
		"Buy":                 2,
		"Outperform":          2,
		"Overweight":          2,
		"Positive":            2,
		"Outperformer":        2,
		"Market Outperform":   2,
		"Sector Outperform":   1,
		"In-Line":             1,
		"Hold":                0,
		"Neutral":             0,
		"Equal Weight":        0,
		"Market Perform":      0,
		"Peer Perform":        0,
		"Sector Perform":      0,
		"Sector Weight":       0,
		"Unchanged":           0,
		"Underperform":        -1,
		"Sector Underperform": -1,
		"Underweight":         -1,
		"Reduce":              -1,
		"Sell":                -2,
	}

	// Calcular el score por cada acción
	for i, d := range datos {
		score := 0

		// Score por Action
		if s, ok := actionScores[d.Action]; ok {
			score += s
		}

		// Score por Rating
		fromScore := ratingScores[d.Rating_from]
		toScore := ratingScores[d.Rating_to]
		score += toScore - fromScore

		// Score por variación de Target
		score += calculateTargetScore(d.Target_from, d.Target_to)

		datos[i].Score = &score
	}

	sort.Slice(datos, func(i, j int) bool {
		if datos[i].Score == nil && datos[j].Score == nil {
			return false
		}
		if datos[i].Score == nil {
			return false
		}
		if datos[j].Score == nil {
			return true
		}
		return *datos[i].Score > *datos[j].Score
	})

	if len(datos) == 0 {
		return nil, fmt.Errorf("no hay datos para procesar")
	}

	if len(datos) == 1 {
		return []DataResponse{datos[0], datos[0]}, nil
	}

	return []DataResponse{datos[0], datos[len(datos)-1]}, nil
}

func calculateTargetScore(from, to float64) int {
	if from == 0 {
		return 0
	}
	change := ((to - from) / from) * 100

	switch {
	case change >= 20:
		return 3
	case change >= 10:
		return 2
	case change > 0:
		return 1
	case change == 0:
		return 0
	case change > -10:
		return -1
	case change > -20:
		return -2
	default:
		return -3
	}
}
