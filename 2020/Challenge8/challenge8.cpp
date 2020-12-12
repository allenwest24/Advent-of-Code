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
	int byr, iyr, eyr, hgt, hcl, ecl, pid, cid;
 	byr = iyr = eyr = hgt = hcl = ecl = pid = cid = 0;
	string hair = "0123456789abcdef";
	vector<string> eyes = { "amb", "blu", "brn", "gry", "grn", "hzl", "oth"};


	// Read the input file into a vector of strings.
	while (getline(file, str)) {
		if (str.empty()) {
			v.push_back(buff);
			buff = "";
		}
		else {
			buff += str += " ";
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

				// Is valid birth year?
				if (curr[jj - 3] == 'b') {
					if (isNumeric(curr.substr(jj + 1, 4))) {
						if (stoi(curr.substr(jj + 1, 4)) >= 1920 && stoi(curr.substr(jj + 1, 4)) <= 2002 && curr[jj + 5] == ' ') {
							byr = 1;
						}
					}
				}

				// Is valid issue year?
				else if (curr[jj - 3] == 'i') {
					if (isNumeric(curr.substr(jj + 1, 4))) {
						if (stoi(curr.substr(jj + 1, 4)) >= 2010 && stoi(curr.substr(jj + 1, 4)) <= 2020 && curr[jj + 5] == ' ') {
							iyr = 1;
						}
					}
				}

				// Is valid expiration year?
				else if (curr[jj - 3] == 'e' && curr[jj - 2] == 'y') {
					if (isNumeric(curr.substr(jj + 1, 4))) {
						if (stoi(curr.substr(jj + 1, 4)) >= 2020 && stoi(curr.substr(jj + 1, 4)) <= 2030 && curr[jj + 5] == ' ') {
							eyr = 1;
						}
					}
				}

				// Is valid height?
				else if (curr[jj - 3] == 'h' && curr[jj - 2] == 'g') {

					// If written in cm:
					if (curr.substr(jj + 4, 2) == "cm") {
						if (isNumeric(curr.substr(jj + 1, 3))) {
							if (stoi(curr.substr(jj + 1, 3)) >= 150 && stoi(curr.substr(jj + 1, 3)) <= 193 && curr[jj + 6] == ' ') {
								hgt = 1;
							}
						}
					}
					
					// If written in in:
					else if (curr.substr(jj + 3, 2) == "in") {
						if (isNumeric(curr.substr(jj + 1, 2))) {
							if (stoi(curr.substr(jj + 1, 2)) >= 59 && stoi(curr.substr(jj + 1, 2)) <= 76 && curr[jj + 5] == ' ') {
								hgt = 1;
							}
						}
					}
				}

				// Is valid hair color?
				else if (curr[jj - 3] == 'h' && curr[jj - 2] == 'c') {
					if (curr[jj + 1] == '#') {
						int tmp = 0;
						for (int kk = 1; kk < 7; kk++) {
							tmp = 0;
							for (int ll = 0; ll < hair.length(); ll++) {
								if (hair[ll] == curr[jj + kk + 1]) {
									tmp = 1;
								}
							}
							if (!tmp || !(curr[jj + 8] == ' ')) {
								hcl = 0;
								break;
							}
							else {
								hcl = 1;
							}
						}	
					}
				}

				// Is valid eye color?
				else if (curr[jj - 3] == 'e' && curr[jj - 2] == 'c') {
					int eyeTmp = 0;
					for (int mm = 0; mm < eyes.size(); mm++) {
						if (curr.substr(jj + 1, 3) == eyes[mm]) {
							eyeTmp = 1;
						}
					}
					if (eyeTmp && !ecl) {
						ecl = 1;
					}
					else {
						ecl = 0;
					}
				}

				// Is valid passport ID?
				else if (curr[jj - 3] == 'p') {
					if (isNumeric(curr.substr(jj + 1, 9)) && curr[jj + 10] == ' '){			
						pid = 1;
					}
				}

				// We currently don't care if the country ID is valid.
				else if (curr[jj - 3] == 'c') {
					cid = 1;
				}
			}
		}

		// Check if all fields are present and are valid values.
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
