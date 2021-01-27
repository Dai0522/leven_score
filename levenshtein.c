#include "levenshtein.h"
#include <stdio.h>

// 检查utf-8字符所占字节数
#define UTF8_CHAR_LEN( byte ) ((( 0xE5000000 >> (( byte >> 3 ) & 0x1e )) & 3 ) + 1)

float leven_score(void *a, size_t a_size, void *b, size_t b_size) {
    uint8_t* a1 = (uint8_t*)a;
    uint8_t* b1 = (uint8_t*)b;
    size_t utf8_a_size = 0;
    size_t utf8_b_size = 0;
    size_t i, j, step = 0;
    for(i=0; i<a_size; i+=step) {
        step = UTF8_CHAR_LEN(a1[i]);
        utf8_a_size++;
    }
    for(i=0; i<b_size; i+=step) {
        step = UTF8_CHAR_LEN(b1[i]);
        utf8_b_size++;
    }

    size_t dist = 0;
    size_t sum_len = utf8_a_size + utf8_b_size;

    int32_t* str1_head = (int32_t*)malloc(utf8_a_size*sizeof(int32_t));
    int32_t* str1 = str1_head;
    int32_t* str2_head = (int32_t*)malloc(utf8_b_size*sizeof(int32_t));
    int32_t* str2 = str2_head;

    j = 0;
    for(i=0; i<a_size; i+=step) {
        step = UTF8_CHAR_LEN(a1[i]);
        int32_t tmp = 0;
        int k;
        for(k=0; k<step; k++) {
            tmp = (tmp<<8) | a1[i+k];
        }
        str1[j] = (int32_t) tmp;
        j++;
    }
    j = 0;
    step = 0;
    for(i=0; i<b_size; i+=step) {
        step = UTF8_CHAR_LEN(b1[i]);
        int tmp = 0;
        int k;
        for(k=0; k<step; k++) {
            tmp = (tmp<<8) | b1[i+k];
        }
        str2[j] = (int32_t) tmp;
        // printf("b--%d---%ld,%ld---\n", (int)str2[j], i, j);
        j++;
    }

    size_t *row;  /* we only need to keep one row of costs */
    size_t *end;

    /* strip common prefix */
    while (utf8_a_size > 0 && utf8_b_size > 0 && *str1 == *str2) {
        utf8_a_size--;
        utf8_b_size--;
        str1++;
        str2++;
    //printf("---%ld---%ld---%d---\n", utf8_a_size, utf8_b_size, wcsncmp(str1, str2, 1));
    }

    /* strip common suffix */
    while (utf8_a_size > 0 && utf8_b_size > 0 && str1[utf8_a_size-1] == str2[utf8_b_size-1]) {
        utf8_a_size--;
        utf8_b_size--;
    }

    //printf("---%ld---%ld---\n", utf8_a_size, utf8_b_size);
    /* catch trivial cases */
    if (utf8_a_size == 0) {
        dist = utf8_b_size;
        goto END;
    }
    if (utf8_b_size == 0) {
        dist = utf8_a_size;
        goto END;
    }

    /* make the inner cycle (i.e. str2) the longer one */
    if (utf8_a_size > utf8_b_size) {
        size_t nx = utf8_a_size;
        int32_t *sx = str1;
        utf8_a_size = utf8_b_size;
        utf8_b_size = nx;
        str1 = str2;
        str2 = sx;
    }
    /* check utf8_a_size == 1 separately */
    if (utf8_a_size == 1) {
        int32_t z = *str1;
        const int32_t *p = str2;
        for (i = utf8_b_size; i; i--) {
            if (*(p++) == z) {
                dist = utf8_b_size - 1;
                goto END;
            }
        }
        dist = utf8_b_size + 1;
        goto END;
    }
    utf8_a_size++;
    utf8_b_size++;

    /* initalize first row */
    row = (size_t*)malloc(utf8_b_size*sizeof(size_t));
    if (!row) {
        dist = (size_t)(-1);
        goto END;
    }
    end = row + utf8_b_size - 1;
    for (i = 0; i < utf8_b_size; i++)
        row[i] = i;
    //for (i = 0; i < utf8_b_size; i++)
    //    printf("---%ld---\n", row[i]);

    for (i = 1; i < utf8_a_size; i++) {
        size_t *p = row + 1;
        const int32_t char1 = str1[i - 1];
        const int32_t *char2p = str2;
        size_t D = i - 1;
        size_t x = i;
        while (p <= end) {
        if (char1 == *(char2p++))
            x = D;
        else
            x++;
        D = *p;
        if (x > D + 1)
            x = D + 1;
        *(p++) = x;
        }
    }
    dist = *end;
    free(row);

END:
    free(str1_head);
    free(str2_head);
    return sum_len == 0 ? 1 : (float)(sum_len - dist) / (float)sum_len;
}
