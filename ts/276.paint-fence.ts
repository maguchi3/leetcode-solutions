function numWays(n: number, k: number): number {
  const memo = new Map<number, number>();

  return totalWays(n, k, memo);
}

function totalWays(i: number, k: number, memo: Map<number, number>): number {
  if (i === 1) return k;
  if (i === 2) return k * k;

  if (memo.has(i)) return memo.get(i);

  const total =
    (k - 1) * (totalWays(i - 1, k, memo) + totalWays(i - 2, k, memo));

  memo.set(i, total);

  return memo.get(i);
}
