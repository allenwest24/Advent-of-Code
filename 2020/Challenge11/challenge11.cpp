#include <iostream>
#include <fstream>
#include <string>
#include <vector>

using namespace std;

// Helper method for deciding whether a string is a number.
bool isNumeric(string str) {
	for (int ii = 0; ii < str.length(); ii++) {
		if (!isdigit(str[ii])) {
			return false;
		}
	}
	return true;
}
			

int main () {
	// Load the file.
	ifstream file("input.txt");

	// Initialize variables.
	string str;
	int valid = 0;
	string buff = "";
	vector<string> v;
	int total = 0;
	
	// Read the input file into a vector of strings.
	while (getline(file, str)) {
		if (str.empty()) {
			v.push_back(buff);
			buff = "";
		}
		else {
			buff += str;
		}
	}

	// Last line will not be an empty line so we have to add it.
	v.push_back(buff);
	buff = "";

	// Count unique answers by group.
	string curr;
	int unique = 1;
	// Each group.
	for (int ii = 0; ii < v.size(); ii++) {
		// Each answer.
		for (int jj = 0; jj < v[ii].size(); jj++) {
			curr = v[ii][jj];
			// Check unique.
			for (int kk = jj + 1; kk < v[ii].size(); kk++) {
				if (v[ii][jj] == v[ii][kk]) {
					unique = 0;
				}
			}
			if (unique) {
				total++;
			}
			unique = 1;
		}
	}

	// Print results.
	cout << total << '\n';
	return 0;
}
