#include <bits/stdc++.h>
using namespace std;

int n;
vector<int> a, seg;

void build(int ind, int low, int high) {
    // Base case: leaf node
    if (low == high) {
        seg[ind] = a[low];
        return;
    }

    int mid = (low + high) / 2;

    build(2 * ind + 1, low, mid);
    build(2 * ind + 2, mid + 1, high);

    seg[ind] = seg[2 * ind + 1] + seg[2 * ind + 2];
}

int main() {
    a = {1, 3, 5, 7, 9, 11};

    n = a.size();
    seg.resize(4 * n);

    build(0, 0, n - 1);

    for (int x : seg) {
        cout << x << " ";
    }

    return 0;
}