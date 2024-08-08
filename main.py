import os
from subprocess import call


def clear() -> None:
    _ = call("clear" if os.name == "posix" else "cls")


def showOptions() -> None:
    size = os.get_terminal_size()[0]
    print(
        ("{:^" + str(size) + "}").format(
            "__  __       _____                                    _"
        )
    )
    print(
        ("{:^" + str(size) + "}").format(
            "|  \/  |     |  __ \                                  | |"
        )
    )
    print(
        ("{:^" + str(size) + "}").format(
            "   | \  / |_   _| |__) |_ _ ___ _____      _____  _ __ __| |___"
        )
    )
    print(
        ("{:^" + str(size) + "}").format(
            "    | |\/| | | | |  ___/ _` / __/ __\ \ /\ / / _ \| '__/ _` / __|"
        )
    )
    print(
        ("{:^" + str(size) + "}").format(
            "    | |  | | |_| | |  | (_| \__ \__ \\\\ V  V / (_) | | | (_| \__ \\"
        )
    )
    print(
        ("{:^" + str(size) + "}").format(
            "    |_|  |_|\__, |_|   \__,_|___/___/ \_/\_/ \___/|_|  \__,_|___/"
        )
    )
    print(
        ("{:^" + str(size) + "}").format(
            "            __/ |                                               "
        )
    )
    print(
        ("{:^" + str(size) + "}").format(
            "           |___/                                                "
        )
    )
    print(("{:^" + str(size) + "}").format("Welcom to MyPassword"))
    print(("{:^" + str(size) + "}").format("Your private password manager"))


if __name__ == "__main__":
    clear()
    showOptions()
    opt = input("")
