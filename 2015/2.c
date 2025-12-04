#include <limits.h>
#include <stdint.h>
#include <stdio.h>
#include <string.h>
#include <ctype.h>

int main(void)
{
    const char* filename = "input.txt";

    FILE* file = fopen(filename, "r");

    char ch;

    unsigned int dims[3] = {0};
    unsigned int idx = 0;

    unsigned int area = 0;
    unsigned int ribbon = 0;

    while ((ch = fgetc(file)) != EOF)
    {
        if (isdigit(ch))
        {
            dims[idx] *= 10; // make space new digit
            dims[idx] += ch - '0';
        } 
        else if (ch == 'x')
        {
            idx = (idx + 1) % 3;
        } 
        else if (ch == '\n')
        {
            unsigned int l = dims[0];
            unsigned int w = dims[1];
            unsigned int h = dims[2];
            
            unsigned int min = UINT_MAX;

            unsigned int lw = l * w;
            if (lw < min) min = lw;

            unsigned int wh = w * h;
            if (wh < min) min = wh;

            unsigned int hl = h * l;
            if (hl < min) min = hl;
            

            area += 2*lw + 2*wh + 2*hl + min;
            idx = 0;
            memset(dims, 0, 3 * sizeof(dims[0]));
        }
        printf("%c", ch);
    }
        printf("\n");

    printf("Length is %d, width is %d, height is %d\n", dims[0], dims[1], dims[2]);
    printf("Area is %d\n", area);

    return 0;
}
