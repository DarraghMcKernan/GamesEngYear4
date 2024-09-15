#include <iostream>

int convertRomanToInt(std::string t_romanNumeral);
bool charactersValid(std::string t_romanNumeral);

int main() {
    std::string s;//users input
    bool romanNumeralValid = false;
    int convertedNumber = 0;

    std::cout << "Enter a Roman Numeral from 1 to 3999 using only capital letters\n";
    std::cout << "The entered value must not exceed 15 characters\n";

    while (romanNumeralValid == false) {
        std::cin >> s;
        if(s.length() <= 15) {
            if(charactersValid(s) == true) {//make sure only I,V,X,L,C,D or M are used as inputs
                romanNumeralValid = true;
            }
            else {
                std::cout << "Enter a Roman Numeral from 1 to 3999 using only capital letters\n";
                std::cout << "The entered value must not exceed 15 characters\n";
            }
        }
        else {
            std::cout << "Enter a Roman Numeral from 1 to 3999 using only capital letters\n";
            std::cout << "The entered value must not exceed 15 characters\n";
        }
    }

    convertedNumber = convertRomanToInt(s);// convert string into integer
    if(convertedNumber == 0) {
        std::cout << "Error incorrect format for Roman Numerals";
    }
    else std::cout << "Converted Number: " <<convertedNumber << "\n";

}

bool charactersValid(std::string t_romanNumeral) {
    for(int i = 0 ; i < t_romanNumeral.length(); i++) {
        if(t_romanNumeral[i] == 'I' || t_romanNumeral[i] == 'V' || t_romanNumeral[i] == 'X' || t_romanNumeral[i] == 'L'
            || t_romanNumeral[i] == 'C'||t_romanNumeral[i] == 'D'||t_romanNumeral[i] == 'M') {
            return true;
        }
    }
    return false;
}


int convertRomanToInt(std::string t_romanNumeral) {
    int convertedNumber = 0;
    bool IUsed = false;
    bool VUsed = false;
    bool XUsed = false;
    bool LUsed = false;
    bool CUsed = false;
    bool DUsed = false;
    bool MUsed = false;

    for(int i = 0 ; i < t_romanNumeral.length(); i++) {
        if(t_romanNumeral[i] == 'I') {
            if(IUsed == true) {
                return 0;
            }
            if(t_romanNumeral[i+1] == 'V' || t_romanNumeral[i+1] == 'X') {//if the current letter is I and the next letter is either V or X
                IUsed = true;//prevent I being used incorrectly such as IIII or IVI by only allowing it to be used once
                XUsed = false;// allows X to be reused if XIX is encountered as after XI it would not be allowed to use X again otherwise
                LUsed = true;//Roman numerals go from largest to smallest and so large numbers are not expected after small numbers
                CUsed = true;
                DUsed = true;
                MUsed = true;
                convertedNumber -= 1; //remove 1 so we get 4 or 9 when the next letter is added into the total
            }
            else if(t_romanNumeral[i+1] == 'I' && t_romanNumeral[i+2] == 'I') {// if II is next after current I do not allow V or X as IIIV and IIIX are invalid
                IUsed = true;
                VUsed = true;
                XUsed = true;
                LUsed = true;
                CUsed = true;
                DUsed = true;
                MUsed = true;
                convertedNumber += 3;
                i+=2;
            }
            else if(t_romanNumeral[i+1] == 'I') {
                IUsed = true;
                VUsed = true;
                XUsed = true;
                LUsed = true;
                CUsed = true;
                DUsed = true;
                MUsed = true;
                convertedNumber += 2;
                i+=1;
            }
            else {
                IUsed = true;//if only one I is by itself it can be followed by V or X and so allow them to still be used
                LUsed = true;
                CUsed = true;
                DUsed = true;
                MUsed = true;
                convertedNumber++;
            }
        }

        if(t_romanNumeral[i] == 'V') {
            if(VUsed == true) {
                return 0;
            }
            VUsed = true;
            XUsed = true;
            LUsed = true;
            CUsed = true;
            DUsed = true;
            MUsed = true;
            convertedNumber += 5;
        }

        if(t_romanNumeral[i] == 'X') {
            if(XUsed == true) {
                return 0;
            }
            XUsed = true;
            if(t_romanNumeral[i+1] == 'L' || t_romanNumeral[i+1] == 'C') {
                XUsed = true;
                CUsed = false;
                DUsed = true;
                MUsed = true;
                convertedNumber -= 10;
            }
            else if(t_romanNumeral[i+1] == 'X' && t_romanNumeral[i+2] == 'X') {
                XUsed = true;
                LUsed = true;
                CUsed = true;
                DUsed = true;
                MUsed = true;
                convertedNumber += 30;
                i+=2;
            }
            else if(t_romanNumeral[i+1] == 'X') {
                XUsed = true;
                LUsed = true;
                CUsed = true;
                DUsed = true;
                MUsed = true;
                convertedNumber += 20;
                i+=1;
            }
            else {
                XUsed = true;
                DUsed = true;
                MUsed = true;
                convertedNumber += 10;
            }
        }

        if(t_romanNumeral[i] == 'L') {
            if(LUsed == true) {
                return 0;
            }
            LUsed = true;
            CUsed = true;
            DUsed = true;
            MUsed = true;
            convertedNumber += 50;
        }

        if(t_romanNumeral[i] == 'C') {
            if(CUsed == true) {
                return 0;
            }
            if(t_romanNumeral[i+1] == 'D' || t_romanNumeral[i+1] == 'M') {
                CUsed = true;
                MUsed = false;//re-enable M as it must be used for CM to be added
                convertedNumber -= 100;
            }
            else if(t_romanNumeral[i+1] == 'C' && t_romanNumeral[i+2] == 'C') {
                CUsed = true;
                DUsed = true;
                MUsed = true;//does not have to be re-enabled here as CM is not expected
                convertedNumber += 300;
                i+=2;
            }
            else if(t_romanNumeral[i+1] == 'C') {
                CUsed = true;
                DUsed = true;
                MUsed = true;
                convertedNumber += 200;
                i+=1;
            }
            else {
                CUsed = true;
                convertedNumber += 100;
            }
        }

        if(t_romanNumeral[i] == 'D') {
            if(DUsed == true) {
                return 0;
            }
            DUsed = true;
            MUsed = true;//DM is not allowed
            convertedNumber += 500;
        }

        if(t_romanNumeral[i] == 'M') {
            if(MUsed == true) {
                return 0;
            }
            if(t_romanNumeral[i+1] == 'M' && t_romanNumeral[i+2] == 'M') {
                MUsed = true;
                convertedNumber += 3000;
                i+=2;
            }
            else if(t_romanNumeral[i+1] == 'M') {
                MUsed = true;
                convertedNumber += 2000;
                i+=1;
            }
            else {
                MUsed = true;
                convertedNumber += 1000;
            }
        }
    }
    return convertedNumber;
}