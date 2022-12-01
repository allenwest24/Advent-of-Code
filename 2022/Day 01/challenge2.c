#include <stdio.h>
#include <stdlib.h>

// Constants.
#define MAX_LINES 4096
#define MAX_LEN 1024

// Main method.
int main(void) {
        // Create static size array.
        char input_data[MAX_LINES][MAX_LEN];

        // Open and read the file.
        FILE *file;
        file = fopen("input.txt", "r");
        if (file == NULL) {
                printf("Could not open file.\n");
                return 1;
        }

        // Store the input data in the array we created.
        int line = 0;
        while (!feof(file) && !ferror(file)) {
                if (fgets(input_data[line], MAX_LEN, file) != NULL) {
                        line++;
                }
        }

        // Close the input file.
        fclose(file);

        // Do work here.
        int curr_sum = 0;
        int first = 0;
        int second = 0;
        int third = 0;
        for (int ii = 0; ii < line; ii++) {
                int curr = atoi(input_data[ii]);
                if (curr == 0) {
                        if (first == 0 || first <= curr_sum) {
                                third = second;
                                second = first;
                                first = curr_sum;
                        }
                        else if (second == 0 || second <= curr_sum) {
                                third = second;
                                second = curr_sum;
                        }
                        else if (third == 0 || third <= curr_sum) {
                                third = curr_sum;
                        }
                        curr_sum = 0;
                }
                else {
                        curr_sum += curr;
                }
        }

        printf("%d", first + second + third);

        return 0;
}
