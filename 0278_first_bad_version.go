package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func firstBadVersion(n int) int {
	var head, tail, cur = 0, n, 0

	for head < tail {
		cur = (tail-head)/2 + head
		if isBadVersion(cur) {
			tail = cur
		} else {
			head = cur + 1
		}
	}
	return tail
}
