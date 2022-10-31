#include <stdio.h>
#include <stdlib.h>
#include <time.h>

typedef unsigned char byte;
typedef byte pow256[256 / 8];
void blst_sk_inverse(pow256 ret, const pow256 a);
int blst_sk_mul_n_check(pow256 ret, const pow256 a, const pow256 b);

void randpow256(pow256 target)
{
    int r = rand();
    for (int i = 0; i < 32; i++)
    {
        target[i] = rand();
    }
}

void benchInverse()
{
    pow256 value;
    randpow256(value);
    pow256 result;

    clock_t t;
    t = clock();
    int n = 1000000;
    for (int i = 0; i < n; i++)
    {
        blst_sk_inverse(result, value);
    }
    t = clock() - t;
    double time_taken = ((double)t) / CLOCKS_PER_SEC;
    printf("Inv(x) takes %.0f ns/op\n", time_taken * 1000000000 / n);
}

void benchMul()
{
    pow256 a, b;
    randpow256(a);
    randpow256(b);

    clock_t t;
    t = clock();
    int n = 1000000;
    for (int i = 0; i < n; i++)
    {
        blst_sk_mul_n_check(a, a, b);
    }
    t = clock() - t;
    double time_taken = ((double)t) / CLOCKS_PER_SEC;
    printf("Mul(a,b) takes %.0f ns/op\n", time_taken * 1000000000 / n);
}

int main()
{
    srand(time(NULL));

    benchInverse();
    benchMul();
    return 0;
}
