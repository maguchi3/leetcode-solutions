function minMeetingRooms(intervals: number[][]): number {
  const compareFunc = (a, b) => a - b;
  const minHeap = new MinHeap(compareFunc);

  intervals.sort(([s1], [s2]) => s1 - s2);

  let maxRooms = 0;

  for (let interval of intervals) {
    if (minHeap.size > 0 && minHeap.peek() <= interval[0]) {
      minHeap.extract();
    }

    minHeap.insert(interval[1]);

    maxRooms = Math.max(maxRooms, minHeap.size);
  }

  return maxRooms;
}

class MinHeap {
  heap: number[];

  constructor(private compareFunc: (a: number, b: number) => number) {
    this.heap = [];
  }

  insert(val: number): void {
    this.heap.unshift(val);
    this.heap.sort(this.compareFunc);
  }

  extract(): number {
    if (this.size === 0) return null;
    return this.heap.shift();
  }

  peek(): number {
    if (this.size === 0) return null;
    return this.heap[0];
  }

  get size(): number {
    return this.heap.length;
  }
}
