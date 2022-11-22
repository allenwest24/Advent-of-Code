// Shell of solutions that sets up reading from a file into an array.
#include <stdio.h>

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
        for (int ii = 0; ii < line; ii++) {
                printf("%s", input_data[ii]);
        }

        return 0;
}
