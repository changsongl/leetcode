//
// Created by Changsong Li on 2021/12/27.
//

#include "stdlib.h"
#include "stdio.h"

int compare(const void *a, const void *b);

int numFriendRequests(int* ages, int agesSize);

int numFriendRequests(int* ages, int agesSize){
    int left, right, i, ans = 0, target, middle, j;

    qsort(ages, agesSize, sizeof(int), compare);

    for (i = agesSize - 1; i >= 0; i--) {
        right = i, left = 0;
        target = ages[right] / 2 + 7;

        while (right > left){
            middle = (right + left) / 2;

            if (ages[middle] > target) {
                right = middle;
            }else{
                left = middle + 1;
            }
        }

        ans += i - right;

        for (j = i - 1; j >= 0 && ages[j] == ages[i]; j--){
            ans += i - right;
        }

        i = j + 1;
    }

    return ans;
}

int compare(const void *a, const void *b) {
    return *((int *) a) - *((int *) b);
}
