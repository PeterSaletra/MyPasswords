import os

from sample.utils import userInput


class StartUp:
    def __init__(self) -> None:
        self.size = os.get_terminal_size()[0]

    def showOptions(self) -> None:
        print(
            ("{:^" + str(self.size) + "}").format(
                "__  __       _____                                    _"
            )
        )
        print(
            ("{:^" + str(self.size) + "}").format(
                "|  \/  |     |  __ \                                  | |"
            )
        )
        print(
            ("{:^" + str(self.size) + "}").format(
                "   | \  / |_   _| |__) |_ _ ___ _____      _____  _ __ __| |___"
            )
        )
        print(
            ("{:^" + str(self.size) + "}").format(
                "    | |\/| | | | |  ___/ _` / __/ __\ \ /\ / / _ \| '__/ _` / __|"
            )
        )
        print(
            ("{:^" + str(self.size) + "}").format(
                "    | |  | | |_| | |  | (_| \__ \__ \\\\ V  V / (_) | | | (_| \__ \\"
            )
        )
        print(
            ("{:^" + str(self.size) + "}").format(
                "    |_|  |_|\__, |_|   \__,_|___/___/ \_/\_/ \___/|_|  \__,_|___/"
            )
        )
        print(
            ("{:^" + str(self.size) + "}").format(
                "            __/ |                                               "
            )
        )
        print(
            ("{:^" + str(self.size) + "}").format(
                "           |___/                                                "
            )
        )
        print(("{:^" + str(self.size) + "}").format("Welcom to MyPassword"))
        print(("{:^" + str(self.size) + "}").format("Your private password manager"))
        print("Options:")
        print("-> login(l) - to login to existing user")
        print("-> signUp(s) - to sign up as  user")
        print("-> quit(q) - exit program\n")

    def login(self) -> None:
        print("Please enter your login:")
        login = userInput("(login)")
        print("Please enter your password")
        password = input("(login)")

    def signUp(self) -> None:
        print("Please enter your database IP(if it is on your local machine type l)")
        dbAdd = userInput("(signUp)")
