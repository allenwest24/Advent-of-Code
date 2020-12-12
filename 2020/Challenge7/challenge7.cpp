#include <iostream>
#include <fstream>
#include <string>
#include <vector>

using namespace std;

int main () {
	// Load the file.
	ifstream file("input.txt");

	// Initialize variables.
	string str;
	int valid = 0;
	string buff = "";
	vector<string> v;
	int byr, iyr, eyr, hgt, hcl, ecl, pid, cid;
 	byr = iyr = eyr = hgt = hcl = ecl = pid = cid = 0;

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

	// Go through each passport and see if they have all the necessary values to be valid.
	for (int ii = 0; ii < v.size(); ii++) {
		string curr = v[ii];
		for (int jj = 0; jj < curr.length(); jj++) {
			if (curr[jj] == ':') {
				if (curr[jj - 3] == 'b') {
					byr = 1;
				}
				else if (curr[jj - 3] == 'i') {
					iyr = 1;
				}
				else if (curr[jj - 3] == 'e' && curr[jj - 2] == 'y') {
					eyr = 1;
				}
				else if (curr[jj - 3] == 'h' && curr[jj - 2] == 'g') {
					hgt = 1;
				}
				else if (curr[jj - 3] == 'h' && curr[jj - 2] == 'c') {
					hcl = 1;
				}
				else if (curr[jj - 3] == 'e' && curr[jj - 2] == 'c') {
					ecl = 1;
				}
				else if (curr[jj - 3] == 'p') {
					pid = 1;
				}
				else if (curr[jj - 3] == 'c') {
					cid = 1;
				}
			}
		}

		// Check if all fields are there.
		if (byr && iyr && eyr && hgt && hcl && ecl && pid) {
			valid++;
		}

		// Zero them out for the next passport.
		byr = iyr = eyr = hgt = hcl = ecl = pid = cid = 0;				
	}

	// Print results.
	cout << valid << '\n';
	return 0;
}
