package q786;

import java.util.PriorityQueue;

public class Question {
    public static void main(String[] args) {
        Question Que = new Question();

        int[] Numbers = {1, 2, 3, 5, 5};
        int[] Result = Que.kthSmallestPrimeFraction(Numbers, 3);

        System.out.printf("%d, %d", Result[0], Result[1]);
    }

    public int[] kthSmallestPrimeFraction(int[] arr, int k) {
        int n = arr.length;
        PriorityQueue<int[]> pq = new PriorityQueue<int[]>((x, y) -> arr[x[0]] * arr[y[1]] - arr[y[0]] * arr[x[1]]);
        for (int j = 1; j < n; ++j) {
            pq.offer(new int[]{0, j});
        }
        for (int i = 1; i < k; ++i) {
            int[] frac = pq.poll();
            int x = frac[0], y = frac[1];
            if (x + 1 < y) {
                pq.offer(new int[]{x + 1, y});
            }
        }
        return new int[]{arr[pq.peek()[0]], arr[pq.peek()[1]]};
    }
}
