import os
from subprocess import call


def clear() -> None:
    _ = call("clear" if os.name == "posix" else "cls")


def userInput(path: str = "") -> str:
    return input(path + "->")
