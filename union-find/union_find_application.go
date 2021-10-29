package unionfind

import (
	"sort"
)

// Detect Cycle in an Undirected Graph.
// Assume that a graph will be given an adjacent list.
func HasCycle(graph [][]int) bool {
	ds := NewDisjointSet(len(graph))
	for u := 0; u < len(graph); u++ {
		for _, v := range graph[u] {
			if v < u {
				// Edges that have been already merged will be excluded. (only if a graph is an undirected graph)
				continue
			}
			setU := ds.Find(u)
			setV := ds.Find(v)
			if setU == setV {
				return true
			}
			ds.Union(setU, setV)
		}
	}
	return false
}

func CountCycle(graph [][]int) int {
	numCycles := 0
	ds := NewDisjointSet(len(graph))
	for u := 0; u < len(graph); u++ {
		for _, v := range graph[u] {
			if v < u {
				// Edges that have been already merged will be excluded. (only if a graph is an undirected graph)
				continue
			}
			setU := ds.Find(u)
			setV := ds.Find(v)
			if setU == setV {
				numCycles++
				continue
			}
			ds.Union(setU, setV)
		}
	}
	return numCycles
}

// Job Scheduling Example Solution
// https://www.geeksforgeeks.org/job-sequencing-using-disjoint-set-union/
// Given an 2D array named by jobs, return the maximum sum of profits from non-overlapping jobs in timeline.
// jobs[i] has the pair (deadline, profit)
// where any job can be arranged anywhere in timeline under deadline and each job spend 1 unit of time to complete.
func MaxProfitSum(jobs [][]int) int {
	sort.Slice(jobs[:], func(i, j int) bool {
		if jobs[i][1] == jobs[j][1] {
			return jobs[i][0] < jobs[j][0]
		}
		return jobs[i][1] > jobs[j][1]
	})

	maxTime := 0
	for _, job := range jobs {
		if job[0] > maxTime {
			maxTime = job[0]
		}
	}

	maxSum := 0
	ds := NewDisjointSet(maxTime + 1)
	for _, job := range jobs {
		availableTime := ds.Find(job[0])
		if availableTime == 0 {
			continue
		}
		if availableTime != ds.Find(availableTime-1) {
			ds.Union(availableTime, ds.Find(availableTime-1))
			maxSum += job[1]
		}
	}
	return maxSum
}
