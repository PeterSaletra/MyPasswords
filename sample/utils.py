import os
import string
import secrets
import subprocess
import platform
import subprocess

from sample.color.color import Color
from sample.db import dbgetPassword

def copy2clip(data: str) -> None:
    if platform.system() == "Linux":
        subprocess.run("xclip -selection clipboard", text=True, input=data, shell=True)
    elif platform.system() == "Windows":
        subprocess.run("clip", text=True, input=data, shell=True)
    elif platform.system() == "Darwin":
        subprocess.run("pbcopy", text=True, input=data)


def clear() -> None:
    _ = subprocess.call("clear" if os.name == "posix" else "cls")


def userInput(path: str = "", username: str = "") -> str:
    if username:
        return input(Color.letters.green + "@" + username + Color.reset + " " + path + "-> ")
    return input(path + "-> ")


def displayPass(passwords: list) -> None:
    for pw in passwords:
        print("\nName: ", pw["name"])
        print("Password: ", pw["password"])
        print("URL: ", pw["url"])
        print()


def generatePass(username: str) -> str:
    print("Enter length of password")
    length = userInput("(add)", username)

    alphabet = string.ascii_letters + string.digits + string.punctuation
    while True:
        password = "".join(secrets.choice(alphabet) for i in range(int(length)))
        if (
            sum(c.islower() for c in password) >= 4
            and sum(c.isupper() for c in password) >= 4
            and sum(c.isdigit() for c in password) >= 4
        ):
            break

    return password


def checkPass(password: str) -> bool:
    if len(password) < 8:
        return False

    if (
        sum(c.islower() for c in password) >= 2
        and sum(c.isupper() for c in password) >= 2
        and sum(c.isdigit() for c in password) >= 2
        and sum(c in string.punctuation for c in password) >= 1
    ):
        return True

    return False


def checkPassExist(passName: str, username: str) -> bool:
    passData = dbgetPassword(passName=passName, username=username)

    if len(passData) == 0:
        return False
    
    return True