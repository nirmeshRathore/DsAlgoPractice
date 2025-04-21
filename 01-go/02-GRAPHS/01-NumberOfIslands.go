/*
Question - https://www.geeksforgeeks.org/problems/find-the-number-of-islands/1
*/

package main
import "fmt"

type Pair struct {
    row, col int
}

func bfs (row, col int, vis*[][]bool, graph[][] byte) {
    n := len(graph)
    m := len(graph[0])
    
    queue := []Pair{{row, col}}
    (*vis)[row][col] = true
    
    for (len(queue) > 0) {
        // Get front from the queue & pop the element
        curr := queue[0]
        queue = queue[1:]
        
        r := curr.row
        c := curr.col

        // Loop through all the 8 valid neighbour
        for deltaRow := -1 ; deltaRow <= 1; deltaRow++  {
            for deltaCol := -1; deltaCol<=1; deltaCol++ {
                newRow := deltaRow + r
                newCol := deltaCol + c
                
                if ( (newRow >= 0 && newRow < n) && (newCol >= 0 && newCol < m) && (!(*vis)[newRow][newCol]) && (graph[newRow][newCol] == '1') ) {
                    (*vis)[newRow][newCol] = true
                    queue = append(queue, Pair{row:newRow, col:newCol})
                }
            }
        }
        
    }
}

func NumberOfIsland(graph[][] byte) (int) {
    n := len(graph)
    m := len(graph[0])
    
    vis := make([][]bool, n)
    for i:= range vis {
        vis[i] = make ([]bool, m)
    }
    numOfIsland := 0
    for r:=0; r<n; r++ {
        for c:=0; c<m; c++ { 
            if ( !vis[r][c] && graph[r][c] == '1') {
                numOfIsland++
                bfs(r, c, &vis, graph)
            }
        }
    }
    
    return numOfIsland
}

func main() {
  graph := [][]byte{
        {'0', '1', '1', '1', '0', '0', '0'},
        {'0', '0', '1', '1', '0', '1', '0'},
  }
  
  fmt.Println("Number of island:", NumberOfIsland(graph))
}
