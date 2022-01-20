/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

package leetcode

// @lc code=start

/* using byte */
func convert(s string, numRows int) string {
	rows := make([][]byte, len(s))

	var row int
	goingDown := false

	for i := 0; i < len(s); i++ {
		curRow := rows[row]
		curRow = append(curRow, s[i])
		rows[row] = curRow
		if row == numRows-1 || row == 0 {
			goingDown = !goingDown
		}

		if goingDown {
			row++
		} else {
			row--
		}
	}

	var flatten []byte

	for _, r := range rows {
		flatten = append(flatten, r...)
	}

	return string(flatten)
}

/* add string (slower than above)
func convert(s string, numRows int) string {
	zigzag := make([]string, len(s))

	var row int
	goingDown := false

	for _, c := range s {
		zigzag[row] += string(c)

		if row == numRows-1 || row == 0 {
			goingDown = !goingDown
		}

		if goingDown {
			row += 1
		} else {
			row -= 1
		}
	}

	var res string

	for i := range zigzag {
		res += zigzag[i]
	}

	return res
}
*/

// @lc code=end
