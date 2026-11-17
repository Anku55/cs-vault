#include <bits/stdc++.h>
using namespace std;

int n;
vector<int> a, seg;

int query(int ind, int low, int high, int l, int r) {

    // No overlap
    if (high < l || low > r) {
        return 0;
    }

    // Complete overlap
    if (low >= l && high <= r) {
        return seg[ind];
    }

    // Partial overlap
    int mid = (low + high) / 2;

    int left = query(2 * ind + 1, low, mid, l, r);
    int right = query(2 * ind + 2, mid + 1, high, l, r);

    return left + right;
}

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
    cout << query(0, 0, n - 1, 1, 3);

    return 0;
}