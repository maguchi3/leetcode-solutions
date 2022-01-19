/*
 * @lc app=leetcode id=127 lang=typescript
 *
 * [127] Word Ladder
 */

// @lc code=start
function ladderLength(
  beginWord: string,
  endWord: string,
  wordList: string[]
): number {
  const wordSet = new Set(wordList);

  let queue = [beginWord];
  let steps = 1;

  while (queue.length) {
    let nq = [];

    for (let word of queue) {
      if (word === endWord) return steps;

      for (let i = 0; i < word.length; i++) {
        for (let j = 0; j < 26; j++) {
          const nw =
            word.slice(0, i) + String.fromCharCode(97 + j) + word.slice(i + 1);

          if (wordSet.has(nw)) {
            nq.push(nw);
            wordSet.delete(nw);
          }
        }
      }
    }

    queue = nq;
    steps++;
  }

  return 0;
}
// @lc code=end
