/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
package main

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || j == 0 || i == m-1 || j == n-1) && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}

}

func dfs(board [][]byte, i, j int) {
	m := len(board)
	n := len(board[0])
	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == 'X' || board[i][j] == '#' {
		return
	}
	board[i][j] = '#'
	dfs(board, i-1, j)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)

}

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || j == 0 || i == m-1 || j == n-1) && board[i][j] == 'O' {
				bfs(board, i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}

}

type pos struct {
	I int
	J int
}

func bfs(board [][]byte, i, j int) {
	m := len(board)
	n := len(board[0])
	if board[i][j] != 'O' {
		return
	}
	board[i][j] = '#'
	queue := make([]*pos, 0)
	queue = append(queue, &pos{i, j})
	for len(queue) > 0 {
		l := len(queue)
		for _, p := range queue[:l] {
			if p.I+1 < m && board[p.I+1][p.J] == 'O' {
				board[p.I+1][p.J] = '#'
				queue = append(queue, &pos{p.I + 1, p.J})
			}
			if p.I-1 >= 0 && board[p.I-1][p.J] == 'O' {
				board[p.I-1][p.J] = '#'
				queue = append(queue, &pos{p.I - 1, p.J})
			}
			if p.J+1 < n && board[p.I][p.J+1] == 'O' {
				board[p.I][p.J+1] = '#'
				queue = append(queue, &pos{p.I, p.J + 1})
			}
			if p.J-1 >= 0 && board[p.I][p.J-1] == 'O' {
				board[p.I][p.J-1] = '#'
				queue = append(queue, &pos{p.I, p.J - 1})
			}
		}
		queue = queue[l:]
	}

}

// @lc code=end
