from sample.startup import StartUp
from sample.utils import clear
from sample.utils import userInput
import time

if __name__ == "__main__":
    startUP = StartUp()
    clear()
    startUP.showOptions()
    opt = userInput()

    while True:
        if opt is "l" or opt is "login":
            startUP.login()
            break
        elif opt is "s" or opt is "signUp":
            startUP.signUp()
            break
        elif opt is "q" or opt is "quit":
            print("Quiting program")
            time.sleep(1)
            clear()
            break
        else:
            print("Invalid input. Try again")
            opt = userInput()
