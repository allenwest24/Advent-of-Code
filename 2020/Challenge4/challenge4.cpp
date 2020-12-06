#include <iostream>
#include <fstream>
#include <string>
#include <vector>

using namespace std;

int main ()
{
	// Load the file.
	ifstream file("input.txt");

	// Initialize variables.
	string str;
	vector<string> v;
	vector<string> parts;
	string index1, index2, letter, pass, tmp;
	int counter = 0;
	int valid = 0;

	// Read the input file into a vector of strings.
	while (getline(file, str)) {
		v.push_back(str);
	}

	// Parse through each line, separate into elements, and check password for validity.
	for (int ii = 0; ii < v.size(); ii++) {
		// Clear the variables.
		index1 = index2 = letter = pass = tmp = "";

		// Separate by elements.
		for (int jj = 0; jj < v[ii].length(); jj++) {
			if (v[ii][jj] == '-') {
				index1 = tmp;
				tmp = "";
			}
			else if (v[ii][jj] == ' ' && v[ii][jj-1] != ':') {
				index2 = tmp;
				tmp = "";
			}
			else if (v[ii][jj] == ':') {
				letter = tmp;
				tmp = "";
			}
			else if (jj == v[ii].length() - 1) {
				tmp += v[ii][jj];
				pass = tmp;
				tmp = "";
			}
			else if (v[ii][jj] != ' ') {
				tmp += v[ii][jj];
			}
		}

		// Check if the letter appears in exaclt ONE of the two provided positions.
		if ((pass[stoi(index1) - 1] == letter[0] && pass[stoi(index2) -1] != letter[0]) || (pass[stoi(index1) -1] != letter[0] && pass[stoi(index2) -1] == letter[0])) {
			valid += 1;
		}
	}

	// Print results.
	cout << valid << '\n';
	return 0;
}
