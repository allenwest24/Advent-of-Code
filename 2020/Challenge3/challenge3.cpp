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
	string min, max, letter, pass, tmp;
	int counter = 0;
	int valid = 0;

	// Read the input file into a vector of strings.
	while (getline(file, str)) {
		v.push_back(str);
	}

	// Parse through each line, separate into elements, and check password for validity.
	for (int ii = 0; ii < v.size(); ii++) {
		// Clear the variables.
		min = max = letter = pass = tmp = "";

		// Separate by elements.
		for (int jj = 0; jj < v[ii].length(); jj++) {
			if (v[ii][jj] == '-') {
				min = tmp;
				tmp = "";
			}
			else if (v[ii][jj] == ' ' && v[ii][jj-1] != ':') {
				max = tmp;
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

		// Count occurances of the given letter in the current password.
		for (int kk = 0; kk < pass.length(); kk++) {
			if (pass[kk] == letter[0]) {
				counter += 1;
			}
		}
		
		// CHeck frequency against the provided rules.
		if (counter >= stoi(min) && counter <= stoi(max)) {
			valid += 1;
		}

		// Zero out the counter.
		counter = 0;
	}

	// Print results.
	cout << valid << '\n';
	return 0;
}
