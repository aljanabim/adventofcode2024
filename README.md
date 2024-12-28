# Advent of Code solved in Golang

Inspiration for graph algorithm of day 6 [https://github.com/APMorto/aoc2024/tree/master](https://github.com/APMorto/aoc2024/tree/master)

## To explore & Learnings

-   Day 6
    -   Part 2: Speed up with Go channels
-   Day 7
    -   Part 1: Improve with generators for combinations
-   Day 16
    -   Learn Dijkstra
-   Day 18
    -   Learned DFS with a queue
-   Day 19
    -   Learned to stick to strings when working with combinatorics, numbers can get high
    -   To find every combination in DFS, put the return outside the for-loop
    -   Cache a recursive function after the recursive call mapping input to its return value.
-   Day 20
    -   When computing an alternative path, remember to subtract the new path cost from the original path to get the savings
    -   In a discrete grid, stick to Abs() to compute distances, don't use continuous approaches like pythagoras theorem which account for diagonal distances.
