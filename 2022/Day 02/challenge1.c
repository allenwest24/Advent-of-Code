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

        int total = 0;
        // Do work here.
        for (int ii = 0; ii < line; ii++) {
                char foe = input_data[ii][0];
                char me = input_data[ii][2];
                if (foe == 'A') {
                        if (me == 'X') {
                                total += 3 + 1;
                        }
                        else if (me == 'Y') {
                                total += 6 + 2;
                        }
                        else if (me == 'Z') {
                                total += 0 + 3;
                        }
                }
                else if (foe == 'B') {
                        if (me == 'X') {
                                total += 0 + 1;
                        }
                        else if (me == 'Y') {
                                total += 3 + 2;
                        }
                        else if (me == 'Z') {
                                total += 6 + 3;
                        }
                }
                else if (foe == 'C') {
                        if (me == 'X') {
                                total += 6 + 1;
                        }
                        else if (me == 'Y') {
                                total += 0 + 2;
                        }
                        else if (me == 'Z') {
                                total += 3 + 3;
                        }
                }
        }

        printf("%d", total);

        return 0;
}
