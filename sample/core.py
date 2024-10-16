import os
import sys
import time
from sample.color.printColor import printBlackOnWhite, printGreen, printRed
from sample.db import (
    createNewDB,
    dbDeletePassword,
    dbChangePasswordName,
    dbgetAllPasswords,
    dbgetPassword,
    dbinsertPassword,
    dbsignUp,
    dblogin,
)

from sample.utils import (
    clear,
    copy2clip,
    displayPass,
    userInput,
    generatePass,
    checkPass,
    checkPassExist
)

SIZE = os.get_terminal_size()[0]
LOGIN = False
USERNAME = ""

def main() -> None:
    clear()
    showASCII()
    showOptions()
    opt = userInput()

    while True:
        opt = opt.split() if opt else "`".split()
        # This ` is for ability to spam enter key to the infinity but must think off other way of doing that(later)

        if opt[0] == "l" or opt[0] == "login":
            login(opt[1::])
        elif opt[0] == "s" or opt[0] == "signup":
            signUp(opt[1::])
        elif opt[0] == "a" or opt[0] == "add":
            addPassword(opt[1::])
        elif opt[0] == "g" or opt[0] == "get":
            getPassword(opt[1::])
        elif opt[0] == "d" or opt[0] == "delete":
            deletePassword(opt[1::])
        elif opt[0] == "c" or opt[0] == "chage":
            changePassword(opt[1::])
        elif opt[0] == "h" or opt[0] == "help":
            showOptions()
        elif opt[0] == "q" or opt[0] == "quit":
            clear()
            print("Quiting program. Thanks for using it")
            time.sleep(1)
            clear()
            sys.exit()
        elif opt[0] == "clear":
            clear()
        else:
            if opt[0] != "`":
                printRed("\nWrong input try again\n")
        opt = userInput(username=USERNAME)


def showASCII() -> None:
    print(
        ("{:^" + str(SIZE) + "}").format(
            "__  __       _____                                    _"
        )
    )
    print(
        ("{:^" + str(SIZE) + "}").format(
            "|  \/  |     |  __ \                                  | |"
        )
    )
    print(
        ("{:^" + str(SIZE) + "}").format(
            "   | \  / |_   _| |__) |_ _ ___ _____      _____  _ __ __| |___"
        )
    )
    print(
        ("{:^" + str(SIZE) + "}").format(
            "    | |\/| | | | |  ___/ _` / __/ __\ \ /\ / / _ \| '__/ _` / __|"
        )
    )
    print(
        ("{:^" + str(SIZE) + "}").format(
            "    | |  | | |_| | |  | (_| \__ \__ \\\\ V  V / (_) | | | (_| \__ \\"
        )
    )
    print(
        ("{:^" + str(SIZE) + "}").format(
            "    |_|  |_|\__, |_|   \__,_|___/___/ \_/\_/ \___/|_|  \__,_|___/"
        )
    )
    print(
        ("{:^" + str(SIZE) + "}").format(
            "            __/ |                                               "
        )
    )
    print(
        ("{:^" + str(SIZE) + "}").format(
            "           |___/                                                "
        )
    )
    print(("{:^" + str(SIZE) + "}").format("Welcom to MyPassword"))
    print(("{:^" + str(SIZE) + "}").format("Your private password manager"))


def showOptions() -> None:
    print("\nOptions:")
    print(
        "-> login(l) - to login to existing user. Type this command alone or followed username and password alread"
    )
    print("-> signUp(s) - to sign up as  user")
    print(
        "-> add(a) - add new password. Type this command alone or followed password name to skip this in next step"
    )
    print(
        "-> get(g) - get password info. Type this command alone, with /all to get all passwords anf thier data or followed by name to skip next step"
    )
    print(
        "-> delete(d) - delete password. Type this command alone or followed by password name"
    )
    print(
        "-> chanege(c) - change password. Type this command alone or followed by password name"
    )
    print("-> clear - clear cmd\n")
    print("-> quit(q) - exit program\n")
    


def login(opt: list) -> None:
    global LOGIN, USERNAME

    if len(opt) == 2:
        if dblogin(opt[0], opt[1]):
            printGreen("\nSuccessfull login\n")
            LOGIN = True
            USERNAME = opt[0]
    else:
        while True:
            printBlackOnWhite("\nPlease enter your login")
            login = userInput("(login) ")
            if  login == "q" or login == "quit":
                break
            printBlackOnWhite("Please enter your password")
            password = userInput("(login) ")
            if password == "q" or password == "quit":
                break
            if dblogin(login, password):
                printGreen("Successfull login\n")
                LOGIN = True
                USERNAME = login
                break


def signUp(opt: list) -> None:
    global LOGIN, USERNAME
    if len(opt) == 1:
        if createNewDB():
            login = opt[0]
            while True:
                printBlackOnWhite("Please enter your password")
                password = userInput("(signUp) ")
                if password == "q" or  password == "quit":
                    break
                printBlackOnWhite("Please repeat your password")
                repeatPass = userInput("(signUp) ")

                if repeatPass == "q" or repeatPass == "quit":
                    break
                if password == repeatPass:
                    break
                else:
                    printRed("Somthing is wrong with password try again")

            if dbsignUp(login, password):
                LOGIN = True
                USERNAME = login
                printGreen("User successfully added")
    else:
        if createNewDB():
            printBlackOnWhite("Please enter your login:")
            login = userInput("(signUP) ")
            if login == "q" or login == "quit":
                return
            while True:
                printBlackOnWhite("Please enter your password")
                password = userInput("(signUp) ")
                if password == "q" or  password == "quit":
                    break
                printBlackOnWhite("Please repeat your password")
                repeatPass = userInput("(signUp) ")
                    
                if repeatPass == "q" or repeatPass == "quit":
                    break
                if password == repeatPass:
                    break
                else:
                    printRed("Something is wrong with password try again")

            if dbsignUp(login, password):
                LOGIN = True
                USERNAME = login
                printGreen("User successfully added")
            else:
                printRed("Something went wrong")


def addPassword(opt: list) -> None:
    global LOGIN, USERNAME
    if len(opt) == 1:
        if LOGIN:
            name = opt[0]

            printBlackOnWhite("\nEnter password or type g to generate one")
            password = userInput("(add) ", USERNAME)
            while not checkPass(password):
                if password == "g":
                    password = generatePass(USERNAME)
                    printBlackOnWhite("Your new password: " + password)
                    break
                else:
                    printRed("Your password is not safe.\nIt should contain min 2 lowercase characters, min 2 uppercase characters, min 2 digits and one special character")
                    password = userInput("(add) ", USERNAME)

            printBlackOnWhite("Entern url for this password")
            url = userInput("(add) ", USERNAME)

            printBlackOnWhite("Enter folder name or leave it blank")
            folder = userInput("(add) ", USERNAME)

            if dbinsertPassword(name, password, url, folder, USERNAME):
                printGreen("Password successfully added\n")
        else:
            printRed("\nYou are not logged. Please log in\n")
    else:
        if LOGIN:
            printBlackOnWhite("\nEnter your pass name")
            name = userInput("(add) ", USERNAME)

            printBlackOnWhite("Enter password or type g to generate one")
            password = userInput("(add) ", USERNAME)
            while not checkPass(password):
                if password == "g":
                    password = generatePass(USERNAME)
                    printBlackOnWhite("Your new password: " + password)
                    break
                else:
                    printRed("Your password is not safe.\nIt should contain min 2 lowercase characters, min 2 uppercase characters, min 2 digits and one special character")
                    password = userInput("(add) ", USERNAME)

            printBlackOnWhite("Enter url for this password")
            url = userInput("(add) ", USERNAME)

            printBlackOnWhite("Enter folder name or leave it blank")
            folder = userInput("(add) ", USERNAME)

            if dbinsertPassword(name, password, url, folder, USERNAME):
                printGreen("Password successfully added\n")
        else:
            printRed("\nYou are not logged. Please log in\n")


def getPassword(opt: list) -> None:
    global LOGIN, USERNAME
    if LOGIN:
        if opt and len(opt) == 1:
            if opt[0] == "/all":
                passData = dbgetAllPasswords(USERNAME)
                displayPass(passData)
            else:
                passData = dbgetPassword(opt[0], USERNAME)
                copy2clip(passData[0]["password"])
                displayPass(passData)
                printGreen("Password already clipped\n")
        else:
            printBlackOnWhite("\nPlease enter password name")
            passName = userInput("(get) ", USERNAME)
            passData = dbgetPassword(passName, USERNAME)
            copy2clip(passData[0]["password"])
            displayPass(passData)
            printGreen("Password already clipped\n")
    else:
        printRed("You are not logged. Please log in\n")
            

def deletePassword(opt: list) -> None:
    global LOGIN, USERNAME

    if LOGIN:
        if opt and len(opt) == 1:
            passName = opt[0]

            if checkPassExist(passName=passName, username=USERNAME):
                if dbDeletePassword(passName, USERNAME):
                    printGreen("Passoword successfully deleted")
                else:
                    printRed("Something wrong went wrong in database")
            else:
                printRed("Wrong password name. Name does not exists in database")
        else:
            passName = userInput("(delete)", USERNAME)

            if checkPassExist(passName=passName, username=USERNAME):
                if dbDeletePassword(passName, USERNAME):
                    printGreen("Passoword successfully deleted")
                else:
                    printRed("Something wrong went wrong in database")
            else:
                printRed("Wrong password name. Name does not exists in database")
    else:
        printRed("You are not logged. Please log in\n")


def changePassword(opt: list) -> None:
    global LOGIN, USERNAME

    if LOGIN:
        if opt and len(opt) == 1:
            passName = opt[0]

            if checkPassExist(passName=passName, username=USERNAME):
                
                if dbChangePasswordName(passName, USERNAME):
                    printGreen("Passoword successfully changed")
                else:
                    printRed("Something wrong with password name. Please try again")
            else:
                printRed("Wrong password name. Name does not exists in database")
        else:
            passName = userInput("(change)", USERNAME)

            if dbChangePasswordName(passName, USERNAME):
                printGreen("Passoword successfully changed")
            else:
                printRed("Something wrong with password name. Please try again")
    else:
        printRed("You are not logged. Please log in\n")
