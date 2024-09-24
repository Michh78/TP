package exo

import "sort"

// Définition de la structure Interval
type Interval struct {
    start int
    end   int
}

// Ft_non_overlap calcule le nombre d'intervalles à supprimer
func Ft_non_overlap(intervals [][]int) int {
    // Conversion en slice d'Interval pour faciliter le tri
    intervalSlice := make([]Interval, len(intervals))
    for i, interval := range intervals {
        intervalSlice[i] = Interval{interval[0], interval[1]}
    }

    // Tri par ordre croissant de début d'intervalle
    sort.Slice(intervalSlice, func(i, j int) bool {
        return intervalSlice[i].start < intervalSlice[j].start
    })

    // Initialisation
    prevEnd := intervalSlice[0].end
    removeCount := 0

    for i := 1; i < len(intervalSlice); i++ {
        // Si l'intervalle courant chevauche l'intervalle précédent
        if intervalSlice[i].start < prevEnd {
            removeCount++
        } else {
            // Met à jour l'intervalle précédent
            prevEnd = intervalSlice[i].end
        }
    }

    return removeCount
}
